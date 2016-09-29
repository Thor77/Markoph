package markoph

import "strings"

// Learn feeds text into the markov-chain
func (m Markoph) Learn(text string) {
	words := strings.Split(text, " ")
	for i, word := range words {
		if len(words) <= i+1 {
			break
		}
		nextWord := words[i+1]
		var nextNextWord string
		if len(words) <= i+2 {
			nextNextWord = end
		} else {
			nextNextWord = words[i+2]
		}
		key := [2]string{word, nextWord}
		m.store.Add(key, nextNextWord)
	}
}
