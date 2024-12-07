package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(fn string) (rules [][]int, seqs [][]int) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", fn, err))
	}
	defer fd.Close()

	n1 := 0
	n2 := 0
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		_, err := fmt.Sscanf(line, "%d|%d\n", &n1, &n2)
		if err != nil {
			panic(fmt.Sprintf("Scan failed %s: %v", fn, err))
		}

		rules = append(rules, []int{n1, n2})
	}
	fmt.Println("Number of rules read:", len(rules))

	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, ",")
		arr := []int{}
		for _, e := range slice {
			x, err := strconv.Atoi(e)
			if err != nil {
				panic(fmt.Sprintf("Pages detection %s: %e", line, e))
			}
			arr = append(arr, x)
		}
		seqs = append(seqs, arr)
	}

	fmt.Println("Number of sequences read:", len(seqs))
	return rules, seqs
}

func check(y, x int, rules [][]int) bool {
	for _, r := range rules {
		if r[0] == x && r[1] == y {
			return false
		}
	}
	return true
}

func main() {
	rules, seqs := readInput("input.txt")
	result := 0
	for _, s := range seqs {
		processed := []int{}
		mark := false
		for i, x := range s {
			for j, y := range processed {
				if !check(y, x, rules) {
					t := s[i]
					s[i] = y
					s[j] = t
					mark = true
				}
			}
			processed = s[:i+1]
		}
		if mark {
			if len(s)%2 == 0 {
				panic("Sequence of even length. Not implemented")
			}
			result += s[len(s)/2]
		}
	}

	fmt.Println("Result: ", result)
}
