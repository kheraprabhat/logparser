package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kheraprabhat/log-parser/parser"
)

func main() {

    args := os.Args
    if len(args) != 2 {
        log.Fatal("Usage `go run main.go [log_file]`")
    }
	
    // Initialse the parser by passing the file path and max number of go routines
    parser, err := parser.NewParser(args[1], 30)
    if err != nil {
        log.Fatalf("While creating parser: %s", err.Error())
    }

    // Parse the file
    report, err := parser.Parse()
    if err != nil {
        log.Fatalf("While parsing the log file: %s", err.Error())
    }

    // Print the reports
    log.Println("^^^^^^^^^^^^^^^^ REPORTS ^^^^^^^^^^^^^^^^^^^^^")
    log.Println(fmt.Sprintf("Top 3 IPs: %s", report.TopThreeIPs()))
    log.Println(fmt.Sprintf("Top 3 URLs: %s", report.TopThreeURLs()))
    log.Println("Number of Unique IPs: ", report.NumUniqueIPs())
    log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
}
