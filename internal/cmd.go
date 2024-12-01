package internal

import "github.com/spf13/cobra"

func NewCommand(name string, short string, run func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:   name,
		Short: short,
		Run:   run,
	}
}
