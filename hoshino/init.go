/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package hoshino

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Clone the Hoshino project",
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
