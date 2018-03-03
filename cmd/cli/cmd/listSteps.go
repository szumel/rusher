package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"rusher/internal/platform/container"
	"rusher/internal/step"
)

func init() {
	rootCmd.AddCommand(listStepsCmd)
}

var listStepsCmd = &cobra.Command{
	Use:   "listSteps",
	Short: "Listing steps",
	Long:  "Listing all available steps",
	Run: func(cmd *cobra.Command, args []string) {
		stepPool, err := container.Get(step.AliasPool)
		if err != nil {
			log.Fatal(err)
		}

		for _, step := range stepPool.(*step.Pool).Steps {
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
