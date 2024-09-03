package worker

import (
	"algvisual/internal/application/consumers"
	"algvisual/internal/application/usecases/layoutgenerator"
	infra "algvisual/internal/infrastructure"
	"algvisual/internal/infrastructure/config"
	"algvisual/internal/infrastructure/database"
	"algvisual/internal/infrastructure/queue"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alitto/pond"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var ticker = time.NewTicker(1 * time.Second)

var (
	quit         = make(chan struct{}, 1)
	gracefulStop = make(chan os.Signal, 1)
)

func init() {
	signal.Notify(gracefulStop, syscall.SIGTERM)
}

func NewWorkerPool(
	client *infra.ImageGeneratorClient,
	queries *database.Queries,
	db *pgxpool.Pool,
	config *config.AppConfig,
	log *zap.Logger,
	sse *infra.ServerSideEventManager,
	wservice WorkerService,
	sqs *queue.SQS,
	adaptationBatchConsumer *consumers.AdaptatipnBatchConsumer,
	layoutJobConsumer *consumers.LayoutJobConsumer,
) (WorkerPool, error) {
	pool := pond.New(10, 100, pond.MinWorkers(5), pond.PanicHandler(func(i interface{}) {
		log.Warn(fmt.Sprintf("[pond] panic in worker: %v", i))
	}))
	return WorkerPool{
		pool:                    pool,
		client:                  client,
		queries:                 queries,
		db:                      db,
		config:                  config,
		log:                     log,
		sse:                     sse,
		sqs:                     sqs,
		serv:                    wservice,
		wg:                      &sync.WaitGroup{},
		adaptationBatchConsumer: adaptationBatchConsumer,
		layoutJobConsumer:       layoutJobConsumer,
	}, nil
}

type WorkerPool struct {
	pool                    *pond.WorkerPool
	client                  *infra.ImageGeneratorClient
	queries                 *database.Queries
	db                      *pgxpool.Pool
	config                  *config.AppConfig
	log                     *zap.Logger
	sse                     *infra.ServerSideEventManager
	serv                    WorkerService
	sqs                     *queue.SQS
	wg                      *sync.WaitGroup
	adaptationBatchConsumer *consumers.AdaptatipnBatchConsumer
	layoutJobConsumer       *consumers.LayoutJobConsumer
}

func (w WorkerPool) Start() error {
	w.log.Info("starting worker pool")
	var wg sync.WaitGroup
	w.wg = &wg
	go func() {
		err := w.AddQueuePoolingWorker(
			context.TODO(),
			w.wg,
			w.config.SQSConfig.AdaptationQueueName,
			w.adaptationBatchConsumer.Execute,
		)
		if err != nil {
			w.log.Error("failed to start adaptation pooling worker")
		}
	}()
	go func() {
		err := w.AddQueuePoolingWorker(
			context.TODO(),
			w.wg,
			w.config.SQSConfig.LayoutJobQueue,
			w.layoutJobConsumer.Execute,
		)
		if err != nil {
			w.log.Error("failed to start layout job pooling worker")
		}
	}()
	return nil
}

func (w WorkerPool) Close() {
	quit <- struct{}{}
	w.wg.Wait()
}

func (w WorkerPool) worker(
	wid int,
	wg *sync.WaitGroup,
	id int32,
) {
	ctx := context.TODO()
	defer func() {
		w.log.Debug("closing worker")
		wg.Done()
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				w.log.Error("panic error in worker", zap.Error(err))
			} else {
				w.log.Error("unknown panic error in worker")
			}
		}
	}()
	err := w.sse.BroadCastEvent(infra.NewEvent("JOB_BATCH_UPDATE"))
	if err != nil {
		w.log.Error("falha ao enviar evento sse", zap.Error(err))
	}
	w.log.Debug(fmt.Sprintf("starting request job %d", wid))
	_, err = w.serv.GenerateJob(ctx, GenerateJobInput{ID: id})
	if err != nil {
		w.log.Error("Falha no processamento da layout request job", zap.Error(err))
	}
	err = w.sse.BroadCastEvent(infra.NewEvent("JOB_BATCH_UPDATE"))
	if err != nil {
		w.log.Error("falha ao enviar evento sse", zap.Error(err))
	}
	w.log.Debug(fmt.Sprintf("finishing request job %d", wid))
}

func (w WorkerPool) createWorkerPool() error {
	out, err := layoutgenerator.ListLayoutRequestJobsNotStartedUseCase(
		context.TODO(),
		w.queries,
		w.config,
	)
	if err != nil {
		w.log.Error("Falha na listagem de jobs não inicializados", zap.Error(err))
		return err
	}
	if len(out.Data) == 0 {
		w.log.Debug("nenhum job para ser processado")
		return nil
	}
	w.log.Debug(fmt.Sprintf("iniciando processamento de %d jobs", len(out.Data)))
	var wg sync.WaitGroup
	for idx, l := range out.Data {
		w.log.Debug("spawning request workers")
		wg.Add(1)
		go w.worker(idx, &wg, l.ID)
	}
	w.log.Debug("waiting request workers")
	wg.Wait()
	w.log.Debug("request workers finished")
	return nil
}
