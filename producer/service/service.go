package service

import (
	"Feed/producer/httpsvr"
	"Feed/producer/service/kafka"
)

type Service struct {
	Kafka   *kafka.Producer
	Httpsvr *httpsvr.HttpServer
}

func NewService() (*Service, error) {

	ka, err := kafka.NewProducer()
	if err != nil {
		return nil, err
	}

	httpsvr := httpsvr.NewHttpServer(nil, ka)

	return &Service{
		Kafka:   ka,
		Httpsvr: httpsvr,
	}, nil
}

func (this *Service) Start() {

	//runtime.GOMAXPROCS(runtime.NumCPU()) // 利用CPU多核来处理HTTP请求
	this.Kafka.Run()
	this.Httpsvr.Start()

}
