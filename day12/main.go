package main

import "bufio"
import "log"
import "os"
import "strings"
import "strconv"

type Position struct {
	D string
	X int
	Y int
}

func main() {
	testdata := readFile("test")
	realdata := readFile("input")

	sp := Position{
		D: "E",
		X: 0,
		Y: 0,
	}

	log.Print("Part 1 test data")
	ept1 := navigate(testdata, sp)
	log.Print(calcMD(ept1))

	log.Print("Part 1 input data")
	epr1 := navigate(realdata, sp)
	log.Print(calcMD(epr1))

	swp := Position{
		D: "X",
		X: 10,
		Y: 1,
	}

	log.Print("Part 2 test data")
	ept2 := navigateWithWaypoint(testdata, sp, swp)
	log.Print(calcMD(ept2))

	log.Print("Part 2 real data")
	epr2 := navigateWithWaypoint(realdata, sp, swp)
	log.Print(calcMD(epr2))

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func calcMD(ep Position) (md int) {
	md = abs(ep.X) + abs(ep.Y)
	return
}

func navigate(data []string, startpos Position) (pos Position) {
	pos = startpos

	for _, line := range data {
		action, steps := parseIns(line)
		move(&pos, action, steps)
		//log.Print(pos)
	}

	return
}

func navigateWithWaypoint(data []string, startpos, startwp Position) (pos Position) {
	pos = startpos
	wp := startwp

	for _, line := range data {
		action, steps := parseIns(line)
		switch {
		case action == "N":
			fallthrough
		case action == "E":
			fallthrough
		case action == "S":
			fallthrough
		case action == "W":
			move(&wp, action, steps)
		case action == "L":
			fallthrough
		case action == "R":
			rotateWaypoint(&wp, action, steps)
		case action == "F":
			moveToWaypoint(&pos, wp, steps)
		}
	}

	return
}

func parseIns(ins string) (action string, steps int) {
	action = strings.Trim(ins, "0123456789")
	steps, _ = strconv.Atoi(strings.TrimLeft(ins, action))

	return
}

func move(pos *Position, action string, steps int) {
	switch {
	case action == "N":
		fallthrough
	case action == "E":
		fallthrough
	case action == "S":
		fallthrough
	case action == "W":
		fallthrough
	case action == "F":
		moveDir(pos, action, steps)
	case action == "L":
		fallthrough
	case action == "R":
		switchDirection(pos, action, steps)
	}

	return
}

func moveDir(pos *Position, dir string, steps int) {
	if dir == "F" {
		moveDir(pos, pos.D, steps)
	}
	switch {
	case dir == "N":
		pos.Y += steps
	case dir == "S":
		pos.Y -= steps
	case dir == "E":
		pos.X += steps
	case dir == "W":
		pos.X -= steps
	}
}

func moveToWaypoint(pos *Position, wp Position, steps int) {
	pos.X += (wp.X * steps)
	pos.Y += (wp.Y * steps)
}

func rotateWaypoint(wp *Position, dir string, deg int) {
	var x, y int

	switch deg {
	case 90:
		if dir == "L" {
			x = -wp.Y
			y = wp.X
		}
		if dir == "R" {
			x = wp.Y
			y = -wp.X
		}
	case 270:
		if dir == "L" {
			x = wp.Y
			y = -wp.X
		}
		if dir == "R" {
			x = -wp.Y
			y = wp.X
		}
	case 180:
		x = -wp.X
		y = -wp.Y
	}

	wp.X = x
	wp.Y = y
}

func switchDirection(pos *Position, dir string, deg int) {
	if deg > 360 {
		deg %= 360
	}
	if dir == "L" {
		deg = 0 - deg
	}
	switch {
	case pos.D == "N":
		deg = 0 + deg
	case pos.D == "S":
		deg = 180 + deg
	case pos.D == "E":
		deg = 90 + deg
	case pos.D == "W":
		deg = 270 + deg
	}

	deg %= 360
	if deg < 0 {
		deg = 360 - abs(deg)
	}

	switch deg {
	case 0:
		pos.D = "N"
	case 90:
		pos.D = "E"
	case 180:
		pos.D = "S"
	case 270:
		pos.D = "W"
	}

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
