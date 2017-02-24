package counter

type Counter interface {
	// Count a fingerprint and return the support for the item.
	// support = count / total
	Count(Countable) float64
	Stats() Stats
}

type Countable interface {
	IsMatch(other Countable) bool
}

type Stats struct {
	UniqueFingerprints int
	Distribution       []int
}
