package types

import (
	"testing"
)

func TestLen(t *testing.T) {

	cases := []struct {
		name string
		d    PairList
		size int
	}{
		{
			name: "size 2",
			d:    PairList{{"john", 1}, {"smith", 3}},
			size: 2,
		},

		{
			name: "size 0",
			d:    PairList{},
			size: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.size != tc.d.Len() {
				t.Errorf("Want number of pairs %d but got %d", tc.size, tc.d.Len())
			}
		})
	}
}

func TestSwap(t *testing.T) {

	d := PairList{{"john", 1}, {"smith", 3}, {"robert", 8}}

	d.Swap(1, 1)
	if d[1].Key != "smith" || d[1].Value != 3 {
		t.Errorf("Swap pairs didn't work for pair %d and %d", 1, 1)
	}

	d.Swap(0, 1)
	if d[0].Key != "smith" || d[0].Value != 3 {
		t.Errorf("Swap pairs didn't work for pair %d and %d", 0, 1)
	}

	d.Swap(1, 2)
	if d[1].Key != "robert" || d[1].Value != 8 {
		t.Errorf("Swap pairs didn't work for pair %d and %d", 1, 2)
	}
}

func TestLess(t *testing.T) {
	d := PairList{{"john", 1}, {"smith", 3}, {"robert", 8}}

	if d.Less(1, 2) {
		t.Errorf("Wanted %t got %t", true, false)
	}

    if !d.Less(2, 1) {
		t.Errorf("Wanted %t got %t", true, false)
	}
}
