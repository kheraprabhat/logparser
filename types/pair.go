package types

import (
	"fmt"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

// PairList implements sort and stringer interface
// It sorts Pair list in decending order with respect to value property
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}

func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairList) Less(i, j int) bool {
	return p[i].Value > p[j].Value
}

func (p PairList) String() string {
	str := make([]string, len(p))
	for i, k := range p {
		str[i] = k.Key + ":" + fmt.Sprintf("%d", k.Value)
	}
	return strings.Join(str, ",")
}
