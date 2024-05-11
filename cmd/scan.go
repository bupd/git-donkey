package cmd

import (
	"fmt"
	"os"

	"github.com/bupd/git-donkey/cmd/program"
	"github.com/bupd/git-donkey/cmd/ui/multiinputs"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "this is short dict",
	Long:  "this is A LONG DIct",

	Run: scan,
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

func scan(cmd *cobra.Command, args []string) {
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

	p := tea.NewProgram(multiinputs.InitialModel(gitInfo), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
