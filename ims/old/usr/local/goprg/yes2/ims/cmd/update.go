package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var update = &cobra.Command{
	Use:   "update",
	Short: "Updates a value of a key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}
