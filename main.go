package main

// Default ngram length
const ngram = 3

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
	t.Ngram = ngram
	return t
}
