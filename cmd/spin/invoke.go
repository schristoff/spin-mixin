package main

import (
	"github.com/schristoff/spin-mixin/pkg/spin"
	"github.com/spf13/cobra"
)

func buildInvokeCommand(m *spin.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Upgrade(cmd.Context())
		},
	}

	// Define a flag for --action so that its presence doesn't cause errors, but ignore it since exec doesn't need it
	var action string
	cmd.Flags().StringVar(&action, "action", "", "Custom action name to invoke.")

	return cmd
}
