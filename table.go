package main

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	purple    = lipgloss.Color("99")
	gray      = lipgloss.Color("245")
	lightGray = lipgloss.Color("241")
)

func makeTable() [][]string {
	rows := [][]string{}
	todos := []string{
		"Clean closet",
		"Buy books",
		"Make my RPI dashboard",
	}
	calendar := []string{
		"voleyball game tonight",
	}
	weather := []string{
		"30 degrees, Rainy",
	}

	// Find the length of the largest slice
	maxLen := len(todos)
	if len(calendar) > maxLen {
		maxLen = len(calendar)
	}
	if len(weather) > maxLen {
		maxLen = len(weather)
	}

	for i := 0; i < maxLen; i++ {
		row := make([]string, 3)
		if i < len(todos) {
			row[0] = todos[i]
		} else {
			row[0] = ""
		}

		if i < len(calendar) {
			row[1] = calendar[i]
		} else {
			row[1] = ""
		}

		if i < len(weather) {
			row[2] = weather[i]
		} else {
			row[2] = ""
		}

		rows = append(rows, row)
	}

	return rows
}

func NewTable(width int) *table.Table {

	re := lipgloss.NewRenderer(os.Stdout)

	var (
		// HeaderStyle is the lipgloss style used for the table headers.
		HeaderStyle = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		// CellStyle is the base lipgloss style used for the table rows.
		CellStyle = re.NewStyle().Padding(1, 1).Width(30).Align(lipgloss.Center)
		// OddRowStyle is the lipgloss style used for odd-numbered table rows.
		OddRowStyle = CellStyle.Copy().Foreground(gray)
		// EvenRowStyle is the lipgloss style used for even-numbered table rows.
		EvenRowStyle = CellStyle.Copy().Foreground(lightGray)
		// BorderStyle is the lipgloss style used for the table border.
		BorderStyle = lipgloss.NewStyle().Foreground(purple)
	)

	rows := makeTable()

	colWidth := (width - 10) / 3

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle
			default:
				style = OddRowStyle
			}

			style = style.Copy().Width(colWidth)

			return style
		}).
		Headers("TO DO", "CALENDAR", "WEATHER").
		Rows(rows...)

	return t
}
