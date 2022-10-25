package lib

import (
	"errors"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"
)

var ErrIsDir = errors.New("cannot generate HTML from directory")

// Fetch a markdown from disk and render as HTML
func FetchMarkdownAsHTML(file string, root string) (template.HTML, error) {
	abs_path := path.Join(root, WithExtension(file, "md"))
	content, err := os.ReadFile(abs_path)
	if err != nil {
		return "", err
	}

	return template.HTML(markdown.ToHTML(content, nil, get_renderer())), nil
}

func FetchDirList(dir string, root string) (good_files []DirEntry, err error) {
	files, err := os.ReadDir(path.Join(root, dir))
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			good_files = append(good_files, DirEntry{
				Name: file.Name(),
				Path: path.Join(dir, file.Name()),
			})
		}
	}

	return good_files, nil
}
