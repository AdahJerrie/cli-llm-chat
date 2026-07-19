package main

import (
	"bufio"
	"cli-chat/config"
	"cli-chat/llm"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	input := scanner.Text()

	fmt.Println("start chat: ", input)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := llm.NewClient(config.BaseURL)

	resp, err := client.Generate(ctx, input, config.Model)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp)
}
