package kncli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kncli",
	Short: "Kilonova CLI",
	Long:  "Interact with Kilonova from the command line.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing kncli '%s'", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

}
