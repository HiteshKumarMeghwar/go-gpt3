package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/vuln/client"
)

func GetResponse(client gpt3.Client, ctx context.Context, question string)  {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{ 
			question,
		},
		MaxTokens: gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
		},func (resp *gpt.CompletionResponse)  {
			fmt.Print(resp.Choises[0].Text)
	})
}

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		panic("Missing API KEY")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	rootCmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "chat with ChatGPT in console",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Say something ('quit' to end):")
				if !scanner.Scan() {
					break
				}
				question := scanner.Text()
				switch question {
				case "quit":
					quit = true
				default:
					GetResponse(client, ctx, question)
				}
			}
		},
	}
	rootCmd.Execute()
}
