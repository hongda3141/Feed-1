package service

import (
	"Feed/producer/httpsvr"
	"Feed/producer/service/kafka"
)

type Service struct {
	Kafka   *kafka.KafkaManager
	Httpsvr *httpsvr.HttpServer
}

func NewService() *Service {

	ka := kafka.NewKafkaManager()

	httpsvr := httpsvr.NewHttpServer(nil)

	return &Service{
		Kafka:   ka,
		Httpsvr: httpsvr,
	}
}

func (this *Service) Start() {

	//runtime.GOMAXPROCS(runtime.NumCPU()) // 利用CPU多核来处理HTTP请求

	this.Httpsvr.Start()

}
