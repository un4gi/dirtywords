package sources

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/un4gi/dirtywords/config"
	"github.com/un4gi/dirtywords/requests"
	"github.com/un4gi/dirtywords/words"
)

// Big shoutout to @tomnomnom and @lc for the inspiration!

// WaybackURLs pulls words from URLs found in the Internet Archive's Wayback Machine
func WaybackURLs(domain string, filename string, minLen int, maxLen int) error {
	pages := getWaybackPagination(domain) // Get total number of pages for given domain

	// Makes n requests based on pagination results to grab archived URLs from the Wayback Machine
	for page := uint(0); page < uint(pages); page++ {
		url := "https://web.archive.org/cdx/search/cdx?url=" + domain + "/*&output=json&collapse=urlkey&fl=original&page=" + fmt.Sprint(page)
		req := requests.MakeGetRequest(url)

		// Unmarshals JSON response
		var result config.Wayback
		err := json.Unmarshal(req, &result)
		if err != nil {
			log.Println(err)
		}

		// Iterates through each URL in response
		for i, entry := range result {
			if i != 0 {

				// URL Encode various special characters to avoid errors with url.Parse in GetWords function
				replacer := strings.NewReplacer(" ", "%20", "$", "%24", "`", "%60", "<", "%3C", "[", "%5B", "]", "%5D", "{", "%7B", "}", "%7D", "%", "%25", "\"", "%22", ";", "%3B", "\\", "%5C", "|", "%7C", ",", "%2C", "'", "%27", "\xa7", "%20", "\x15", "")
				word := replacer.Replace(entry[0])

				// Parses each URL and writes words to file
				words.GetWords(word, filename, minLen, maxLen)
			}
		}
	}
	return nil
}

// getWaybackPagination gets the total number of pages (of results) for the given domain
func getWaybackPagination(domain string) uint {
	getPages := "https://web.archive.org/cdx/search/cdx?url=" + domain + "/*&output=json&collapse=urlkey&fl=original&page=0&showNumPages=true"

	req := requests.MakeGetRequest(getPages)

	// Unmarshals JSON response
	var numPages uint
	err := json.Unmarshal(req, &numPages)
	if err != nil {
		log.Println(err)
	}

	// Grabs number of pages from response
	var pagination uint
	bodyString := string(req)
	if len(bodyString) > 0 {
		pagination := numPages
		return uint(pagination)
	}
	return pagination
}
