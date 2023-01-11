package dao

import (
	"log"

	"github.com/simminni/appmetadata-go/model"
)

var app_metadatas = []model.App_metadata{}

// GetAppMetadata responds with the list of all app metadata persisted in memory
func GetAppMetadata() []model.App_metadata {
	return app_metadatas
}

// GetAppMetadataByTiTle locates the app metadata by TiTle value
// then returns all the app metadata that match TiTle.
func GetAppMetadataByTiTle(title string) []model.App_metadata {
	var matching_app_metadatas []model.App_metadata

	// Loop over the list of appmetadata, looking for
	// an app metadata whose name value matches the parameter.
	for _, a := range app_metadatas {
		log.Println("'" + a.Title + "'")
		log.Println("'" + title + "'")
		if a.Title == title {
			matching_app_metadatas = append(matching_app_metadatas, a)
		}
	}
	return matching_app_metadatas
}

// PostAppMetadata persists app metadata entry
func PostAppMetadata(new_app_metadata model.App_metadata) model.App_metadata {

	// Add the new album to the slice.
	app_metadatas = append(app_metadatas, new_app_metadata)
	return new_app_metadata
}
