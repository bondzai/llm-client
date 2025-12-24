# LLM Client POC (Go)

This is a Proof of Concept (POC) demonstrating how to integrate a free LLM API (Google Gemini 1.5 Flash) using Go.

## Prerequisites

- [Go](https://go.dev/dl/) installed (version 1.21+ recommended).
- A Google account to generate a free API Key.

## Setup

### 1. Get a Free API Key

1. Valid visit [Google AI Studio](https://aistudio.google.com/app/apikey).
2. Click **Create API key**.
3. Copy the key.

### 2. Clone/Initialize

```bash
git clone https://github.com/YOUR_USERNAME/llm-client.git
cd llm-client
go mod tidy
```

### 3. Run the Application

You can set the `GEMINI_API_KEY` in a `.env` file or export it as an environment variable.

**Option A: Using .env file (Recommended)**
1. Copy the example file:
   ```bash
   cp .env.example .env
   ```
2. Edit `.env` and paste your API Key.
3. Run the app:
   ```bash
   go run main.go
   ```

**Option B: Environment Variable**
**Mac/Linux:**
```bash
export GEMINI_API_KEY="your_api_key_here"
go run main.go
```

**Windows (PowerShell):**
```powershell
$env:GEMINI_API_KEY="your_api_key_here"
go run main.go
```

## How it Works

- Uses the official `github.com/google/generative-ai-go` SDK.
- Connects to the `gemini-1.5-flash` model (fast and free-tier eligible).
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
