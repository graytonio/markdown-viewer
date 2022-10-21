package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"grayton.jfrog.com/markdown-viewer/lib"
)

func main() {
	router := gin.Default()

	router.GET("/*note_path", func(ctx *gin.Context) {
		path := ctx.Param("note_path")
		content, err := lib.FetchMarkdownAsHTML(path, lib.GetEnvD("MD_ROOT", "/"))
		if err != nil {
			// Handle not found
			if os.IsNotExist(err) {
				// TODO Handle directory listing
				ctx.AbortWithStatus(404)
				return
			}

			ctx.AbortWithError(500, err)
			return
		}
		ctx.Data(http.StatusOK, "text/html", content)
	})

	// Read in port from ENV
	router.Run(fmt.Sprintf(":%s", lib.GetEnvD("PORT", "9090")))
}
