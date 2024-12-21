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

package program

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Sessioniser(selected string) {
	// Check if selected directory is empty
	if selected == "" {
		os.Exit(0)
	}

	// Get the base name of the selected directory
	selectedName := filepath.Base(selected)
	selectedName = strings.Replace(selectedName, ".", "_", -1)

	// Check if tmux is running
	cmd := exec.Command("pgrep", "tmux")
	if err := cmd.Run(); err != nil {
		if err.Error() == "exit status 1" {
			fmt.Println("Tmux is not running.")
			// Start new tmux session if tmux is not running and not in tmux session
			if os.Getenv("TMUX") == "" {
				cmd = exec.Command("tmux", "new-session", "-s", selectedName, "-c", selected)
				if err := cmd.Run(); err != nil {
					fmt.Println("Error starting new tmux session:", err)
					os.Exit(1)
				}
			}
		} else {
			fmt.Println("Error checking if tmux is running:", err)
			os.Exit(1)
		}
	}

	// Check if session exists
	cmd = exec.Command("tmux", "has-session", "-t="+selectedName)
	if err := cmd.Run(); err != nil {
		if err.Error() == "exit status 1" {
			// Create new tmux session if session does not exist
			cmd = exec.Command("tmux", "new-session", "-ds", selectedName, "-c", selected)
			if err := cmd.Run(); err != nil {
				fmt.Println("Error creating new tmux session:", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Error checking if tmux session exists:", err)
			os.Exit(1)
		}
	}

	// Switch to the tmux session
	cmd = exec.Command("tmux", "switch-client", "-t", selectedName)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error switching to tmux session:", err)
		os.Exit(1)
	}
}
