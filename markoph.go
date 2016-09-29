package markoph

const end = "STOP"

// Markoph describes a Markov-Chain-Instance
type Markoph struct {
	store Store
}

// NewMarkoph initializes a Markov-Chain-Instance
func NewMarkoph(store Store) Markoph {
	return Markoph{store}
}
