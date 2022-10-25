package systems

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/gin-gonic/gin"
	"grayton.jfrog.io/markdown-viewer/lib"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	router.Static("/static", "./static")
	router.FuncMap = template.FuncMap{
		"join": func(paths ...string) string {
			return path.Join(paths...)
		},
	}
	router.LoadHTMLGlob("./templates/*")

	router.GET("/note/*note_path", handle_note_request)
	log.Println("Webserver Initialized")
}

func WebServer(wg *sync.WaitGroup) {
	defer wg.Done()
	port := lib.GetConfig().Port
	log.Printf("Web server started and listening on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}

// Returns true if an error needed to be handled
func handle_error_response(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}

	if os.IsNotExist(err) {
		ctx.AbortWithStatus(404)
		return true
	}

	ctx.AbortWithError(500, err)
	return true
}

// Base handler for a request for a note
func handle_note_request(ctx *gin.Context) {
	log.Println("HELLO")
	path := ctx.Param("note_path")
	is_dir, _ := lib.IsDir(path, lib.GetConfig().MDRoot)
	if is_dir {
		handle_render_directory(ctx)
	} else {
		handle_render_note(ctx)
	}
}

// Fetches data and renders note template
func handle_render_note(ctx *gin.Context) {
	path := ctx.Param("note_path")
	content, err := lib.FetchMarkdownAsHTML(path, lib.GetConfig().MDRoot)
	if handle_error_response(ctx, err) {
		return
	}

	ctx.HTML(http.StatusOK, "note.html", gin.H{
		"content": content,
	})
}

// Fetches data and renders directory template
func handle_render_directory(ctx *gin.Context) {
	path := ctx.Param("note_path")
	files, err := lib.FetchDirList(path, lib.GetConfig().MDRoot)
	if handle_error_response(ctx, err) {
		return
	}

	ctx.HTML(http.StatusOK, "directory.html", gin.H{
		"notes": files,
	})
}
