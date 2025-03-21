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

func readInput(fn string) (s state) {
	fd, err := os.Open(fn)
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fn, err))
	}
	defer fd.Close()

	for i := range rC + 1 {
		if i < rA {
			s.ds[i] = i
		} else {
			str := fmt.Sprintf("Register %s: %%d", d_names[i])
			_, err := fmt.Fscanf(fd, str, &s.ds[i])
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

	return s
}

func dump_insn(s state, i int) {
	fmt.Printf("  %s\t%s\n", i_names[s.cs[i]], d_names[s.cs[i+1]])
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

func main() {
	s := readInput("input.txt")

	dump(s)
	res := ""

	for i := 0; i < len(s.cs); {
		op := s.cs[i+1]
		switch s.cs[i] {
		case adv:
			s.ds[rA] >>= s.ds[op]
		case bdv:
			s.ds[rB] = s.ds[rA] >> s.ds[op]
		case cdv:
			s.ds[rC] = s.ds[rA] >> s.ds[op]
		case bxl:
			s.ds[rB] ^= int(op)
		case bst:
			s.ds[rB] = s.ds[op] % 8
		case jnz:
			if s.ds[rA] != 0 {
				i = int(op)
				continue
			}
		case bxc:
			s.ds[rB] ^= s.ds[rC]
		case out:
			res += fmt.Sprintf("%d,", s.ds[op]%8)
		default:
			panic("Impossible opcode")
		}
		i += 2
	}

	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	fmt.Println("Result:", res)
}
