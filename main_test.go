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

func BenchmarkStringToNgram(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StringToNgram("1234567890", 3)
	}
}
