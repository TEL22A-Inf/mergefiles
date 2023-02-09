package main

import (
	"flag"
	"log"

	"github.com/reinerhuechting/mergefiles/files"
	"github.com/reinerhuechting/mergefiles/stringlists"
)

func main() {

	outfilename := flag.String("out", "out.txt", "The name of the output file to write.")
	flag.Parse()
	infilenames := flag.Args()
	if len(infilenames) < 1 {
		log.Fatal("Not enough arguments. At least one input file name must be given.")
	}

	filecontents := [][]string{}
	for _, filename := range infilenames {
		contents := files.ReadLines(filename)
		filecontents = append(filecontents, contents)
	}

	result := stringlists.RemoveDuplicates(stringlists.Join(filecontents...))
	files.WriteLines(result, *outfilename)
}
