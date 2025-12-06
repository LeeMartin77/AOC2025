package solution

import (
	"strings"
)

func ComputeSolutionOne(data []byte) int64 {
	mapstr := string(data)
	mp := map[int]map[int]bool{}
	lines := strings.Split(mapstr, "\n")
	for y, ln := range lines {
		for x, v := range ln {
			if y == 0 {
				mp[x] = map[int]bool{}
			}
			if v == '@' {
				mp[x][y] = true
			}
		}
	}
	adjlmt := 4
	moveables := [][]int{}
	for x, mpln := range mp {
		for y := range mpln {
			adj := 0
			for _, xoff := range []int{-1, 0, 1} {

				for _, yoff := range []int{-1, 0, 1} {

					if xoff == 0 && yoff == 0 {
						continue
					}
					if mp[x+xoff][y+yoff] {
						adj += 1
					}
				}
			}
			if adj < adjlmt {
				moveables = append(moveables, []int{x, y})
			}
		}
	}
	return int64(len(moveables))
}

func ComputeSolutionTwo(data []byte) int64 {
	mapstr := string(data)
	mp := map[int]map[int]bool{}
	lines := strings.Split(mapstr, "\n")
	for y, ln := range lines {
		for x, v := range ln {
			if y == 0 {
				mp[x] = map[int]bool{}
			}
			if v == '@' {
				mp[x][y] = true
			}
		}
	}
	adjlmt := 4

	count := int64(0)
	moved := true
	for moved {
		moved = false
		for x, mpln := range mp {
			for y := range mpln {
				adj := 0
				for _, xoff := range []int{-1, 0, 1} {

					for _, yoff := range []int{-1, 0, 1} {

						if xoff == 0 && yoff == 0 {
							continue
						}
						if mp[x+xoff][y+yoff] {
							adj += 1
						}
					}
				}
				if adj < adjlmt {
					count += 1
					delete(mp[x], y)
					moved = true
				}
			}
		}
	}

	return count
}
