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

	for k := range 200 {
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
		if k < 0 {
			continue
		}
	}
	var x = [...]int{0, 0, 0, 0}
	for _, r := range rr {
		i := 0
		if r.x == sz_x/2 || r.y == sz_y/2 {
			continue
		}
		if r.x > sz_x/2 {
			i += 1
		}
		if r.y > sz_y/2 {
			i += 2
		}
		x[i]++
	}
	res := x[0] * x[1] * x[2] * x[3]
	fmt.Println(x[0], x[1], x[2], x[3])
	fmt.Println("Number of robots parsed:", len(rr))
	fmt.Println("Result is:", res)
}
