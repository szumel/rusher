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
			title := `
## %s [%s]
### %s`
			fmt.Printf(title, step.Name(), step.Code(), step.Description())
			if len(step.Params()) > 0 {
				fmt.Println("\nParams:")
				for name, value := range step.Params() {
					fmt.Println("* " + name + " -> " + value)
				}
			}

			fmt.Println("\n-------------------------------")
		}
	},
}
