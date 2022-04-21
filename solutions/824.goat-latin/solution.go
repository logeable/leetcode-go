package solution

import (
	"strings"
)

func toGoatLatin(sentence string) string {
	words := strings.Fields(sentence)

	for i, word := range words {
		fc := word[0] | 32

		if fc == 'a' || fc == 'e' || fc == 'i' || fc == 'o' || fc == 'u' {
			word = word + "ma"
		} else {
			word = word[1:] + string(word[0]) + "ma"
		}
		for j := 0; j < i+1; j++ {
			word = word + "a"
		}
		words[i] = word
	}
	return strings.Join(words, " ")
}
