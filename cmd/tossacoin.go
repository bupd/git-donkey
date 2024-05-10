package cmd

import (
	"fmt"
	"os"

	"github.com/bupd/git-donkey/cmd/ui/tossacoin"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var tossCmd = &cobra.Command{
	Use:   "tossacoin",
	Short: "this is short dict for the toss",
	Long:  "this is A LONG DIct for the toss you idiot",

	Run: toss,
}

type Options struct {
	ProjectName string
	ProjectType string
}

func init() {
	rootCmd.AddCommand(tossCmd)
}

func toss(cmd *cobra.Command, args []string) {
	fmt.Println("tossing a coin")
	if _, err := tea.NewProgram(tossacoin.InitialModel()).Run(); err != nil {
		fmt.Printf("could not start program: %s\n", err)
		os.Exit(1)
	}
}
