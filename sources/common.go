package sources

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"

	"github.com/un4gi/dirtywords/config"
	"github.com/un4gi/dirtywords/requests"
	"github.com/un4gi/dirtywords/words"
)

// Big shoutout to @lc for the inspiration!

// CommonCrawl retrieves a list of archived URLs from commoncrawl.org
func CommonCrawl(domain string, filename string, minLen int, maxLen int) {
	req := requests.MakeGetRequest("http://index.commoncrawl.org/collinfo.json")

	// Unmarshals JSON response
	var cdxAPI config.CollInfo
	err := json.Unmarshal(req, &cdxAPI)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(req)

	if len(bodyString) > 0 {
		index := cdxAPI[0].CDXAPI + fmt.Sprintf("?url=%s/*&output=json&fl=url", domain)
		pages := getPagination(index)
		fmt.Println("Generating wordlist. This could take a while...")

		// For each page of results, grab individual URLs
		for page := uint(0); page <= pages; page++ {
			url := index + fmt.Sprintf("&page=%d", page)

			resp, err := requests.PlainGetRequest(url)
			if err != nil {
				fmt.Println("Something went wrong...")
			}

			// Parse response body for URLs
			s := bufio.NewScanner(resp.Body)
			for s.Scan() {
				var getURLs config.UrlInfo
				if err := json.Unmarshal(s.Bytes(), &getURLs); err != nil {
					resp.Body.Close()
					fmt.Println("Could not unmarshal data...")
				}
				if getURLs.Error != "" {
					fmt.Println("Received an error from API.")
				}
				// Parse URLs and write words to file
				words.GetWords(getURLs.URL, filename, minLen, maxLen)
			}
			resp.Body.Close()
		}
	}
}

// getPagination gets the total number of pages (of results) for the given domain
func getPagination(index string) (pagination uint) {
	getPages := index + "&page=0&showNumPages=true"
	req := requests.MakeGetRequest(getPages)

	// Unmarshals JSON response
	var pageNum config.PageInfo
	err := json.Unmarshal(req, &pageNum)
	if err != nil {
		log.Println(err)
	}

	// Grabs number of pages
	bodyString := string(req)
	if len(bodyString) > 0 {
		pagination := pageNum.Pages
		return uint(pagination)
	}
	return pagination
}
