package lib

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

var internal_link_regex = regexp.MustCompile(`\[\[(.*?)\]\]`)

func get_note_link(name string) string {
	link_format := `<a href="/note/%s">%s</a>`
	if strings.Contains(name, "/") {
		parts := strings.Split(name, "/")
		return fmt.Sprintf(link_format, name, parts[len(parts)-1])
	}

	note_info := GetNoteInfo(name)
	if note_info == nil {
		return ""
	}
	return fmt.Sprintf(link_format, note_info.Path, name)
}

func custom_link_render(w io.Writer, node ast.Text, entering bool) (ast.WalkStatus, bool) {
	// new_content := internal_link_regex.ReplaceAllString(string(node.Literal), "<a>$1</a>")

	new_content := internal_link_regex.ReplaceAllStringFunc(string(node.Literal), func(b string) string {
		trimmed_name := strings.TrimSuffix(strings.TrimPrefix(b, "[["), "]]")
		return get_note_link(trimmed_name)
	})

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
