package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"rusher/internal/platform/container"
	"rusher/internal/rollback"
	"rusher/internal/step"
	"rusher/internal/step/schema"
)

var schemaFlag string
var envFlag string

func init() {
	rushCmd.Flags().StringVarP(&schemaFlag, "schema", "s", "schema.xml", "Location of schema")
	rushCmd.Flags().StringVarP(&envFlag, "env", "e", "test", "Specify which environment from schema will be executed")
	rootCmd.AddCommand(rushCmd)
}

var rushCmd = &cobra.Command{
	Use:   "rush",
	Short: "Executes rusher's steps",
	Long:  "Parses schema, validates and executes steps. On error rollbacks.",
	Run: func(cmd *cobra.Command, args []string) {
		if os.Getuid() == 0 {
			fmt.Println("Rusher should not be executed by root user!!!")
		}

		fmt.Println("Trying to rush " + envFlag + " environment from " + schemaFlag)
		s, err := schema.New(schemaFlag)
		if err != nil {
			log.Fatal("Could not fetch schema. Please ensure that given schema path is correct [" + err.Error() + "]")
		}
		currentConfig, err := schema.GetCurrentConfig(s, envFlag)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Rushing...")
		err = step.Rush(currentConfig, envFlag)

		switch err.(type) {
		case *step.ErrInvalidStep:
			log.Fatal(err)
			break
		case error:
			err := doRollback(err, currentConfig)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	},
}

func doRollback(err error, currentConfig *schema.Config) error {
	fmt.Println("error [" + err.Error() + "] starts rollback...")
	rollbacker, err := container.Get(rollback.AliasRollbacker)
	if err != nil {
		return err
	}
	fmt.Println("Starting rollback...")
	rollbackers := rollbacker.(*rollback.Pool).Rollbackers
	for _, rollbacker := range rollbackers {
		for _, stepConfig := range currentConfig.Steps {
			if rollbacker.Code() == stepConfig.Code {
				rollbacker.Rollback()
			}
		}
	}
	if err != nil {
		return err
	}
	fmt.Println("Rollbacker has been done successfully")

	return nil
}
