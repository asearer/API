package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestQuoteFields validates the fields of the Quote struct
func TestQuoteFields(t *testing.T) {
	q := Quote{ID: 1, Text: "Test", Author: "Tester"}
	assert.Equal(t, 1, q.ID)
	assert.Equal(t, "Test", q.Text)
	assert.Equal(t, "Tester", q.Author)
}

// TestQuoteStoreInitialization checks that QuoteStore has initial quotes
func TestQuoteStoreInitialization(t *testing.T) {
	assert.GreaterOrEqual(t, len(QuoteStore), 2, "QuoteStore should have at least 2 initial quotes")
}

// TestQuoteStoreAppend ensures we can append a new quote without errors
func TestQuoteStoreAppend(t *testing.T) {
	initialLen := len(QuoteStore)
	newQuote := Quote{ID: initialLen + 1, Text: "Another one", Author: "Author"}
	QuoteStore = append(QuoteStore, newQuote)

	last := QuoteStore[len(QuoteStore)-1]
	assert.Equal(t, "Another one", last.Text)
	assert.Equal(t, "Author", last.Author)
	assert.Equal(t, initialLen+1, last.ID)
}
