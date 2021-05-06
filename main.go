package main

import "encoding/base64"

// Default ngram length
const ngramDefault = 3

type NgramIndex struct {
	// Index of trigrams
	// string = trigram
	// map[int] = index value
	NgramMap map[string]map[int]struct{}

	// Map of ALL index values in NgramMap
	IndexesMap map[int]struct{}

	// Length of n-grams to use
	// (recommended number is '3', trigram)
	Ngram int
}

/*
 * Returns a new ngram index
 * using default values.
 * Use NgramIndex{} for custom ngram lengths
 */
func NewNgramIndex() *NgramIndex {
	t := new(NgramIndex)
	t.NgramMap = make(map[string]map[int]struct{})
	t.IndexesMap = make(map[int]struct{})
	t.Ngram = ngramDefault
	return t
}

/*
 * Ngram slice
 * splits a string into groups of 3,
 * used for indexing and fast searching
 */
func StringToNgram(s string, ngram int) []string {
	trigrams := make([]string, 0)

	if len(s) == 0 {
		return trigrams
	}

	for i := 0; i < len(s)-(ngram-1); i++ {
		// Encode string to base 64,
		// this cleans up map keys with
		// special chars
		s64 := base64.StdEncoding.EncodeToString([]byte(s[i : i+ngram]))
		trigrams = append(trigrams, s64)
	}

	return trigrams
}
