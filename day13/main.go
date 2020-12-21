package main

import "bufio"

//import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

func main() {

	testdata1 := readFile("test1")
	log.Printf("Testdata1 result p1: %d", findBusP1(testdata1))

	realdata := readFile("input")
	log.Printf("Realdata result p1: %d", findBusP1(realdata))

	log.Printf("Testdata1 result p2: %d", findBusP2(testdata1))

	testdata2 := readFile("test2")
	log.Printf("Testdata2 result p2: %d", findBusP2(testdata2))

	testdata3 := readFile("test3")
	log.Printf("Testdata3 result p2: %d", findBusP2(testdata3))

	testdata4 := readFile("test4")
	log.Printf("Testdata4 result p2: %d", findBusP2(testdata4))

	testdata5 := readFile("test5")
	log.Printf("Testdata5 result p2: %d", findBusP2(testdata5))

	testdata6 := readFile("test6")
	log.Printf("Testdata6 result p2: %d", findBusP2(testdata6))

	log.Printf("Realdata result p2: %d", findBusP2(realdata))
}

func findBusP1(data []string) (sol int) {
	var buses []int

	ts, _ := strconv.Atoi(data[0])
	bs := strings.Split(data[1], ",")

	for _, b := range bs {
		if b != "x" {
			busid, _ := strconv.Atoi(b)
			buses = append(buses, busid)
		}
	}

	//log.Print(buses)

	for i := 0; true; i++ {
		for _, b := range buses {
			if (ts+i)%b == 0 {
				//log.Printf("Take bus %d at %d", b, ts + i)
				return i * b
			}
		}
	}

	return

}

func findBusP2(data []string) (sol int) {
	buses := make(map[int]int)

	bs := strings.Split(data[1], ",")

	for i, b := range bs {
		if b == "x" {
			continue
		}
		bid, _ := strconv.Atoi(b)
		buses[bid] = i
	}

	log.Print(buses)

	var hid, ho int = 0, 0

	for k, v := range buses {
		if k > hid {
			hid = k
			ho = v
		}
	}

	_ = ho

	for i := 1; true; i++ {
		c := 0
		ts := (hid * i) - ho
		for b, o := range buses {
			if (ts+o)%b == 0 {
				c += 1
			}
		}
		if c == len(buses) {
			sol = ts
			return
		}
	}

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
