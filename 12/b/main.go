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

func create_empty_map(lx, ly int) (m [][]byte) {
	for range ly {
		m = append(m, []byte{})
		for range lx {
			m[len(m)-1] = append(m[len(m)-1], 0)
		}
	}
	return
}

func fill_shadow(m, shadow [][]byte, x, y int) {
	lx := len(m[0])
	ly := len(m)
	if x < 0 || y < 0 || x >= lx || y >= ly || shadow[y][x] == 1 {
		return
	}
	shadow[y][x] = 1
	if x > 0 && m[y][x] == m[y][x-1] {
		fill_shadow(m, shadow, x-1, y)
	}
	if x+1 < lx && m[y][x] == m[y][x+1] {
		fill_shadow(m, shadow, x+1, y)
	}
	if y > 0 && m[y][x] == m[y-1][x] {
		fill_shadow(m, shadow, x, y-1)
	}
	if y+1 < ly && m[y][x] == m[y+1][x] {
		fill_shadow(m, shadow, x, y+1)
	}
}

func main() {
	m := readInput("input.txt")
	lx := len(m[0])
	ly := len(m)

	visited := create_empty_map(lx, ly)

	result := 0
	for {
		x := -1
		y := -1
	scan_unvisited:
		for j := range ly {
			for i := range lx {
				if visited[j][i] == 0 {
					x = i
					y = j
					break scan_unvisited
				}
			}
		}

		if x == -1 {
			break
		}

		shadow := create_empty_map(lx, ly)
		fill_shadow(m, shadow, x, y)
		area := 0
		perimeter := 0
		for j := range ly {
			for i := range lx {
				if shadow[j][i] == 1 {
					if i > 0 && shadow[j][i-1] == 0 ||
						i == 0 {
						if j < 1 || shadow[j-1][i] != 1 ||
							i > 0 && shadow[j-1][i-1] == 1 {
							perimeter++
						}
					}
					if i+1 < lx && shadow[j][i+1] != 1 ||
						i+1 == lx {
						if j < 1 || shadow[j-1][i] != 1 ||
							i+1 < lx && shadow[j-1][i+1] == 1 {
							perimeter++
						}
					}
					if j > 0 && shadow[j-1][i] == 0 ||
						j == 0 {
						if i < 1 || shadow[j][i-1] != 1 ||
							j > 0 && shadow[j-1][i-1] == 1 {
							perimeter++
						}
					}
					if j+1 >= ly || shadow[j+1][i] == 0 ||
						j+1 == ly {
						if i < 1 || shadow[j][i-1] != 1 ||
							j+1 < ly && shadow[j+1][i-1] == 1 {
							perimeter++
						}
					}
					area++
					visited[j][i] = 1
				}
			}
		}
		// fmt.Println("v=", string(m[y][x]), "a=", area, "p=", perimeter)
		result += area * perimeter
	}

	fmt.Println("Result is:", result)
}
