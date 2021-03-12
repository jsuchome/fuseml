package client

import (
	"github.com/fuseml/fuseml/cli/paas"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var ()

// CmdInfo implements the carrier info command
var CmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Shows information about the Carrier environment",
	Long:  `Shows status and version for Kubernetes, Gitea, Tekton, Quarks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, cleanup, err := paas.NewCarrierClient(cmd.Flags(), nil)
		defer func() {
			if cleanup != nil {
				cleanup()
			}
		}()

		if err != nil {
			return errors.Wrap(err, "error initializing cli")
		}

		err = client.Info()
		if err != nil {
			return errors.Wrap(err, "error retrieving Carrier environment information")
		}

		return nil
	},
	SilenceErrors: true,
	SilenceUsage:  true,
}
