package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init()  {
	rootCmd.AddCommand(deployCmd)
}

var rootCmd = &cobra.Command{
	Use: "mhall [command]",
	Short: "Deploy mhall env.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return  nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}