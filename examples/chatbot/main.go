package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/bitcomputing/go-openai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		return
	}
	config := openai.DefaultConfig(os.Getenv("OPENAI_API_KEY"))
	config.BaseURL = "https://proxy-dev.ai.opencord.so/openai/v3"
	client := openai.NewClientWithConfig(config)

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
		},
	}
	fmt.Println("Conversation")
	fmt.Println("---------------------")
	fmt.Print("> ")
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		//fmt.Printf("%+v\n\n", resp)
		fmt.Printf("jobID:%s\n\n", resp.JobID)
		fmt.Print("> ")
	}
}
