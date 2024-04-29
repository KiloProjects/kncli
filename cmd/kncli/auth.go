package kncli

import (
	"fmt"
	"github.com/KiloProjects/kncli/pkg/kncli"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Kilonova",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		res := kncli.GenerateRandomNumber()
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
