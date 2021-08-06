package words

import (
	"bufio"
	"os"
	"strings"

	"github.com/mpvl/unique"
)

// Thanks to incidrthreat for help with building the sort function!

// SortWordList opens the unsorted wordlist and sorts it uniquely
func SortWordList(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	unique.Sort(unique.StringSlice{&lines}) // Sorts the slice of strings (unsorted word list) uniquely

	// Opens unsorted wordlist to truncate (overwrite) with sorted list
	f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	w := bufio.NewWriter(f)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		} else {
			w.WriteString(line + "\n")
		}
	}

	w.Flush()
	f.Close()
}
