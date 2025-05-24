package tui

import (
	"context"
	"fmt"
	"log"

	store "github.com/Dhairya3124/ReaderCLI/internal/store"
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
	store       *store.Store
	articles    []store.Article
	currArticle store.Article
	listIndex   int
	textarea    textarea.Model
	textinput   textinput.Model
}

func NewModel(store *store.Store) Model {
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
		cmd  tea.Cmd
	)
	m.textinput, cmd = m.textinput.Update(msg)
	cmds = append(cmds, cmd)
	m.textarea, cmd = m.textarea.Update(msg)
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
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.currArticle = store.Article{}
				m.state = titleview
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
				//TODO: add more fields to SetValue
				m.textarea.SetValue(m.currArticle.Description)
				m.textarea.Focus()
				m.textarea.CursorEnd()
				m.state = bodyview

			}
		case titleview:
			switch key {
			case "enter":
				title := m.textinput.Value()
				if title != "" {
					m.currArticle.Title = title
					m.textarea.SetValue("")
					m.textarea.Focus()
					m.textarea.CursorEnd()
					m.state = bodyview
				}

			case "esc":
				m.state = listview

			}
		case bodyview:
			switch key {
			case "ctrl+s":
				body := m.textarea.Value()
				m.currArticle.Description = body
				ctx := context.Background()
				if err := m.store.Create(ctx, &m.currArticle); err != nil {
					//Todo: handle error
					fmt.Println(err)
					return m, tea.Quit
				}
				var err error
				m.articles, err = m.store.GetArticles(ctx)
				if err != nil {
					//Todo: handle error
					fmt.Println(err)
					return m, tea.Quit
				}
				m.currArticle = store.Article{}
				m.state = listview

			case "esc":
				m.state = listview

			}

		}
	}
	return m, tea.Batch(cmds...)
}
