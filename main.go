package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/simminni/appmetadata-go/service"
)

// controller
func main() {
	log.SetFlags(1111)
	router := gin.Default()
	router.GET("/appmetadata", service.GetAppMetadata)
	router.GET("/appmetadata/:title", service.GetAppMetadataByTiTle)
	router.POST("/appmetadata", service.PostAppMetadata)
	router.Run("localhost:8080")
}

/*
// Sample app_metadata data structure to persist inmemory
var app_metadatas = []model.App_metadata{

	{
		Title:   "Valid App 1",
		Version: "0.0.1",
		Maintainers: []model.Maintainer{
			{
				Name:  "firstmaintainer app1",
				Email: "firstmaintainer@hotmail.com",
			},
			{
				Name:  "secondmaintainer app1",
				Email: "secondmaintainer@gmail.com",
			},
		},
		Company: "Random Inc.",
		Website: "https://website.com",
		Source:  "https://github.com/random/repo",
		License: "Apache-2.0",
		Description: `
		 		### Interesting Title
		 		Some application content, and description`,
	},
	{
		Title:   "Valid App 2",
		Version: "1.0.1",
		Maintainers: []model.Maintainer{
			{
				Name:  "AppTwo Maintainer",
				Email: "apptwo@hotmail.com",
			},
		},
		Company: "Upbound Inc.",
		Website: "https://upbound.io",
		Source:  "https://github.com/upbound/repo",
		License: "Apache-2.0",
		Description: `
		 		### Why app 2 is the best
				 Because it simply is...`,
	},
}
*/
