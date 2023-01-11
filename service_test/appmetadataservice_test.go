package service_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/simminni/appmetadata-go/controller"
	"github.com/stretchr/testify/assert"
)

var valid_app_metadata_1 = `title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
  ### Interesting Title
  Some application content, and description
`

var valid_app_metadata_2 = `title: Valid App 2
version: 1.0.1
maintainers:
- name: AppTwo Maintainer
  email: apptwo@hotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |
  ### Why app 2 is the best
  Because it simply is...
`

var invalid_payload_1 = `title: App w/ Invalid maintainer email
version: 1.0.1
maintainers:
- name: Firstname Lastname
  email: apptwohotmail.com
company: Upbound Inc.
website: https://upbound.io
source: https://github.com/upbound/repo
license: Apache-2.0
description: |
  ### blob of markdown
  More markdown
`

var invalid_payload_2 = `title: App w/ missing version
maintainers:
  - name: first last
    email: email@hotmail.com
  - name: first last
    email: email@gmail.com
company: Company Inc.
website: https://website.com
source: https://github.com/company/repo
license: Apache-2.0
description: |
  ### blob of markdown
  More markdown
`

func TestGetAppMetadataOnLoad(t *testing.T) {
	router := controller.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/appmetadata", nil)

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]\n", w.Body.String())
}

func TestPostAppMetadataValidRecord1(t *testing.T) {
	router := controller.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/appmetadata", strings.NewReader(valid_app_metadata_1))
	req.Header.Add("Content-Type", "application/x-yaml")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, valid_app_metadata_1, string(responseData))
}

func TestPostAppMetadataValidRecord2(t *testing.T) {
	router := controller.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/appmetadata", strings.NewReader(valid_app_metadata_2))
	req.Header.Add("Content-Type", "application/x-yaml")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, valid_app_metadata_2, string(responseData))
}

func TestPostAppMetadataInvalidEmail(t *testing.T) {
	var expectedError string = "message: 'Invalid payload.Key: ''App_metadata.Maintainers[0].Email'' Error:Field validation\n" +
		"  for ''Email'' failed on the ''email'' tag'\n"

	router := controller.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/appmetadata", strings.NewReader(invalid_payload_1))
	req.Header.Add("Content-Type", "application/x-yaml")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectedError, string(responseData))
}

func TestPostAppMetadataMissingVersion(t *testing.T) {
	var expectedError string = "message: 'Invalid payload.Key: ''App_metadata.Version'' Error:Field validation for\n" +
		"  ''Version'' failed on the ''required'' tag'\n"

	router := controller.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/appmetadata", strings.NewReader(invalid_payload_2))
	req.Header.Add("Content-Type", "application/x-yaml")

	router.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, expectedError, string(responseData))
}
