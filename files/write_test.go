package files

import (
	"fmt"
	"testing"
	"time"
)

func TestWriteLines(t *testing.T) {
	// Fix a file name for the test output.
	testoutfile := "testoutput.txt"

	// Generate some lines to write.
	now := time.Now().Nanosecond()
	lines := []string{
		fmt.Sprintf("%v", now),
		fmt.Sprintf("%v", now+1),
		fmt.Sprintf("%v", now+2),
	}

	// Write the lines.
	WriteLines(lines, testoutfile)

	// Read the file again.
	readLines := ReadLines(testoutfile)

	// Assert that the read lines are identical to those written.
	if len(readLines) != len(lines) {
		t.Errorf("unexpected length of read lines: %d", len(readLines))
	}
	for i, readline := range readLines {
		if readline != lines[i] {
			t.Errorf("unexpected line at position %d: %s, expected %s", i, readline, lines[i])
		}
	}
}
