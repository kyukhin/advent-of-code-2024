package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readInput(fn string) (data []int) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fd, err))
	}
	defer fd.Close()

	var c rune
	for {
		_, err := fmt.Fscanf(fd, "%c", &c)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				break
			} else {
				panic(fmt.Sprintf("Scan failed %s: %v", fd, err))
			}
		}
		if c == '\n' {
			continue
		}
		d, err := strconv.Atoi(string(c))
		if err != nil {
			panic(fmt.Sprintf("Conversion failed %s: %v", string(c), err))
		}
		data = append(data, d)
	}

	return data
}

func find_empty(d []int, prev, max int) int {
	for i := prev + 1; i < max; i++ {
		if d[i] == -1 {
			return i
		}
	}
	return -1
}

func main() {
	data_compressed := readInput("input.txt")

	data_raw := []int{}
	fid := 0
	for i, x := range data_compressed {
		d := -1
		if i%2 == 0 {
			d = fid
			fid++
		}
		for range x {
			data_raw = append(data_raw, d)
		}
	}

	ptr_r := len(data_raw) - 1
	ptr_w := find_empty(data_raw, -1, len(data_raw))
	for ptr_w != -1 {
		data_raw[ptr_w] = data_raw[ptr_r]
		data_raw[ptr_r] = -1
		for data_raw[ptr_r] == -1 {
			ptr_r--
		}
		ptr_w = find_empty(data_raw, ptr_w, ptr_r)
	}

	result := 0
	for i := 0; i < len(data_raw); i++ {
		if data_raw[i] != -1 {
			result += i * data_raw[i]
		} else {
			break
		}
	}

	fmt.Println("Result is:", result)
}
