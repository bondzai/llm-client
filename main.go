package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	// 1. Get API Key from environment variable
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: GEMINI_API_KEY environment variable is not set.\nPlease set it using: export GEMINI_API_KEY=your_key_here")
	}

	ctx := context.Background()

	// 2. Create a new Gemini client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	defer client.Close()

	// 3. Select the model (gemini-1.5-flash is free-tier eligible and fast)
	model := client.GenerativeModel("gemini-1.5-flash")

	// Optional: Configure generation settings
	model.SetTemperature(0.7)
	topK := int32(40)
	model.TopK = &topK
	topP := float32(0.95)
	model.TopP = &topP
	maxTokens := int32(8192)
	model.MaxOutputTokens = &maxTokens
	model.ResponseMIMEType = "text/plain"

	fmt.Println("---------------------------------------------------------")
	fmt.Println("  Gemini Go Client POC (gemini-1.5-flash)")
	fmt.Println("  Type 'quit' or 'exit' to stop.")
	fmt.Println("---------------------------------------------------------")

	// 4. Start a chat session (maintains history)
	cs := model.StartChat()
	cs.History = []*genai.Content{} // Start with empty history

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nResult > ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "quit" || userInput == "exit" {
			break
		}
		if userInput == "" {
			continue
		}

		fmt.Print("Gemini > ")

		// 5. Send message using streaming for better UX
		iter := cs.SendMessageStream(ctx, genai.Text(userInput))
		for {
			resp, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Printf("\nError generating response: %v", err)
				break
			}
			for _, part := range resp.Candidates[0].Content.Parts {
				if txt, ok := part.(genai.Text); ok {
					fmt.Print(string(txt))
				}
			}
		}
		fmt.Println() // Newline after response
	}
}
