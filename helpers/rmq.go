package helpers

import (
	"goweb-docker-cg/conf"

	"goweb-docker-cg/pkg/golib/v2/rmq"
	"goweb-docker-cg/pkg/golib/v2/zlog"
)

func InitRmq() {
	for _, produceConf := range conf.RConf.Rmq.Producer {
		zlog.Debugf(nil, "register Rmq producer: %s", produceConf.Service)
		if err := rmq.InitProducer(produceConf); err != nil {
			panic("register Rmq producer[" + produceConf.Service + "] error: " + err.Error())
		}

		if err := rmq.StartProducer(produceConf.Service); err != nil {
			panic("Rmq StartProducer[" + produceConf.Service + "] error: " + err.Error())
		}
	}
}
