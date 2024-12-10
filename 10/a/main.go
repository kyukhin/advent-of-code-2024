package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(fn string) (data [][]int) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		l := []int{}
		for _, c := range line {
			x := int(c - '0')
			if x < 0 || x > 9 {
				panic(fmt.Sprintf("Unable to recognize height %c: %d", c, x))
			}
			l = append(l, x)
		}
		data = append(data, l)
	}

	return
}

type point struct {
	x int
	y int
}

func is_fit(rect point, x, y int) bool {
	if x >= 0 && y >= 0 &&
		x < rect.x && y < rect.y {
		return true
	} else {
		return false
	}
}

const (
	up = iota
	right
	down
	left
)

func get(data [][]int, x, y int) int {
	return data[y][x]
}

func next_c(dir int) int {
	dir += 1
	if dir > left {
		dir = up
	}
	return dir
}

func step(data [][]int, r point, x, y, d int) (new_x, new_y int) {
	v := get(data, x, y)
	switch d {
	case up:
		if is_fit(r, x, y-1) &&
			get(data, x, y-1)-v == 1 {
			return x, y - 1
		}
	case right:
		if is_fit(r, x+1, y) &&
			get(data, x+1, y)-v == 1 {
			return x + 1, y
		}
	case down:
		if is_fit(r, x, y+1) &&
			get(data, x, y+1)-v == 1 {
			return x, y + 1
		}
	case left:
		if is_fit(r, x-1, y) &&
			get(data, x-1, y)-v == 1 {
			return x - 1, y
		}
	}

	return -1, -1
}

func search(data, visited [][]int, r point, x, y, d, score int) int {
	if get(data, x, y) == 9 {
		if get(visited, x, y) == -1 {
			visited[y][x] = 0
			return 1
		} else {
			return 0
		}
	}

	s := 0
	for nd := 0; nd < left+1; nd++ {
		if d != -1 && nd == next_c(next_c(d)) {
			continue // Do not go back.
		}
		nx, ny := step(data, r, x, y, nd)
		if nx != -1 {
			s += search(data, visited, r, nx, ny, nd, score)
		}
	}
	return score + s
}

func main() {
	data := readInput("input.txt")

	r := point{x: len(data[0]), y: len(data)}
	visited := [][]int{}
	for i := range r.y {
		visited = append(visited, []int{})
		for range r.x {
			visited[i] = append(visited[i], -1)
		}
	}

	result := 0
	for i := range r.x {
		for j := range r.y {
			if get(data, i, j) == 0 {
				result += search(data, visited, r, i, j, -1, 0)
				for ii := range r.x {
					for jj := range r.y {
						visited[jj][ii] = -1
					}
				}

			}
		}
	}

	fmt.Println("Result is:", result)
}
