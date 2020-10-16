package main

import (
	"bufio"
	"fmt"
	"github.com/ugol/dawg"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: create-dictionary dictionary.txt dictionary.bin")
		return
	}

	textDictionary := os.Args[1]
	binaryDictionary := os.Args[2]

	builder := dawg.New()

	file, _ := os.Open(textDictionary)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		fmt.Println(word)
		builder.Add(word)
	}

	if scanner.Err() != nil {
		fmt.Printf(" > Failed!: %v\n", scanner.Err())
	}

	builder.Finish()
	_, err := builder.Save(binaryDictionary)

	if err != nil {
		fmt.Println("Error in saving dictionary")
		fmt.Println(err)
	} else {
		fmt.Printf("Dictionary correctly created: %v\n", binaryDictionary)
	}

}
