package types

import (
	"sort"
)

// Report keeps the IPs and URLs
type Report struct {
	IPs  MapStrInt
	URLs MapStrInt
}

// TopThreeIPs returns top 3 used IPs
func (r *Report) TopThreeIPs() PairList {
	ips := r.IPs.d.GetPairList()
	sort.Sort(ips)

	if len(ips) > 3 {
		return ips[:3]
	}
	return ips
}

// TopThreeIPs returns top 3 used URLs
func (r *Report) TopThreeURLs() PairList {
	urls := r.URLs.d.GetPairList()
	sort.Sort(urls)
	//log.Printf("%v", urls)

	if len(urls) > 3 {
		return urls[:3]
	}
	return urls
}

// TopThreeIPs returns number of unique IPs
func (r *Report) NumUniqueIPs() int {
	return len(r.IPs.d)
}
