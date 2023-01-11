# appmetadata-go
To build this app, I have referred following docs,
1. https://go.dev/doc/tutorial/web-service-gin.html
2. https://gin-gonic.com/docs/testing/

Code for this application is organized into model, controller, dao and service packages. Tests are in their corresponding test packages.
This app facilitates following API calls,
1. GET /appmetadata 
    This returns current appmetadata persisted in memory
2. GET /appmetadata/:title
    This return all the metadata records that match with title
3. POST /appmetadata
    This takes payload of type model.App_metadata and after validations, persist it in existing in memory App_metadata datastructre

