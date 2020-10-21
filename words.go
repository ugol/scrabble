package main

import (
	"fmt"
	"github.com/ugol/dawg"
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
	fmt.Printf(	"Usage:\n" +
						"   words string [min] [max]\n" +
						"   words animals\n" +
						"   words animals 4\n" +
						"   words animals 4 5\n" +
						"   words AniMals 5\n");
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
	maxLength := len(word)

	if len(os.Args) >= 3 {
		minLength, err = strconv.Atoi(os.Args[2])
		if err != nil {
			usage()
			return
		}
	}

	if len(os.Args) == 4 {
		maxLength, err = strconv.Atoi(os.Args[3])
		if err != nil {
			usage()
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
		results := finder.FindAllPrefixesOf(strings.ToLower(string(permutation)), minLength, maxLength)
		for _, result := range results {
			fmt.Printf("%v\n", result.Word)
		}
	})
	//*/

}
