package parser

import (
	"log"
	"regexp"
	"sync"
	"testing"

	"github.com/kheraprabhat/log-parser/types"
)

func TestParseLine(t *testing.T) {
	// Prepare a report object
	report := types.Report{IPs: types.MapStrInt{}, URLs: types.MapStrInt{}}

	var wg sync.WaitGroup
	// create a guard to control number of threads
	guard := make(chan bool, 1)

	reIP = regexp.MustCompile(RegxIP)
	reURL = regexp.MustCompile(RegxURL)

	wg.Add(1)
	guard <- true
	line := `168.41.191.40 - - [09/Jul/2018:10:10:38 +0200] "GET http://example.net/blog/category/meta/ HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/534.24 (KHTML, like Gecko) RockMelt/0.9.58.494 Chrome/11.0.696.71 Safari/534.24"`
	parseLine(line, 1, &report, &wg, &guard)

	if report.NumUniqueIPs() != 1 {
		t.Errorf("Number of unique IPs shuould be %d", 1)
	}

	wg.Add(1)
	guard <- true
	line = `168.41.191.41 - - [09/Jul/2018:10:10:38 +0200] "GET /blog HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/534.24 (KHTML, like Gecko) RockMelt/0.9.58.494 Chrome/11.0.696.71 Safari/534.24"`
	parseLine(line, 2, &report, &wg, &guard)

	if report.NumUniqueIPs() != 2 {
		t.Errorf("Number of unique IPs shuould be %d", 2)
	}

	ips := report.TopThreeIPs()
	if !(ips[0].Key == "168.41.191.41" || ips[0].Key == "168.41.191.40") {
		t.Errorf("IP should be one of %s or %s", "168.41.191.41", "168.41.191.40")
	}

	wg.Add(1)
	guard <- true
	line = `168.41.191.41 - - [09/Jul/2018:10:10:38 +0200] "GET /blog HTTP/1.1" 200 3574 "-" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/534.24 (KHTML, like Gecko) RockMelt/0.9.58.494 Chrome/11.0.696.71 Safari/534.24"`
	parseLine(line, 2, &report, &wg, &guard)

	if report.NumUniqueIPs() != 2 {
		t.Errorf("Number of unique IPs shuould be %d", 2)
	}

	ips = report.TopThreeIPs()
	urls := report.TopThreeURLs()
	if ips[0].Key != "168.41.191.41" {
		t.Errorf("First IP should be %s", "168.41.191.41")
	}

	if ips[0].Value != 2 {
		t.Errorf("First IP count should be %d", 2)
	}

	if urls[0].Key != "/blog" {
		t.Errorf("First URL should be %s", "/blog")
	}

	if urls[0].Value != 2 {
		t.Errorf("First URL count should be %d", 2)
	}

	// Should not parse the line
	wg.Add(1)
	guard <- true
	line = `"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_7) AppleWebKit/534.24 (KHTML, like Gecko) RockMelt/0.9.58.494 Chrome/11.0.696.71 Safari/534.24"`
	parseLine(line, 3, &report, &wg, &guard)

	if report.NumUniqueIPs() != 2 {
		t.Errorf("Number of unique IPs shuould be %d", 2)
	}
}

func BenchmarkParse(b *testing.B) {
	
    p, _ := NewParser("../log/access.log", 20)
    for i := 0; i < b.N; i++ {
        _, err := p.Parse()
        log.Println(err)
    }
}

