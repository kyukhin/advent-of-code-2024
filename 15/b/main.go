package main

import (
	"bufio"
	"fmt"
	"os"
)

type field struct {
	d  [][]byte
	p  []byte
	rx int
	ry int
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
		f.d = append(f.d, []byte{})
		fl := len(f.d) - 1
		// fmt.Println(line)
		for _, b := range line {
			switch b {
			case '#':
				fallthrough
			case '.':
				f.d[fl] = append(f.d[fl], byte(b), byte(b))
			case '@':
				f.d[fl] = append(f.d[fl], byte(b), '.')
			case 'O':
				f.d[fl] = append(f.d[fl], '[', ']')
			default:
				panic("Unexpected symbol")
			}
		}
	}

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

func dump(title string, f field) {
	fmt.Println(title)
	for _, y := range f.d {
		for _, x := range y {
			fmt.Printf(string(x))
		}
		fmt.Println()
	}
}

func step_h(x int, y int, dir int, f *field) bool {
	switch f.d[y][x] {
	case '.':
		return true
	case '#':
		return false
	case '@':
		fallthrough
	case '[':
		fallthrough
	case ']':
		if step_h(x+dir, y, dir, f) {
			if f.d[y][x] == '@' {
				f.rx += dir
			}
			f.d[y][x+dir] = f.d[y][x]
			f.d[y][x] = '.'
			return true
		} else {
			return false
		}
	default:
		panic("Impossible")
	}
	return false
}

func check_step_v(x int, y int, dir int, f *field) bool {
	switch f.d[y][x] {
	case '.':
		return true
	case '#':
		return false
	case '@':
		if check_step_v(x, y+dir, dir, f) {
			return true
		}
	case '[':
		if check_step_v(x, y+dir, dir, f) && check_step_v(x+1, y+dir, dir, f) {
			return true
		}
	case ']':
		if check_step_v(x, y+dir, dir, f) && check_step_v(x-1, y+dir, dir, f) {
			return true
		}
	default:
		panic("Impossible")
	}
	return false
}

func do_step_v(x int, y int, dir int, f *field) {
	switch f.d[y][x] {
	case '@':
		do_step_v(x, y+dir, dir, f)
		f.d[y+dir][x] = f.d[y][x]
		f.d[y][x] = '.'
		f.ry += dir
	case '[':
		do_step_v(x, y+dir, dir, f)
		f.d[y+dir][x] = f.d[y][x]
		f.d[y][x] = '.'

		do_step_v(x+1, y+dir, dir, f)
		f.d[y+dir][x+1] = f.d[y][x+1]
		f.d[y][x+1] = '.'
	case ']':
		do_step_v(x, y+dir, dir, f)
		f.d[y+dir][x] = f.d[y][x]
		f.d[y][x] = '.'

		do_step_v(x-1, y+dir, dir, f)
		f.d[y+dir][x-1] = f.d[y][x-1]
		f.d[y][x-1] = '.'
	}
}

func main() {
	f := readInput("input.txt")
	fmt.Printf("Field size: %d x %d, prog. len: %d, robot start: x=%d, y=%d\n",
		len(f.d[0]), len(f.d), len(f.p), f.rx, f.ry)

	// dump("Init: ", f)
	for _, c := range f.p {
		switch c {
		case '^':
			if check_step_v(f.rx, f.ry, -1, &f) {
				do_step_v(f.rx, f.ry, -1, &f)
			}
		case 'v':
			if check_step_v(f.rx, f.ry, 1, &f) {
				do_step_v(f.rx, f.ry, 1, &f)
			}
		case '<':
			step_h(f.rx, f.ry, -1, &f)
		case '>':
			step_h(f.rx, f.ry, 1, &f)
		}
		// dump(string(c), f)
	}

	res := 0
	for y, l := range f.d {
		for x, e := range l {
			if e == '[' {
				res += 100*y + x
			}
		}
	}
	// dump("Finish: ", f)
	fmt.Println("Result:", res)
}
