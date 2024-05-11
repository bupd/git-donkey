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

type helloMsg string

func waitASec() tea.Cmd {
	return func() tea.Msg {
		return "kumarlsadjflsajdlfjasldjflasdfads;;fkasdjl;fa"
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case helloMsg:
		// We caught our message like a Pokémon!
		// From here you could save the output to the model
		// to display it later in your view.
		m.choices[3] = string(msg)
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
			targetDirectory := "/home/bupd/code/" // Replace this with the directory you want to navigate to

			// Navigate to the specified directory
			if err := os.Chdir(targetDirectory); err != nil {
				fmt.Println("Error navigating to directory:", err)
				return nil, nil
			}

			// Confirm the current working directory after navigation
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				return nil, nil
			}

			fmt.Println("Successfully navigated to:", cwd)

			fmt.Printf("cd /home/bupd/code/")

			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
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
