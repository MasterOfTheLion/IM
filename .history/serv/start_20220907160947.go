package serv

import (
	"context"

	"github.com/spf13/cobra"
)

type ServerStartOptions struct {
	id     string
	listen string
}

func NewServerStartCmd(ctx context.Context, version string) *cobra.Command {
	opts := &ServerStartOptions{}

	cmd := &cobra.Command {
		Use: "chat",
		Short: "Starts a chat server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunServerStart(ctx, opts, version)
		},
	}
}