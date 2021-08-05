package sources

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/un4gi/dirtywords/config"
	"github.com/un4gi/dirtywords/requests"
	"github.com/un4gi/dirtywords/words"
)

func OTX(domain string, filename string) {
	var url string
	for page := 0; ; page++ {
		url = fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/domain/%s/url_list?limit=%d&page=%d", domain, config.OTXResultsLimit, page)
		req := requests.MakeGetRequest(url)

		var result config.OTXResult
		err := json.Unmarshal(req, &result)
		if err != nil {
			log.Println(err)
		}

		for _, entry := range result.URLList {
			words.GetWords(entry.URL, domain, filename)
		}

		if !result.HasNext {
			break
		}
	}

}
