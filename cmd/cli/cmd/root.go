package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "rusher",
	Short: "Rusher helps you with all kind of deployment",
	Long: "Rusher helps you with all kind of deployment",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
