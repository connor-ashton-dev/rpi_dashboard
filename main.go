package main

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	t := NewTable(80)
	p := tea.NewProgram(model{table: t}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
