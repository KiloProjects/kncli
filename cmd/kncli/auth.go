package kncli

import (
	"fmt"
	"github.com/KiloProjects/kncli/pkg/kncli"
	"github.com/spf13/cobra"
)

func init() {
	var authCmd = &cobra.Command{
		Use:   "auth",
		Short: "Authenticate with Kilonova",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			res := kncli.GenerateRandomNumber()
			fmt.Println(res)
		},
	}

	rootCmd.AddCommand(authCmd)
}
