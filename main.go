package main

import (
	"bufio"
	"os"

	log "github.com/sirupsen/logrus"
)

var endl = []byte("\n")

func endsWithNewLine(filename string) bool {
	logf := log.WithFields(log.Fields{
		"filename": filename,
	})

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		logf.Error("could not read file")
		return false
	}

	_, err = f.Seek(-1, 2)
	if err != nil {
		logf.Info("appears to be empty")
		return false
	}

	eof := make([]byte, 1)
	f.Read(eof)

	if eof[0] != endl[0] {
		logf.Error("missing newline")
		return false
	}

	return true
}

func main() {
	errors := false
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if !endsWithNewLine(scanner.Text()) {
			errors = true
		}
	}

	if errors {
		os.Exit(1)
	}

	os.Exit(0)
}
