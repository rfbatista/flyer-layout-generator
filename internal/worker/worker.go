package worker

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var ticker = time.NewTicker(5 * time.Second)

var (
	quit         = make(chan struct{})
	gracefulStop = make(chan os.Signal, 1)
)

func NewWorkerPool(
	client *infra.ImageGeneratorClient,
	queries *database.Queries,
	db *pgxpool.Pool,
	config *infra.AppConfig,
	log *zap.Logger,
	sse *infra.ServerSideEventManager,
) (WorkerPool, error) {
	return WorkerPool{
		client:  client,
		queries: queries,
		db:      db,
		config:  config,
		log:     log,
		sse:     sse,
	}, nil
}

type WorkerPool struct {
	client  *infra.ImageGeneratorClient
	queries *database.Queries
	db      *pgxpool.Pool
	config  *infra.AppConfig
	log     *zap.Logger
	sse     *infra.ServerSideEventManager
}

func (w WorkerPool) Start() {
	w.log.Info("starting worker pool")
	signal.Notify(gracefulStop, syscall.SIGTERM)
	go func() {
		defer func() {
			w.log.Warn("closing worker thread")
			if r := recover(); r != nil {
				err, ok := r.(error)
				if ok {
					w.log.Error("panic error in worker", zap.Error(err))
				} else {
					w.log.Error("unknown panic error in worker")
				}
			}
		}()
		for {
			select {
			case <-ticker.C:
				createWorkerPool(w.client, w.queries, w.db, w.config, w.log, w.sse)
			case <-quit:
				ticker.Stop()
				w.log.Info("closing worker pool")
				return
			case <-gracefulStop:
				w.log.Info("closing worker pool")
				ticker.Stop()
				return
			}
		}
	}()
}

func (w WorkerPool) Close() {
	quit <- struct{}{}
}

func worker(
	wid int,
	wg *sync.WaitGroup,
	id int32,
	client *infra.ImageGeneratorClient,
	queries *database.Queries,
	db *pgxpool.Pool,
	config *infra.AppConfig,
	log *zap.Logger,
	sse *infra.ServerSideEventManager,
) {
	defer func() {
		log.Debug("closing worker")
		wg.Done()
		if r := recover(); r != nil {
			err, ok := r.(error)
			if ok {
				log.Error("panic error in worker", zap.Error(err))
			} else {
				log.Error("unknown panic error in worker")
			}
		}
	}()
	err := sse.BroadCastEvent(infra.NewEvent("JOB_BATCH_UPDATE"))
	if err != nil {
		log.Error("falha ao enviar evento sse", zap.Error(err))
	}
	log.Debug(fmt.Sprintf("starting request job %d", wid))
	err = layoutgenerator.StartRequestJobUseCase(
		client,
		queries,
		db,
		*config,
		log,
		layoutgenerator.StartRequestJobInput{ID: id},
	)
	if err != nil {
		log.Error("Falha no processamento da layout request job", zap.Error(err))
	}
	err = sse.BroadCastEvent(infra.NewEvent("JOB_BATCH_UPDATE"))
	if err != nil {
		log.Error("falha ao enviar evento sse", zap.Error(err))
	}
	log.Debug(fmt.Sprintf("finishing request job %d", wid))
}

func createWorkerPool(
	client *infra.ImageGeneratorClient,
	queries *database.Queries,
	db *pgxpool.Pool,
	config *infra.AppConfig,
	log *zap.Logger,
	sse *infra.ServerSideEventManager,
) error {
	log.Debug("buscando novo bactch de jobs")
	out, err := layoutgenerator.ListLayoutRequestJobsNotStartedUseCase(
		context.TODO(),
		queries,
		config,
	)
	if err != nil {
		log.Error("Falha na listagem de jobs não inicializados", zap.Error(err))
		return err
	}
	if len(out.Data) == 0 {
		log.Debug("nenhum job para ser processado")
		return nil
	}
	log.Debug(fmt.Sprintf("iniciando processamento de %d jobs", len(out.Data)))
	var wg sync.WaitGroup
	for idx, l := range out.Data {
		log.Debug("spawning request workers")
		wg.Add(1)
		go worker(idx, &wg, l.ID, client, queries, db, config, log, sse)
	}
	log.Debug("waiting request workers")
	wg.Wait()
	log.Debug("request workers finished")
	return nil
}
