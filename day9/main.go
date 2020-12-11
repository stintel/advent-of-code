package main

import "bufio"
import "log"
import "os"
import "strconv"

func main() {
	testdata := readFile("test")
	realdata := readFile("input")

	log.Print("Part 1 test data")
	invalidt := findInvalid(testdata, 5)

	log.Print("Part 1 input data")
	invalidr := findInvalid(realdata, 25)

	log.Print("Part 2 test data")
	contiguoust := findContiguous(testdata, invalidt)
	sumt := addLowestHighest(contiguoust)
	log.Print(sumt)

	log.Print("Part 2 real data")
	contiguousr := findContiguous(realdata, invalidr)
	sumr := addLowestHighest(contiguousr)
	log.Print(sumr)

}

func addLowestHighest(numbers []int) (sum int) {
	var h, l int = 0, 0

	l = numbers[0]

	for _, i := range(numbers) {
		if i < l {
			l = i
		} else if i > h {
			h = i
		}
	}

	sum = l + h

	return
}

func findContiguous(data []string, num int) (contiguous []int) {
	var i, j, sum, t, x int = 0, 0, 0, 0, 0

	for i < len(data) {
		//log.Printf("Starting at line %d: %s", i, data[i])
		x, _ = strconv.Atoi(data[i])
		if x == num {
			break
		}
		contiguous = append(contiguous, x)
		j = i
		for sum < num {
			t, _ = strconv.Atoi(data[j])
			sum += t
			//log.Printf("Adding line %d: %d, sum is now %d", j, t, sum)
			j++
			contiguous = append(contiguous, t)
		}
		if sum == num {
			return
		}
		contiguous = nil
		sum = 0
		i++
	}

	return
}

func findInvalid(data []string, preamble int) (invalid int) {
	var num, x, y, z int
	var valid bool

	i := preamble

	for i < len(data) {
		num, _ = strconv.Atoi(data[i])
		valid = false
		//log.Printf("Testing line %d: %d", i, num)
		for j := 0; j < preamble; j++ {
			for k := j + 1; k < preamble; k++ {
				x, _ = strconv.Atoi(data[i - preamble + j])
				y, _ = strconv.Atoi(data[i - preamble + k])
				z = x + y
				//log.Printf("Trying %d + %d = %d", x, y, z)

				if num == z {
					//log.Printf("Valid: %d + %d = %d", x, y, z)
					valid = true
				}
			}
		}
		if !valid {
			log.Printf("Line %d is invalid: %d", i, num)
			invalid = num
		}
		i++
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
