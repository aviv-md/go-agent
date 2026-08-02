package model

import (
	"testing"

	"github.com/openai/openai-go/v3/responses"
)

func TestConvertRoleToOpenAIRole(t *testing.T) {
	// Let's test all roles.
	ra, err := convertRoleToOpenAIRole(RoleAssistant)
	if err != nil {
		t.Errorf("Error converting role to OpenAI role: %v", err)
	}
	rs, err := convertRoleToOpenAIRole(RoleSystem)
	if err != nil {
		t.Errorf("Error converting role to OpenAI role: %v", err)
	}
	ru, err := convertRoleToOpenAIRole(RoleUser)
	if err != nil {
		t.Errorf("Error converting role to OpenAI role: %v", err)
	}

	if ra != responses.EasyInputMessageRoleAssistant {
		t.Errorf("Expected assistant, got %s", ra)
	}
	if rs != responses.EasyInputMessageRoleSystem {
		t.Errorf("Expected system, got %s", rs)
	}
	if ru != responses.EasyInputMessageRoleUser {
		t.Errorf("Expected user, got %s", ru)
	}

	rt, err := convertRoleToOpenAIRole(RoleTool)
	if rt != "" {
		t.Errorf("Expected an empty string, got %s", rt)
	}

	if err == nil {
		t.Errorf("Expected an error, got nil")
	}

	rj, err := convertRoleToOpenAIRole(Role("invalid"))
	if rj != "" {
		t.Errorf("Expected an empty string, got %s", rj)
	}

	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}

func TestConvertMessageToOpenAIInput(t *testing.T) {

	message := NewUserMessage(
		"psst kid wanna write some tests?",
	)

	openAIInput, err := convertMessageToOpenAIInput(message)
	if err != nil {
		t.Fatalf("Error converting message to OpenAI input: %v", err)
	}

	for _, conv := range openAIInput {
		r := conv.GetRole()
		if r == nil {
			t.Fatalf("Expected a role, got nil")
		}
		if *r != string(responses.EasyInputMessageRoleUser) {
			t.Errorf("Expected user, got %s", *r)
		}

		c := conv.OfMessage

		if c == nil {
			t.Fatalf("Expected a message, got nil")
		}
		if !c.Content.OfString.Valid() {
			t.Fatalf("Expected c.Content.OfString.Valid() to be true, instead got false.")
		}

		if c.Content.OfString.Value != "psst kid wanna write some tests?" {
			t.Fatalf("Expected 'psst kid wanna write some tests?', got %s", c.Content.OfString.String())
		}
	}
}
