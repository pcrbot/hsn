## 插件适配指南

### Git仓库设置

如果您的仓库像[H-K-Y/Genshin_Impact_bot](https://github.bajins.com/H-K-Y/Genshin_Impact_bot) 一样支持
在`HoshinoBot\hoshino\modules` 目录下使用`git clone`指令安装，只需要像下面这样填写(不能有注释)

```
{
  "name": "genshin", // 生成的文件夹名
  "version": "1.0.0", // 版本号
  "description": "原神相关娱乐和信息查询功能", // 插件描述
  "plugin": {
    "git": "H-K-Y/Genshin_Impact_bot.git", // 插件仓库地址，不需要 https://github.com
    "requirements": []
  }
}
```
hsn 会根据用户设置的github镜像生成git地址并clone仓库到指定目录

如果您的仓库不能直接git clone(比如一些插件库),您需要像下面这样填写
(hsn会拼接链接从jsdelivr上下载),其中 `files` 可以使用 `hsn export`
指令自动生成.

```
{
  "name": "generator",
  "version": "1.0.0",
  "description": "生成器插件",
  "plugin": {
    "git": "pcrbot/cappuccilo_plugins/generator",
    "files": [
      "config.json",
      "data.json",
      "diary.png",
      "diary_data.json",
      "generator.py",
      "jichou.jpg"
    ],
    "requirements": [],
    "res": {
      "hoshino/modules/generator/simhei.ttf": "https://github.com/pcrbot/cappuccilo_plugins/releases/download/v1.0/simhei.ttf"
    }
  }
}
```

### Python依赖

工作原理是拼接pip 指令从清华源上安装
和requirements.txt基本一致,例如:

```
{
  "name": "rss",
  "version": "1.0.0",
  "description": "rss订阅插件,默认使用地河云RSSHub",
  "plugin": {
    "git": "zyujs/rss",
    "files": [
      "rss.py",
      "README.md"
    ],
    "requirements": [
      "feedparser~=5.2",
      "pillow~=6.2"
    ]
  }
}
```

### 资源文件

这部分处理的比较暴力，例如：
```
"res": {
  "arial.ttf": "http://d.xiazaiziti.com/en_fonts/fonts/a/Arial.ttf"
}
```
这样填写会在Hoshino根目录(run.py所在的路径)下, 从`http://d.xiazaiziti.com/en_fonts/fonts/a/Arial.ttf` 
下载文件`arial.ttf`

```
"res": {
  "hoshino/modules/generator/simhei.ttf": "https://github.com/pcrbot/cappuccilo_plugins/releases/download/v1.0/simhei.ttf"
}
```
`hoshino/modules/generator/simhei.ttf` 是目标文件对于Hoshino根目录的相对路径,
 `https://github.com/pcrbot/cappuccilo_plugins/releases/download/v1.0/simhei.ttf`
 是目标文件的下载直链
 
 支持下载多个资源文件 (某些需要套两层文件夹的插件可以全部当作资源文件处理)
