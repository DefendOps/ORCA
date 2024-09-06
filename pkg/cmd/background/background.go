package background

import (
	"fmt"

	"github.com/defendops/orca/pkg/cmd/factory"
	"github.com/spf13/cobra"
)

func NewBackgroundCmd(f *factory.CmdFactory) *cobra.Command {
	return &cobra.Command{
		Use:   "background",
		Short: "Manage background services",
		Long:  "Commands to manage background services for ORCA.",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := f.ORCAConfiguration
			fmt.Println("Background command executed with config:", cfg)
		},
	}
}
