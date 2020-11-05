# Hoshino-cli

## 安装方法

Windows的简单粗暴方法：下载release中的对应文件，改名为hsn.exe扔到`C:\Windows\System32`，并在`C:\Users\Administrator`目录下创建文件`.hoshino.yml`

## 使用方法

### 设置Hoshino工程目录 
`hsn set `

- --path : 填写Hoshino工程目录的绝对路径，或者不填写将使用当前目录

例如 `hsn set --path=/root/hoshino`

### 安装Hoshino插件
`hsn install `

请确保安装了`git`, `python`

例如 `hsn install music`
