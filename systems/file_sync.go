package systems

import (
	"log"
	"sync"

	"github.com/go-co-op/gocron"
	"grayton.jfrog.io/markdown-viewer/lib"
)

var vault_type lib.VaultType
var update_rate string

func init_git_sync() {
	err := lib.GitClone()
	if err != nil {
		log.Printf("Error cloning git repo: %s", err.Error())
	}
}

func init() {
	vault_type = lib.GetConfig().Vault
	update_rate = lib.GetConfig().GitUpdateRate

	if vault_type == lib.GitVault {
		init_git_sync()
	}

	log.Printf("File sync system for %s initialized", vault_type)
}

func FileSync(wg *sync.WaitGroup, cron *gocron.Scheduler) {
	defer wg.Done()
	// Local vault does not use any internal syncing so exit service
	if vault_type == lib.LocalVault {
		log.Println("Using local vault with no sync nothing to do")
		return
	}

	if vault_type == lib.GitVault {
		_, err := cron.Cron(update_rate).Do(func() {
			err := lib.GitClone()
			if err != nil {
				log.Printf("Error cloning git repo: %s", err.Error())
			}
		})

		if err != nil {
			log.Fatalf("Error starting file sync: %s", err.Error())
		}
	}

	log.Println("File Sync Started Successfully")
}
