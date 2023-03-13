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

  下面是适用于Go语言的Hello World程序：

    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, world!")
    }

  可以在本地编译运行，也可以在很多在线编辑器中编写和运行。


使用GoLang格式化当前时间为“yyyy-mm-dd hh:ii:ss”格式

  您可以使用以下方式在Go中以“yyyy-mm-dd”格式格式化当前日期：

    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        t := time.Now()
        fmt.Println(t.Format("2006-01-02"))
    }

  输出将为形如“2022-02-22”的字符串。

  要以“h:ii:ss”格式格式化当前时间，您可以使用以下方式在Go中：

    package main

    import (
        "fmt"
        "time"
    )

    func main() {
        t := time.Now()
        fmt.Println(t.Format("15:04:05"))
    }

  输出将为形如“14:35:55”的字符串。这里使用了 15:04:05
  的格式来表示小时、分钟和秒。


quit
ChatGPT 已退出
```
