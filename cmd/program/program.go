package program

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
)

// recursively find the Directories that contains .git directory
func findGitDirectories(rootDir string) ([]string, error) {
	var gitDirs []string

	// Walk through the directory and its subdirectories
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the current directory contains a .git directory
		if info.IsDir() && filepath.Base(path) == ".git" {
			// Add the parent directory to the gitDirs slice
			gitDirs = append(gitDirs, filepath.Dir(path))
			// Do not continue traversing this directory
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return gitDirs, nil
}

func GitDirs() (int, []string) {
	timeNow := time.Now()
	// Example usage
	rootDir := "."

	gitDirs, err := findGitDirectories(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
	}
	notTracked := len(gitDirs)

	timeSince := time.Since(timeNow)
	fmt.Println(timeSince)

	return notTracked, gitDirs
}

// Function to check for untracked changes in a Git repository
func hasUntrackedChanges(repoPath string) (bool, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return false, err
	}

	// Get the worktree of the repository
	worktree, err := repo.Worktree()
	if err != nil {
		return false, err
	}

	// Check for untracked files
	status, err := worktree.Status()
	if err != nil {
		return false, err
	}

	return !status.IsClean(), nil
}

func UntrackedChanges(dirs []string) []string {
	var untracked []string
	for _, path := range dirs {

		hasChanges, err := hasUntrackedChanges(path)
		if err != nil {
			log.Printf("\n\n\n kuma error %v", err)
		} else if hasChanges {
			untracked = append(untracked, path)
		}
	}
	return untracked
}

func hasUncommittedChanges(repoPath string) (bool, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return false, err
	}

	// Get the worktree of the repository
	worktree, err := repo.Worktree()
	if err != nil {
		return false, err
	}

	// Check for uncommitted changes
	status, err := worktree.Status()
	if err != nil {
		return false, err
	}

	// if the status.string comes true then no file is modified.
	// Check if there are changes that are not staged
	for _, entry := range status {
		if entry.Staging == git.Added || entry.Staging == git.Modified {
			return true, nil
		}
	}

	return false, nil
}

func UncommittedChanges(dirs []string) []string {
	var uncommitted []string
	for _, path := range dirs {

		hasChanges, err := hasUncommittedChanges(path)
		if err != nil {
			log.Printf("\n kumaru error %v", err)
		} else if hasChanges {
			uncommitted = append(uncommitted, path)
		}
	}
	return uncommitted
}

func hasUnpushedChanges(repoPath string) (bool, error) {
	// Define the paths to the refs directories
	localRefPath := filepath.Join(repoPath, ".git/refs/heads/main")
	originRefPath := filepath.Join(repoPath, ".git/refs/remotes/origin/main")

	// Read the contents of localRefPath
	localRefContent, err := readFileContents(localRefPath)
	if err != nil {
		return false, err
	}

	// Read the contents of originRefPath
	originRefContent, err := readFileContents(originRefPath)
	if err != nil {
		return false, err
	}

	// Compare the contents of the two ref files
	isEqual := compareRefContents(localRefContent, originRefContent)

	// Print the result
	return isEqual, nil
}

// Function to read file contents
func readFileContents(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// Function to compare ref contents
func compareRefContents(ref1, ref2 string) bool {
	return ref1 == ref2
}

func UnpushedChanges(dirs []string) []string {
	var unpushed []string
	for _, path := range dirs {

		hasChanges, err := hasUnpushedChanges(path)
		if err != nil {
			_ = err
		} else if !hasChanges {
			unpushed = append(unpushed, path)
		}
	}
	return unpushed
}
