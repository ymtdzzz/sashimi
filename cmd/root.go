package cmd

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-plugin"
	"github.com/spf13/cobra"
	"github.com/ymtdzzz/sashimi/shared"
)

var rootCmd = &cobra.Command{
	Use:     "sashimi [job command]",
	Short:   "sashimi is a tool to run distributed jobs",
	Args:    cobra.MinimumNArgs(1),
	Example: "sashimi batch_name [args]",
	RunE: func(cmd *cobra.Command, args []string) error {
		exCmd := exec.Command(args[0])
		if len(args) > 1 {
			exCmd = exec.Command(args[0], args[1:]...)
		}
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: shared.Handshake,
			Plugins:         shared.PluginMap,
			Cmd:             exCmd,
			AllowedProtocols: []plugin.Protocol{
				plugin.ProtocolGRPC,
			},
		})
		defer client.Kill()

		rpcClient, err := client.Client()
		if err != nil {
			return err
		}

		raw, err := rpcClient.Dispense("job")
		if err != nil {
			return err
		}

		job := raw.(shared.Job)
		cmds, err := job.SplitJob(context.Background())
		if err != nil {
			return err
		}
		fmt.Printf(strings.Join(cmds, ","))

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
}
