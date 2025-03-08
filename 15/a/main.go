package main

import (
	"bufio"
	"fmt"
	"os"
)

type field struct {
	sx int
	sy int
	d  [][]byte
	rx int
	ry int
	p  []byte
}

func readInput(fn string) (f field) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	// Read the field
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		// fmt.Println(line)
		line_b := []byte(line)
		f.d = append(f.d, line_b)
	}
	f.sx = len(f.d[0])
	f.sy = len(f.d)

	// Find the robot
	for y, l := range f.d {
		found := false
		for x, e := range l {
			if e == '@' {
				f.rx = x
				f.ry = y
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// Read program
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			f.p = append(f.p, byte(c))
		}
	}

	return
}

func main() {
	f := readInput("input.txt")
	fmt.Printf("Field size: %d x %d, prog. len: %d, robot start: x=%d, y=%d\n",
		f.sx, f.sy, len(f.p), f.rx, f.ry)

	for _, c := range f.p {
		switch c {
		case '^':
			found := false
			ny := f.ry
			for ny >= 0 {
				if f.d[ny][f.rx] == '#' {
					break
				}
				if f.d[ny][f.rx] == '.' {
					found = true
					break
				}
				ny--
			}
			if found {
				for ny < f.ry {
					f.d[ny][f.rx] = f.d[ny+1][f.rx]
					ny++
				}
				f.d[ny][f.rx] = '.'
				f.ry--
			}
			break
		case 'v':
			found := false
			ny := f.ry
			for ny < f.sy {
				if f.d[ny][f.rx] == '#' {
					break
				}
				if f.d[ny][f.rx] == '.' {
					found = true
					break
				}
				ny++
			}
			if found {
				for ny > f.ry {
					f.d[ny][f.rx] = f.d[ny-1][f.rx]
					ny--
				}
				f.d[ny][f.rx] = '.'
				f.ry++
			}
			break
		case '<':
			found := false
			nx := f.rx
			for nx >= 0 {
				if f.d[f.ry][nx] == '#' {
					break
				}
				if f.d[f.ry][nx] == '.' {
					found = true
					break
				}
				nx--
			}
			if found {
				for nx < f.rx {
					f.d[f.ry][nx] = f.d[f.ry][nx+1]
					nx++
				}
				f.d[f.ry][nx] = '.'
				f.rx--
			}
			break
		case '>':
			found := false
			nx := f.rx
			for nx < f.sx {
				if f.d[f.ry][nx] == '#' {
					break
				}
				if f.d[f.ry][nx] == '.' {
					found = true
					break
				}
				nx++
			}
			if found {
				for nx > f.rx {
					f.d[f.ry][nx] = f.d[f.ry][nx-1]
					nx--
				}
				f.d[f.ry][nx] = '.'
				f.rx++
			}
			break
		}
	}

	res := 0
	for y, l := range f.d {
		for x, e := range l {
			if e == 'O' {
				res += 100*y + x
			}
		}
	}
	fmt.Println("Result is:", res)

}
