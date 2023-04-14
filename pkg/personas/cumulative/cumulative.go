// Copyright 2023 dvonthenen ChatGPT Proxy contributors. All Rights Reserved.
// SPDX-License-Identifier: Apache License 2.0

package cumulative

import (
	"context"

	openai "github.com/sashabaranov/go-openai"

	advanced "github.com/dvonthenen/chat-gpeasy/pkg/personas/advanced"
	interfaces "github.com/dvonthenen/chat-gpeasy/pkg/personas/interfaces"
)

func New(client *openai.Client) (*Persona, error) {
	if client == nil {
		return nil, interfaces.ErrInvalidInput
	}
	advanced, err := advanced.New(client)
	if err != nil {
		return nil, err
	}
	return &Persona{
		persona: advanced,
	}, nil
}

func (c *Persona) Init(level interfaces.SkillType, model string) error {
	return c.persona.Init(interfaces.SkillType(level), model)
}

func (c *Persona) Query(ctx context.Context, statement string) ([]openai.ChatCompletionChoice, error) {
	return c.persona.Query(ctx, openai.ChatMessageRoleUser, statement)
}

func (c *Persona) AddDirective(directives string) error {
	return c.persona.AddDirective(directives)
}

func (c *Persona) CommitResponse(index int) error {
	return c.persona.CommitResponse(index)
}
