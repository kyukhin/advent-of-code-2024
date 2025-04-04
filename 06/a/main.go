package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(fn string) (data [][]byte) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("ope %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		l := []byte(line)
		data = append(data, l)
	}

	return
}

const (
	up = iota
	right
	down
	left
)

func main() {
	data := readInput("input.txt")

	x := -1
	y := -1
	dir := -1
	for i, l := range data {
		for j, e := range l {
			if e == '^' {
				x = j
				y = i
				dir = up
			}
		}
	}

	lx := len(data[0])
	ly := len(data)
	for {
		data[y][x] = 'X'
		if (dir == up && y == 0) ||
			(dir == down && y+1 == ly) ||
			(dir == left && x == 0) ||
			(dir == right && x+1 == lx) {
			break
		}
		for {
			var next byte
			switch dir {
			case up:
				next = data[y-1][x]
			case down:
				next = data[y+1][x]
			case right:
				next = data[y][x+1]
			case left:
				next = data[y][x-1]
			}
			if next == '#' {
				dir++
				if dir > left {
					dir = up
				}
			} else {
				break
			}
		}

		switch dir {
		case up:
			y -= 1
		case down:
			y += 1
		case left:
			x -= 1
		case right:
			x += 1
		}
	}

	count := 0
	for _, l := range data {
		for _, e := range l {
			if e == 'X' {
				count++
			}
		}
	}

	fmt.Println("Result is:", count)
}
