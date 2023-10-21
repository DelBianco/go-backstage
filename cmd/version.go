package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "go-backstage version",
	Long:  `Print the version number of go-backstage`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-backstage v0.0.1 dev")
	},
}
