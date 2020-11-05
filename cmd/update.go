package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pcrbot/hsn/utils"
	"github.com/spf13/cobra"
	"runtime"
)

var Version = "unknown"

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update cli or package",
	Long:  `Update cli or package`,
	Run: func(cmd *cobra.Command, args []string) {
		p := PluginPackage{}
		rsp, err := utils.Download(fmt.Sprint("https://cdn.jsdelivr.net/gh/pcrbot/hsn@main/package.json"))
		if err != nil {
			fmt.Println("获取版本信息失败！")
			return
		}
		err = json.Unmarshal(rsp, &p)
		if err != nil {
			fmt.Println("获取版本信息失败！")
			return
		}

		if Version == p.Version {
			fmt.Println("已是最新版本,无需更新!")
			return
		}

		fmt.Printf("当前版本: %v\n最新版本: %v", Version, p.Version)

		_ = utils.DownloadFile(
			fmt.Sprintf(
				"https://github.com/pcrbot/hsn/releases/download/%v/hsn-%v-%v-%v.zip",
				p.Version,
				p.Version,
				runtime.GOOS,
				runtime.GOARCH,
			),
			"111.zip",
		)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

}
