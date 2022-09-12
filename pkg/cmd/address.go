package cmd

import (
	"github.com/spf13/cobra"
	"cgts-client/pkg/utils"
)

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "IP address information",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()	
	},
}

// subcommands
var addressCmdShow = &cobra.Command {
	Use:	"show",
	Short:	"Show IP address attributes",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()	
	},
}

var addressCmdList= &cobra.Command {
	Use:	"list",
	Short:	"List IP addresses on host",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()	
	},
}

var addressCmdDelete = &cobra.Command {
	Use:	"delete",
	Short:	"Delete an IP address",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()	
	},
}

var addressCmdAdd = &cobra.Command {
	Use:	"add",
	Short:	"Add an IP address",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()	
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)

	// subcommands
	addressCmd.AddCommand(addressCmdShow)
	addressCmd.AddCommand(addressCmdList)
	addressCmd.AddCommand(addressCmdDelete)
	addressCmd.AddCommand(addressCmdAdd)
}
