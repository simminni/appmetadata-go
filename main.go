package main

import (
	"log"

	"github.com/simminni/appmetadata-go/controller"
)

// controller
func main() {
	log.SetFlags(1111)
	router := controller.SetupRouter()

	router.Run("localhost:8080")
}
