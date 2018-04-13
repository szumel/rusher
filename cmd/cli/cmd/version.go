package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szumel/rusher/internal/platform/version"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version in format x.x.x",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Get())
	},
}
