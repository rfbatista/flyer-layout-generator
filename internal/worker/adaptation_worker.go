package worker

import (
	"algvisual/internal/infra/sqs"
	"algvisual/internal/shared"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

func (w WorkerPool) StartAdaptationWorker(wg *sync.WaitGroup) error {
	w.log.Info("starting adaptation worker")
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			w.log.Warn("closing adaptation worker thread")
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
			case <-quit:
				ticker.Stop()
				w.log.Info("closing worker pool")
				return
			case <-gracefulStop:
				w.log.Info("closing worker pool")
				ticker.Stop()
				return
			case <-ticker.C:
				err := w.AdaptationWorker(0)
				if err != nil {
					w.log.Info("adaptation worker failed", zap.Error(err))
				}
			}
		}
	}()
	return nil
}

func (w WorkerPool) AdaptationWorker(stack int) error {
	stack += 1
	if stack > 1000000 {
		return nil
	}
	select {
	case <-quit:
		ticker.Stop()
		w.log.Info("closing worker pool")
		//TODO: ao retornar deve se fechar a rotina anterior tambem
		return nil
	case <-gracefulStop:
		w.log.Info("closing worker pool")
		ticker.Stop()
		return nil
	default:
		event, err := w.sqs.PullAdaptationEvent()
		switch err := err.(type) {
		case *shared.AppError:
			if err.ErrorCode == sqs.NO_NEW_EVENTS {
				w.log.Info("adaptation queue is empty")
			}
			return err
		case error:
			w.log.Error("error pulling adaptation event", zap.Error(err))
			return err
		default:
			w.log.Info(fmt.Sprintf("adp wok %d", event.ID))
		}
		return w.AdaptationWorker(stack)
	}
}
