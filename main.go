package main

import (
	"net/http"
	"encoding/json"
	"cli-chat/llm"
)

func main() {
	mux := http.NewServeMux()
	mux.S("POST /api/generate", llm.NewClient)

	http.ListenAndServe(":11434", mux)
}