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
	"grayton.jfrog.com/markdown-viewer/lib"
)

var router *gin.Engine
var notes_root string

func init() {
	router = gin.Default()
	notes_root = lib.GetEnvD("MD_ROOT", "/markdown")
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

func StartWebServer(wg *sync.WaitGroup) {
	defer wg.Done()
	port := lib.GetEnvD("PORT", "9090")
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
	path := ctx.Param("note_path")
	is_dir, err := lib.IsDir(path, notes_root)
	if handle_error_response(ctx, err) {
		return
	}

	if is_dir {
		handle_render_directory(ctx)
	} else {
		handle_render_note(ctx)
	}
}

// Fetches data and renders note template
func handle_render_note(ctx *gin.Context) {
	path := ctx.Param("note_path")
	content, err := lib.FetchMarkdownAsHTML(path, notes_root)
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
	files, err := lib.FetchDirList(path, notes_root)
	if handle_error_response(ctx, err) {
		return
	}

	ctx.HTML(http.StatusOK, "directory.html", gin.H{
		"notes": files,
	})
}
