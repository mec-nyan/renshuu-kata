package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mec-nyan/kana-master/pkg/kana"
)

var round = kana.List["a"]

type model struct {
	width, height int
}

var screenCentered = lipgloss.NewStyle().Margin(1)

var titleStyle = lipgloss.NewStyle().
	Padding(1, 2).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("4")).
	AlignHorizontal(lipgloss.Center)

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
	var output string
	var paddingLeft int

	// Show a welcoming message at the top.
	welcomeMsg := "Welcome to Renshuu Kata!"
	paddingLeft = max(0, (m.width-lipgloss.Width(welcomeMsg))/2)
	output += screenCentered.MarginLeft(paddingLeft).Italic(true).Foreground(lipgloss.Color("8")).Render(welcomeMsg)

	// This lesson's title.
	title := titleStyle.Render("ひらがな : Kata I")
	paddingLeft = max(0, (m.width-lipgloss.Width(title))/2)
	output += screenCentered.MarginLeft(paddingLeft).Render(title)

	// Print the Hiragana chart.
	output += "\n\n"
	for _, row := range kana.Rows {
		var kanas string
		for i, kana := range row {
			if kana.Hiragana != "" {
				kanas += kana.Hiragana
			} else {
				kanas += "  "
			}
			if i != 4 {
				kanas += "  "
			}
		}
		paddingLeft = max(0, (m.width-lipgloss.Width(kanas))/2)
		output += screenCentered.MarginLeft(paddingLeft).Foreground(lipgloss.Color("2")).Render(kanas)
	}
	output += "\n\n"

	quitMsg := "Hit q!"
	paddingLeft = max(0, (m.width-lipgloss.Width(quitMsg))/2)
	output += screenCentered.MarginLeft(paddingLeft).Italic(true).Foreground(lipgloss.Color("6")).Render(quitMsg)

	return output
}

func main() {
	if _, err := tea.NewProgram(model{}, tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("error starting the app:", err)
		os.Exit(1)
	}
}
