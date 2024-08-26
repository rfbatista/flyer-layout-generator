package config

type SQSConfig struct {
	AdaptationQueueName  string
	ReplicationQueueName string
	LayoutJobQueue       string
	RendererJobQueue     string
}
