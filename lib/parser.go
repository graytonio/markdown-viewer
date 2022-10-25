package lib

import (
	"html/template"
	"os"
	"path"

	"github.com/gomarkdown/markdown"
)

// Fetch a markdown from disk and render as HTML
func FetchMarkdownAsHTML(file string, root string) (template.HTML, error) {
	abs_path := path.Join(root, WithExtension(file, "md"))
	content, err := os.ReadFile(abs_path)
	if err != nil {
		return "", err
	}

	return template.HTML(markdown.ToHTML(content, nil, get_renderer())), nil
}
