package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/ymtdzzz/sashimi/examples/go/base"
	"github.com/ymtdzzz/sashimi/shared"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"job": &shared.JobGRPCPlugin{Impl: &base.Job{
				Name: "job_a",
				Unit: 5,
			}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
