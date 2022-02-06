package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ramziChbl/gic-server/pkg/lakefs"
)

type Response struct {
	Status  int
	Message []string
	Error   []string
}

func SendResponse(c *gin.Context, response Response) {
	if len(response.Message) > 0 {
		c.JSON(response.Status, map[string]interface{}{"message": strings.Join(response.Message, "; ")})
	} else if len(response.Error) > 0 {
		c.JSON(response.Status, map[string]interface{}{"error": strings.Join(response.Error, "; ")})
	}
}

func main() {
	initServer()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Static("/", "public/html")

	router.POST("/upload", uploadImage)

	// Start and run the server
	router.Run(":3000")
}

func initServer() {
	_, err := lakefs.SetupLakeFS()
	if err != nil {
		fmt.Printf("%v", err)
	}

	_, err = lakefs.CreateRepo("images", "local://images", "main")
	if err != nil {
		fmt.Printf("%v", err)
	}

}

func uploadImage(c *gin.Context) {

	//file, _, err: = c.Request.FormFile("file")
	imageQuality, err := strconv.Atoi(c.PostForm("quality"))
	if err != nil {
		SendResponse(c, Response{
			Status:  http.StatusBadRequest,
			Message: []string{},
			Error:   []string{"quality parameter is wrong"},
		})

		return
	}
	if err != nil {
		c.String(http.StatusNotAcceptable, "Couldn't parse required field : %s", string(imageQuality))
		return
	}

	uploadedFile, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusNotAcceptable, "Couldn't read uploaded file")
		return
	}

	//  Ensure our file does not exceed 10MB
	if uploadedFile.Size > 10*1024*1024 {
		fmt.Printf("File size (%s) exceeds limit of 10MB", string(uploadedFile.Size))
		return
	}

	// Upload the file to specific dst.
	c.SaveUploadedFile(uploadedFile, filepath.Join("web/files/uncompressed", uploadedFile.Filename))

	SendResponse(c, Response{
		Status:  http.StatusAccepted,
		Message: []string{fmt.Sprintf("%s uploaded", uploadedFile.Filename)},
		Error:   []string{},
	})
}
