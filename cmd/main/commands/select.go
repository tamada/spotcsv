package commands

import (
	"errors"

	"github.com/spf13/cobra"
)

var selectCommand = &cobra.Command{
	Use:   "index",
	Short: "Select specified indexes items from the CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires column indexes")
		}
		return nil
	},
}

var exceptCommand = &cobra.Command{
	Use:   "except",
	Short: "Except specified indexes items from the CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires column indexes")
		}
		return nil
	},
}

func init() {
	rootCommand.AddCommand(selectCommand)
	rootCommand.AddCommand(exceptCommand)
}
