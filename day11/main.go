package main

import "bufio"
import "log"
import "os"
import "reflect"

//import "strings"
//import "strconv"

func main() {
	testdata := toNumeric(readFile("test"))
	realdata := toNumeric(readFile("input"))

	log.Print("Part 1 test data")
	log.Print(countOccupied(stoelenDans(testdata, 1, 4)))

	log.Print("Part 1 input data")
	log.Print(countOccupied(stoelenDans(realdata, 1, 4)))

	log.Print("Part 2 test data")
	log.Print(countOccupied(stoelenDans(testdata, -1, 5)))

	log.Print("Part 2 real data")
	log.Print(countOccupied(stoelenDans(realdata, -1, 5)))

}

func countOccupied(data [][]int) (occupied int) {
	occupied = 0

	for i, _ := range data {
		for j, _ := range data[i] {
			if data[i][j] == 1 {
				occupied += 1
			}
		}
	}

	return
}

func toNumeric(data []string) (result [][]int) {
	result = make([][]int, len(data))
	for i := range result {
		result[i] = make([]int, len(data[0]))
	}

	for i, line := range data {
		for j, _ := range line {
			switch {
			case line[j] == '.':
				result[i][j] = -1
			case line[j] == 'L':
				result[i][j] = 0
			case line[j] == '#':
				result[i][j] = 1
			}
		}
	}

	return
}

func stoelenDans(data [][]int, depth int, maxoccupied int) (result [][]int) {
	var i, j, occupied int = 0, 0, 0
	var directions [8]string = [8]string{"n", "ne", "e", "se", "s", "sw", "w", "nw"}

	result = make([][]int, len(data))
	for i := range result {
		result[i] = make([]int, len(data[0]))
	}

	for i, _ = range data {
		for j, _ = range data[i] {
			if data[i][j] == -1 {
				result[i][j] = data[i][j]
				continue
			}
			occupied = 0

			for _, dir := range directions {
				if testDirection(data, j, i, dir, depth) {
					occupied += 1
				}
			}

			if data[i][j] == 0 && occupied == 0 {
				result[i][j] = 1
			} else if data[i][j] == 1 && occupied >= maxoccupied {
				result[i][j] = 0
			} else {
				result[i][j] = data[i][j]
			}
		}
	}

	if !reflect.DeepEqual(data, result) {
		result = stoelenDans(result, depth, maxoccupied)
	}

	return
}

func testDirection(data [][]int, x int, y int, direction string, depth int) (occupied bool) {
	occupied = false

	i := 1

	for true {
		switch {
		case direction == "n":
			if y-i < 0 || data[y-i][x] == 0 {
				return
			}
			if data[y-i][x] == 1 {
				occupied = true
				return
			}

		case direction == "s":
			if y+i >= len(data) || data[y+i][x] == 0 {
				return
			}
			if data[y+i][x] == 1 {
				occupied = true
				return
			}

		case direction == "e":
			if x+i >= len(data[y]) || data[y][x+i] == 0 {
				return
			}
			if data[y][x+i] == 1 {
				occupied = true
				return
			}

		case direction == "w":
			if x-i < 0 || data[y][x-i] == 0 {
				return
			}
			if data[y][x-i] == 1 {
				occupied = true
				return
			}

		case direction == "nw":
			if x-i < 0 || y-i < 0 || data[y-i][x-i] == 0 {
				return
			}
			if data[y-i][x-i] == 1 {
				occupied = true
				return
			}

		case direction == "ne":
			if x+i >= len(data[y]) || y-i < 0 || data[y-i][x+i] == 0 {
				return
			}
			if data[y-i][x+i] == 1 {
				occupied = true
				return
			}

		case direction == "se":
			if x+i >= len(data[y]) || y+i >= len(data) || data[y+i][x+i] == 0 {
				return
			}
			if data[y+i][x+i] == 1 {
				occupied = true
				return
			}

		case direction == "sw":
			if x-i < 0 || y+i >= len(data) || data[y+i][x-i] == 0 {
				return
			}
			if data[y+i][x-i] == 1 {
				occupied = true
				return
			}
		}

		i++

		if depth > 0 && i > depth {
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
