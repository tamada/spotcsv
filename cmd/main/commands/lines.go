package commands

import (
	"errors"

	"github.com/spf13/cobra"
)

var linesCommand = &cobra.Command{
	Use:   "lines",
	Short: "Show specified lines in the CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires line numbers")
		}
		return nil
	},
}

var headCommand = &cobra.Command{
	Use:   "head",
	Short: "Show specified lines from head in the CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires line numbers")
		}
		return nil
	},
}

var tailCommand = &cobra.Command{
	Use:   "tail",
	Short: "Show specified lines from tail in the CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("requires line numbers")
		}
		return nil
	},
}

func init() {
	rootCommand.AddCommand(linesCommand)
	rootCommand.AddCommand(headCommand)
	rootCommand.AddCommand(tailCommand)
}
