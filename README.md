# LLM Client POC (Go)

This is a Proof of Concept (POC) demonstrating how to integrate a free LLM API (Google Gemini 2.0 Flash) using Go.
...
## How it Works

- Uses the official `github.com/google/generative-ai-go` SDK.
- Connects to the `gemini-2.0-flash` model (fast and free-tier eligible).
- Maintains a chat history within the session.
- Streams responses for a better user experience.

## Push to GitHub

To make this repository public on GitHub:

1. Create a new repository on [GitHub](https://github.com/new).
2. Follow the commands:
   ```bash
   git add .
   git commit -m "Initial commit: Gemini API POC"
   git branch -M main
   git remote add origin https://github.com/YOUR_USERNAME/YOUR_REPO_NAME.git
   git push -u origin main
   ```
