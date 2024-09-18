package metrics

type StubMetrics struct {
}

func NewStubMetrics() *StubMetrics {
	return &StubMetrics{}
}

func (s *StubMetrics) Increment(name string) {}
