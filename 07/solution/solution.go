package solution

import (
	"fmt"
	"strings"
)

func beamAtCord(beams [][]int, x, y int) bool {
	for _, beam := range beams {
		if beam[0] == x && beam[1] == y {
			return true
		}
	}
	return false
}

func renderMap(splitters [][]rune, beams [][]int) {
	lines := make([]string, len(splitters[0]))
	for x, ys := range splitters {
		for y, rn := range ys {
			if x == 0 {
				if beamAtCord(beams, x, y) {

					lines[y] = string('|')
				} else {

					lines[y] = string(rn)
				}
			} else {
				if beamAtCord(beams, x, y) {

					lines[y] += string('|')
				} else {

					lines[y] += string(rn)
				}
			}
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

func ComputeSolutionOne(data []byte) int64 {
	mapstr := string(data)
	splitters := [][]rune{}
	lines := strings.Split(mapstr, "\n")
	beams := [][]int{}
	maxy := 0
	for y, ln := range lines {
		for x, v := range ln {
			if y == 0 {
				splitters = append(splitters, []rune{})
			}
			splitters[x] = append(splitters[x], v)

			if v == 'S' {
				beams = append(beams, []int{x, y})
			}
		}
		maxy = y
	}

	splits := 0

	for len(beams) > 0 {
		new_beams := [][]int{}
		for _, b := range beams {
			b[1] += 1
			if splitters[b[0]][b[1]] == '^' {
				didsplit := false
				if !beamAtCord(new_beams, b[0]-1, b[1]) {
					didsplit = true
					new_beams = append(new_beams, []int{b[0] - 1, b[1]})
				}
				if !beamAtCord(new_beams, b[0]+1, b[1]) {
					didsplit = true

					new_beams = append(new_beams, []int{b[0] + 1, b[1]})
				}
				if didsplit {

					splits++
				}
			} else if b[1] < maxy {
				new_beams = append(new_beams, b)
			}
		}
		beams = new_beams
	}

	return int64(splits)
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}
