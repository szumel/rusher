package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szumel/rusher/internal/step"
)

func init() {
	rootCmd.AddCommand(listStepsCmd)
}

var listStepsCmd = &cobra.Command{
	Use:   "listSteps",
	Short: "Listing steps",
	Long:  "Listing all available steps",
	Run: func(cmd *cobra.Command, args []string) {
		for _, step := range step.StepsPool.Steps {
			fmt.Println("\n#COMMAND")
			fmt.Println("	Name: " + step.Name())
			fmt.Println("	Code: " + step.Code())
			fmt.Println("	Description: " + step.Description())
			if len(step.Params()) > 0 {
				fmt.Println("	Params:")
				for name, value := range step.Params() {
					fmt.Println("		" + name + " -> " + value)
				}
			}
			fmt.Println("\n#END")
		}
	},
}
