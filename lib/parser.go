package lib

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"
)

// Make sure a given path has passed file extension
func with_extension(file string, extension string) string {
	if !strings.HasSuffix(file, extension) {
		return fmt.Sprintf("%s.%s", file, extension)
	}
	return file
}

// Fetch a markdown from disk and render as HTML
func FetchMarkdownAsHTML(file string, root string) ([]byte, error) {
	content, err := os.ReadFile(path.Join(root, with_extension(file, "md")))
	if err != nil {
		return nil, err
	}

	return markdown.ToHTML(content, nil, get_renderer()), nil
}
