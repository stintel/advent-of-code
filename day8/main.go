package main

import "bufio"
import "log"
import "os"
import "strconv"
import "strings"

type Ins struct {
	Arg int
	Op string
}

func detectLoop(data []string) (isLoop bool, acc int) {
	acc = 0
	isLoop = false
	hist := make([]int, len(data))

	i := 0
	for i < len(data) {
		if hist[i] != 1 {
			hist[i] = 1
			i += execOp(&acc, parseIns(data[i]))
		} else {
			isLoop = true
			return
		}

	}

	return
}

func execOp(acc *int, ins Ins) (next int) {
	next = 1
	switch {
		case ins.Op == "acc":
			*acc += ins.Arg
		case ins.Op == "jmp":
			next = ins.Arg
	}
	return
}

func parseIns(s string) (ins Ins) {
	t := strings.Fields(s)

	ins.Op = t[0]
	ins.Arg, _ = strconv.Atoi(t[1])

	return
}

func readFile(path string) (data []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	var datap2 []string

	data := readFile("input")

	_, ins := detectLoop(data)
	log.Printf("Loop detected at instruction %d", ins)

	i := 0
	for i < len(data) {
		datap2 = nil

		for j, line := range(data) {
			if i == j {
				// strings.Contains(line, "jmp") {
				line = strings.Replace(line, "jmp", "noop", 1)
			}
			datap2 = append(datap2, line)
		}

		i++

		isLoop, acc := detectLoop(datap2)
		if !isLoop {
			log.Printf("Value of acc: %d", acc)
		}
	}
}
