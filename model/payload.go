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
