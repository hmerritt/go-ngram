package ngram

import (
	"encoding/base64"
	"sort"
)

// Trigram
// Used for NewNgramIndex()
const DefaultNgramLength = 3

// Index value, the value which an ngram points to.
//
// Gets returned when using Search() or GetMatches().
//
// Do NOT modify 'Matches' -> this value is set automatically
// when searching and will only mess things up!
type IndexValue struct {
	Index   int
	Matches int
	Text    string

	// Debating to include the following types,
	// as far as i'm aware, there is no downsides.
	// SliceInt        []int
	// SliceString     []string
	// MapStringInt    map[string]int
	// MapStringString map[string]string
}

// Ngram index, uses a reverse-index map (NgramMap)
// to store and search through items.
//
// Ngram length is decided here.
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

// Returns a new ngram index using default values.
//
// Use NgramIndex{} for custom ngram lengths
func NewNgramIndex() *NgramIndex {
	t := new(NgramIndex)
	t.NgramMap = make(map[string]map[int]struct{})
	t.IndexesMap = make(map[int]struct{})
	t.Ngram = ngramDefault
	return t
}

// Ngram slice
// splits a string into groups of N length,
// used for adding items to the index and fast searching
func StringToNgram(s string, ngram int) []string {
	if len(s) < ngram {
		return []string{}
	}

	ngrams := make([]string, 0, len(s))

	for i := 0; i < len(s)-(ngram-1); i++ {
		// Encode string to base 64,
		// this cleans up map keys with
		// special chars
		s64 := base64.StdEncoding.EncodeToString([]byte(s[i : i+ngram]))
		ngrams = append(ngrams, s64)
	}

	return ngrams
}

// Add a string and an index value to the store
//
// string will be indexed as an ngram
// and the index value will be stored in each
// ngram - this means the index value is accessible
// through any part of the original string
func (n *NgramIndex) Add(str string, index int) {
	// Add index to main map
	// index *should* always be unique
	n.IndexesMap[index] = struct{}{}

	// Get ngram slice from input string
	ngram := StringToNgram(str, n.Ngram)

	// Add each ngram into store
	for _, ng := range ngram {
		// Check if ng does NOT exist
		if _, exist := n.NgramMap[ng]; !exist {
			// Create ng
			newNg := make(map[int]struct{})
			n.NgramMap[ng] = newNg
		}

		// Add index value to ng
		n.NgramMap[ng][index] = struct{}{}
	}
}

// Search for all matches AND sorts
// the matches into 'best match'
//
// Alias of GetMatches + SortMatches
func (n *NgramIndex) Search(str string) [][]int {
	match := n.GetMatches(str)
	return n.SortMatches(match)
}

// Search the NgramMap and return
// an array of all the stored index values
// that matched the input string
func (n *NgramIndex) GetMatches(str string) map[int]int {
	// Create map of indexes plus how often
	// each one matched. This is used to detirmine
	// an indexes 'weight'.
	matches := make(map[int]int)

	// Get ngram slice
	ngram := StringToNgram(str, n.Ngram)

	// Loop each ngram looking for matches
	for _, tg := range ngram {
		// Check if tg exists
		if _, exist := n.NgramMap[tg]; exist {
			// MATCH!

			// Add all indexes to matched
			for index := range n.NgramMap[tg] {
				// Has index been added already,
				// increment match weight
				if _, exist := matches[index]; exist {
					matches[index] = matches[index] + 1
				} else {
					matches[index] = 1
				}
			}
		}
	}

	return matches
}

// Sort matched items from GetMatches()
// into a slice with 'best match' first
// decending into 'weakest' match last
//
// best match = most matches
func (n *NgramIndex) SortMatches(matches map[int]int) [][]int {
	// Create slice of index + weight
	sortMatches := make([][]int, 0, len(matches))
	for k := range matches {
		sortMatches = append(sortMatches, []int{k, matches[k]})
	}

	// Sort slice
	sort.Slice(sortMatches, func(i, j int) bool {
		return sortMatches[i][1] > sortMatches[j][1]
	})

	return sortMatches
}
