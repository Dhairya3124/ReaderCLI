package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	m := NewModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatalf("unable to run program: %v", err)
	}
}
