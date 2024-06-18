package parsers

import (
	"context"
	"fmt"
	"strings"

	"github.com/claytoncasey01/open-recipe-gin/dto"
	"github.com/sashabaranov/go-openai"
)

type OpenAIParser struct {
	Client *openai.Client
}

var PROMPTS = []openai.ChatCompletionMessage{
	{
		Role: openai.ChatMessageRoleUser,
		Content: `You are a helpful AI assistant that takes unstructured and possibly partial recipe data and turns it into complete structured recipes using the below format.
{
"name": string,
"description": string,
"difficulty": uint,
"total_calories": uint,
"total_prep_time": string,
"ingredients": [
{
"name": string,
"quantity": string,
"measurement_unit": string
}
],
"directions": [
{
"description": string,
"time": null,
"order": uint
}
]
}`,
	},
	{
		Role:    openai.ChatMessageRoleAssistant,
		Content: "Sure! Please provide the unstructured recipe data, and I'll convert it into the structured format.",
	},
}

func (p OpenAIParser) Parse(content string) (*dto.SuggestedRecipeDTO, error) {
	resp, err := p.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    openai.GPT4o,
		Messages: append(PROMPTS, openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: content}),
	})
	if err != nil {
		return nil, err
	}

	// Strip the markdown annotations from the response
	var jsonString string

	respContent := resp.Choices[0].Message.Content

	// Get the first { and the last } to get the JSON string
	start := strings.Index(respContent, "{")
	end := strings.LastIndex(respContent, "}")

	// Check if both '{' and '}' were found
	if start != -1 && end != -1 && start < end {
		// Extract the JSON string
		jsonString = respContent[start : end+1]
	} else {
		return nil, fmt.Errorf("Failed to extract the JSON string from the response")
	}

	fmt.Println(jsonString)

	// Convert the resp to a SuggestedRecipeDTO
	suggestedRecipe, err := dto.DerserializeSuggestedRecipeDTO(jsonString)
	if err != nil {
		return nil, fmt.Errorf("Failed to deserialize the response: %v", err)
	}

	return suggestedRecipe, nil
}
