package lib

import (
	"io"
	"regexp"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

var internal_link_regex = regexp.MustCompile(`\[\[(.*?)\]\]`)

func render_reference_link(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	block, ok := node.(*ast.Text) // Try to parse node into text node
	if !ok {                      // If not a text node do not modify
		return ast.GoToNext, false
	}

	io.WriteString(w, string(block.Literal))

	links := internal_link_regex.FindAllStringIndex(string(block.Literal), -1)

	return ast.GoToNext, true
}

func get_renderer() *html.Renderer {
	opts := html.RendererOptions{
		RenderNodeHook: render_reference_link,
	}
	return html.NewRenderer(opts)
}
