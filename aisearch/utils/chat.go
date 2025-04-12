package utils

// Package utils provides utility functions for the AI search system.
// This file contains OpenAI chat message related utilities.

import "github.com/sashabaranov/go-openai"

// ChatMsgFromUser creates a user chat message for OpenAI API.
//
// Parameters:
//   txt: The message text content
//
// Returns:
//   A ChatCompletionMessage with user role
//
// Example:
//   msg := ChatMsgFromUser("What is AI?")
func ChatMsgFromUser(txt string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: txt,
	}
}

// ChatMsgFromSystem creates a system chat message for OpenAI API.
//
// Parameters:
//   txt: The system instruction/message
//
// Returns:
//   A ChatCompletionMessage with system role
//
// Example:
//   msg := ChatMsgFromSystem("You are a helpful AI assistant")
func ChatMsgFromSystem(txt string) openai.ChatCompletionMessage {
	return openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: txt,
	}
}

// ModelGPT35Request creates a default GPT-3.5 chat completion request.
//
// Parameters:
//   msg: Conversation history/messages
//
// Returns:
//   A ChatCompletionRequest with default settings:
//   - Temperature: 0.1 (low randomness)
//   - N: 1 (single response)
//   - Model: GPT3Dot5Turbo0125
//
// Example:
//   msgs := []openai.ChatCompletionMessage{
//       ChatMsgFromSystem("You are helpful"),
//       ChatMsgFromUser("Hello"),
//   }
//   req := ModelGPT35Request(msgs)
func ModelGPT35Request(msg []openai.ChatCompletionMessage) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Temperature: 0.1,
		N:           1,
		Model:       openai.GPT3Dot5Turbo0125,
		Messages:    msg,
	}
}
