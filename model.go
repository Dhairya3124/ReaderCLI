package main

import (
	"context"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listview  uint = iota
	titleview      = 1
	bodyview       = 2
)

type Model struct {
	state       uint
	store       *Store
	articles    []Article
	currArticle Article
	listIndex   int
	textarea    textarea.Model
	textinput   textinput.Model
}

func NewModel(store *Store) Model {
	ctx := context.Background()
	articles, err := store.GetArticles(ctx)
	if err != nil {
		log.Fatalf("unable to get articles: %v", err)
	}
	return Model{state: listview, store: store, articles: articles, textarea: textarea.New(), textinput: textinput.New()}
}
func (m Model) Init() tea.Cmd {
	return nil

}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd tea.Cmd
	)
	m.textinput,cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)
	m.textarea,cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)


	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch m.state {
		case listview:
			switch key {
			case "q":
				return m, tea.Quit
			case "a": // for creating a new article
				m.state = titleview
				// TODO: add input
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex < len(m.articles) {
					m.listIndex++
				}
			case "enter":
				m.currArticle = m.articles[m.listIndex]
				m.state = bodyview
				// show textarea

			}
		}
	}
	return m, tea.Batch(cmds...)
}
