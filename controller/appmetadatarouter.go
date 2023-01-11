package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/simminni/appmetadata-go/service"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/appmetadata", service.GetAppMetadata)
	router.GET("/appmetadata/:title", service.GetAppMetadataByTiTle)
	router.POST("/appmetadata", service.PostAppMetadata)
	return router
}
