package words

import (
	"fmt"
	"os"
)

func WriteWordlist(domain string, word string, filename string) {
	wordlist, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error writing to file.")
		os.Exit(1)
	}

	wordlist.WriteString(word + "\r\n")
	wordlist.Close()
}
