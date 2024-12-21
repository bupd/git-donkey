/*
Copyright © 2024 Prasanth Bupd <bupdprasanth@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

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
