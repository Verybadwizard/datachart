package chart

type chart struct {
	counter int32
	data    map[int32]string
}

func create() *chart {
	return &chart{
		counter: 0,
		data:    make(map[int32]string),
	}
}

func (m *chart) Add(s string) int32 {
	m.counter++
	m.data[m.counter] = s
	return m.counter
}

func (m *chart) Get(i int32) (string, bool) {
	s, exists := m.data[i]
	return s, exists
}
