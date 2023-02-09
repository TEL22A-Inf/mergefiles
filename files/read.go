package files

import (
	"bufio"
	"log"
	"os"
)

// ReadLines opens a given path and reads that file line by line.
// Returns the resulting list of strings.
func ReadLines(path string) []string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	result := []string{}
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	return result
}
