package agent

import (
	"github.com/defendops/orca/pkg/cmd/factory"
	listAgents "github.com/defendops/orca/pkg/cmd/agent/list"
	"github.com/spf13/cobra"
)

func NewAgentCmd(f *factory.CmdFactory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agents",
		Short: "Manage agents",
		Long:  "Commands to manage and interact with agents.",
	}

	cmd.AddCommand(listAgents.NewCmdRun(f))

	return cmd 
}