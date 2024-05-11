package multiinputs

import (
	"fmt"
	"os/exec"

	"github.com/bupd/git-donkey/cmd/program"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	done        bool
}

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(1)
	selectedItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	selectedStyle     = lipgloss.NewStyle().PaddingLeft(0).Foreground(lipgloss.Color("81"))
	chosenStyle       = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("244"))
	chosenCheckStyle  = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("81"))
	headingStyle      = lipgloss.NewStyle().
				Background(lipgloss.Color("#8ae234")).
				Foreground(lipgloss.Color("#000000")).
				Bold(true).
				Padding(0, 1, 0)

	notPushStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#ef2929")).
			Foreground(lipgloss.Color("#000000")).
			Bold(true).
			Padding(0, 1, 0)
)

type editorFinishedMsg struct{ err error }

// change the dir and execute lazygit on the specified directory
func changeDir(dir string) tea.Cmd {
	c := exec.Command("lazygit")
	c.Dir = dir
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return editorFinishedMsg{err}
	})
}

// InitialModel to render the TUI
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

		case "c":
			dir := m.choices[m.cursor]
			return m, changeDir(dir)
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
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
				dir := m.choices[m.cursor]
				return m, changeDir(dir)
			}
			return m, nil
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
	notTrackedStyled := headingStyle.Render("Not Tracked Changes:")
	view += fmt.Sprintf("%s \n\n", notTrackedStyled)
	for i, choice := range m.choices {

		if m.notTracked < 1 {
			nothing := chosenStyle.Render("nothing to show...")
			view += fmt.Sprintf("%s\n", nothing)
		}

		choiceStyled := itemStyle.Render(choice)
		checked := itemStyle.Render("[ ]")
		cursor := " "
		if m.cursor == i {
			cursor = selectedStyle.Render(">")
			choiceStyled = selectedItemStyle.Render(choice)
		}

		if _, ok := m.selected[i]; ok {
			checked = chosenCheckStyle.Render("[x]")
			choiceStyled = chosenStyle.Render(choice)
		}

		if m.notCommited < 1 {
		} else {
			if i == (m.notTracked) {
				notCommitedStyled := headingStyle.Render("Not Commited Changes:")
				view += fmt.Sprintf("\n%s\n\n", notCommitedStyled)
				if m.notCommited < 1 {
					nothing := chosenStyle.Render("nothing to show...")
					view += fmt.Sprintf("%s \n %v", nothing, m.notTracked)
				}
			}
		}

		if m.notPushed < 1 {
		} else {
			if i == (m.notCommited + m.notTracked) {
				notPushedStyled := notPushStyle.Render("Not Pushed Changes:")
				view += fmt.Sprintf("\n%s\n\n", notPushedStyled)
				if m.notPushed < 1 {
					nothing := chosenStyle.Render("nothing to show...")
					view += fmt.Sprintf("%s \n", nothing)
				}
			}
		}

		view += fmt.Sprintf("%s %s %s\n\n", cursor, checked, choiceStyled)
	}

	// Footer
	view += "\nPress q to quit.\n"

	return view
}
