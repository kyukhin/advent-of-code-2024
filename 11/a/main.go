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

func iterate(l *list.List) {
	for i := l.Front(); i != nil; i = i.Next() {
		v := i.Value.(int)
		// 1. If the stone is engraved with the number 0, it
		//    is replaced by a stone engraved with the number 1.
		if v == 0 {
			i.Value = 1
			continue
		}

		// 2. If the stone is engraved with a number that has
		//    an even number of digits, it is replaced by two stones.
		//    The left half of the digits are engraved on the new left
		//    stone, and the right half of the digits are engraved on the
		//    new right stone. (The new numbers don't keep extra leading
		//    zeroes: 1000 would become stones 10 and 0.)
		n := count_digits(v)
		if n%2 == 0 {
			n /= 2
			x := v / pow10(n)
			y := v % pow10(n)

			i.Value = x
			i = l.InsertAfter(y, i)
			continue
		}

		// 3. If none of the other rules apply, the stone is replaced by
		//    a new stone; the old stone's number multiplied by 2024 is
		//    engraved on the new stone.
		i.Value = v * 2024
	}
}

func main() {
	list := readInput("input.txt")

	for range 25 {
		iterate(list)
	}

	len := list.Len()
	fmt.Println("Result is:", len)
}
