package model

// payload structure
type App_metadata struct {
	//title: Valid App 1
	Title string `yaml:"title,omitempty" binding:"required"`
	//version: 0.0.1
	Version string `yaml:"version,omitempty" binding:"required"`
	/*maintainers:
			- name: firstmaintainer app1
	  		  email: firstmaintainer@hotmail.com
			- name: secondmaintainer app1
	  		  email: secondmaintainer@gmail.com
	*/
	Maintainers []Maintainers `yaml:"maintainers,omitempty" binding:"dive"`
	//company: Random Inc.
	Company string `yaml:"company,omitempty" binding:"required"`
	//website: https://website.com
	Website string `yaml:"website,omitempty" binding:"required"`
	//source: https://github.com/random/repo
	Source string `yaml:"source,omitempty" binding:"required"`
	//license: Apache-2.0
	License string `yaml:"license,omitempty" binding:"required"`
	/*
			description: |
		 		### Interesting Title
		 		Some application content, and description
	*/
	Description string `yaml:"description" binding:"required"`
}

type Maintainers struct {
	Name  string `yaml:"name,omitempty" binding:"required"`
	Email string `yaml:"email,omitempty" binding:"required,email"`
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
