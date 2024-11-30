package cli

import (
	"core2/cli/commands"
	"github.com/spf13/cobra"
)

// rootCmd is the root of the CLI
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "CLI for core2 tasks",
	Long:  "CLI for running various core2 tasks such as migrations and seeding",
}

// RootCmd exposes the root command for extensions
func RootCmd() *cobra.Command {
	return rootCmd
}

// Init initializes core commands and adds additional commands
func Init(additionalCmds []*cobra.Command) {
	commands.InitCommands(rootCmd)
	rootCmd.AddCommand(additionalCmds...)
}
