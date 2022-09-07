package serv

import "context"

type ServerStartOptions struct {
	id     string
	listen string
}

func NewServerStartCmd(ctx context.Context, version string)