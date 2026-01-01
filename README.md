# LLM Client POC (Go)

This is a Proof of Concept (POC) demonstrating how to integrate a free LLM API (Google Gemini Flash) using Go.
...
## How it Works

- Uses the official `github.com/google/generative-ai-go` SDK.
- Connects to the `gemini-flash-latest` model (fast and free-tier eligible).
- Maintains a chat history within the session.
- Streams responses for a better user experience.
