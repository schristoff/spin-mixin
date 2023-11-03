package main

import (
	"github.com/schristoff/spin-mixin/pkg/spin"
	"github.com/spf13/cobra"
)

func buildUpgradeCommand(m *spin.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute(cmd.Context())
		},
	}
	return cmd
}
