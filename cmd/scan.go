package cmd

import (
	"fmt"

	"github.com/bupd/git-donkey/cmd/ui/listdirs"
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
	listdirs.RunModel()
	fmt.Println("kumaaruuu")
}
