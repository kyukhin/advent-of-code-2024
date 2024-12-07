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

func check(b1, b2, b3, b4 byte) bool {
	if b1 == 'X' && b2 == 'M' && b3 == 'A' && b4 == 'S' {
		return true
	}
	return false
}

func main() {
	d := readInput("input.txt")

	counter := 0
	for i, e := range d {
		for j := 0; j < len(e); j++ {
			if j+3 < len(e) {
				// Right
				if check(e[j], e[j+1], e[j+2], e[j+3]) {
					counter++
				}
				// Left
				if check(e[j+3], e[j+2], e[j+1], e[j]) {
					counter++
				}
			}
			if i+3 < len(d) {
				// Down
				if check(d[i][j], d[i+1][j], d[i+2][j], d[i+3][j]) {
					counter++
				}
				// Up
				if check(d[i+3][j], d[i+2][j], d[i+1][j], d[i][j]) {
					counter++
				}
			}
			if i+3 < len(d) && j+3 < len(e) {
				// Down-right
				if check(d[i][j], d[i+1][j+1], d[i+2][j+2], d[i+3][j+3]) {
					counter++
				}
				// Down-left
				if check(d[i+3][j+3], d[i+2][j+2], d[i+1][j+1], d[i][j]) {
					counter++
				}
				// Up-right
				if check(d[i][j+3], d[i+1][j+2], d[i+2][j+1], d[i+3][j]) {
					counter++
				}
				// Up-left
				if check(d[i+3][j], d[i+2][j+1], d[i+1][j+2], d[i][j+3]) {
					counter++
				}
			}
		}
	}

	fmt.Println("Result: ", counter)
}
