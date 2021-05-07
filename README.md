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

| Methods         | Parameters  | Description                                                                                 |
| ------------- | ---------- | ------------------------------------------------------------------------------- |
| `NewNgramIndex` | -           | Returns a new trigram index with the default settings                                       |
| `Add`           | string, int | Add a string AND your own index value                                                       |
| `Search`        | string      | Returns slice of all matched strings which contains your index value and the match strength, `[][index, weight]` | 
|                 |             |                                                                                             | 
| Helper functions |            |                                                                                             | 
|                 |             |                                                                                             | 
| `StringToNgram` | string, int | Returns an ngram of length N - Splits a string into groups of N length                      | 
| `GetMatches`    | string      | Get all ngram matches, returns an unsorted map of indexes along with their match weight     | 
| `SortMatches`   | `map[int]int` | Sorts output from `GetMatches` into a slice, first index [0] = best match                 | 

### General
```go
// Create a new index
// Default ngram size is '3' (tri-gram)
ni := ngram.NewNgramIndex()

// Add string + add your OWN index value (this index value is what will be returned when a search matches)
ni.Add(string, int)
ni.Add("I just think they're neat.", 0)
ni.Add("Don’t eat me. I have a wife and kids. Eat them.", 1)
ni.Add("Words to match above quotes... think eat neat", 2)

// Usually, you would loop your own data-set 
// and add the index from that as the ngram index value
for index, value := myFiles {
	ni.Add(value, index)
}

// Search
// Returns a sorted array of matches (sorts based on number of matches)
result := ni.Search("Eat them")
// [0, 1, 2]

result := ni.Search("near") // near will match [nea]t
// [0, 2]
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
	Name    string
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
		Content: "This data could be anything, a file, a string",
	})
	
	/*
	 * Create N-gram here!
	 */

	// Create a new index
	// Default ngram size is '3' (tri-gram)
	ni := ngram.NewNgramIndex()
	
	// Add each File in our example data-set
	for index, file := range files {

		// Add a string + the index of a struct
		ni.Add(file.Content, index)
		
		// What this is actually doing
		ni.Add("This data could be anything, a file, a string", 1)

	}

	// Search indexed data
	// returns a sorted slice
	ret := ti.Search("anything")
	fmt.Println("Search ret=", ret)
	// [1]
}
```
