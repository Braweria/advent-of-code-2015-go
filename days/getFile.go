package days

import (
	"log"
	"os"
)

func GetFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return file
}
