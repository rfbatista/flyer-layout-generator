package worker

import (
	"algvisual/internal/database"
	"algvisual/internal/infra"
	"algvisual/internal/layoutgenerator"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var ticker = time.NewTicker(5 * time.Second)

var quit = make(chan struct{})

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
	go func() {
		for {
			select {
			case <-ticker.C:
				createWorkerPool(w.client, w.queries, w.db, w.config, w.log, w.sse)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func worker(
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
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	err := sse.BroadCastEvent(infra.NewEvent("JOB_BATCH_UPDATE"))
	if err != nil {
		log.Error("falha ao enviar evento sse", zap.Error(err))
	}
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
	wg.Done()
}

func createWorkerPool(
	client *infra.ImageGeneratorClient,
	queries *database.Queries,
	db *pgxpool.Pool,
	config *infra.AppConfig,
	log *zap.Logger,
	sse *infra.ServerSideEventManager,
) error {
	log.Info("buscando novo bactch de jobs")
	out, err := layoutgenerator.ListLayoutRequestJobsNotStartedUseCase(context.TODO(), queries)
	if err != nil {
		log.Error("Falha na listagem de jobs não inicializados", zap.Error(err))
		return err
	}
	if len(out.Data) == 0 {
		log.Info("nenhum job para ser processado")
		return nil
	}
	log.Info(fmt.Sprintf("iniciando processamento de %d jobs", len(out.Data)))
	var wg sync.WaitGroup
	wg.Add(len(out.Data))
	for _, l := range out.Data {
		go worker(&wg, l.ID, client, queries, db, config, log, sse)
	}
	wg.Wait()
	return nil
}
