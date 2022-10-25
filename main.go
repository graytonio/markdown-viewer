package main

import (
	"grayton.jfrog.io/markdown-viewer/systems"
)

func main() {
	systems.AddService(systems.FileSync)
	systems.AddService(systems.NoteTable)
	systems.AddService(systems.WebServer)
	systems.StartSystem()
}
