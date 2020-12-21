package main

import "bufio"
//import "fmt"
import "log"
import "os"
import "strconv"
import "strings"

func main() {
	testdata1 := readFile("test1")
	td1r := game(toSlice(testdata1), 30000000)
	log.Printf("Testdata1 result: %d", td1r)

	testdata2 := readFile("test2")
	td2r := game(toSlice(testdata2), 30000000)
	log.Printf("Testdata2 result: %d", td2r)

	testdata3 := readFile("test3")
	td3r := game(toSlice(testdata3), 30000000)
	log.Printf("Testdata3 result: %d", td3r)

	testdata4 := readFile("test4")
	td4r := game(toSlice(testdata4), 30000000)
	log.Printf("Testdata4 result: %d", td4r)

	testdata5 := readFile("test5")
	td5r := game(toSlice(testdata5), 30000000)
	log.Printf("Testdata5 result: %d", td5r)

	testdata6 := readFile("test6")
	td6r := game(toSlice(testdata6), 30000000)
	log.Printf("Testdata6 result: %d", td6r)

	testdata7 := readFile("test7")
	td7r := game(toSlice(testdata7), 30000000)
	log.Printf("Testdata7 result: %d", td7r)

	realdata := readFile("input")
	rdr := game(toSlice(realdata), 30000000)
	log.Printf("Real data result: %d", rdr)
}

func game(ndata []int, times int) (n int) {
	var res = make([]int, times)
	var places = make(map[int]int)

	copy(res, ndata)

	for i, d := range(ndata) {
		places[d] = i
	}
	//log.Print(places)

	cur := len(ndata)

	for i := cur; i < times; i++ {
		if val, found := places[res[i-1]]; found {
			res[i] = i - 1 - val
		} else {
			res[i] = 0
		}
		places[res[i-1]] = i - 1
		//log.Print(res)
	}

	n = res[times-1]

	return
}

func toSlice(data []string) (ndata []int) {
	for _, line := range data {
		ts := strings.Split(line, ",")
		for i, _ := range ts {
			ti, _ := strconv.Atoi(ts[i])
			ndata = append(ndata, ti)
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
