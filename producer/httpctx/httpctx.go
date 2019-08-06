package httpctx

import (
	"Feed/producer/service/elastic"
	"Feed/producer/service/kafka"
)

type ServerContext struct {
	EsManager *elastic.ElasticManager
	Kafka     *kafka.Producer
}
