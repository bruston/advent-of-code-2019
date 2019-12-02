package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	OpAdd = iota + 1
	OpMultiply
	OpHalt = 99
)

type machine struct {
	ip   int
	code []int
}

func (m *machine) exec() (int, error) {
	if len(m.code) < 4 {
		return 0, nil
	}
	m.code[1] = 12
	m.code[2] = 2
	for {
		if m.ip+3 >= len(m.code) {
			return 0, errors.New("unexpected EOF")
		}
		switch m.code[m.ip] {
		case OpAdd:
			a, b, dst := m.code[m.ip+1], m.code[m.ip+2], m.code[m.ip+3]
			m.code[dst] = m.code[a] + m.code[b]
			m.ip += 4
		case OpMultiply:
			a, b, dst := m.code[m.ip+1], m.code[m.ip+2], m.code[m.ip+3]
			m.code[dst] = m.code[a] * m.code[b]
			m.ip += 4
		case OpHalt:
			return m.code[0], nil
		default:
			return 0, fmt.Errorf("unsupported operation: %d", m.code[m.ip])
		}
	}
	return m.code[0], nil
}

func main() {
	input := flag.String("f", "input.txt", "path to input.txt")
	flag.Parse()

	b, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	b = bytes.TrimSpace(b)

	split := strings.Split(string(b), ",")
	vm := &machine{code: make([]int, len(split))}
	for i, v := range split {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatalf("input format is incorrect, could not convert %s to int", string(v))
		}
		vm.code[i] = n
	}

	result, err := vm.exec()
	if err != nil {
		log.Fatalf("error processing bytecode: %v", err)
	}
	fmt.Println(result)
}
