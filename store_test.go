package markoph

import (
	"testing"

	"github.com/thor77/markoph/stores"
)

var (
	availableStores = []Store{
		stores.NewMap(make(map[[2]string]map[string]int)),
	}

	testKey  = [2]string{"a", "b"}
	testWord = "c"
)

func TestAdd(t *testing.T) {
	for _, store := range availableStores {
		store.Add(testKey, testWord)
	}
}

func TestGet(t *testing.T) {
	for _, store := range availableStores {
		for word, rate := range store.Get(testKey) {
			if word != testWord {
				t.Fatal(word, "!=", testWord)
			}
			if rate != 1 {
				t.Fatal(rate, "!=", 1)
			}
		}
	}
}
