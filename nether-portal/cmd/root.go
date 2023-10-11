package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:           "np",
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE:          rootRun,
}

func rootRun(cmd *cobra.Command, args []string) error {
	return errors.New("error")
}
