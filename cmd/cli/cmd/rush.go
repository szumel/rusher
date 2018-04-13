package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szumel/rusher/internal/platform/container"
	"github.com/szumel/rusher/internal/platform/log"
	"github.com/szumel/rusher/internal/platform/rollback"
	"github.com/szumel/rusher/internal/platform/schema"
	"github.com/szumel/rusher/internal/step"
	golog "log"
	"os"
)

var schemaFlag string
var envFlag string

func init() {
	rushCmd.Flags().StringVarP(&schemaFlag, "schema", "s", "schema.xml", "Location of schema")
	rushCmd.Flags().StringVarP(&envFlag, "env", "e", "test", "Specify which environment from schema will be executed")
	rootCmd.AddCommand(rushCmd)
}

var rushCmd *cobra.Command = &cobra.Command{
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
			log.Logger.Println(err.Error())
			golog.Fatal(err)
		}
		currentConfig, err := schema.GetCurrentConfig(s, envFlag)
		if err != nil {
			log.Logger.Println(err.Error())
			golog.Fatal(err)
		}

		fmt.Println("Rushing...")
		err = step.Rush(currentConfig, envFlag)

		switch err.(type) {
		case *step.ErrInvalidStep:
			log.Logger.Println(err.Error())
			golog.Fatal(err)
			break
		case error:
			if err := doRollback(err, currentConfig); err != nil {
				log.Logger.Println(err.Error())
				golog.Fatal(err)
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
