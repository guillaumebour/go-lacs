package server

import (
	"github.com/gin-gonic/gin"
)

type LacsOptions struct {
	UploadFileSizeLimit int64
}

func NewLacsRouter(options LacsOptions) *gin.Engine {
	router := gin.Default()
	router.MaxMultipartMemory = options.UploadFileSizeLimit
	router.POST("/compile", CompilationHandler)
	return router
}
