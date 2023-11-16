package helpers

import (
	"goweb-docker-cg/conf"

	"goweb-docker-cg/pkg/golib/v2/kafka"
)

var DemoPubClient *kafka.PubClient

func InitKafkaProducer() {
	DemoPubClient = kafka.InitKafkaPub(conf.RConf.KafkaPub["demo"])

	if DemoPubClient == nil {
		panic("init redis failed!")
	}
}

func CloseKafkaProducer() {
	_ = DemoPubClient.CloseProducer()
}
