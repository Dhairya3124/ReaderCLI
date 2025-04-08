package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listview  uint = iota
	titleview      = 1
	bodyview       = 2
)

type model struct {
	state uint
	// store Store
	// textarea.Model
}

func NewModel() model {
	return model{state: listview}
}
func (m model) Init() tea.Cmd {
	return nil

}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}
