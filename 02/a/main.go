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
		panic(fmt.Sprintf("ope %s: %v", fn, err))
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

func main() {
	data := readInput("input.txt")

	cnt := 0
	for _, arr := range data {
		is_safe := true
		is_asc := true
		for idx, _ := range arr {
			if idx == len(arr)-1 {
				break
			}

			if idx == 0 {
				if arr[idx] > arr[idx+1] {
					is_asc = false
				}
			}

			if abs(arr[idx]-arr[idx+1]) < 1 ||
				abs(arr[idx]-arr[idx+1]) > 3 {
				is_safe = false
				break
			}
			if is_asc != (arr[idx] < arr[idx+1]) {
				is_safe = false
				break
			}
		}
		if is_safe {
			cnt++
		}
	}
	fmt.Println("Cnt: ", cnt)
}
