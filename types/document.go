package types

import "sync"

type Document map[string]int

func (d Document) GetPairList() PairList {
	pairs := make(PairList, len(d))
	i := 0
	for k, v := range d {
		pairs[i] = Pair{k, v}
		i++
	}
	return pairs
}

type MapStrInt struct {
	d Document
	sync.RWMutex
}

func (m *MapStrInt) Write(key string) {
	m.Lock()
	if m.d == nil {
		m.d = Document{}
	}
	m.d[key] = m.d[key] + 1
	m.Unlock()
}
