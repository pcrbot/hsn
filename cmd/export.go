package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export your plugin to json",
	Long:  `快速生成发布插件的json文件`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
