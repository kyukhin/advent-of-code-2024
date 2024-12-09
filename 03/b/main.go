package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func dec_pow(v int) (r int) {
	r = 1
	for i := 0; i < v; i++ {
		r = r * 10
	}
	return r
}

func parse_num(s string) (val int, l int, success bool) {
	val = 0
	digits := make([]int, 0)
	for i := 0; i < 3; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			v, err := strconv.Atoi(s[i : i+1])
			if err != nil {
				panic(fmt.Sprintf("Unable to convert %s: %v", s[i:i+1], v))
			}
			digits = append(digits, v)
		} else {
			break
		}
	}
	l = len(digits)
	if l > 0 {
		val = 0
		for i := 0; i < l; i++ {
			val += dec_pow(l-i-1) * digits[i]
		}
		return val, l, true
	} else {
		return 0, 0, false
	}
}

func main() {
	fd, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("Open %s: %v", fd, err))
	}
	defer fd.Close()

	var c rune
	var data string
	for {
		_, err := fmt.Fscanf(fd, "%c", &c)

		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				break
			} else {
				panic(fmt.Sprintf("Scan failed %s: %v", fd, err))
			}
		}
		data = data + string(c)
	}

	res := 0
	enabled := true
	const lit_len = len("mul(,)")
	for i := 0; i < len(data); i++ {
		if i+4 < len(data) && data[i:i+4] == "do()" {
			enabled = true
			i += len("do()") - 1
		}
		if i+7 < len(data) && data[i:i+7] == "don't()" {
			enabled = false
			i += len("don't") - 1
		}
		if enabled &&
			i+lit_len+2 < len(data) &&
			data[i] == 'm' &&
			data[i+1] == 'u' &&
			data[i+2] == 'l' &&
			data[i+3] == '(' {
			// Match 1st number
			x, dsp1, succ := parse_num(data[i+4 : i+7])
			if succ {
				if data[i+4+dsp1] == ',' {
					// Match 2nd num
					y, dsp2, succ := parse_num(data[i+5+dsp1 : i+5+dsp1+3])
					if succ {
						if data[i+5+dsp1+dsp2] != ')' {
							continue
						}
						res += x * y
						i += lit_len + dsp1 + dsp2 - 1 // Account loop ++
					} else {
						continue
					}
				} else {
					continue
				}
			} else {
				continue
			}
		} else {
			continue
		}
	}

	fmt.Println("Result is:", res)
}
