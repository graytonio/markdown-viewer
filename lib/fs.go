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

// Make sure a given path has passed file extension
func WithExtension(file string, extension string) string {
	if !strings.HasSuffix(file, extension) {
		return fmt.Sprintf("%s.%s", file, extension)
	}
	return file
}

// Check if given file path is a directory
func IsDir(file string, root string) (bool, error) {
	stat, err := os.Stat(path.Join(root, file))
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
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

func GetNoteInfo(name string) *Note {
	return name_path_map[strings.ToLower(name)]
}

type DirEntry struct {
	Name string
	Path string
}

// Generate a hash map from the directory structure to allow lookup of notes
func GenerateNoteTree() error {
	log.Println("Generating New Note Table...")
	name_path_map = make(map[string]*Note)
	notes_root := GetConfig().MDRoot
	err := filepath.Walk(notes_root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore hidden files/directories
		if strings.HasPrefix(info.Name(), ".") || info.IsDir() {
			return nil
		}

		// Pull out important information
		rel_path := strings.TrimPrefix(path, notes_root)                     // Trim notes_root from path
		file_name := strings.ToLower(strings.TrimSuffix(info.Name(), ".md")) // Remove extension from files

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
