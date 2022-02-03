package solution

import "strings"

func uniqueMorseRepresentations(words []string) int {
	morse := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	morseDict := make(map[rune]string)
	for i := 'a'; i <= 'z'; i++ {
		morseDict[i] = morse[int(i-'a')]
	}

	counter := make(map[string]int)
	var m strings.Builder
	for _, w := range words {
		for _, b := range w {
			m.WriteString(morseDict[b])
		}
		counter[m.String()]++
		m.Reset()
	}
	return len(counter)
}
