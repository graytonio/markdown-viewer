package lib

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Note struct {
	Duplicates bool
	Path       string
	Paths      []string
}

var name_path_map map[string]*Note
var notes_root string

func init() {
	notes_root = GetEnvD("MD_ROOT", "/markdown")
	GenerateNoteTree()
}

// Make sure a given path has passed file extension
func WithExtension(file string, extension string) string {
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

func GetNoteInfo(name string) *Note {
	return name_path_map[name]
}

type DirEntry struct {
	Name string
	Path string
}

// TODO Allow nested root directories for duplicate internal links
func GenerateNoteTree() error {
	log.Println("Generating New Note Table...")
	name_path_map = make(map[string]*Note)
	err := filepath.Walk(notes_root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore hidden files/directories
		if strings.HasPrefix(info.Name(), ".") || info.IsDir() {
			return nil
		}

		rel_path := strings.TrimPrefix(path, notes_root)
		file_name := strings.TrimSuffix(info.Name(), ".md")

		if note, ok := name_path_map[file_name]; ok {
			// If previously not duplicate
			if !note.Duplicates {
				note.Duplicates = true                     // Set duplicate true
				note.Paths = append(note.Paths, note.Path) // Add single path to array
				note.Path = ""                             // Remove path property
			}

			note.Paths = append(note.Paths, rel_path) // Add new path to array
		} else {
			name_path_map[file_name] = &Note{
				Duplicates: false,
				Path:       rel_path,
			}
		}
		return nil
	})
	log.Println("Completed Note Table")
	return err
}
