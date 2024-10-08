package queue

import "algvisual/internal/domain/entities"

type MessageAtributte string

const (
	AdaptationBatchID MessageAtributte = "adaptation_batch_id"
)

type AdaptationBatchEvent struct {
	ID         string
	Adaptation entities.Job
}

type LayoutJobEvent struct {
	ID int64
}

type RendererJobEvent struct {
	ID int64
}

type ReplicationBatchEvent struct {
	ID int64
}

type ApplicationEvent struct {
	ID   string
	Body string
}
