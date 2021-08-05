package requests

import (
	"io/ioutil"
	"log"
	"net/http"
)

func MakeGetRequest(url string) (bodyBytes []byte) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}

		return bodyBytes
	}
	return bodyBytes
}

func PlainGetRequest(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	SetHeaders(req)

	resp, e := http.DefaultClient.Do(req)
	return resp, e

}

func SetHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0")
}
