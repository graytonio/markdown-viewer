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
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./templates/*")

	router.GET("/note/*note_path", func(ctx *gin.Context) {
		path := ctx.Param("note_path")
		content, err := lib.FetchMarkdownAsHTML(path, lib.GetEnvD("MD_ROOT", "/markdown"))
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

		ctx.HTML(http.StatusOK, "note.html", gin.H{
			"content": content,
		})
	})

	// Read in port from ENV
	router.Run(fmt.Sprintf(":%s", lib.GetEnvD("PORT", "9090")))
}
