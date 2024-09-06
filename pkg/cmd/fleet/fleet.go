package fleet

import (
	"fmt"

	"github.com/defendops/orca/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewFleetCmd(f *factory.CmdFactory) *cobra.Command {
	return &cobra.Command{
		Use:   "fleet",
		Short: "Manage fleets",
		Long:  "Commands to manage and deploy fleets of agents.",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := f.ORCAConfiguration
			fmt.Println("Fleet command executed with config:", cfg)
		},
	}
}
