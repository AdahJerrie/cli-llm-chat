package main

import (
	"bufio"
	"cli-chat/config"
	"cli-chat/llm"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, using system environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}

	fmt.Print(">>")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	input := scanner.Text()

	fmt.Println("start chat: ", input)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := llm.NewClient(cfg.BaseURL)

	resp, err := client.Generate(ctx, input, cfg.Model)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp)
}
