package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	Version        = "1.1.0"
	GitCommit      = ""
	BuildTime      = ""
	ReleaseChannel = "dev"
	GoVersion      = ""
	OS             = func() string {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "GOOS" {
					return setting.Value
				}
			}
		}

		return ""
	}
	Arch = func() string {
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, setting := range info.Settings {
				if setting.Key == "GOARCH" {
					return setting.Value
				}
			}
		}

		return ""
	}
	System = OS() + "/" + Arch()
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Version of git-donkey",
	Long:    `Get git-donkey version, git commit, go version, build time, release channel, os/arch, etc.`,
	Example: `  git-donkey version`,

	Run: version,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) {
	fmt.Printf("Version:      %s\n", Version)
	fmt.Printf("Go version:   %s\n", GoVersion)
	fmt.Printf("Git commit:   %s\n", GitCommit)
	fmt.Printf("Built:        %s\n", BuildTime)
	fmt.Printf("OS/Arch:      %s\n", System)
}
