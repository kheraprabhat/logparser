package parser

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/kheraprabhat/log-parser/types"
)

// const for IP and URL RegEx
const RegxIP = `^([0-9]{0,3}\.){3}[0-9]{0,3}`
const RegxURL = `[GET|PUT|POST] (.*) HTTP`

var reIP *regexp.Regexp
var reURL *regexp.Regexp

// Parser holds filePath and maxThreads to control number of threads per execution
type Parser struct {
	filePath   string
	maxThreads int
}

// NewParser receives filePath and returns Parser
func NewParser(filePath string, maxThreads int) (*Parser, error) {
	if filePath == "" {
		return nil, errors.New("file path cannot be blank")
	}
	if maxThreads == 0 {
		maxThreads = 1
	}
	return &Parser{filePath: filePath, maxThreads: maxThreads}, nil
}

func parseLine(line string, lineNum int, report *types.Report, wg *sync.WaitGroup, guard *chan (bool)) {

	defer wg.Done()

	ip := reIP.FindString(line)
	url := reURL.FindSubmatch([]byte(line))

	if ip != "" {
		report.IPs.Write(ip)
	} else {
		log.Println(fmt.Sprintf("[WARN] No IP found at line %d", lineNum))
	}

	if len(url) > 1 {
		report.URLs.Write(string(url[1]))
	} else {
		log.Println(fmt.Sprintf("[WARN] No URL found at line %d", lineNum))
	}

	<-*guard
}

// Parse method parse reads and parse the file and returns Reports
func (p *Parser) Parse() (*types.Report, error) {

	// Open file to read
	file, err := os.Open(p.filePath)
	if err != nil {
		return nil, err
	}
	// close the file when function returns
	defer file.Close()

	log.Println("[INFO] file openned")

	// create a scanner to read from the file
	scanner := bufio.NewScanner(file)

	// Prepare a report object
	report := types.Report{IPs: types.MapStrInt{}, URLs: types.MapStrInt{}}

	var wg sync.WaitGroup
	// create a guard to control number of threads
	guard := make(chan bool, p.maxThreads)

	reIP = regexp.MustCompile(RegxIP)
	reURL = regexp.MustCompile(RegxURL)

	lines := 0
	// read the file line by line
	for scanner.Scan() {

		line := scanner.Text()

		wg.Add(1)
		guard <- true
		// do the parsing in go routine
		go parseLine(line, lines, &report, &wg, &guard)

		lines++
	}

	// Wait to finish all the go routines
	wg.Wait()

	log.Println(fmt.Sprintf("[INFO] Total lines in the log file %d", lines))
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &report, nil
}

