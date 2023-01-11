package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/simminni/appmetadata-go/dao"
	"github.com/simminni/appmetadata-go/model"
)

// service: getAppMetadata
// getAppMetadata responds with the list of all app metadata as YAML.
func GetAppMetadata(c *gin.Context) {
	c.YAML(http.StatusOK, dao.GetAppMetadata())
}

// service: getAppMetadataByTiTle
// getAppMetadataByTiTle locates the app metadata by TiTle value
// then returns all the app metadata that match TiTle as a response.
func GetAppMetadataByTiTle(c *gin.Context) {
	title := c.Param("title")
	var matching_app_metadatas = dao.GetAppMetadataByTiTle(title)

	if len(matching_app_metadatas) != 0 {
		c.YAML(http.StatusOK, matching_app_metadatas)
		return
	} else {
		c.YAML(http.StatusNotFound, gin.H{"message": "app metadata not found"})
	}
}

// service: postAppMetadata
// postAppMetadata adds an app metadata from YAML received in the request body.
func PostAppMetadata(c *gin.Context) {
	var new_app_metadata model.App_metadata

	// Call BindYAML to bind the received YAML to
	// new_app_metadata.
	if err := c.BindYAML(&new_app_metadata); err != nil {
		log.Printf("Invalid payload received.")
		c.YAML(http.StatusBadRequest, gin.H{"message": "Invalid payload." + err.Error()})
		return
	}

	// persist metadata
	dao.PostAppMetadata(new_app_metadata)
	c.YAML(http.StatusCreated, new_app_metadata)
}
