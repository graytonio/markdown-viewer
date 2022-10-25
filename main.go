package main

import (
	"log"
	"sync"

	"grayton.jfrog.com/markdown-viewer/lib"
	"grayton.jfrog.com/markdown-viewer/systems"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go systems.StartWebServer(&wg) // Start webserver system
	log.Printf("Webserver Listening on :%s", lib.GetEnvD("PORT", "9000"))
	wg.Wait()
}
