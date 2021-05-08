# n-gram Index and Search
[![Go Reference](https://pkg.go.dev/badge/github.com/hmerritt/go-ngram.svg)](https://pkg.go.dev/github.com/hmerritt/go-ngram)  [![Go](https://github.com/hmerritt/go-ngram/actions/workflows/go.yml/badge.svg)](https://github.com/hmerritt/go-ngram/actions/workflows/go.yml)  [![Coverage Status](https://coveralls.io/repos/github/hmerritt/go-ngram/badge.svg?branch=master)](https://coveralls.io/github/hmerritt/go-ngram?branch=master&kill_cache=1)

Create n-grams, index items, and search through them.


## Install
```bash
go get github.com/hmerritt/go-ngram
```


## What is an N-gram reverse-index?
I highly recommend reading the following:
- [Russ Cox - Google Code Search, Trigram Index](https://swtch.com/~rsc/regexp/regexp4.html)

The `N` in n-gram refers to a set number. The most common n-gram is the 3-gram, usually called a "tri-gram".

A tri-gram index stores each item in sets of 3.

```
Hello, World  ->  [Hel ell llo lo, o,_ _Wo Wor orl rld]

Search("rld") -> Hello, World
        ^^^ any 3 (or more) characters will find a match
```

#### Why use N-gram indexing?
- Used to search through large data sets quickly.
- Has built-in spelling correction (only 3 chars need to match)


## Usage

| Methods          | Description                                                                                                      |
| ---------------- | ---------------------------------------------------------------------------------------------------------------- |
| `NewNgramIndex`  | Returns a new trigram index with the default settings                                                            |
| `NewIndexValue`  | Returns a new index value using custom input values (use this for `Add()`)                                       |
| `Add`            | Add a string AND your own index value                                                                            |
| `Search`         | Returns slice of all matched strings which contains your index value and the match strength, `[]*IndexValue`     |
|                  |                                                                                                                  |
| Helper functions |                                                                                                                  |
| `StringToNgram`  | Returns an ngram of length N - Splits a string into groups of N length                                           |
| `GetMatches`     | Get all ngram matches, returns an unsorted map of indexes along with their Index Value struct                    |
| `SortMatches`    | Sorts output from `GetMatches` into a slice, first index [0] = best match                                        |

### Example
```go
// Create a new index
// Default ngram size is '3' (tri-gram)
ni := ngram.NewNgramIndex()

// Add string + add your OWN index value
// (the index value is what will be returned when a search matches)
ni.Add(string, *IndexValue)
ni.Add("I just think they're neat.", NewIndexValue(0, "This additional index string can be anything"))
ni.Add("Don’t eat me. I have a wife and kids. Eat them", NewIndexValue(1, "It could be a filepath"))
ni.Add("Words to match above quotes... string eat neat", NewIndexValue(2, "Or maybe a URL"))

// Search
// Returns a sorted slice of matches (sorts based on number of matches)
search := ni.Search("Eat")
// -> [1, 2]
// -> [IndexValue{1, "It could be a filepath"}, IndexValue{2, "Or maybe a URL"}]

search = ni.Search("near") // near will match [nea]t
// -> [0, 2]

// Usually, you would loop your own data-set 
// and add the index from that as the ngram index value
for index, file := myFiles {
	ni.Add(file.content, NewIndexValue(index, file.path)
}
```

### Long-Winded Example
```go
package main

import (
	"fmt"
	ngram "github.com/hmerritt/go-ngram"
)

// Example structure
type File struct {
	Path    string
	Content string
}

// Example data-set
var files = make([]File, 0)

func main() {
	// Add Files to example data-set
	files = append(files, File{
		Path:    "/home/me/nice.txt",
		Content: "Example file data, nice",
	})
	files = append(files, File{
		Path:    "/home/me/example-file.txt",
		Content: "More example file data, neat",
	})
	
	/*
	 * Create N-gram here!
	 */

	// Create a new index
	// Default ngram size is '3' (tri-gram)
	ni := ngram.NewNgramIndex()
	
	// Add each File in our example data-set
	for index, file := range files {

		// Add a string + the index value we want to store in assosiation with the input string
		ni.Add(file.Content, NewIndexValue(index, file.Path))
		
		// What this is actually doing
		ni.Add("More example file data, neat", NewIndexValue(1, "/home/me/example-file.txt"))

	}

	// Search indexed data
	// returns a sorted slice
	search := ti.Search("neat")
	fmt.Println("Search =", search)
	// -> [1]
	// -> [IndexValue{
	//      Index: 1,
	//      Data: "/home/me/example-file.txt"
	//    }]
}
```
