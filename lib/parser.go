package lib

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/gomarkdown/markdown"
)

var ErrIsDir = errors.New("cannot generate HTML from directory")

// Make sure a given path has passed file extension
func with_extension(file string, extension string) string {
	if !strings.HasSuffix(file, extension) {
		return fmt.Sprintf("%s.%s", file, extension)
	}
	return file
}

func IsDir(file string, root string) (bool, error) {
	stat, err := os.Stat(path.Join(root, file))
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

// Fetch a markdown from disk and render as HTML
func FetchMarkdownAsHTML(file string, root string) (template.HTML, error) {
	isDir, err := IsDir(file, root)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	} else if isDir {
		return "", ErrIsDir
	}

	abs_path := path.Join(root, with_extension(file, "md"))
	content, err := os.ReadFile(abs_path)
	if err != nil {
		return "", err
	}

	return template.HTML(markdown.ToHTML(content, nil, get_renderer())), nil
}

type DirEntry struct {
	Name string
	Path string
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
