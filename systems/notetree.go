package systems

import (
	"log"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"grayton.jfrog.io/markdown-viewer/lib"
)

func NoteTable(wg *sync.WaitGroup) {
	defer wg.Done()
	s := gocron.NewScheduler(time.UTC)
	// TODO Configurable update time
	_, err := s.Every(10).Minutes().Do(func() {
		err := lib.GenerateNoteTree()
		if err != nil {
			log.Printf("Error parsing note tree: %s", err.Error())
		}
	})

	if err != nil {
		log.Fatalf("Error scheduling note tree parser")
	}

	log.Println("Note Tree Parser Started Successfully")
	s.StartBlocking()
}
