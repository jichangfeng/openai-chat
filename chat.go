package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/glamour"
	figure "github.com/common-nighthawk/go-figure"
	openai "github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func main() {
	// 获取 OpenAI Url
	apiUrl := os.Getenv("OPENAI_URL")
	if apiUrl == "" {
		fmt.Println("请设置 OPENAI_URL 环境变量")
		fmt.Println("    Bash (Linux or macOS): export OPENAI_URL=\"https://api.openai.com\"")
		fmt.Println("    PowerShell (Windows): $env:OPENAI_URL=\"https://api.openai.com\"")
		return
	}
	// 获取 OpenAI API Key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("请设置 OPENAI_API_KEY 环境变量")
		fmt.Println("    Bash (Linux or macOS): export OPENAI_API_KEY=\"XXXXXX\"")
		fmt.Println("    PowerShell (Windows): $env:OPENAI_API_KEY=\"XXXXXX\"")
		return
	}

	// 初始化 Glamour 渲染器
	renderStyle := glamour.WithEnvironmentConfig()
	mdRenderer, err := glamour.NewTermRenderer(
		renderStyle,
	)
	if err != nil {
		fmt.Println("初始化 Markdown 渲染器失败")
		return
	}

	// 输出欢迎语(命令行应用启动界面)
	myFigure := figure.NewFigure("ChatGPT", "", true)
	myFigure.Print()
	fmt.Println("输入 start 启动应用，输入 quit 退出应用")

	// 创建 ChatGPT 客户端
	//client := openai.NewClient(apiKey)
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = apiUrl + "/v1"
	client := openai.NewClientWithConfig(config)
	if err != nil {
		fmt.Printf("创建客户端失败: %s\n", err.Error())
		return
	}

	messages := []openai.ChatCompletionMessage{
		{
			Role:    "system",
			Content: "你是ChatGPT, OpenAI训练的大型语言模型, 请尽可能简洁地回答我的问题",
		},
	}

	// 读取用户输入并交互
	var userInput string
	for {
		fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "start" {
			fmt.Println("ChatGPT 启动成功，请输入您的问题：")
		} else if strings.ToLower(userInput) == "quit" {
			fmt.Println("ChatGPT 已退出")
			return
		} else if userInput != "" {
			messages = append(
				messages, openai.ChatCompletionMessage{
					Role:    "user",
					Content: userInput,
				},
			)
			// 调用 ChatGPT API 接口生成回答
			resp, err := client.CreateChatCompletion(
				context.Background(),
				openai.ChatCompletionRequest{
					Model:       openai.GPT3Dot5Turbo,
					Messages:    messages,
					MaxTokens:   1024,
					Temperature: 0,
					N:           1,
				},
			)
			if err != nil {
				fmt.Printf("ChatGPT 接口调用失败: %s\n", err.Error())
				continue
			}

			// 格式化输出结果
			output := resp.Choices[0].Message.Content
			mdOutput, err := mdRenderer.Render(output)
			if err != nil {
				fmt.Printf("Markdown 渲染失败: %s\n", err.Error())
				continue
			}
			fmt.Println(mdOutput)
			messages = append(
				messages, openai.ChatCompletionMessage{
					Role:    "assistant",
					Content: output,
				},
			)
		}
	}
}
