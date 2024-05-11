package multiinputs

import (
	"fmt"

	"github.com/bupd/git-donkey/cmd/program"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices     []string
	cursor      int              // which to-do list item our cursor is pointing at
	selected    map[int]struct{} // which to-do items are selected
	notTracked  int
	notCommited int
	notPushed   int
	gitDirs     []string
	totalGits   int
}

func InitialModel(gitInfo program.GitInfo) model {
	choices := append(gitInfo.Untracked, gitInfo.Uncommitted...)
	choices = append(choices, gitInfo.Unpushed...)

	return model{
		choices: choices,

		notTracked:  gitInfo.TotalUntracked,
		notCommited: gitInfo.TotalUncommitted,
		notPushed:   gitInfo.TotalUnpushed,
		// A map which indicates which choices are selected. We're using
		// the map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			fmt.Printf("value %v", m.choices[m.cursor])
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// Initialize the view string
	view := ""

	// Not Tracked Changes section
	view += "\nNot Tracked Changes\n\n"
	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if m.cursor == i {
			checked = "x"
		}

		if i == (m.notTracked + m.notCommited - 1) {
			view += fmt.Sprintf("\nNot Commited Changes\n\n")
		}

		if i == (m.notPushed + m.notCommited + m.notTracked - 1) {
			view += "\nNot Pushed Changes\n\n"
		}

		view += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// Footer
	view += "\nPress q to quit.\n"

	return view
}