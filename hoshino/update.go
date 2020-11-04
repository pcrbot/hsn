package hoshino

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "unknown"

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update cli or package",
	Long:  `Update cli or package`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}
