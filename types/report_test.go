package types

import (
	"testing"
)

func TestTopThreeIPs(t *testing.T) {

	cases := []struct {
		name    string
		r       Report
		size    int
		topPair Pair
	}{
		{
			name:    "One IP",
			r:       Report{IPs: MapStrInt{d: Document{"127.0.0.1": 2}}},
			size:    1,
			topPair: Pair{"127.0.0.1", 2},
		},

		{
			name:    "Four IP",
			r:       Report{IPs: MapStrInt{d: Document{"127.0.0.1": 2, "127.0.0.2": 7, "127.0.0.3": 9, "127.0.0.4": 5}}},
			size:    3,
			topPair: Pair{"127.0.0.3", 9},
		},

		{
			name:    "Three IP",
			r:       Report{IPs: MapStrInt{d: Document{"127.0.0.2": 7, "127.0.0.3": 9}}},
			size:    2,
			topPair: Pair{"127.0.0.3", 9},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			ips := tc.r.TopThreeIPs()
			if tc.size != len(ips) {
				t.Errorf("Want size %d but got %d", tc.size, len(ips))
			}

			if tc.topPair != ips[0] {
				t.Errorf("Want top pair %v but got %v", tc.topPair, ips[0])
			}
		})
	}
}

func TestTopThreeURLs(t *testing.T) {

	cases := []struct {
		name    string
		r       Report
		size    int
		topPair Pair
	}{
		{
			name:    "One URL",
			r:       Report{URLs: MapStrInt{d: Document{"/test": 2}}},
			size:    1,
			topPair: Pair{"/test", 2},
		},

		{
			name:    "Four URLs",
			r:       Report{URLs: MapStrInt{d: Document{"/home": 2, "/test": 7, "/user": 9, "/listings": 5}}},
			size:    3,
			topPair: Pair{"/user", 9},
		},

		{
			name:    "Three URLs",
			r:       Report{URLs: MapStrInt{d: Document{"/user": 9, "/listings": 5}}},
			size:    2,
			topPair: Pair{"/user", 9},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			urls := tc.r.TopThreeURLs()
			if tc.size != len(urls) {
				t.Errorf("Want size %d but got %d", tc.size, len(urls))
			}

			if tc.topPair != urls[0] {
				t.Errorf("Want top pair %v but got %v", tc.topPair, urls[0])
			}
		})
	}
}

func TestNumUniqueIPs(t *testing.T) {

	cases := []struct {
		name  string
		r     Report
		count int
	}{
		{
			name: "One IP",
			r:    Report{IPs: MapStrInt{d: Document{"127.0.0.1": 2}}},
			count: 1,
		},

		{
			name: "Four IP",
			r:    Report{IPs: MapStrInt{d: Document{"127.0.0.1": 2, "127.0.0.2": 7, "127.0.0.3": 9, "127.0.0.4": 5}}},
			count: 4,
		},

		{
			name:    "Three IP",
			r:       Report{IPs: MapStrInt{d: Document{"127.0.0.2": 7, "127.0.0.3": 9}}},
			count:    2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			count := tc.r.NumUniqueIPs()
			if tc.count != count {
				t.Errorf("Want count %d but got %d", tc.count, count)
			}
		})
	}
}
