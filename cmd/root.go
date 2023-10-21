package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "go-backstage",
	Short: "go-backstage is a go alternative for backstage.io",
	Long: `A fast and flexible cli that aims to serve as a backstage to enterprise services 
	serving as a API Catalog, Management and Observability resource`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
