package cmd 

import (
	"fmt"
	"github.com/bupd/git-donkey/cmd/program"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "this is for the version",
    
	Run: version,

}


func init() { 
	rootCmd.AddCommand(versionCmd)
}

func version (cmd *cobra.Command, args [] string){ 
    fmt.Printf("The version of Git-Donkey...\n")

	if len(args) > 0 {
		program.GetVersion(args[0])
	} else {
		fmt.Println("No version argument provided")
	}
}