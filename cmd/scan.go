package cmd

import (
	"fmt"
	"os"

	"github.com/bupd/git-donkey/cmd/program"
	"github.com/bupd/git-donkey/cmd/ui/multiinputs"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "this is short dict",
	Long:  "this is A LONG DIct",

	Run: scan,
}

var _ = `
                          /\          /\
                         ( \\        // )
                          \ \\      // /
                           \_\\||||//_/
                            \/ _  _ \
                           \/|(O)(O)|
          ❤  bupd         \/ |      |
      ___________________\/  \      /
     //                //     |____|
    //                ||     /      \
   //|                \|     \ 0  0 /
  // \       )         V    / \____/
 //   \     /        (     /
""     \   /_________|  |_/
       /  /\   /     |  ||
      /  / /  /      \  ||
      | |  | |        | ||
      | |  | |        | ||
      |_|  |_|        |_||
       \_\  \_\        \_\\

  `

var logo2 = `
git-donkey

      \\__//
       /OO\\_______
       \__/\       )\/\
           ||----/ |
   ❤  bupd ||     ||

  `

var logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("81")).Bold(true)

func init() {
	rootCmd.AddCommand(scanCmd)
}

func scan(cmd *cobra.Command, args []string) {
	fmt.Printf("Scanning through repos...\n")
	fmt.Printf("This may take a while Please wait...\n\n")
	// Initialize GitInfo struct
	gitInfo := program.GitInfo{}
	// get the git dirs
	gitInfo.TotalGits, gitInfo.GitDirs = program.GitDirs()
	gitInfo.Untracked = program.UntrackedChanges(gitInfo.GitDirs)
	gitInfo.TotalUntracked = len(gitInfo.Untracked)
	gitInfo.Uncommitted = program.UncommittedChanges(gitInfo.GitDirs)
	gitInfo.TotalUncommitted = len(gitInfo.Uncommitted)
	gitInfo.Unpushed = program.UnpushedChanges(gitInfo.GitDirs)
	gitInfo.TotalUnpushed = len(gitInfo.Unpushed)

	fmt.Printf("%s\n", logoStyle.Render(fmt.Sprint(gitInfo.TotalGits)))
	fmt.Printf("%s\n", logoStyle.Render(logo2))

	p := tea.NewProgram(multiinputs.InitialModel(gitInfo))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
