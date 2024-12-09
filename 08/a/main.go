package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInput(fn string) (data [][]byte) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		l := []byte(line)
		data = append(data, l)
	}

	return
}

type point struct {
	x int
	y int
}

func is_fit(rect, p point) bool {
	if p.x >= 0 && p.y >= 0 &&
		p.x < rect.x && p.y < rect.y {
		return true
	} else {
		return false
	}
}

func flip(p1, p2 point) point {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return point{x: p1.x - dx, y: p1.y - dy}
}

func fill(a_locs [][]byte, r, p1, p2 point) {
	p := flip(p1, p2)
	if is_fit(r, p) {
		a_locs[p.y][p.x] = '*'
	}
}

func main() {
	data := readInput("input.txt")

	result := 0

	r := point{x: len(data[0]), y: len(data)}
	antennas := map[byte][]point{}
	for i, l := range data {
		for j, e := range l {
			if e != '.' {
				antennas[e] = append(antennas[e], point{x: j, y: i})
			}
		}
	}

	a_locs := make([][]byte, r.y)
	for i := range a_locs {
		a_locs[i] = make([]byte, r.x)
		for j := range a_locs[i] {
			a_locs[i][j] = '.'
		}
	}

	for _, a := range antennas {
		for i, e1 := range a {
			for _, e2 := range a[i+1:] {
				fill(a_locs, r, e1, e2)
				fill(a_locs, r, e2, e1)
			}
		}
	}

	for _, l := range a_locs {
		for _, e := range l {
			if e == '*' {
				result++
			}
		}
	}

	fmt.Println("Result is:", result)
}
