package main

import (
	"fmt"
	"io"
	"os"
)

func readInput(fn string) (arr1, arr2 []int) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("ope %s: %v", fn, err))
	}

	var n1, n2 int
	for {
		_, err := fmt.Fscanf(fd, "%d   %d\n", &n1, &n2)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				return
			}

			panic(fmt.Sprintf("Scan failed %s: %v", fn, err))
		}

		arr1 = append(arr1, n1)
		arr2 = append(arr2, n2)
	}

	return
}

func main() {
	arr1, arr2 := readInput("input.txt")

	dst := 0

	for _, elt1 := range arr1 {
		for _, elt2 := range arr2 {
			if elt1 == elt2 {
				dst += elt1
			}
		}
	}
	fmt.Println("Distance: ", dst)
}
