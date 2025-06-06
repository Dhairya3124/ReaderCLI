package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Foreground(lipgloss.Color("230")).Padding(0, 1).Bold(true)
	faintStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Faint(true)
	docStyle     = lipgloss.NewStyle().Margin(1, 2)
	titleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true).MarginBottom(1)
	errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).MarginTop(1)
	inputStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62")).Padding(1)
)

func (m Model) View() string {
	var s strings.Builder

	s.WriteString(appNameStyle.Render("ReaderCLI"))
	s.WriteString("\n\n")

	if m.err != nil {
		s.WriteString(errorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		s.WriteString("\n\n")
	}

	switch m.state {
	case titleview:
		s.WriteString(titleStyle.Render("Create New Article"))
		s.WriteString("\n")
		s.WriteString("Enter article title:\n\n")
		s.WriteString(inputStyle.Render(m.textinput.View()))
		s.WriteString("\n\n")
		s.WriteString(helpStyle.Render("enter - continue to body • esc - cancel"))

	case bodyview:
		title := "Edit Article"
		if m.currArticle.Title != "" {
			title = fmt.Sprintf("Edit: %s", m.currArticle.Title)
		}
		s.WriteString(titleStyle.Render(title))
		s.WriteString("\n")
		s.WriteString("Article content:\n\n")
		s.WriteString(m.textarea.View())
		s.WriteString("\n\n")
		s.WriteString(helpStyle.Render("ctrl+s - save article • esc - cancel"))

	case listview:
		s.WriteString(titleStyle.Render("Your Articles"))
		s.WriteString("\n")

		if len(m.articles) == 0 {
			s.WriteString(faintStyle.Render("No articles yet. Press 'a' to create your first article."))
			s.WriteString("\n\n")
		} else {
			s.WriteString(docStyle.Render(m.list.View()))
		}

		s.WriteString(helpStyle.Render("a - new article • enter - edit selected • q - quit"))
	}

	return s.String()
}

func (m Model) ViewWithPreview() string {
	var s strings.Builder

	s.WriteString(appNameStyle.Render("ReaderCLI"))
	s.WriteString("\n\n")

	if m.err != nil {
		s.WriteString(errorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
		s.WriteString("\n\n")
	}

	switch m.state {
	case titleview:
		s.WriteString(titleStyle.Render("Create New Article"))
		s.WriteString("\n")
		s.WriteString("Enter article title:\n\n")
		s.WriteString(inputStyle.Render(m.textinput.View()))
		s.WriteString("\n\n")
		s.WriteString(helpStyle.Render("enter - continue to body • esc - cancel"))

	case bodyview:
		title := "Edit Article"
		if m.currArticle.Title != "" {
			title = fmt.Sprintf("Edit: %s", m.currArticle.Title)
		}
		s.WriteString(titleStyle.Render(title))
		s.WriteString("\n")
		s.WriteString("Article content:\n\n")
		s.WriteString(m.textarea.View())
		s.WriteString("\n\n")
		s.WriteString(helpStyle.Render("ctrl+s - save article • esc - cancel"))

	case listview:
		s.WriteString(titleStyle.Render("Your Articles"))
		s.WriteString("\n")

		if len(m.articles) == 0 {
			s.WriteString(faintStyle.Render("No articles yet. Press 'a' to create your first article."))
			s.WriteString("\n\n")
		} else {
			for i, article := range m.articles {
				prefix := "  "
				itemStyle := lipgloss.NewStyle().MarginBottom(1)

				if i == m.list.Index() {
					prefix = "▶ "
					itemStyle = itemStyle.Background(lipgloss.Color("235")).Padding(0, 1)
				}

				shortDesc := strings.ReplaceAll(article.Description, "\n", " ")
				if len(shortDesc) > 50 {
					shortDesc = shortDesc[:50] + "..."
				}

				articleLine := fmt.Sprintf("%s%s",
					prefix,
					lipgloss.NewStyle().Bold(true).Render(article.Title))

				if shortDesc != "" {
					articleLine += "\n   " + faintStyle.Render(shortDesc)
				}

				s.WriteString(itemStyle.Render(articleLine))
				s.WriteString("\n")
			}
		}

		s.WriteString(helpStyle.Render("↑/↓ or j/k - navigate • a - new article • enter - edit • q - quit"))
	}

	return s.String()
}
