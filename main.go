package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"grayton.jfrog.com/markdown-viewer/lib"
)

var notes_root string

func render_note_template(ctx *gin.Context, content template.HTML) {
	ctx.HTML(http.StatusOK, "note.html", gin.H{
		"content": content,
	})
}

func render_dir_template(ctx *gin.Context, files []lib.DirEntry) {
	ctx.HTML(http.StatusOK, "directory.html", gin.H{
		"notes": files,
	})
}

func init_web_server(router *gin.Engine) {
	router.Static("/static", "./static")
	router.FuncMap = template.FuncMap{
		"join": func(paths ...string) string {
			return path.Join(paths...)
		},
	}
	router.LoadHTMLGlob("./templates/*")

	router.GET("/note/*note_path", func(ctx *gin.Context) {
		path := ctx.Param("note_path")
		content, err := lib.FetchMarkdownAsHTML(path, notes_root)
		if err != nil {
			// Render The Directory List
			if errors.Is(err, lib.ErrIsDir) {
				files, err := lib.FetchDirList(path, notes_root)
				if err != nil {
					ctx.AbortWithError(500, err)
					return
				}

				// Render Directory Template
				render_dir_template(ctx, files)
				return
			}

			// Handle Not Found
			if os.IsNotExist(err) {
				ctx.AbortWithStatus(404)
				return
			}

			// Server Error
			ctx.AbortWithError(500, err)
			return
		}

		// Render Note Template
		render_note_template(ctx, content)
	})
}

func start_web_server(router *gin.Engine) {
	router.Run(fmt.Sprintf(":%s", lib.GetEnvD("PORT", "9090")))
}

func init() {
	notes_root = lib.GetEnvD("MD_ROOT", "/markdown")
}

func main() {
	router := gin.Default()
	init_web_server(router)
	start_web_server(router)
}
