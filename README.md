# OpenAI-Chat

使用 Go 语言实现命令行版 ChatGPT 应用

## Build

```
shell> go mod tidy
shell> go build ./chat.go
```

## Usage

启动这个应用，如果没有设置 OPENAI_URL 系统环境变量，会提示你设置

```PowerShell
shell> ./chat.exe
请设置 OPENAI_URL 环境变量
    Bash (Linux or macOS): export OPENAI_URL="https://api.openai.com"
    PowerShell (Windows): $env:OPENAI_URL="https://api.openai.com"
#
# 备注：可以设置 OpenAI 官方域名，或者已实现科学上网的代理域名
shell> $env:OPENAI_URL="https://api.openai.com"
```

启动这个应用，如果没有设置 OPENAI_API_KEY 系统环境变量，会提示你设置

```PowerShell
shell> ./chat.exe
请设置 OPENAI_API_KEY 环境变量
    Bash (Linux or macOS): export OPENAI_API_KEY="XXXXXX"
    PowerShell (Windows): $env:OPENAI_API_KEY="XXXXXX"
#
# 备注：把 XXXXXX 替换为自己 OpenAI 账户的 API KEY
shell> $env:OPENAI_API_KEY="XXXXXX"
```

测试使用

```PowerShell
shell> ./chat.exe
   ____   _               _      ____   ____    _____
  / ___| | |__     __ _  | |_   / ___| |  _ \  |_   _|
 | |     | '_ \   / _` | | __| | |  _  | |_) |   | |
 | |___  | | | | | (_| | | |_  | |_| | |  __/    | |
  \____| |_| |_|  \__,_|  \__|  \____| |_|       |_|
输入 start 启动应用，输入 quit 退出应用
start
ChatGPT 启动成功，请输入您的问题：
使用golang写一个输出helloworld的程序

  以下是使用golang编写的输出helloworld程序：

    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, world!")
    }

  通过运行这个程序，终端会输出一个简单的"helloworld"。


使用GoLang格式化当前时间为“yyyy-mm-dd hh:ii:ss”格式

  以下是使用GoLang将当前时间格式化为“yyyy-mm-dd hh:ii:ss”的代码：

    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        currentTime := time.Now()
        formattedTime := currentTime.Format("2006-01-02 15:04:05")
        fmt.Println(formattedTime)
    }

  运行这段代码会输出当前时间的格式化版本，类似于这样： "2021-08-08 15:33:34"
  (会根据当前时间变化)。

  在这个例子中，我们使用了GoLang的time包来获取当前时间并使用Format()函数将其格式化到指定的布局字符串(例如："2006-
  01-02 15:04:05")。


quit
ChatGPT 已退出
```
