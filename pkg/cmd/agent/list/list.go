package agent

import (
	"github.com/defendops/orca/pkg/cmd/factory"
	"github.com/spf13/cobra"
)


func NewCmdRun(f *factory.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List agents running on providers",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if f.ORCAConfiguration.CloudProviders.DigitalOcean.Enabled {
				f.ProviderClients.DOClient.ListAgents()
			}

			return nil
		},
	}

	return cmd
}