package markoph

import (
	"math/rand"
	"strings"
)

// matchTo chooses a word, matching to key
func (m Markoph) matchTo(key [2]string) string {
	possibleWords := m.store.Get(key)
	if len(possibleWords) <= 1 {
		for word := range possibleWords {
			return word
		}
	}
	// build a list with all possibleWords x rate
	var weightedWords []string
	for word, rate := range possibleWords {
		for i := 0; i < rate; i++ {
			weightedWords = append(weightedWords, word)
		}
	}
	if len(weightedWords) == 1 {
		return weightedWords[0]
	}
	// choose a random word from the list
	return weightedWords[rand.Intn(len(weightedWords)-1)]
}

// Generate generates a random answer based on a start-key
func (m Markoph) Generate(start [2]string, length int) string {
	sentence := make([]string, length)
	// choose a random key to start with
	sentence[0] = start[0]
	sentence[1] = start[1]
	// start actual generation
	for i := 2; i <= len(sentence)-1; i++ {
		sentence[i] = m.matchTo([2]string{sentence[i-2], sentence[i-1]})
		if sentence[i] == end {
			sentence[i] = ""
			break
		}
	}
	return strings.Join(sentence, " ")
}

// GenerateRand generates a random answer based on the existing dictionary
func (m Markoph) GenerateRand(length int) string {
	keys := m.store.Keys()
	// choose random key
	randKey := keys[rand.Intn(len(keys)-1)]
	return m.Generate(randKey, length)
}
