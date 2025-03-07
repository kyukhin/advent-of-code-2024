package main

import (
	"fmt"
	"io"
	"os"
)

const sz_x int = 101
const sz_y int = 103

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func readInput(fn string) (rr []robot) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	for {
		var r robot
		_, err := fmt.Fscanf(fd, "p=%d,%d v=%d,%d\n", &r.x, &r.y, &r.dx, &r.dy)
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				break
			} else {
				panic(fmt.Sprintf("Reading robot params: %v", err))
			}
		}
		// fmt.Println("Robot:", r.x, r.y, r.dx, r.dy)
		rr = append(rr, r)
	}
	return
}

func main() {
	rr := readInput("input.txt")

	max_l := 0
	idx := 0
	for k := range sz_x * sz_y {
		for i := range rr {
			if rr[i].dx > 0 {
				rr[i].x = (rr[i].x + rr[i].dx) % sz_x
			} else {
				rr[i].x -= (-rr[i].dx) % sz_x
				if rr[i].x < 0 {
					rr[i].x += sz_x
				}
			}

			if rr[i].dy > 0 {
				rr[i].y = (rr[i].y + rr[i].dy) % sz_y
			} else {
				rr[i].y -= (-rr[i].dy) % sz_y
				if rr[i].y < 0 {
					rr[i].y += sz_y
				}
			}
		}

		l := 0
		for x := range sz_x {
			start := false
			tmp := 0
			for y := range sz_y {
				found := false
				for _, r := range rr {
					if start && r.x == x && r.y == y {
						found = true
						tmp++
						break
					}
					if !start && r.x == x && r.y == y {
						start = true
						found = true
						tmp++
						break
					}
				}
				if found == false && start == true {
					if tmp > l {
						l = tmp
					}
					start = false
					tmp = 0
				}
			}
		}

		if max_l < l {
			max_l = l
			idx = k
		}
	}
	fmt.Println("Result is:", idx+1, ", height: ", max_l)
}
