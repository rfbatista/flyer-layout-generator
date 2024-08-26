package sqs

type MessageAtributte string

const (
	AdaptationBatchID MessageAtributte = "adaptation_batch_id"
)

type AdaptationBatchEvent struct {
	ID int64
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
