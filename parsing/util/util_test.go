package util

import "testing"

func BenchmarkFirstWord(b *testing.B) {
	text := `this is a sentence`
	for n := 0; n < b.N; n++ {
		FirstWord(text)
	}
}
