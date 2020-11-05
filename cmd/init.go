package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化工程",
	Long: `A simple way to clone the Hoshino project
`,
	Run: func(cmd *cobra.Command, args []string) {
		git := exec.Command("git", "clone", "https://github.com/Ice-Cirno/HoshinoBot.git")
		git.Stdout = os.Stdout
		git.Stderr = os.Stderr
		err := git.Run()
		fmt.Println("[ERROR] ", err)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
