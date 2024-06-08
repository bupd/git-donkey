#!/bin/bash

## THIS SCRIPT SERVES AS A PROOF OF CONCEPT OF THE GIT DONKEY PROJECT. ##
# Function to check if the directory is a git repository
check_git_repo() {
    if [ -d .git ]; then
        echo "This directory is a git repository."
    else
        echo "This directory is not a git repository."
        exit 1
    fi
}

# Function to check if there are any unstaged changes
check_unstaged_changes() {
    if ! git diff --quiet; then
        echo "There are unstaged changes in this repository."
    fi
}

# Function to check if there are any staged but uncommitted changes
check_staged_changes() {
    if ! git diff --cached --quiet; then
        echo "There are staged but uncommitted changes in this repository."
    fi
}

# Function to check if there are any committed but not pushed changes
check_unpushed_commits() {
    if [ $(git rev-list --count @{u}..HEAD) -gt 0 ]; then
        echo "There are committed but not pushed changes in this repository."
    fi
}

# Function to check for untracked changes
check_untracked_changes() {
    if [ $(git ls-files --others --exclude-standard | wc -l) -gt 0 ]; then
        echo "There are untracked changes in this repository."
    fi
}

# Function to check for any other potential problems
check_other_problems() {
    # Add additional checks here if needed
    echo "No other problems found."
}

# Main function to call all checks
main() {
    check_git_repo
    check_unstaged_changes
    check_staged_changes
    check_untracked_changes
    check_unpushed_commits
    check_other_problems
}

# Call the main function
main
