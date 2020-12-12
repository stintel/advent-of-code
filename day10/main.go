package main

import "bufio"
import "log"
import "os"
import "sort"
import "strconv"

func main() {
	testdata1 := readFile("test1")
	testdata2 := readFile("test2")
	realdata := readFile("input")

	log.Print("Part 1 test data1")
	t1d1, t1d3 := chainAdapters(testdata1)
	log.Printf("diff1: %d, diff3: %d", t1d1, t1d3)

	log.Print("Part 1 test data2")
	t2d1, t2d3 := chainAdapters(testdata2)
	log.Printf("diff1: %d, diff3: %d", t2d1, t2d3)

	log.Print("Part 1 input data")
	rd1, rd3 := chainAdapters(realdata)
	log.Printf("diff1: %d, diff3: %d, product: %d", rd1, rd3, rd1 * rd3)

	log.Print("Part 2 test data")

	log.Print("Part 2 real data")

}

func chainAdapters(adapters sort.IntSlice) (diff1, diff3 int) {
	diff1 = 0
	diff3 = 0

	adapters.Sort()

	if adapters[0] == 1 {
		diff1 += 1
	}
	if adapters[0] == 3 {
		diff3 += 1
	}

	for i, adapter := range adapters {
		if i == 0 {
			continue
		}
		if adapter - adapters[i-1] == 1 {
			diff1 += 1
		}
		if adapter - adapters[i-1] == 3 {
			diff3 += 1
		}
	}

	diff3 += 1

	return
}

func readFile(path string) (data sort.IntSlice) {
	var t int

        file, err := os.Open(path)
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
		t, _ = strconv.Atoi(scanner.Text())
                data = append(data, t)
        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }

        return
}
