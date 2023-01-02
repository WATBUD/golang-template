package main

import (
	"log"
	"os"
)

func main() {
	src := "C:\\Users\\WATBUD\\Desktop\\Test.txt"

	dest := "C:\\Users\\WATBUD\\Desktop\\words2.txt"

	bytesRead, err := os.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
