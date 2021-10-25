package types

import (
	"sort"
	"testing"
)

func TestGetPairList(t *testing.T) {

	cases := []struct {
		name         string
		d            Document
		wantUnsorted string
		wantSorted   string
		size         int
	}{
		{
			name:         "matched document",
			d:            Document{"john": 1, "smith": 3},
			wantUnsorted: "john:1,smith:3",
			wantSorted:   "smith:3,john:1",
			size:         2,
		},

		{
			name:         "blank document",
			d:            Document{},
			wantUnsorted: "",
			wantSorted:   "",
			size:         0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			res := tc.d.GetPairList()

			if tc.size != len(res) {
				t.Errorf("Want number of pairs %d but got %d", tc.size, len(res))
			}

			if tc.wantUnsorted != res.String() {
				t.Errorf("Want %s but got %s", tc.wantUnsorted, res.String())
			}

			sort.Sort(res)
			if tc.wantSorted != res.String() {
				t.Errorf("Want %s but got %s", tc.wantSorted, res.String())
			}
		})
	}
}
