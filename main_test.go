package main

import (
	"testing"
)

func TestNewNgramIndex(t *testing.T) {
	ni := NewNgramIndex()

	// Default ngam is '3'
	if ni.Ngram != 3 {
		t.Errorf("NewNgramIndex failed, Ngram value expect 3, got %d\n", ni.Ngram)
	}
}
