package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	width, height int
}

var fullScreenStyle = lipgloss.NewStyle().Margin(1)

var titleStyle = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("4"))

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	title := titleStyle.Render("Welcome to Renshuu Kata!")
	paddingLeft := max(0, (m.width-lipgloss.Width(title))/2)
	fullScreenStyle := fullScreenStyle.MarginLeft(paddingLeft)
	return fullScreenStyle.Render(title)
}

func main() {
	if _, err := tea.NewProgram(model{}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("error starting the app:", err)
		os.Exit(1)
	}
}
