package day5

type Mapper struct {
	counter int
	items   map[string]int
}

func (m *Mapper) GetItem(name string) int {
	val, ok := m.items[name]
	if ok {
		return val
	}

	m.counter++
	m.items[name] = m.counter

	return m.counter
}

func NewMapper() *Mapper {
	m := &Mapper{
		counter: -1,
		items:   map[string]int{},
	}

	return m
}
