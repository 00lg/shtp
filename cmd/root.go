package cmd

import (
	"github.com/spf13/cobra"
)

// shtp: self host tooy platform

func Execute() {
	cmd := &cobra.Command{
		Use:   "shtp",
		Short: "shtp is a simple static app deployer.",
	}

	cmd.AddCommand(NewRunCommand())

	cobra.CheckErr(cmd.Execute())
}
