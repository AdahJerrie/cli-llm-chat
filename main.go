package main

import (
	"cli-chat/llm"
	"fmt"
)

func main() {
	client := llm.NewClient("http://localhost:11434")

	resp, err := client.Generate("what is the shape of the earth", "llama3")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(resp)
}
