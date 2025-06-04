package crawler

import (
	"context"
	"testing"
)

func TestPDFExtracting(t *testing.T) {
	page, err := ParsePDF(context.Background(), "https://rdi.berkeley.edu/llm-agents/assets/llm-reasoning.pdf")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(page)
}
