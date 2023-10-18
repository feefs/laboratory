package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "np",
	SilenceErrors: true, // Don't use cobra's error printer
	SilenceUsage:  true, // Don't print usage if an error occurs
	RunE:          rootRun,
}

func rootRun(cmd *cobra.Command, args []string) error {
	return errors.New("error")
}

func init() {
	rootCmd.AddCommand(blazeCmd)
}
