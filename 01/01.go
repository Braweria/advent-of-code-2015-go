package main

import (
	"fmt"
	"log"
	"os"
)

func Day01() {
	fmt.Println("Day 01")
	content, err := os.Readlink("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
