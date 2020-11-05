package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pcrbot/hsn/utils"
	"github.com/spf13/cobra"
	"net/http"
	"runtime"

	"github.com/inconshreveable/go-update"
	"github.com/manifoldco/promptui"
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

		fmt.Printf("当前版本: %v\n最新版本: %v\n", Version, p.Version)
		prompt := promptui.Prompt{
			Label:     "你确定要更新吗?",
			IsConfirm: true,
		}

		_, err = prompt.Run()

		if err != nil {
			fmt.Printf("更新取消: %v\n", err)
			return
		}

		fmt.Println("正在更新Hohsino-cli,请稍等...")

		url := fmt.Sprintf(
			"%v/pcrbot/hsn/releases/download/%v/hsn-%v-%v-%v",
			GetGitHubImage(),
			p.Version,
			p.Version,
			runtime.GOOS,
			runtime.GOARCH,
		)
		if runtime.GOOS == "windows" {
			url = url + ".exe"
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("更新失败!")
			return
		}
		err = update.Apply(resp.Body, update.Options{})
		if err != nil {
			fmt.Println("更新失败!")
			return
		}
		fmt.Println("更新完成！")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
