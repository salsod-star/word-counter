package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\n word3 \nword4")

	exp, res := 4, count(b, false)

	if exp != res {
		t.Errorf("Expected %d, Got %d", exp, res)
	}
}