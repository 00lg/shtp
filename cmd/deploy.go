package cmd

import (
	"errors"

	"github.com/00lg/shtp/internal"
	"github.com/spf13/cobra"
)

func NewRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Builds and runs your app from a Dockerfile.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic(errors.New("requires at least one arg"))
			}
			internal.Deploy(cmd.Context(), args[0])
		},
	}
	return cmd
}
