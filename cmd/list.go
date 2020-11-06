package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pcrbot/hsn/utils"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "可用插件列表",
	Long:  `list all of the available plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		p := PluginPackage{}
		rsp, err := utils.Download("https://cdn.jsdelivr.net/gh/pcrbot/hsn@main/package.json")
		if err != nil {
			fmt.Println("获取插件插件失败！")
			return
		}
		err = json.Unmarshal(rsp, &p)
		if err != nil {
			fmt.Println("获取插件列表失败！")
			return
		}

		for index, plugin := range p.PluginList {
			fmt.Println(index, ": ", plugin)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

type PluginPackage struct {
	Version    string   `json:"version"`
	PluginList []string `json:"plugin_list"`
}
