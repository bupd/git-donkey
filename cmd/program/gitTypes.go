package program

type GitInfo struct {
	TotalGits        int
	GitDirs          []string
	Untracked        []string
	TotalUntracked   int
	Uncommitted      []string
	TotalUncommitted int
	Unpushed         []string
	TotalUnpushed    int
}
