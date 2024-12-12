package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(fn string) *list.List {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fd, err))
	}
	defer fd.Close()

	l := list.New()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		slice := strings.Split(line, " ")
		for _, e := range slice {
			d, err := strconv.Atoi(e)
			if err != nil {
				panic(fmt.Sprintf("Unable to convert to num %s: %d", e, err))
			}
			l.PushBack(d)
		}
	}

	return l
}

func count_digits(d int) (c int) {
	for ; d > 0; c++ {
		d = d / 10
	}
	return c
}

func pow10(n int) (r int) {
	r = 1
	for range n {
		r *= 10
	}
	return r
}

func map_get(m map[int]map[int]int, num, iters int) int {
	if arr, ok := m[num]; ok {
		if x, ok := arr[iters]; ok {
			return x
		}
	}
	return -1
}

func map_set(m map[int]map[int]int, num, iters, val int) {
	if arr, ok := m[num]; ok {
		arr[iters] = val
	} else {
		m[num] = make(map[int]int)
		m[num][iters] = val
	}
}

func iterate(m map[int]map[int]int, l *list.List, num int) (len int) {
	if num == 0 {
		len = l.Len()
	} else {
		for i := l.Front(); i != nil; i = i.Next() {
			if x := map_get(m, i.Value.(int), num-1); x != -1 {
				len += x
			} else {
				l1 := list.New()

				v := i.Value.(int)
				// 1. If the stone is engraved with the number 0, it
				//    is replaced by a stone engraved with the number 1.
				if v == 0 {
					l1.PushBack(1)
				} else {
					// 2. If the stone is engraved with a number that has
					//    an even number of digits, it is replaced by two
					//    stones.
					//    The left half of the digits are engraved on the new
					//    left stone, and the right half of the digits are
					//    engraved on the new right stone. (The new numbers
					//    don't keep extra leading zeroes: 1000 would become
					//    stones 10 and 0.)
					if n := count_digits(v); n%2 == 0 {
						n /= 2
						x := v / pow10(n)
						y := v % pow10(n)
						l1.PushBack(x)
						l1.PushBack(y)
					} else {
						// 3. If none of the other rules apply, the stone is
						//    replaced by a new stone; the old stone's number
						//    multiplied by 2024 is engraved on the new stone.
						l1.PushBack(v * 2024)
					}
				}
				l := iterate(m, l1, num-1)
				map_set(m, v, num-1, l)
				len += l
			}
		}
	}
	return len
}

func main() {
	l := readInput("input.txt")

	// Number -> Iterations -> Len
	calc_map := map[int]map[int]int{}
	len := iterate(calc_map, l, 75)

	fmt.Println("Result is:", len)
}
