package cli

import (
	"core2/cli/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A core2 application",
}

// Expose RootCmd for extension
func RootCmd() *cobra.Command {
	return rootCmd
}

// Init adds additional commands during initialization
func Init(additionalCmds []*cobra.Command) {
	rootCmd.AddCommand(commands.MigrateCmd)
	rootCmd.AddCommand(commands.SeederCmd)

	rootCmd.AddCommand(additionalCmds...)
}

// Execute starts the CLI application
func Execute() {
	rootCmd.Execute()
}
