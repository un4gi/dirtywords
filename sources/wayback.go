package sources

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/un4gi/dirtywords/config"
	"github.com/un4gi/dirtywords/requests"
	"github.com/un4gi/dirtywords/words"
)

func WaybackURLs(domain string, filename string) error {
	pages := getWaybackPagination(domain)
	for page := uint(0); page < uint(pages); page++ {
		url := "https://web.archive.org/cdx/search/cdx?url=" + domain + "/*&output=json&collapse=urlkey&fl=original&page=" + fmt.Sprint(page)
		req := requests.MakeGetRequest(url)

		var result config.Wayback
		err := json.Unmarshal(req, &result)
		if err != nil {
			log.Println(err)
		}

		for i, entry := range result {
			if i != 0 {
				words.GetWords(entry[0], domain, filename)
			}
		}
	}
	return nil
}

func getWaybackPagination(domain string) uint {
	getPages := "https://web.archive.org/cdx/search/cdx?url=" + domain + "/*&output=json&collapse=urlkey&fl=original&page=0&showNumPages=true"

	req := requests.MakeGetRequest(getPages)

	var numPages uint
	err := json.Unmarshal(req, &numPages)
	if err != nil {
		log.Println(err)
	}

	var pagination uint
	bodyString := string(req)
	if len(bodyString) > 0 {
		pagination := numPages
		return uint(pagination)
	}
	return pagination
}
