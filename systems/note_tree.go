package systems

import (
	"log"
	"sync"

	"github.com/go-co-op/gocron"
	"grayton.jfrog.io/markdown-viewer/lib"
)

func init() {
	err := lib.GenerateNoteTree()
	if err != nil {
		log.Printf("Error parsing note tree: %s", err.Error())
	}
}

func NoteTable(wg *sync.WaitGroup, cron *gocron.Scheduler) {
	defer wg.Done()

	update_rate := lib.GetConfig().UpdateRate    // Get update rate from config
	_, err := cron.Cron(update_rate).Do(func() { // Set cron job to generate note tree at set rate
		err := lib.GenerateNoteTree()
		if err != nil {
			log.Printf("Error parsing note tree: %s", err.Error())
		}
	})

	if err != nil {
		log.Fatalf("Error starting note tree parser: %s", err.Error())
	}

	log.Println("Note Tree Parser Started Successfully")
}
