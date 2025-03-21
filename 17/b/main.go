package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	adv = 0
	bxl = 1
	bst = 2
	jnz = 3
	bxc = 4
	out = 5
	bdv = 6
	cdv = 7
)

var i_names = []string{"adv", "bxl", "bst", "jnz", "bxc",
	"out", "bdv", "cdv"}

const (
	// imm == 0..3
	rA = 4
	rB = 5
	rC = 6
)

var d_names = []string{"c0", "c1", "c2", "c3", "A", "B", "C"}

type state struct {
	// Let's just use 0..3 as constants 0..3 making indexing trivial
	ds [7]int
	// Have no idea how to call it, let it be code segment
	cs []byte
}

func readInput(fn string) (s state, ref string) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	for i := range rC + 1 {
		if i < rA {
			s.ds[i] = i
		} else {
			// We will brute-force value of rA, rest are 0
			str := fmt.Sprintf("Register %s: %%d", d_names[i])
			var unused int
			_, err := fmt.Fscanf(fd, str, &unused)
			if err != nil {
				panic(fmt.Sprintf("Reading %s register: %v", d_names[i], err))
			}
		}
	}

	var str string
	fmt.Fscanf(fd, "\nProgram: %s", &str)
	slice := strings.Split(str, ",")
	for _, e := range slice {
		s.cs = append(s.cs, byte(e[0])-byte('0'))
	}

	return s, str
}

func dump_insn(s state, i int) {
	if s.cs[i] == bxl || s.cs[i] == jnz {
		fmt.Printf("  %s\t%d\n", i_names[s.cs[i]], s.cs[i+1])
	} else {
		fmt.Printf("  %s\t%s\n", i_names[s.cs[i]], d_names[s.cs[i+1]])
	}
}

func dump_regs(s state) {
	fmt.Printf("A: %d,\tB: %d,\tC:%d\n", s.ds[rA], s.ds[rB], s.ds[rC])
}

func dump(s state) {
	fmt.Println("DS:")
	for i := 0; i <= rC; i++ {
		fmt.Printf("  %s: %d\n", d_names[i], s.ds[i])
	}

	fmt.Println("CS:")
	for i := 0; i < len(s.cs); i += 2 {
		dump_insn(s, i)
	}
}

func search_r(cur int, r []int) int {
	if len(r) == 0 {
		return cur
	}
	for j := 0; j < 8; j++ {
		A := (cur << 3) | j
		t := A & 7
		val := (t ^ (A >> (t ^ 3)) ^ 6) & 7
		if val == r[0] {
			// fmt.Println("Match:", A, r)
			res := search_r(A, r[1:len(r)])
			if res > 0 {
				return res
			}
		}
	}
	return -1
}

func main() {
	s, ref := readInput("input.txt")
	dump(s)
	fmt.Println("Ref: ", ref)
	//  Initially tried to optimize the interpreter, but then
	//  understood that the number of iterations is around 2^45 which
	//  looks like a joke.
	//  That is why, whole program was reduced from:
	//    bst	A
	//    bxl	3
	//    cdv	B
	//    bxl	5
	//    adv	c3
	//    bxc	c1
	//    out	B
	//    jnz	0
	//  to (on start rB == rC == 0):
	//    t := rA & 7
	//	  val := (t ^ (rA >> (t ^ 3)) ^ 6) & 7
	//    loop -2 if rA > 0
	//  brute-forcing won't work here as well.
	//  Finally, it can be seen that calculation is uses 4-bit chunks at the
	//  end of rA + some upper bits.  After output such chunk is discarded.
	//  Hence, if walk backward, we can assert that upper bits of A are zeroes
	//  at the end and get those 3 bits, then having them proceed to next
	//  element in the output (backwards) and so on.  Recursion is needed here
	//  as there may be matches on given level which will led to dead end.
	//  Really disappointing task. I was hoping to do some interpreter tuning.

	var r []int
	for i := len(ref) - 1; i >= 0; {
		r = append(r, int(ref[i])-int('0'))
		i -= 2
	}
	fmt.Println("Result:", search_r(0, r))
}
