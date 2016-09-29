package markoph

// Store stores key-value pairs for the markov-chain
type Store interface {
	Add([2]string, string)
	Get([2]string) map[string]int
	Keys() [][2]string // BUG(thor77) probably use RandomKey or similar to allow optimizations on Store-side
}
