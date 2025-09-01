package models

// Quote represents a single inspirational quote.
type Quote struct {
	ID     int    `json:"id"`     // Unique identifier
	Text   string `json:"text"`   // Quote text
	Author string `json:"author"` // Person who said it
}

// QuoteStore holds all quotes in memory
var QuoteStore = []Quote{
	{ID: 1, Text: "The best way to get started is to quit talking and begin doing.", Author: "Walt Disney"},
	{ID: 2, Text: "Don’t let yesterday take up too much of today.", Author: "Will Rogers"},
}
