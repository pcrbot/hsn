package hoshino

import (
	"fmt"
	"github.com/pcrbot/Hoshino-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "设置工程目录",
	Long:  `Set the project config`,
	Run: func(cmd *cobra.Command, args []string) {
		path := cmd.Flag("path").Value.String()

		if path == "." {
			path, _ = os.Getwd()
		}

		if utils.IsExist(path + "/run.py") {
			viper.Set("HOSHINO_PATH", path)
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println("write config error: ", err)
			}
		} else {
			fmt.Println("Can't find run.py, please check the path.")
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("path","p", ".", "set hoshino project path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
