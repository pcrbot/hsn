package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pcrbot/hsn/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export your plugin to json",
	Long:  `快速生成发布插件的json文件`,
	Run: func(cmd *cobra.Command, args []string) {
		var exportName string
		if args == nil || len(args) == 0 {
			fmt.Println("请输入您要导出的插件名(文件夹名)...")
			prompt := promptui.Prompt{
				Label: "PluginName",
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("无法获取您的输入: %v\n", err)
				return
			}
			exportName = result
		} else {
			exportName = args[0]
		}

		hsnPath, err := GetHoshinoPath()
		if err != nil {
			fmt.Println(err)
			return
		}

		exportFolder := fmt.Sprint(hsnPath, "/hoshino/modules/", exportName)
		if !utils.IsDir(exportFolder) {
			fmt.Printf("无法找到您要导出的文件夹...")
			return
		}

		exportFiles := utils.GetFiles(exportFolder)
		export := pluginInfo{
			Name:        exportName,
			Version:     "1.0.0",
			Description: "",
			Plugin: &plugin{
				Git:          "",
				Files:        []string{},
				Requirements: []string{},
			},
		}
		for _, file := range exportFiles {
			export.Plugin.Files = append(export.Plugin.Files, file[len(exportFolder)+1:])
		}

		exportData, err := json.Marshal(export)
		if err != nil {
			fmt.Printf("导出失败: %v\n", err)
			return
		}

		err = ioutil.WriteFile(exportName+".json", exportData, os.ModePerm)
		if err != nil {
			fmt.Printf("导出失败: %v\n", err)
			return
		}
		fmt.Println("导出成功！")
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
