package main

import (
	"bufio"
	"fmt"
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

func match_r(expected, cur int, data []int, trace []byte) bool {
	if len(data) == 1 {
		if cur+data[0] == expected {
			fmt.Println("Match ", expected, string(append(trace, '+')))
			return true
		}
		if cur*data[0] == expected {
			fmt.Println("Match ", expected, string(append(trace, '*')))
			return true
		}
		return false
	} else {
		return match_r(expected, cur+data[0], data[1:], append(trace, '+')) ||
			match_r(expected, cur*data[0], data[1:], append(trace, '*'))
	}
}

func main() {
	data := readInput("input.txt")

	result := 0
	for _, d := range data {
		if match_r(d[0], d[1], d[2:], []byte{}) {
			fmt.Println("  ==>", d[1:])
			result += int(d[0])
		}
	}

	fmt.Println("Result is:", result)
}
