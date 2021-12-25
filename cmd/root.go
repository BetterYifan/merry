package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "iptrans",
	Short: "iptrans is a proj to transfer decimal ip addr to binary",
	Long:  "iptrans is a proj to transfer decimal ip addr to binary",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(d2bCommand)
	rootCmd.AddCommand(b2dCommand)
}
