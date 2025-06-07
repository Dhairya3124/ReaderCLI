package tui

import (
	"context"
	"log"

	store "github.com/Dhairya3124/ReaderCLI/internal/store"
	"github.com/charmbracelet/bubbles/list"
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
	textarea    textarea.Model
	textinput   textinput.Model
	list        list.Model
	err         error
}

type ArticleItem struct {
	store.Article
}

func (a ArticleItem) Title() string       { return a.Article.Title }
func (a ArticleItem) Description() string { return a.Article.Description }
func (a ArticleItem) FilterValue() string { return a.Article.Title }

func NewModel(store *store.Store) Model {
	ctx := context.Background()
	articles, err := store.GetArticles(ctx)
	if err != nil {
		log.Fatalf("unable to get articles: %v", err)
	}

	items := make([]list.Item, len(articles))
	for i, a := range articles {
		items[i] = ArticleItem{a}
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Articles"

	ta := textarea.New()

	ti := textinput.New()

	return Model{
		state:     listview,
		store:     store,
		articles:  articles,
		textarea:  ta,
		textinput: ti,
		list:      l,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) refreshArticles(ctx context.Context) error {
	articles, err := m.store.GetArticles(ctx)
	if err != nil {
		return err
	}

	m.articles = articles
	items := make([]list.Item, len(articles))
	for i, a := range articles {
		items[i] = ArticleItem{a}
	}
	m.list.SetItems(items)
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch m.state {
	case listview:
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	case titleview:
		m.textinput, cmd = m.textinput.Update(msg)
		cmds = append(cmds, cmd)
	case bodyview:
		m.textarea, cmd = m.textarea.Update(msg)
		cmds = append(cmds, cmd)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch m.state {
		case listview:
			switch key {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "a":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.currArticle = store.Article{}
				m.state = titleview
			case "enter":
				if len(m.articles) > 0 {
					selectedIndex := m.list.Index()
					if selectedIndex < len(m.articles) {
						m.currArticle = m.articles[selectedIndex]
						m.textarea.SetValue(m.currArticle.Description)
						m.textarea.Focus()
						m.textarea.CursorEnd()
						m.state = bodyview
					}
				}
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
				m.textinput.Blur()
				m.state = listview
			}

		case bodyview:
			switch key {
			case "ctrl+s":
				body := m.textarea.Value()
				m.currArticle.Description = body
				ctx := context.Background()

				if err := m.store.Create(ctx, &m.currArticle); err != nil {
					m.err = err
					return m, nil
				}

				if err := m.refreshArticles(ctx); err != nil {
					m.err = err
					return m, nil
				}

				m.currArticle = store.Article{}
				m.textarea.Blur()
				m.state = listview

			case "esc":
				m.textarea.Blur()
				m.currArticle = store.Article{}
				m.state = listview
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		m.textarea.SetWidth(msg.Width - h)
		m.textarea.SetHeight(msg.Height - v - 5)
	}

	return m, tea.Batch(cmds...)
}
