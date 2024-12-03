package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput(fn string) (data [][]int) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		var arr []int
		for _, w := range strings.Fields(line) {
			var v int
			fmt.Sscanf(w, "%d", &v)
			arr = append(arr, v)
		}
		data = append(data, arr)
	}

	return
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func is_report_safe(r []int) bool {
	is_safe := true
	is_asc := true
	for idx, _ := range r {
		if idx == len(r)-1 {
			break
		}

		if idx == 0 {
			if r[idx] > r[idx+1] {
				is_asc = false
			}
		}

		if abs(r[idx]-r[idx+1]) < 1 ||
			abs(r[idx]-r[idx+1]) > 3 {
			is_safe = false
			break
		}
		if is_asc != (r[idx] < r[idx+1]) {
			is_safe = false
			break
		}
	}

	return is_safe
}

func main() {
	data := readInput("input.txt")

	cnt := 0
	for _, arr := range data {
		if is_report_safe(arr) {
			cnt++
		} else {
			// Len is > 1 for sure, as 1 is always safe
			s := make([]int, len(arr)-1)
			for i, _ := range arr {
				if i == 0 || i == len(arr)-1 {
					if i == 0 {
						copy(s, arr[i+1:])
					}
					if i == len(arr)-1 {
						copy(s, arr[:i])
					}
					if is_report_safe(s) {
						cnt++
						break
					}
				} else {
					sl1 := make([]int, len(arr[:i]))
					copy(sl1, arr[:i])

					sl2 := make([]int, len(arr[i+1:]))
					copy(sl2, arr[i+1:])

					copy(s, append(sl1, sl2...))
					if is_report_safe(s) {
						cnt++
						break
					}
				}
			}
		}
	}
	fmt.Println("Safe reports count: ", cnt)
}
