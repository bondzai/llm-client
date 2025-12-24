package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/viper"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	apiKey := loadConfig()
	ctx := context.Background()

	client := createClient(ctx, apiKey)
	defer client.Close()

	model := configureModel(client)

	runChatSession(ctx, model)
}

func loadConfig() string {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	apiKey := viper.GetString("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: GEMINI_API_KEY environment variable is not set.\nPlease set it using: export GEMINI_API_KEY=your_key_here")
	}
	return apiKey
}

func createClient(ctx context.Context, apiKey string) *genai.Client {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}
	return client
}

func configureModel(client *genai.Client) *genai.GenerativeModel {
	model := client.GenerativeModel("gemini-flash-latest")
	model.SetTemperature(0.7)
	topK := int32(40)
	model.TopK = &topK
	topP := float32(0.95)
	model.TopP = &topP
	maxTokens := int32(8192)
	model.MaxOutputTokens = &maxTokens
	model.ResponseMIMEType = "text/plain"
	return model
}

func runChatSession(ctx context.Context, model *genai.GenerativeModel) {
	printForIntro()

	cs := model.StartChat()
	cs.History = []*genai.Content{}

	reader := bufio.NewReader(os.Stdin)

	for {
		input := readInput(reader)
		if shouldExit(input) {
			break
		}
		if input == "" {
			continue
		}

		streamResponse(ctx, cs, input)
	}
}

func printForIntro() {
	fmt.Println("---------------------------------------------------------")
	fmt.Println("  Gemini Go Client POC (gemini-flash-latest)")
	fmt.Println("  Type 'quit' or 'exit' to stop.")
	fmt.Println("---------------------------------------------------------")
}

func readInput(reader *bufio.Reader) string {
	fmt.Print("\nResult > ")
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}

func shouldExit(input string) bool {
	return input == "quit" || input == "exit"
}

func streamResponse(ctx context.Context, cs *genai.ChatSession, input string) {
	fmt.Print("Gemini > ")

	iter := cs.SendMessageStream(ctx, genai.Text(input))
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("\nError generating response: %v", err)
			break
		}
		printCandidates(resp)
	}
	fmt.Println()
}

func printCandidates(resp *genai.GenerateContentResponse) {
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			fmt.Print(string(txt))
		}
	}
}
