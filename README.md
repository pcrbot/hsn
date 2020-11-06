# Hoshino-cli

## 安装方法

Windows的简单粗暴方法：下载release中的对应文件，改名为hsn.exe扔到`C:\Windows\System32`
(只要给`hsn.exe`添加到环境变量Path路径就行了)

Linux的简单粗暴方法：下载release中的对应文件，改名为hsn扔到`\bin`文件夹下 (只要有环境变量就行了)

## 使用方法

### 设置CLI
`hsn set `

- -p (--path) : 填写Hoshino工程目录的绝对路径，或者填写 `.` 将使用当前目录
- -i (--image) : 设置使用Github的镜像源地址,默认为 `https://hub.fastgit.org`

例如 
- `hsn set --path=/root/hoshino`
- `hsn set -i=https://github.com` (好家伙，不用镜像)

### 安装Hoshino插件
`hsn install `

请确保安装了`git`, `python`

例如 `hsn install music`

### 获取可用插件列表
`hsn list`

### 更新CLI
`hsn update`

### 生成发布插件的json文件 

`hsn export [plugin name]`

让后填写仓库地址,所需依赖和资源文件就行了
