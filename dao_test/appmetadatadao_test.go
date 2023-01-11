package dao_test

import (
	"testing"

	"github.com/simminni/appmetadata-go/dao"
	"github.com/simminni/appmetadata-go/model"
	"github.com/stretchr/testify/assert"
)

// Test for GetAppMetadata when app is first initialized
func TestGetAppMetadataOnLoad(t *testing.T) {
	assert.Equal(t, []model.App_metadata{}, dao.GetAppMetadata())
}

// Test for GetAppMetadata when app is first initialized
func TestPostAppMetadata(t *testing.T) {
	var new_app_metadata = model.App_metadata{
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
	}
	dao.PostAppMetadata(new_app_metadata)
	assert.Equal(t, []model.App_metadata{new_app_metadata}, dao.GetAppMetadata())
}
