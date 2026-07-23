# cli-llm-chat

A simple Go CLI application that demonstrates the fundamentals of integrating with an LLM over HTTP.

This project currently uses a **mock LLM server** instead of a real model (such as Ollama) to focus on learning the request/response lifecycle and building clean, maintainable Go code.

## Learning Objectives

* Build an HTTP client in Go
* Send JSON requests
* Parse JSON responses
* Use `context` for request cancellation
* Manage configuration with environment variables
* Organize code into reusable packages
* Practice clean error handling

---

## Project Structure

```text
cli-chat/
│
├── config/
│   └── config.go          # Loads application configuration
│
├── llm/
│   └── client.go          # LLM client implementation
│
├── mock-server/
│   └── main.go            # Mock LLM HTTP server
│
├── .env.example           # Example environment variables
├── .gitignore
├── go.mod
└── main.go                # CLI entry point
```

---

## Architecture

```text
User
  │
  ▼
CLI Application
  │
  ▼
LLM Client
  │
HTTP POST
  │
  ▼
Mock LLM Server
  │
JSON Response
  │
  ▼
CLI Output
```

The CLI communicates only with the `llm` package. The `llm` package is responsible for all HTTP communication with the server.

---

## Configuration

The application reads its configuration from environment variables.

Example:

```env
BASE_URL=http://localhost:8080
MODEL=llama3
API_KEY=
```

Required variables:

* `BASE_URL`
* `MODEL`

`API_KEY` is optional for now and will be used when integrating with authenticated LLM providers.

---

## Running the Project

### 1. Start the mock server

```bash
go run ./mock-server
```

### 2. Configure environment variables

Example:

```bash
export BASE_URL=http://localhost:8080
export MODEL=llama3
```

(Windows PowerShell)

```powershell
$env:BASE_URL="http://localhost:8080"
$env:MODEL="llama3"
```

### 3. Run the CLI

```bash
go run .
```

Enter a prompt when prompted.

---

## Current Features

* HTTP client for LLM communication
* Mock LLM server
* JSON request/response handling
* Context-based request timeout
* HTTP client timeout
* Environment-based configuration
* Basic error handling
* Modular project structure

---

## Current Limitations

This is intentionally a minimal implementation.

It does **not** yet support:

* Conversation memory
* Streaming responses
* Tool calling
* Structured outputs
* Authentication
* Multiple LLM providers

These features will be introduced in later projects.

---

## Next Milestone

The next project builds on this foundation by introducing **conversation memory**, allowing the CLI to maintain context across multiple user interactions.
