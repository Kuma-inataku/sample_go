package router

import (
	sampleHandler "gin-api-sample/internal/sample/interface/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Sampleハンドラーの設定
	sampleHandler := sampleHandler.NewSampleHandler()
	r.GET("/samples", sampleHandler.GetSamples)
	r.POST("/samples", sampleHandler.CreateSample)
	r.PUT("/samples/:id", sampleHandler.UpdateSample)
	r.DELETE("/samples/:id", sampleHandler.DeleteSample)

	return r
}
