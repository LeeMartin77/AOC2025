package solution

import (
	"fmt"
	"strings"
	"sync"
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

func setup(mapstr string) ([][]rune, [][]int, int) {
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
	return splitters, beams, maxy
}

func ComputeSolutionOne(data []byte) int64 {
	mapstr := string(data)
	splitters, beams, maxy := setup(mapstr)

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

func progressBeam(wg *sync.WaitGroup, splitters [][]rune, b []int, timecount chan int) {
	if b[1] >= len(splitters[0]) {
		wg.Done()
		return
	}

	if splitters[b[0]][b[1]] == '^' {
		timecount <- 1
		wg.Add(2)
		wg.Done()
		go progressBeam(wg, splitters, []int{b[0] - 1, b[1] - 1}, timecount)
		go progressBeam(wg, splitters, []int{b[0] + 1, b[1] - 1}, timecount)
	} else {
		go progressBeam(wg, splitters, []int{b[0], b[1] + 1}, timecount)
	}
}

func ComputeSolutionTwo(data []byte) int64 {
	mapstr := string(data)
	splitters, beams, _ := setup(mapstr)

	start := beams[0]
	timelines := 0

	wg := sync.WaitGroup{}
	wg.Add(1)
	timecount := make(chan int)

	go func() {
		for cnt := range timecount {
			timelines += cnt
		}
	}()

	go progressBeam(&wg, splitters, start, timecount)

	wg.Wait()
	close(timecount)

	return int64(timelines + 1)
}
