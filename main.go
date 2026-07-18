package main

import (
	"bufio"
	"cli-chat/llm"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	input := scanner.Text()

	fmt.Println("start chat: ", input)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := llm.NewClient("http://localhost:8080")

	resp, err := client.Generate(ctx, input, "llama3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp)
}
