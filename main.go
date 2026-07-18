package main

import (
	"bufio"
	"cli-chat/llm"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Println("Error reading input:", scanner.Err())
		return
	}
	input := scanner.Text()

	fmt.Println("start chat: ", input)

	client := llm.NewClient("http://localhost:8080")

	resp, err := client.Generate(input, "llama3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp)
}
