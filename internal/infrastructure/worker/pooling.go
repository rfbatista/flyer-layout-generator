package worker

import (
	"algvisual/internal/shared"
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

func (w *WorkerPool) AddQueuePoolingWorker(ctx context.Context, wg *sync.WaitGroup, queueURL string, consumer shared.Consumer) error {
	w.log.Info(fmt.Sprintf("starting to pull events from %s", queueURL))
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			w.log.Warn("closing queue pooling thread")
			if r := recover(); r != nil {
				err, ok := r.(error)
				if ok {
					w.log.Error("panic error in worker", zap.Error(err))
				} else {
					w.log.Error("unknown panic error in queue pooling")
				}
			}
		}()
		backOff := false
		for {
			select {
			case <-quit:
				ticker.Stop()
				w.log.Info("closing queue thread")
				return
			case <-gracefulStop:
				w.log.Info("closing queue thread")
				ticker.Stop()
				return
			case <-ticker.C:
				w.log.Info(fmt.Sprintf("[tick] pulling event from %s", queueURL))
				events, err := w.sqs.PullEvent(queueURL)
				if err != nil || len(events) == 0 {
					backOff = true
				} else {
					backOff = false
				}
				for _, event := range events {
					ev := event
					w.pool.Submit(func() {
						err := consumer(ctx, &ev)
						if err != nil {
							backOff = true
						} else {
							w.log.Debug("removing event from queue")
							err := w.sqs.DeleteEvent(ctx, queueURL, event.Receipt)
							if err != nil {
								w.log.Warn("failed to remove evento from queue")
							}
						}
					})
				}
			default:
				if !backOff {
					w.log.Info(fmt.Sprintf("[default] pulling event from %s", queueURL))
					events, err := w.sqs.PullEvent(queueURL)
					if err != nil || len(events) == 0 {
						backOff = true
					} else {
						backOff = false
					}
					for _, event := range events {
						ev := event
						w.log.Info("submiting event", zap.String("event_id", event.ID))
						w.pool.Submit(func() {
							err := consumer(ctx, &ev)
							if err != nil {
								backOff = true
							} else {
								w.log.Debug("removing event from queue")
								err := w.sqs.DeleteEvent(ctx, queueURL, event.Receipt)
								if err != nil {
									w.log.Warn("failed to remove evento from queue")
								}
							}
						})
					}
				}
			}
		}
	}()
	return nil
}
