package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/job"
	"goweb-docker-cg/pkg/golib/v2/command"
)

var Job *job.Job

func InitJob(g *gin.Engine) {
	Job = command.NewJob(g)
}
