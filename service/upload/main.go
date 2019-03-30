package main

import (
	"go-filestore-server/config"
	"go-filestore-server/route"
)

func main() {
	// gin framework
	router := route.Router()
	router.Run(config.UploadServiceHost)
}
