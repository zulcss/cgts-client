package cmd

import (
	"github.com/spf13/cobra"
	"cgts-client/pkg/utils"
)

var addressPoolCmd = &cobra.Command{
	Use:   "addresspool",
	Short: "Address Pool",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()
	},
}

var addressPoolShow = &cobra.Command{
	Use:   "show",
	Short: "Show address pool",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()
	},
}

var addressPoolList = &cobra.Command{
	Use:   "list",
	Short: "List IP address pool",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()
	},
}

var addressPoolDelete = &cobra.Command{
	Use:   "delete",
	Short: "Delete IP address pool",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()
	},
}

var addressPoolAdd = &cobra.Command{
	Use:   "add",
	Short: "Add IP address pool",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()
	},
}

var addressPoolModify = &cobra.Command{
	Use:   "modify",
	Short: "Modify IP address pool",
	Run: func(cmd *cobra.Command, args []string) {
		utils.NotImplemented()
	},
}

func init() {
	rootCmd.AddCommand(addressPoolCmd)

	// subcommands
	addressPoolCmd.AddCommand(addressPoolShow)
	addressPoolCmd.AddCommand(addressPoolList)
	addressPoolCmd.AddCommand(addressPoolDelete)
	addressPoolCmd.AddCommand(addressPoolAdd)
	addressPoolCmd.AddCommand(addressPoolModify)
}
