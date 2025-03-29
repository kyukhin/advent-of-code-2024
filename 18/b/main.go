package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

type cell struct {
	v byte
	c int
}

type field struct {
	d  [73][73]cell
	sx int
	sy int
	ex int
	ey int
}

type queue_elt struct {
	cost int
	idx  int
	x    int
	y    int
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

func readInput(fn string) (list []point) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y int

		if len(line) == 0 {
			break
		}
		_, err := fmt.Sscanf(line, "%d,%d", &x, &y)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse %v", err))
		}
		list = append(list, point{x: x, y: y})
	}

	return list
}

func resetField(list []point) (f field) {
	f.sx = 0 + 1
	f.sy = 0 + 1
	f.ex = 70 + 1
	f.ey = 70 + 1
	// 71 + 1 + 1 for borders
	for y := range 73 {
		for x := range 73 {
			if y == 0 || y == 72 || x == 0 || x == 72 {
				f.d[y][x] = cell{v: '#', c: -1}
			} else {
				f.d[y][x] = cell{v: '.', c: -1}
			}
		}
	}

	for _, e := range list {
		f.d[e.y+1][e.x+1].v = '#'
	}
	return f
}

func dump(title string, f field) {
	fmt.Println(title)
	for i, y := range f.d {
		for _, x := range y {
			if x.v == '#' {
				fmt.Printf("##")
			} else {
				fmt.Printf("%2d", x.c)
			}
		}
		fmt.Println()
		if i > 80 {
			break
		}
	}
}

func isPathExist(f *field) bool {
	pq := make(p_queue, 1)
	pq[0] = &queue_elt{x: f.sx, y: f.sy, cost: 0}
	heap.Init(&pq)
	f.d[f.sy][f.sx].c = 0

	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*queue_elt)
		// Pair is {0: y, 1: x}
		dir := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		for k := range 4 {
			nx := i.x + dir[k][1]
			if nx >= len(f.d[0]) || nx < 0 {
				continue
			}

			ny := i.y + dir[k][0]
			if ny >= len(f.d) || ny < 0 {
				continue
			}

			if f.d[ny][nx].v == '.' &&
				(f.d[ny][nx].c > i.cost+1 ||
					f.d[ny][nx].c < 0) {
				f.d[ny][nx].c = i.cost + 1
				heap.Push(&pq, &queue_elt{x: nx, y: ny, cost: i.cost + 1})
			}
		}
	}
	return f.d[f.ey][f.ex].c > 0
}

func findR(list []point, l, r int) int {
	if r == l {
		return r // We are interested in first brick which breaks the path
	}
	m := (l + r) / 2
	f := resetField(list[0:m])
	if isPathExist(&f) {
		// fmt.Println("True", m, l, r)
		return findR(list, m+1, r)
	} else {
		// fmt.Println("False", m, l, r)
		return findR(list, l, m-1)
	}
}

func main() {
	list := readInput("input.txt")

	res := findR(list, 0, len(list)-1)
	//fmt.Println(list[0:res])

	// dump("Finish: ", f)
	fmt.Println("Result:", list[res])
}
