package lib

import (
	"errors"
	"io"
	"log"

	"gopkg.in/src-d/go-git.v4"
)

func GitPull() error {
	root := GetConfig().MDRoot

	r, err := git.PlainOpen(root)
	if err != nil {
		return err
	}

	w, err := r.Worktree()
	if err != nil {
		return err
	}

	err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		return err
	}

	return nil
}

func GitClone() error {
	url := GetConfig().GitURL
	root := GetConfig().MDRoot

	log.Printf("Syncing %s to %s", url, root)

	_, err := git.PlainClone(root, false, &git.CloneOptions{
		URL:      url,
		Progress: io.Discard, // Printing this out to the stdout makes logs look unreadable
	})

	if errors.Is(err, git.ErrRepositoryAlreadyExists) { // If cloning does not work because it already exists pull
		return GitPull()
	}

	return err
}
