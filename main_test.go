package main

import (
	"os"
	"testing"
)

/*
 * Test helper functions
 */

func openFile(p string) []byte {
	data, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return data
}

func openFileAsString(p string) string {
	data := openFile(p)
	return string(data)
}

func mapIntIntExists(m map[int]int, i int) bool {
	if _, ok := m[i]; ok {
		return true
	}
	return false
}

func ngramMapKeyExists(m *NgramIndex, s string) bool {
	if _, ok := m.NgramMap[s]; ok {
		return true
	}
	return false
}

func ngramMapValueExists(m *NgramIndex, s string, i int) bool {
	if _, ok := m.NgramMap[s][i]; ok {
		return true
	}
	return false
}

/*
 * Tests + Benchmarks
 */

func TestNewNgramIndex(t *testing.T) {
	ni := NewNgramIndex()

	// Default ngam is '3'
	if ni.Ngram != 3 {
		t.Errorf("NewNgramIndex failed, Ngram value expect 3, got %d\n", ni.Ngram)
	}
}

func TestStringToNgram(t *testing.T) {
	// Test di-gram, 2 chars
	gram2 := StringToNgram("to", 2)
	gram2_many := StringToNgram("two chars", 2)
	if gram2[0] != "dG8=" { // to
		t.Errorf("2-gram failed, expect 'dG8='\n")
	} else if gram2_many[0] != "dHc=" { // tw
		t.Errorf("2-gram failed, expect 'dHc='\n")
	} else if gram2_many[2] != "byA=" { // 'o '
		t.Errorf("2-gram failed, expect 'byA='\n")
	}

	// Test tri-gram, 3 chars
	gram3 := StringToNgram("3ry", 3)
	gram3_many := StringToNgram("three chars", 3)
	if gram3[0] != "M3J5" { // 3ry
		t.Errorf("2-gram failed, expect 'M3J5'\n")
	} else if gram3_many[1] != "aHJl" { // hre
		t.Errorf("2-gram failed, expect 'aHJl'\n")
	} else if gram3_many[4] != "ZSBj" { // 'e c'
		t.Errorf("2-gram failed, expect 'ZSBj'\n")
	}

	// Test 4-gram, 4 chars
	gram4 := StringToNgram("four", 4)
	gram4_many := StringToNgram("four chars", 4)
	if gram4[0] != "Zm91cg==" { // four
		t.Errorf("2-gram failed, expect 'Zm91cg=='\n")
	} else if gram4_many[1] != "b3VyIA==" { // 'our '
		t.Errorf("2-gram failed, expect 'b3VyIA=='\n")
	} else if gram4_many[4] != "IGNoYQ==" { // ' cha'
		t.Errorf("2-gram failed, expect 'IGNoYQ=='\n")
	}
}

func BenchmarkStringToDigram(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToNgram("1234567890", 2)
	}
}

func BenchmarkStringToTrigram(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToNgram("1234567890", 3)
	}
}

func BenchmarkStringToTrigram_long(b *testing.B) {
	// Fetch 'main.go' file as a string
	file := openFileAsString("main.go")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Ngram of an entire file 'main.go'
		StringToNgram(file, 3)
	}
}

func TestAdd(t *testing.T) {
	// Create new index
	ni := NewNgramIndex()

	// Add a few items
	ni.Add("My first index item", 0)
	ni.Add("Second item", 1)
	ni.Add("Thired item too", 2)

	// Check the index got added
	if ni.IndexesMap[0] != struct{}{} || ni.IndexesMap[1] != struct{}{} || ni.IndexesMap[2] != struct{}{} {
		t.Errorf("IndexMap does not match items added'\n")
	}

	// Check if ngrams added correctly
	if !ngramMapKeyExists(ni, "TXkg") { // 'My '
		t.Errorf("NgramMap trigram does not match string added. Expected TXkg'\n")
	}

	if !ngramMapKeyExists(ni, "c3Qg") { // 'st '
		t.Errorf("NgramMap trigram does not match string added. Expected c3Qg'\n")
	}

	if !ngramMapKeyExists(ni, "aXJl") { // 'ire'
		t.Errorf("NgramMap trigram does not match string added. Expected aXJl'\n")
	}

	if !ngramMapKeyExists(ni, "dG9v") { // 'too'
		t.Errorf("NgramMap trigram does not match string added. Expected dG9v'\n")
	}

	// Check if n-gram index values match up
	if !ngramMapValueExists(ni, "TXkg", 0) { // 'My '
		t.Errorf("NgramMap index value not found. Expected 0'\n")
	}

	if !ngramMapValueExists(ni, "dG9v", 2) { // 'too'
		t.Errorf("NgramMap index value not found. Expected 2'\n")
	}

	if !(ngramMapValueExists(ni, "aXRl", 0) && ngramMapValueExists(ni, "aXRl", 1) && ngramMapValueExists(ni, "aXRl", 2)) { // 'ite'
		t.Errorf("NgramMap index value not found. Expected 0, 1 and 2'\n")
	}
}

func BenchmarkAdd(b *testing.B) {
	ni := NewNgramIndex()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ni.Add("1234567890", i)
	}
}

func BenchmarkAdd_long(b *testing.B) {
	ni := NewNgramIndex()

	// Fetch 'main.go' file as a string
	file := openFileAsString("main.go")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Ngram of an entire file 'main.go'
		ni.Add(file, i)
	}
}

func TestGetMatches(t *testing.T) {
	// Create new index
	ni := NewNgramIndex()

	// Add a few items
	ni.Add("My first index item", 0)
	ni.Add("Second item", 1)
	ni.Add("Thired item too", 2)
	ni.Add("Fourth item woowop", 3)

	// Get search results
	res1 := ni.GetMatches("first")
	res2 := ni.GetMatches("Second")
	res3 := ni.GetMatches("all items")
	res4 := ni.GetMatches("count first item")

	if !mapIntIntExists(res1, 0) || res1[0] != 3 { // 'first' matches 3 times when using a trigram
		t.Errorf("Match count for 'first' is wrong. Expected 3'\n")
	}

	if !mapIntIntExists(res2, 1) || res2[1] != 4 { // 'Second' matches 4 times when using a trigram
		t.Errorf("Match count for 'Second' is wrong. Expected 4'\n")
	}

	if !mapIntIntExists(res3, 0) || !mapIntIntExists(res3, 1) || !mapIntIntExists(res3, 2) || !mapIntIntExists(res3, 3) {
		t.Errorf("Match for 'all items' failed. Expected 0, 1, 2, 3\n")
	}

	if !mapIntIntExists(res4, 0) || res4[0] != 9 {
		t.Errorf("Match for 'count first item' failed. Expected 9 matches for first item\n")
	}
}
