package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

func read_in_mardown_file_content(path string) ([]byte, error) {
	// Check for file extension
	var fixed_path string = path
	if !strings.HasSuffix(path, ".md") {
		fixed_path = fmt.Sprintf("%s.md", path)
	}

	content, err := os.ReadFile(fixed_path)
	if err != nil {
		return nil, err
	}
	return markdown.ToHTML(content, nil, nil), nil
}

func main() {
	router := gin.Default()

	router.GET("*note_path", func(ctx *gin.Context) {
		path := ctx.Param("note_path")
		log.Println(path)
		content, err := read_in_mardown_file_content(path)
		if err != nil {
			// Handle not found
			if os.IsNotExist(err) {
				// Handle directory listing
				ctx.AbortWithStatus(404)
				return
			}

			ctx.AbortWithError(500, err)
			return
		}
		ctx.Data(http.StatusOK, "text/html", content)
	})

	// TODO Read in port from ENV
	router.Run(":9090")
}
