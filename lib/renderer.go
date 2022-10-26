package lib

import (
	"fmt"
	"io"
	"path"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
)

func get_note_link(name string) string {
	link_format := `<a href="%s">%s</a>`

	parts := strings.Split(name, "/")
	note_name := parts[len(parts)-1]
	note_info := GetNoteInfo(note_name)
	if note_info == nil {
		fmt.Printf("ERROR: No note found for %s", note_name)
		return fmt.Sprintf(link_format, "#", name)
	}
	note_path := note_info.Path

	if note_info.Duplicates { // Find correct path if note is duplicate
		for _, path := range note_info.Paths {
			if strings.HasSuffix(WithExtension(name, "md"), path) {
				note_path = path
				break
			}
		}
	}

	return fmt.Sprintf(link_format, path.Join("/note", note_path), note_name)
}

var internal_link_regex = regexp.MustCompile(`\[\[(.*?)\]\]`)

func custom_link_render(w io.Writer, node ast.Text, entering bool) (ast.WalkStatus, bool) {
	// Find all internal links marked by [[name]] and replace them with the formated link
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
