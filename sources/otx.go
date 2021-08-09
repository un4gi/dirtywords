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

// Big shoutout to @lc for the inspiration!

// OTX pulls words from URLs found in AlienVault Open Threat Exchange
func OTX(domain string, filename string, minLen int, maxLen int) {
	var url string
	for page := 0; ; page++ {
		url = fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/domain/%s/url_list?limit=%d&page=%d", domain, config.OTXResultsLimit, page)
		req := requests.MakeGetRequest(url)

		// Unmarshal JSON response
		var result config.OTXResult
		err := json.Unmarshal(req, &result)
		if err != nil {
			log.Println(err)
		}

		// Iterates through each URL in the results
		for _, entry := range result.URLList {

			// URL Encode various special characters to avoid errors with url.Parse in GetWords function
			replacer := strings.NewReplacer(" ", "%20", "$", "%24", "`", "%60", "<", "%3C", "[", "%5B", "]", "%5D", "{", "%7B", "}", "%7D", "%", "%25", "\"", "%22", ";", "%3B", "\\", "%5C", "|", "%7C", ",", "%2C", "'", "%27", "\xa7", "%20", "\x15", "")
			word := replacer.Replace(entry.URL)

			// Parses each URL and writes words to file
			words.GetWords(word, filename, minLen, maxLen)
		}

		if !result.HasNext {
			break
		}
	}

}
