package main

import (
	"fmt"
	"io"
	"os"
)

type machine struct {
	a_x    int
	a_y    int
	b_x    int
	b_y    int
	dest_x int
	dest_y int
}

func readInput(fn string) (mm []machine) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	for {
		var m machine
		_, err := fmt.Fscanf(fd, "Button A: X+%d, Y+%d\n", &m.a_x, &m.a_y)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				break
			} else {
				panic(fmt.Sprintf("Reading A button params: %v", err))
			}
		}
		_, err = fmt.Fscanf(fd, "Button B: X+%d, Y+%d\n", &m.b_x, &m.b_y)
		if err != nil {
			panic(fmt.Sprintf("Reading B button params: %v", err))
		}
		_, err = fmt.Fscanf(fd, "Prize: X=%d, Y=%d\n\n", &m.dest_x, &m.dest_y)
		if err != nil {
			panic(fmt.Sprintf("Reading prize button params: %v", err))
		}
		// fmt.Println("Machine:", m.step_a_x, m.step_a_y, m.step_b_x, m.step_b_y, m.dest_x, m.dest_y)
		mm = append(mm, m)
	}
	return
}

func main() {
	mm := readInput("input.txt")

	res := 0
	for _, m := range mm {
		if (m.dest_x*m.b_y-m.b_x*m.dest_y)%(m.a_x*m.b_y-m.b_x*m.a_y) == 0 {
			tmp := (m.dest_x*m.b_y - m.b_x*m.dest_y) / (m.a_x*m.b_y - m.b_x*m.a_y)
			if (m.dest_y-m.a_y*tmp)%m.b_y == 0 {
				fmt.Println("Match!", m)
				res += 3 * tmp
				res += (m.dest_y - m.a_y*tmp) / m.b_y
			}
		}
	}

	fmt.Println("Number of machines parsed:", len(mm))

	fmt.Println("Result is:", res)
}
