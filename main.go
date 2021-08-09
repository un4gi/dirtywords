package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/un4gi/dirtywords/sources"
	"github.com/un4gi/dirtywords/words"
)

func main() {

	domain := flag.String("d", "", "domain to build a wordlist for")
	minLen := flag.Int("minLen", 3, "minimum word length - defaults to 3")
	maxLen := flag.Int("maxLen", 12, "maximum word length - defaults to 12")
	noSubs := flag.Bool("nosubs", false, "build wordlist without gathering data from subdomains")
	outFile := flag.String("o", "", "filename to generate as wordlist")
	sort := flag.Bool("s", true, "uniquely sort the results")
	flag.Parse()

	if *domain == "" {
		fmt.Println("A domain is needed! Exiting")
		os.Exit(1)
	}

	var rootDomain string
	if *noSubs {
		rootDomain = *domain
	} else {
		rootDomain = "*." + *domain
	}

	var filename string
	if *outFile == "" {
		filename = *domain + ".txt"
	} else {
		filename = *outFile
	}

	fmt.Println("Gathering data...")
	sources.CommonCrawl(rootDomain, filename, *minLen, *maxLen) // Common Crawl
	sources.WaybackURLs(rootDomain, filename, *minLen, *maxLen) // Wayback Machine
	sources.OTX(rootDomain, filename, *minLen, *maxLen)         // Open Threat Exchange

	//os := runtime.GOOS
	if *sort {
		fmt.Println("Sorting...")
		words.SortWordList(filename)
	}
}
