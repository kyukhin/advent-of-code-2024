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

func check_cycle(data [][]byte, x, y, dir int) bool {
	lx := len(data[0])
	ly := len(data)
	count := 0
	for {
		count++
		if count > lx*ly {
			return true
		}
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

	return false
}

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

	count := 0
	for i, l := range data {
		for j, _ := range l {
			// Prohibited to but obstacle just in front of
			if (j == x && i == y-1) ||
				data[i][j] == '#' ||
				data[i][j] == '^' {
				continue
			}
			data[i][j] = '#'
			if check_cycle(data, x, y, dir) {
				count++
			}
			data[i][j] = '.'
		}
	}

	fmt.Println("Result is:", count)
}
