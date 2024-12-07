package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(fn string) (data []string) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("ope %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		data = append(data, line)
	}

	return
}

func check(b1, b2 byte) bool {
	if b1 == 'M' && b2 == 'S' {
		return true
	}
	return false
}

func main() {
	d := readInput("input.txt")

	counter := 0
	for i, e := range d {
		for j := 0; j < len(e); j++ {
			if i+2 < len(d) && j+2 < len(e) {
				if d[i+1][j+1] == 'A' &&
					(check(d[i][j], d[i+2][j+2]) ||
						check(d[i+2][j+2], d[i][j])) &&
					(check(d[i+2][j], d[i][j+2]) ||
						check(d[i][j+2], d[i+2][j])) {
					counter++
				}
			}
		}
	}

	fmt.Println("Result: ", counter)
}
