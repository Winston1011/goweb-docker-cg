package helpers

import (
	"goweb-docker-cg/pkg/golib/v2/kms"
)

var Kms kms.Kms

func InitKms() {
	Kms = kms.Init()
}
