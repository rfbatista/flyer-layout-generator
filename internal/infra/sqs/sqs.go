package sqs

import (
	"algvisual/internal/entities"
	"algvisual/internal/infra/config"
	"algvisual/internal/shared"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func NewSQS(cfg config.AppConfig) (*SQS, error) {
	waitTime := 10
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
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
}

func (s SQS) PublishAdaptation(a entities.AdaptationBatch) error {
	raw, err := json.Marshal(a)
	if err != nil {
		return err
	}
	_, err = s.svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(string(raw)),
		QueueUrl:     s.adaptationURL,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s SQS) PublishReplication() error {
	return nil
}

func (s SQS) PublishLayoutJob() error {
	return nil
}

func (s SQS) PublishImageJob() error {
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
	var batch entities.AdaptationBatch
	err = json.Unmarshal([]byte(*body), &batch)
	if err != nil {
		return nil, errors.New("error in unmarshal adaptation evento")
	}
	return &AdaptationBatchEvent{ID: batch.ID}, nil
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
