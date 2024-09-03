package queue

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/shared"
	"context"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
)

func NewSQS(cfg config.AppConfig, log *zap.Logger) (*SQS, error) {
	waitTime := 5
	cred := credentials.NewStaticCredentials(cfg.AWS.AccessKey, cfg.AWS.SecretKey, "")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: cred,
	})
	if err != nil {
		return nil, err
	}
	svc := sqs.New(sess)
	adaptationResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &cfg.SQSConfig.AdaptationQueueName,
	})
	if err != nil {
		return nil, err
	}
	replicationResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &cfg.SQSConfig.ReplicationQueueName,
	})
	if err != nil {
		return nil, err
	}
	rendererResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &cfg.SQSConfig.RendererJobQueue,
	})
	if err != nil {
		return nil, err
	}
	layoutResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &cfg.SQSConfig.LayoutJobQueue,
	})
	if err != nil {
		return nil, err
	}
	_, err = svc.SetQueueAttributes(&sqs.SetQueueAttributesInput{
		QueueUrl: adaptationResult.QueueUrl,
		Attributes: aws.StringMap(map[string]string{
			"ReceiveMessageWaitTimeSeconds": strconv.Itoa(aws.IntValue(&waitTime)),
		}),
	})
	if err != nil {
		return nil, err
	}
	return &SQS{
		sess:           sess,
		svc:            svc,
		adaptationURL:  adaptationResult.QueueUrl,
		replicationURL: replicationResult.QueueUrl,
		rendererURL:    rendererResult.QueueUrl,
		layoutURL:      layoutResult.QueueUrl,
		waitTime:       int64(waitTime),
		log:            log,
	}, nil
}

type SQS struct {
	sess           *session.Session
	svc            *sqs.SQS
	adaptationURL  *string
	layoutURL      *string
	rendererURL    *string
	replicationURL *string
	waitTime       int64
	log            *zap.Logger
}

func (s SQS) PublishBatch(queueURL string, bodies []ApplicationEvent) error {
	var divided [][]ApplicationEvent
	// sqs have a limit of 10 messages per batch
	chunkSize := 10
	for i := 0; i < len(bodies); i += chunkSize {
		end := i + chunkSize
		if end > len(bodies) {
			end = len(bodies)
		}
		divided = append(divided, bodies[i:end])
	}
	for _, msg := range divided {
		var messages []*sqs.SendMessageBatchRequestEntry
		for _, body := range msg {
			messages = append(messages, &sqs.SendMessageBatchRequestEntry{
				Id:          &body.ID,
				MessageBody: aws.String(string(body.Body)),
			})
		}
		_, err := s.svc.SendMessageBatch(&sqs.SendMessageBatchInput{
			QueueUrl: &queueURL,
			Entries:  messages,
		})
		if err != nil {
			return err
		}
		messages = nil
	}
	return nil
}

func (s SQS) Publish(queueURL string, body []byte) error {
	_, err := s.svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(1),
		MessageBody:  aws.String(string(body)),
		QueueUrl:     &queueURL,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s SQS) PullEvent(queueURL string) ([]shared.ApplicationEvent, error) {
	var events []shared.ApplicationEvent
	// tempo para evento reaparecer na fila após ser recebido e não removido da fila após processamento
	var timeout int64 = 300
	msgResult, err := s.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
		WaitTimeSeconds:     &s.waitTime,
	})
	if err != nil {
		return nil, err
	}
	for _, msg := range msgResult.Messages {
		if msg.Body == nil {
			s.log.Error("message with empty body")
			continue
		}
		if msg.MessageId == nil {
			s.log.Error("message with empty id")
			continue
		}
		if msg.ReceiptHandle == nil {
			s.log.Error("message with empty receipt")
			continue
		}
		events = append(events, shared.ApplicationEvent{
			ID:      *msg.MessageId,
			Receipt: *msg.ReceiptHandle,
			Body:    *msg.Body,
		})
	}
	return events, nil
}

func (s SQS) DeleteEvent(ctx context.Context, queueURL string, id string) error {
	_, err := s.svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &queueURL,
		ReceiptHandle: &id,
	})
	return err
}

func (s SQS) PublishAdaptation(a entities.Job) error {
	s.log.Debug("publishing adaptation")
	raw, err := json.Marshal(a)
	if err != nil {
		return err
	}
	_, err = s.svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(1),
		MessageBody:  aws.String(string(raw)),
		QueueUrl:     s.adaptationURL,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s SQS) PullAdaptationEvent() (*AdaptationBatchEvent, error) {
	var timeout int64 = 10
	msgResult, err := s.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            s.adaptationURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
		WaitTimeSeconds:     &s.waitTime,
	})
	if err != nil {
		return nil, err
	}
	if len(msgResult.Messages) == 0 {
		return nil, shared.NewError(NO_NEW_EVENTS, "no new events in queue", "")
	}
	body := msgResult.Messages[0].Body
	if body == nil {
		return nil, errors.New("empy body in event")
	}
	var batch entities.Job
	err = json.Unmarshal([]byte(*body), &batch)
	if err != nil {
		return nil, errors.New("error in unmarshal adaptation evento")
	}
	return &AdaptationBatchEvent{ID: *msgResult.Messages[0].MessageId, Adaptation: batch}, nil
}

func (s SQS) PullLayoutEvent() (*LayoutJobEvent, error) {
	var timeout int64 = 10
	msgResult, err := s.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            s.layoutURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
		WaitTimeSeconds:     &s.waitTime,
	})
	if err != nil {
		return nil, err
	}
	batchId := msgResult.Messages[0].Attributes[string(AdaptationBatchID)]
	i, err := strconv.ParseInt(*batchId, 10, 64)
	if err != nil {
		return nil, err
	}
	return &LayoutJobEvent{ID: i}, nil
}

func (s SQS) PullRendererEvent() (*RendererJobEvent, error) {
	var timeout int64 = 10
	msgResult, err := s.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            s.layoutURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
		WaitTimeSeconds:     &s.waitTime,
	})
	if err != nil {
		return nil, err
	}
	batchId := msgResult.Messages[0].Attributes[string(AdaptationBatchID)]
	i, err := strconv.ParseInt(*batchId, 10, 64)
	if err != nil {
		return nil, err
	}
	return &RendererJobEvent{ID: i}, nil
}

func (s SQS) PullReplicationEvent() (*ReplicationBatchEvent, error) {
	var timeout int64 = 10
	msgResult, err := s.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            s.layoutURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   &timeout,
		WaitTimeSeconds:     &s.waitTime,
	})
	if err != nil {
		return nil, err
	}
	batchId := msgResult.Messages[0].Attributes[string(AdaptationBatchID)]
	i, err := strconv.ParseInt(*batchId, 10, 64)
	if err != nil {
		return nil, err
	}
	return &ReplicationBatchEvent{ID: i}, nil
}
