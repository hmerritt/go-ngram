package main

import (
	"encoding/base64"
)

// Default ngram length
const ngramDefault = 3

type NgramIndex struct {
	// Index of ngrams
	// string = ngram
	// map[int] = index value
	NgramMap map[string]map[int]struct{}

	// Map of ALL index values in NgramMap
	IndexesMap map[int]struct{}

	// Length of n-grams to use
	// (recommended number is '3', ngram)
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
 * splits a string into groups of N length,
 * used for identifying indexed items and fast searching
 */
func StringToNgram(s string, ngram int) []string {
	ngrams := make([]string, 0)

	if len(s) < ngram {
		return ngrams
	}

	for i := 0; i < len(s)-(ngram-1); i++ {
		// Encode string to base 64,
		// this cleans up map keys with
		// special chars
		s64 := base64.StdEncoding.EncodeToString([]byte(s[i : i+ngram]))
		ngrams = append(ngrams, s64)
	}

	return ngrams
}

/*
 * Add a string and an index value to the store
 *
 * string will be indexed as an ngram
 * and the index value will be stored in each
 * ngram - this means the index value is accessible
 * through any part of the original string
 */
func (t *NgramIndex) Add(str string, index int) {
	// Add index to main map
	// index *should* always be unique
	t.IndexesMap[index] = struct{}{}

	// Get ngram slice from input string
	ngram := StringToNgram(str, t.Ngram)

	// Add each ngram into store
	for _, ng := range ngram {
		// Check if ng does NOT exist
		if _, exist := t.NgramMap[ng]; !exist {
			// Create ng
			newNg := make(map[int]struct{})
			t.NgramMap[ng] = newNg
		}

		// Add index value to ng
		t.NgramMap[ng][index] = struct{}{}
	}
}
