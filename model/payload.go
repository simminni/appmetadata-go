package model

// payload structure
type App_metadata struct {
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
	Company string `yaml:"company" binding:"required"`
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

/* Sample app_metadata data structure to persist inmemory
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
