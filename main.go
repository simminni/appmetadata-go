package main

import (
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	log.SetPrefix("main: ")
	log.SetFlags(0)

	router := gin.Default()
	router.GET("/appmetadata", getAppMetadata)
	router.GET("/appmetadata/:title", getAppMetadataByTiTle)
	router.POST("/appmetadata", postAppMetadata)
	router.Run("localhost:8080")
}

type app_metadata struct {
	//title: Valid App 1
	Title string `yaml:"title" binding:"required"`
	//version: 0.0.1
	Version string `yaml:"version" binding:"required"`
	/*maintainers:
			- name: firstmaintainer app1
	  		  email: firstmaintainer@hotmail.com
			- name: secondmaintainer app1
	  		  email: secondmaintainer@gmail.com
	*/
	Maintainers []Maintainer `yaml:"maintainers,omitempty" binding:"required"`
	//company: Random Inc.
	Company string `yaml:company binding:"required"`
	//website: https://website.com
	Website string `yaml:"website" binding:"required"`
	//source: https://github.com/random/repo
	Source string `yaml:"source" binding:"required"`
	//license: Apache-2.0
	License string `yaml:"license" binding:"required"`
	/*
			description: |
		 		### Interesting Title
		 		Some application content, and description
	*/
	Description string `yaml:"description" binding:"required"`
}

type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

var app_metadatas = []app_metadata{

	{
		Title:   "Valid App 1",
		Version: "0.0.1",
		Maintainers: []Maintainer{
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
		Maintainers: []Maintainer{
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

// getAppMetadata responds with the list of all app metadata as JSON.
func getAppMetadata(c *gin.Context) {
	log.Printf("in getAppMetadata")
	c.YAML(http.StatusOK, app_metadatas)
}

// postAppMetadata adds an app metadata from YAML received in the request body.
func postAppMetadata(c *gin.Context) {
	log.Printf("in postAppMetadata")
	var new_app_metadata app_metadata

	// Call BindYAML to bind the received YAML to
	// new_app_metadata.
	if err := c.BindYAML(&new_app_metadata); err != nil {
		log.Printf("Invalid payload received.")
		c.YAML(http.StatusBadRequest, gin.H{"message": "Invalid payload. " + err.Error()})
		return
	}

	// Add the new album to the slice.
	app_metadatas = append(app_metadatas, new_app_metadata)
	c.YAML(http.StatusCreated, new_app_metadata)
}

// getAppMetadataByTiTle locates the app metadata by TiTle value
// then returns all the app metadata that match TiTle as a response.
func getAppMetadataByTiTle(c *gin.Context) {
	log.Printf("in getAppMetadataByTiTle")
	title := c.Param("title")
	var matching_app_metadatas []app_metadata

	// Loop over the list of appmetadata, looking for
	// an app metadata whose name value matches the parameter.
	for _, a := range app_metadatas {
		log.Println("'" + a.Title + "'")
		log.Println("'" + title + "'")
		if a.Title == title {
			matching_app_metadatas = append(matching_app_metadatas, a)
		}
	}

	if len(matching_app_metadatas) != 0 {
		c.YAML(http.StatusOK, matching_app_metadatas)
		return
	} else {
		c.YAML(http.StatusNotFound, gin.H{"message": "app metadata not found"})
	}
}
