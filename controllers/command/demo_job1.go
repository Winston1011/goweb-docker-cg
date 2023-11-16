package command

import (
	"goweb-docker-cg/models/demo"

	"github.com/gin-gonic/gin"
)

func DemoJob1(ctx *gin.Context) error {
	_, err := demo.GetDemoByName(ctx, []string{"goweb-docker-cg"})
	if err != nil {
		return err
	}

	return nil
}
