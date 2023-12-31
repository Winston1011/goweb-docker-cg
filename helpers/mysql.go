package helpers

import (
	"goweb-docker-cg/conf"

	"gorm.io/gorm"
	"goweb-docker-cg/pkg/golib/v2/base"
)

var (
	MysqlClientDemo *gorm.DB
	MysqlClientTest *gorm.DB
)

func InitMysql() {
	var err error
	for name, dbConf := range conf.RConf.Mysql {
		switch name {
		case "gc-hk":
			MysqlClientDemo, err = base.InitMysqlClient(dbConf)
		case "test":
			MysqlClientTest, err = base.InitMysqlClient(dbConf)
		}

		if err != nil {
			panic("mysql connect error: %v" + err.Error())
		}
	}
}

func CloseMysql() {
}
