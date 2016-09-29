package stores

import "fmt"

// Map store for markoph-chain, stored in memory
// not usable for large collections or where persistence is required
type Map struct {
	internal map[[2]string]map[string]int
}

// NewMap initializes a new Map
func NewMap(internalStore map[[2]string]map[string]int) Map {
	return Map{internalStore}
}

// Add adds a key and value to the store
func (m Map) Add(key [2]string, value string) {
	if _, ok := m.internal[key]; ok {
		m.internal[key][value]++
	} else {
		m.internal[key] = map[string]int{
			value: 1,
		}
	}
}

// Get returns the values for key
func (m Map) Get(key [2]string) map[string]int {
	return m.internal[key]
}

// Keys returns a list of all available keys
func (m Map) Keys() [][2]string {
	// gen list
	keys := make([][2]string, len(m.internal))
	i := 0
	for k := range m.internal {
		keys[i] = k
		i++
	}
	return keys
}

// Print prints the internal storage
func (m Map) Print() {
	fmt.Println(m.internal)
}
