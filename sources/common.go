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

func CommonCrawl(domain string, filename string) {
	req := requests.MakeGetRequest("http://index.commoncrawl.org/collinfo.json")

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
		for page := uint(0); page <= pages; page++ {
			url := index + fmt.Sprintf("&page=%d", page) // Example of data stored in url: https://index.commoncrawl.org/CC-MAIN-2021-25-index?url=example.com/*&output=json&fl=url&page=0

			resp, err := requests.PlainGetRequest(url)
			if err != nil {
				fmt.Println("Something went wrong...")
			}
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
				words.GetWords(getURLs.URL, domain, filename)
			}
			resp.Body.Close()
		}
	}
}

func getPagination(index string) (pagination uint) {
	getPages := index + "&page=0&showNumPages=true"
	req := requests.MakeGetRequest(getPages)

	var pageNum config.PageInfo
	err := json.Unmarshal(req, &pageNum)
	if err != nil {
		log.Println(err)
	}

	bodyString := string(req)
	if len(bodyString) > 0 {
		pagination := pageNum.Pages
		return uint(pagination)
	}
	return pagination
}
