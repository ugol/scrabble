package main

import (
	"fmt"
	"github.com/smhanov/dawg"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}

func shouldSwap(a []rune, start int, curr int) bool {

	if unicode.IsUpper(a[curr]) || unicode.IsUpper(a[start]) {
		return false
	}

	for i := start; i < curr; i++ {
		if a[i] == a[curr] {
			return false
		}
	}
	return true
}

func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		if shouldSwap(a, i, j) {
			a[i], a[j] = a[j], a[i]
			perm(a, f, i+1)
			a[i], a[j] = a[j], a[i]
		}
	}
}

func usage() {
	fmt.Println("Usage: words string 4")
	return
}

func main() {

	finder, err := dawg.Load("full-ita.bin")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) == 1 {
		usage()
		return
	}

	word := os.Args[1]
	minLength := 3

	if len(os.Args) == 3 {
		minLength, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Usage: words string 4")
			return
		}

	}

	/*

		finder.NumNodes()
		Perm([]rune(word), func(a []rune) {
			fmt.Println(string(a))
		})

		//*/

	//*
	Perm([]rune(word), func(permutation []rune) {
		results := finder.FindAllPrefixesOf(strings.ToLower(string(permutation)), minLength)
		for _, result := range results {
			fmt.Printf("%v\n", result.Word)
		}
	})
	//*/

}
