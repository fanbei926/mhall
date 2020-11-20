package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init()  {
	deployCmd.AddCommand(helpCmd)
}

var helpCmd = &cobra.Command{
	Use: "help",
	Short: "Show deploy command info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`mahll deploy -m <gh|Ghall|basic|hall|ghall|Video|Video>`)
	},
}
