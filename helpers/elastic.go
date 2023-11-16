package helpers

import (
	"github.com/olivere/elastic/v7"
	"goweb-docker-cg/conf"

	"goweb-docker-cg/pkg/golib/v2/base"
)

var ElasticClient *elastic.Client

func InitEs() {
	var err error
	demoConf := conf.RConf.Elastic["demo"]
	ElasticClient, err = base.NewESClientV7(demoConf)
	if err != nil {
		panic("elastic init error: " + err.Error())
	}
}
