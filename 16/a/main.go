package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type cell struct {
	v byte
	c int
}

type field struct {
	d  [][]cell
	sx int
	sy int
	ex int
	ey int
}

type queue_elt struct {
	cost int
	dir  int
	idx  int
	x    int
	y    int
	px   int
	py   int
}

type p_queue []*queue_elt

func (pq p_queue) Len() int           { return len(pq) }
func (pq p_queue) Less(i, j int) bool { return pq[i].cost > pq[j].cost }
func (pq p_queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].idx = i
	pq[j].idx = j
}
func (pq *p_queue) Push(x any) {
	n := len(*pq)
	item := x.(*queue_elt)
	item.idx = n
	*pq = append(*pq, item)
}
func (pq *p_queue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.idx = -1
	*pq = old[0 : n-1]
	return item
}

func readInput(fn string) (f field) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		f.d = append(f.d, []cell{})
		fl := len(f.d) - 1
		// fmt.Println(line)
		for _, b := range line {
			switch b {
			case 'S':
				f.d[fl] = append(f.d[fl], cell{v: '.', c: -1})
				f.sx = len(f.d[fl]) - 1
				f.sy = fl
			case 'E':
				f.d[fl] = append(f.d[fl], cell{v: '.', c: -1})
				f.ex = len(f.d[fl]) - 1
				f.ey = fl
			case '#':
				fallthrough
			case '.':
				f.d[fl] = append(f.d[fl], cell{v: byte(b), c: -1})
			default:
				panic("Unexpected symbol")
			}
		}
	}

	return
}

func dump(title string, f field) {
	fmt.Println(title)
	for _, y := range f.d {
		for _, x := range y {
			if x.v == '#' {
				fmt.Printf("#####")
			} else {
				fmt.Printf("%5d", x.c)
			}
		}
		fmt.Println()
	}
}

func main() {
	f := readInput("input.txt")
	sx := len(f.d[0])
	sy := len(f.d)
	fmt.Printf("Field size: %d x %d, start: {%d, %d}, finish = {%d, %d}\n",
		sx, sy, f.sx, f.sy, f.ex, f.ey)

	// dump("Field: ", f)

	// Pair is {0: y, 1: x}
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	pq := make(p_queue, 1)
	pq[0] = &queue_elt{x: f.sx, y: f.sy, cost: 0, dir: 0}
	heap.Init(&pq)

	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*queue_elt)
		if f.d[i.y][i.x].v == '.' {
			if f.d[i.y][i.x].c < 0 ||
				f.d[i.y][i.x].c > i.cost {
				f.d[i.y][i.x].c = i.cost
			}
		}
		nc := i.cost + 1
		for k := range 4 {
			ndir := (i.dir + k) % 4

			nx := i.x + dir[ndir][1]
			if nx >= sx || nx < 0 {
				continue
			}

			ny := i.y + dir[ndir][0]
			if ny >= sy || ny < 0 {
				continue
			}

			if f.d[ny][nx].v == '.' &&
				(f.d[ny][nx].c > nc ||
					f.d[ny][nx].c < 0) {
				heap.Push(&pq, &queue_elt{
					x: nx, y: ny,
					px: i.x, py: i.y,
					cost: nc, dir: ndir})
			}
			if k < 2 {
				nc += 1000
			} else {
				nc -= 1000
			}
		}

	}

	res := f.d[f.ey][f.ex].c
	// dump("Finish: ", f)
	fmt.Println("Result:", res)
}
