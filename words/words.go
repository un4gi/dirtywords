package words

import (
	"log"
	"sort"
	"strings"

	"net/url"
)

// GetWords takes a URL passed from another function and parses it. Each resulting string (file/folder) is written to the output file.
func GetWords(urls string, filename string, minLen int, maxLen int) {

	u, err := url.Parse(urls)
	//u.RawQuery = ""
	//u.Fragment = ""

	// Handle errors from invalid characters in URL
	if err != nil {
		log.Println(err)
	}

	a := u.EscapedPath()            // Stores escaped file path
	b := strings.TrimPrefix(a, "/") // Strips "/" from beginning of file path
	c := strings.TrimSuffix(b, "§") // Strips "§" from end of file path due to encountered errors during testing
	d := strings.Split(c, "/")      // Splits the string into a slice of strings to separate files from directories
	e := sort.StringSlice(d)        // Sorts the slice of strings

	// Iterates over each string in the slice
	for f := len(e) - 1; f >= 0; f-- {

		// Only retrieve words that match min/max length
		if len(e[f]) >= minLen && len(e[f]) <= maxLen {
			WriteWordlist(e[f], filename) // Writes word to file
		}
	}
}
