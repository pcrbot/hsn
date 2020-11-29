package cmd

import (
	"fmt"
	"github.com/pcrbot/hsn/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "设置工程目录",
	Long:  `Set the project config`,
	Run: func(cmd *cobra.Command, args []string) {
		writeConfig := func() {
			err := viper.WriteConfig()
			if err != nil {
				fmt.Println("write config error: ", err)
			}
		}
		path := cmd.Flag("path").Value.String()
		if path != "" {
			if path == "." {
				path, _ = os.Getwd()
			}
			if utils.IsExist(path + "/run.py") {
				viper.Set("HOSHINO_PATH", path)
				writeConfig()
			} else {
				fmt.Println("Can't find run.py, please check the path.")
			}
		}

		path = cmd.Flag("image").Value.String()
		if path != "" {
			if strings.HasSuffix(path, "/") { // 去地址尾部
				path = path[:len(path)-1]
			}
			viper.Set("GITHUB_IMAGE", path)
		}

		pip := cmd.Flag("pip").Value.String()
		if pip != "" {
			viper.Set("PIP_COMMAND", pip)
		}
		writeConfig()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("path", "p", "", "set hoshino project path")
	setCmd.Flags().StringP("image", "i", "", "set github image path")
	setCmd.Flags().String("pip", "", "set pip command")
}

func GetGitHubImage() string {
	if image, ok := viper.Get("GITHUB_IMAGE").(string); ok {
		return image
	}
	return "https://github.com" // 默认不使用镜像保证兼容性
}

func GetPipCommand() []string {
	if pip, ok := viper.Get("PIP_COMMAND").(string); ok {
		return strings.Split(pip," ")
	}
	return []string{"pip3"}
}
