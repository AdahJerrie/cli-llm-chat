Project Overview
Objective

The objective of this project is not to build a production LLM client.

The objective is to understand the fundamental mechanics behind LLM integration while applying sound software engineering principles.

Since access to a local LLM (Ollama) was not available due to device restrictions, we intentionally introduced a mock HTTP server that simulates an LLM endpoint.

This allows development to continue independently of the actual AI model.

This follows a common engineering practice:

Replace external dependencies with controlled simulations so application development can continue independently.

This means that throughout the project we are learning:

HTTP communication
Request/Response lifecycle
JSON serialization
Context handling
Configuration
Error handling

without introducing the complexity of a real model.

High-Level Architecture
                User
                  │
                  ▼
            CLI Application
                  │
                  ▼
            LLM Client Package
                  │
          HTTP POST Request
                  │
                  ▼
            Mock LLM Server
                  │
          Fixed JSON Response
                  │
                  ▼
            LLM Client Package
                  │
                  ▼
            CLI Application
                  │
                  ▼
             Console Output

Notice that the CLI never communicates directly with the server.

Everything goes through the LLM client.

That separation is intentional.

Why a dedicated LLM package?

Instead of writing HTTP requests inside main.go, we extracted that responsibility into its own package.

main.go

↓

llm.Client

↓

HTTP Server

This follows the Single Responsibility Principle.

main.go is responsible for application flow.

The llm package is responsible for communicating with an LLM service.

If tomorrow we build:

Discord Bot
REST API
Slack Bot
AI Documentation Generator

they all reuse the same client package.

The HTTP logic exists only once.

Why create a mock server?

Initially, the project was intended to communicate with Ollama.

However, because the development machine prevented software installation, we replaced Ollama with our own HTTP server.

Instead of treating this as a limitation, we used it as an opportunity to learn the protocol itself.

The mock server simulates the API contract.

CLI

↓

POST /api/generate

↓

JSON Request

↓

JSON Response

From the client's perspective, it does not matter whether the server is:

Ollama
OpenAI
Anthropic
or our mock implementation.

As long as the API contract remains consistent, the client behaves identically.

This mirrors how integration testing is often performed in industry.

Request Lifecycle

The Generate method follows a well-defined sequence.

User Input

↓

Build Request Struct

↓

Marshal JSON

↓

Create HTTP Request

↓

Attach Context

↓

Send Request

↓

Receive Response

↓

Read Response Body

↓

Unmarshal JSON

↓

Return Result

Each stage has one responsibility.

This makes debugging significantly easier.

Why use structs instead of maps?

The request and response are represented using Go structs.

type GenerateRequest struct {
    Model string
    Prompt string
    Stream bool
}

instead of

map[string]any

because structs provide:

compile-time type safety,
self-documenting APIs,
easier maintenance,
IDE autocomplete,
clearer JSON contracts.

As the project grows, this becomes increasingly valuable.

Why introduce a Config package?

Initially the application contained:

"http://localhost:8080"

"llama3"

inside main.go.

These values are not business logic.

They are configuration.

So we introduced a dedicated configuration layer.

Now:

Environment Variables

↓

Config Package

↓

Main Application

The application no longer cares where configuration comes from.

Today:

.env

Tomorrow:

Docker
Kubernetes
AWS Secrets Manager

The application remains unchanged.

This is an important separation between behavior and configuration.

Why validate configuration?

Instead of allowing failures later:

HTTP request fails

↓

Model missing

↓

Application crashes

we fail immediately during startup.

Load configuration

↓

Validate

↓

Start application

This follows the Fail Fast principle.

Errors should be discovered as early as possible.

Why use Context?

Every request is wrapped in:

context.WithTimeout(...)

The purpose is not only timeout control.

Context establishes a mechanism for cancellation.

If:

the server hangs,
the user interrupts,
the application shuts down,

the request can be cancelled cleanly.

Context is the standard mechanism used throughout Go's networking libraries.

Why add an HTTP client timeout?

The client now owns its own timeout.

This protects against failures such as:

DNS lookup delays
TCP connection hangs
TLS negotiation
slow servers
incomplete responses

Even if the caller forgets to apply a context timeout, the client still has a defensive timeout.

Multiple layers of protection improve robustness.

Error Handling Philosophy

Instead of assuming success, every stage checks for failure.

Marshal

↓

Create Request

↓

Send Request

↓

Read Body

↓

Decode JSON

Every operation can fail independently.

Each error is wrapped with contextual information.

For example:

encoding request body

instead of

json error

This provides far more useful diagnostics.

Why not add more features?

During development we intentionally avoided introducing:

streaming
interfaces
retries
middleware
abstractions
dependency injection frameworks
multiple providers

This decision was deliberate.

The goal of Project 1 is to understand:

How does software communicate with an LLM?

Adding unrelated features would increase complexity without improving understanding.

Complexity should only be introduced when a project requires it.

Why no interfaces?

Many tutorials immediately introduce:

type LLM interface {
    Generate(...)
}

We intentionally avoided this.

There is only one implementation.

Introducing an interface now would be speculative abstraction.

Interfaces should emerge when multiple implementations exist.

For example:

MockClient

OllamaClient

OpenAIClient

At that point an interface solves a real problem.

Today it does not.

Project Architecture Philosophy

Throughout development we followed one guiding principle:

Every package should have one clear responsibility.

Current responsibilities are:

main.go

Application orchestration.

config

Application configuration.

llm

Communication with an LLM endpoint.

mock-server

Simulation of an external AI provider.

This separation improves readability, testability, and maintainability.

Lessons Learned

Technically, this project is about much more than sending an HTTP request. It introduces the foundational building blocks of AI systems engineering:

Designing software around clear responsibilities.
Separating configuration from behavior.
Building reusable clients instead of embedding HTTP logic in the application.
Using mock services to decouple development from external dependencies.
Applying Go's context package for request lifecycles and cancellation.
Structuring request and response data with typed contracts rather than unstructured maps.
Handling errors explicitly and failing early when configuration is invalid.
Writing code that can evolve incrementally without premature abstraction.
Why this project matters in the bigger curriculum

From a curriculum perspective, this project establishes the infrastructure that every later project will reuse.

Project 1
│
├── HTTP Client
├── Configuration
├── Error Handling
├── Mock Server
└── Clean Architecture
        │
        ▼
Project 2
Conversation Memory
        │
        ▼
Project 3
Streaming
        │
        ▼
Project 4
Structured Output
        │
        ▼
Project 5
AI Developer Assistant

Each subsequent project introduces one major concept while building on a stable, well-understood foundation rather than repeatedly rewriting the basics.

One final observation

There's something subtle but important about the path you've chosen.

Many AI tutorials begin with, "Here's how to call the OpenAI API." They teach the API before they teach the software.

This project does the opposite. It teaches software engineering first, using an LLM as the external service. That means when you eventually switch from your mock server to Ollama, or from Ollama to OpenAI, the change is primarily one of configuration and API details—not of architecture.

In my view, that's a much stronger foundation for becoming a production AI systems engineer than learning a specific provider's SDK first.