package files

import (
	"log"
	"os"
)

// WriteLines expects a list of lines and a file path.
// Opens that file and writes the list of lines.
// Overwrites any existing file content.
func WriteLines(lines []string, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}
