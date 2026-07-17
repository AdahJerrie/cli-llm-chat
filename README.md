# cli-llm-chat
The real-world problem we're solving
You have a program (your CLI app) and a completely separate program (Ollama, running as a background service on your machine) that holds an LLM. They are two different processes that don't share memory. The only way for your CLI to say "hey, here's a prompt, give me a response" is to talk over a network protocol — even though it's all happening on localhost.
That's the entire reason HTTP exists in this context: it's a shared language two independent programs use to exchange data.

┌──────────────┐                              ┌──────────────┐
│  Your CLI    │──── 1. HTTP Request ───────▶ │ Ollama Server │
│  (Go program)│      (method, URL, headers,   │ (localhost:   │
│              │       body = JSON prompt)      │  11434)       │
│              │◀─── 2. HTTP Response ──────── │               │
└──────────────┘      (status code, headers,    └──────────────┘
                       body = JSON with reply)