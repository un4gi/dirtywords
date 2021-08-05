package words

import (
	"sort"
	"strings"

	"net/url"
)

func GetWords(urls string, domain string, filename string) {

	u, _ := url.Parse(urls)
	u.RawQuery = ""
	u.Fragment = ""

	a := strings.TrimPrefix(u.String(), "https://")
	b := strings.TrimPrefix(a, "http://")
	c := strings.TrimPrefix(b, u.Host+"/")
	d := strings.Split(c, "/")
	e := sort.StringSlice(d)
	for f := len(e) - 1; f >= 0; f-- {
		WriteWordlist(domain, e[f], filename)
	}

}
