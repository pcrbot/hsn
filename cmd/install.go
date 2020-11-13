package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pcrbot/hsn/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装插件",
	Long:  `A quick way to install plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) == 0 {
			fmt.Println("请输入你要安装的插件，例如 hsn install eqa")
			return
		}
		hoshinoPath, err := GetHoshinoPath()
		if err != nil {
			fmt.Println(err)
			return
		}

		pluginName := args[0]
		if i, err := strconv.Atoi(pluginName);err == nil {
			pluginName = GetPluginName(i)
		}
		p := pluginInfo{}
		rsp, err := utils.Download(fmt.Sprint("https://cdn.jsdelivr.net/gh/pcrbot/hsn@main/bucket/", pluginName, ".json"))
		if err != nil {
			fmt.Println("获取插件信息失败！")
			return
		}
		err = json.Unmarshal(rsp, &p)
		if err != nil {
			fmt.Println("获取插件信息失败！")
			return
		}
		fmt.Println("插件名: ", pluginName)
		fmt.Println("包名: ", p.Name)
		fmt.Println("当前版本: ", p.Version)
		fmt.Println("插件说明: ", p.Description)
		prompt := promptui.Prompt{
			Label:     "你确定要安装该插件吗?",
			IsConfirm: true,
		}

		_, err = prompt.Run()

		if err != nil {
			fmt.Printf("已取消安装...")
			return
		}

		for index, requirement := range p.Plugin.Requirements { // 安装依赖
			fmt.Println("正在安装依赖", index, ": ", requirement)
			pip := exec.Command("pip3", "install", requirement, "-i", "https://pypi.tuna.tsinghua.edu.cn/simple")
			pip.Stdout = os.Stdout
			pip.Stderr = os.Stderr
			err := pip.Run()
			if err != nil {
				fmt.Println("[ERROR] ", err)
			}
		}

		func() { // 下载文件
			path := hoshinoPath + "/hoshino/modules/" + p.Name
			if !utils.IsExist(path) {
				_ = os.Mkdir(path, os.ModePerm)
			}

			if p.Plugin.Files != nil && len(p.Plugin.Files) != 0 {
				for _, file := range p.Plugin.Files {
					paths := strings.Split(file, "/")

					var ph = path
					for _, f := range paths[:len(paths)-1] {
						ph = ph + "/" + f
						if !utils.IsExist(ph) {
							_ = os.Mkdir(ph, os.ModePerm)
						}
					}

					fmt.Println(fmt.Sprint("https://cdn.jsdelivr.net/gh/", p.Plugin.Git, "/", file))
					err = utils.DownloadFile(
						fmt.Sprint("https://cdn.jsdelivr.net/gh/", p.Plugin.Git, "/", file),
						fmt.Sprint(path+"/"+file),
					)

					if err != nil {
						fmt.Println("下载文件: ", file, " 失败: ", err, "!")
					}
				}
				return
			}

			git := exec.Command("git", "clone", fmt.Sprint(GetGitHubImage(), "/", p.Plugin.Git), path)
			git.Stdout = os.Stdout
			git.Stderr = os.Stderr
			err := git.Run()
			if err != nil {
				fmt.Println("[ERROR] ", err)
			}
		}()

		configPath := fmt.Sprint(hoshinoPath, "/hoshino/config/__bot__.py")
		data, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Println("Can't read hoshino config.")
			return
		}
		data = append(data, []byte(fmt.Sprintf("\nMODULES_ON.add('%v')", p.Name))...)
		err = ioutil.WriteFile(configPath, data, os.ModePerm)
		if err != nil {
			fmt.Println("Failed to write config: ", err)
			return
		}
		if p.Plugin.Res != nil {
			fmt.Println("正在下载资源文件...")
			res := gjson.ParseBytes(p.Plugin.Res)
			res.ForEach(func(key, value gjson.Result) bool {
				file := hoshinoPath + "/" + key.String()
				fmt.Println("正在下载文件 ", file, "...")

				paths := strings.Split(file, "/") // 创建路径
				var ph = hoshinoPath
				for _, f := range paths[:len(paths)-1] {
					ph = ph + "/" + f
					if !utils.IsExist(ph) {
						_ = os.Mkdir(ph, os.ModePerm)
					}
				}

				err = utils.DownloadFileWithProgress(file, value.String())
				if err != nil {
					fmt.Println("下载文件 ", key.Value(), "失败...")
				}
				return true
			})
		}

		fmt.Println("Success to install plugin ", p.Name, ".")
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

type pluginInfo struct {
	Name        string  `json:"name"`
	Version     string  `json:"version"`
	Description string  `json:"description"`
	Plugin      *plugin `json:"plugin"`
}

type plugin struct {
	Git          string          `json:"git"`
	Files        []string        `json:"files"`
	Requirements []string        `json:"requirements"`
	Res          json.RawMessage `json:"res"`
}

func GetHoshinoPath() (string, error) {
	if ph, ok := viper.Get("HOSHINO_PATH").(string); ok {
		return ph, nil
	} else {
		return "", errors.New("can't find HOSHINO_PATH, use 'hsn set --path=' to set HOSHINO_PATH")
	}
}

func GetPluginName(index int) string {
	p := PluginPackage{}
	rsp, err := utils.Download("https://cdn.jsdelivr.net/gh/pcrbot/hsn@main/package.json")
	if err != nil {
		fmt.Println("获取插件列表失败！")
		return ""
	}
	err = json.Unmarshal(rsp, &p)
	if err != nil {
		fmt.Println("获取插件列表失败！")
		return ""
	}

	if index < len(p.PluginList) {
		return p.PluginList[index]
	}
	return ""
}
