package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(fn string) (data [][]int) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		l := scanner.Text()
		v, err := strconv.Atoi(l[:strings.Index(l, ":")])
		if err != nil {
			panic("Unexpected format. Each line should be <number: ...>")
		}
		a := []int{v}
		for _, s := range strings.Split(l[strings.Index(l, ":")+2:], " ") {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic("Unexpected format. Each line should be <number: ...>")
			}
			a = append(a, v)
		}
		data = append(data, a)
	}

	return data
}

func concat(x, y int) int {
	pow := int(math.Log10(float64(y))) + 1
	for range pow {
		x *= 10
	}
	return x + y
}

func match_r(expected, cur int, data []int, trace []byte) bool {
	if len(data) == 0 {
		if cur == expected {
			fmt.Println("Match", expected, string(trace))
			return true
		} else {
			return false
		}
	}

	return match_r(expected, cur+data[0], data[1:], append(trace, '+')) ||
		match_r(expected, cur*data[0], data[1:], append(trace, '*')) ||
		match_r(expected, concat(cur, data[0]), data[1:], append(trace, 'W'))
}

func main() {
	data := readInput("input.txt")

	result := 0
	for _, d := range data {
		if match_r(d[0], d[1], d[2:], []byte{}) {
			fmt.Println("  -> ", d[1:])
			result += d[0]
		}
	}

	fmt.Println("Result is:", result)
}
