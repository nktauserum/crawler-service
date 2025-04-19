package crawler

import (
	"context"
	"testing"
	"time"
)

func TestMakeReadable(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	raw_content, err := GetContent(ctx, "https://rdi.berkeley.edu/llm-agents/assets/llm-reasoning.pdf")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(raw_content.Content)
}
