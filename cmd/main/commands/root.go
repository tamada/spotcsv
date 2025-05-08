package commands

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "spotcsv",
	Short: "Manipulating the CSV files",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no command specified")
	},
}

func Execute(args []string) error {
	rootCommand.SetOut(os.Stdout)
	if err := rootCommand.ParseFlags(args); err != nil {
		return err
	}
	return rootCommand.Execute()
}
