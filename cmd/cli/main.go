package main

import (
	"log"

	store "github.com/Dhairya3124/ReaderCLI/internal/store"
	"github.com/Dhairya3124/ReaderCLI/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	store := new(store.Store)
	if _, err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)

	}
	m := tui.NewModel(store)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run program: %v", err)
	}

}
