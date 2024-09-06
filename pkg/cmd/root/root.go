package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/defendops/orca/pkg/cmd/agent"
	"github.com/defendops/orca/pkg/cmd/background"
	"github.com/defendops/orca/pkg/cmd/factory"
	"github.com/defendops/orca/pkg/cmd/fleet"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "orca",
	Short: "ORCA is a CLI tool for External Attack Surface Management (EASM)",
	Long: heredoc.Doc(`
		ORCA is a CLI tool designed for scalable External Attack Surface Management (EASM).
		It enables you to manage and deploy agents, run background services, and manage fleets of agents.
	`),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func ExecuteRootCmd(f *factory.CmdFactory) error {
	f.LoadConfiguration()

	rootCmd.AddCommand(agent.NewAgentCmd(f))
	rootCmd.AddCommand(background.NewBackgroundCmd(f))
	rootCmd.AddCommand(fleet.NewFleetCmd(f))

	return rootCmd.Execute()
}
