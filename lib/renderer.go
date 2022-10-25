package lib

import (
	"io"
	"regexp"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

var internal_link_regex = regexp.MustCompile(`\[\[(.*?)\]\]`)

func custom_link_render(w io.Writer, node ast.Text, entering bool) (ast.WalkStatus, bool) {
	new_content := internal_link_regex.ReplaceAllString(string(node.Literal), "<strong>$1</strong>")
	io.WriteString(w, new_content)
	return ast.GoToNext, true
}

func render_hook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	// If text block use custom text rendering hook
	if block, ok := node.(*ast.Text); ok {
		return custom_link_render(w, *block, entering)
	}

	// Default to use default
	return ast.GoToNext, false
}

func get_renderer() *html.Renderer {
	opts := html.RendererOptions{
		RenderNodeHook: render_hook,
	}
	return html.NewRenderer(opts)
}
