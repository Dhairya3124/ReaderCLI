package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle    = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)
	faintStyle      = lipgloss.NewStyle().Background(lipgloss.Color("255")).Faint(true)
	enumeratorStyle = lipgloss.NewStyle().Background(lipgloss.Color("255")).MarginRight(1)
)

func (m Model) View() string {
	s := appNameStyle.Render("ReaderCLI") + "\n\n"
	if m.state == titleview {
		s += "Article title:\n\n"
		s += m.textinput.View() + "\n\n"
		s += faintStyle.Render("enter - save,esq - discard")
	}
	if m.state == listview {
		for i, a := range m.articles {
			prefix := " "
			if i == m.listIndex {
				prefix = ">"
			}
			shortBody := strings.ReplaceAll(a.Description, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += enumeratorStyle.Render(prefix) + a.Title + " | " +
				faintStyle.Render(shortBody) + "\n\n"
		}

	}
	s += faintStyle.Render("a - new article, q - quit")

	return s

}
