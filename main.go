package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ramziChbl/gic-server/pkg/lakefs"
)

func main() {
	initServer()
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	os.Exit(1)

	// Serve frontend static files
	router.Static("/", "public/html")

	// Start and run the server
	router.Run(":3000")
}

func initServer() {
	_, err := lakefs.SetupLakeFS()
	if err != nil {
		fmt.Printf("%v", err)
	}
}
