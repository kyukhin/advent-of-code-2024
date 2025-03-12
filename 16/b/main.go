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
				fmt.Printf("  #  ")
			} else {
				fmt.Printf("%5d", x.c)
			}
		}
		fmt.Println()
	}
}

func dijkstra(f *field) {
	// Pair is {0: y, 1: x}
	dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	pq := make(p_queue, 4)
	pq[0] = &queue_elt{x: f.sx, y: f.sy, cost: 0, dir: 0}
	pq[1] = &queue_elt{x: f.sx, y: f.sy, cost: 0, dir: 1}
	pq[2] = &queue_elt{x: f.sx, y: f.sy, cost: 0, dir: 2}
	pq[3] = &queue_elt{x: f.sx, y: f.sy, cost: 0, dir: 3}
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
			if nx >= len(f.d[0]) || nx < 0 {
				continue
			}

			ny := i.y + dir[ndir][0]
			if ny >= len(f.d) || ny < 0 {
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
}

func main() {
	f1 := readInput("input.txt")
	f2 := readInput("input.txt")
	fmt.Printf("Field start: {%d, %d}, finish = {%d, %d}\n",
		f1.sx, f1.sy, f1.ex, f1.ey)

	dijkstra(&f1)

	f2.sx, f2.ex = f2.ex, f2.sx
	f2.sy, f2.ey = f2.ey, f2.sy
	dijkstra(&f2)

	if f1.d[f1.ey][f1.ex].c != f2.d[f2.ey][f2.ex].c {
		panic("Impossible")
	}

	sum := f1.d[f1.ey][f1.ex].c

	res := 0
	for y := range len(f1.d) {
		for x := range len(f1.d[0]) {
			k := 0
			// Check for cut corners as they are different depending
			// of direction. We also should ignore walls.
			if f1.d[y][x].v == '#' {
				continue
			}
			// No sense to check on borders, we know that
			// we are guarded by the frame.
			if x != 0 && y != 0 &&
				y < len(f1.d)-1 &&
				x < len(f1.d[0])-1 {
				dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
				for i := range 4 {
					if f1.d[y+dir[i][0]][x+dir[i][1]].c+
						f2.d[y+dir[i][0]][x+dir[i][1]].c == sum {
						k++
					}
				}
			}
			if f1.d[y][x].c+f2.d[y][x].c == sum ||
				k >= 2 { // corner is possible for k == 2, 3 and 4
				res++
			}
		}
	}

	// Useful for small fields like input2
	// dump("Finish1: ", f1)
	// dump("Finish2: ", f2)

	fmt.Println("Result:", res)
}
