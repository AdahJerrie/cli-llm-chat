package main

import (
	"cli-chat/llm"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/generate", generateHandler)

	fmt.Printf("starting mock-server on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Print("starting server:", err)
	}
}

func generateHandler(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	reqbyte, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}

	var genReq llm.GenerateRequest
	err = json.Unmarshal(reqbyte, &genReq)
	if err != nil {
		http.Error(w, "failed to unmarshal request body", http.StatusInternalServerError)
		return
	}

	resp := llm.GenerateResponse{
		Model:      genReq.Model,
		Response:   "The earth was first proposed to be spherical",
		Created_at: http.TimeFormat,
		Done:       true,
	}

	respbyte, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respbyte)
}
