package main

import (
	"log"
	"os"

	"github.com/reinerhuechting/mergefiles/files"
	"github.com/reinerhuechting/mergefiles/stringlists"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("Not enough arguments.")
	}
	outfilename := os.Args[1]
	infilenames := os.Args[2:]

	filecontents := [][]string{}
	for _, filename := range infilenames {
		contents := files.ReadLines(filename)
		filecontents = append(filecontents, contents)
	}

	result := stringlists.RemoveDuplicates(stringlists.Join(filecontents...))
	files.WriteLines(result, outfilename)
}
