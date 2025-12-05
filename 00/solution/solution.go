package solution

import (
	"strconv"
	"strings"
)

func ComputeSolutionOne(data []byte) int64 {
	strdta := string(data)

	lines := strings.Split(strdta, "\n")
	position := 50
	min := 0
	max := 99

	count := 0

	for _, l := range lines {
		dirbyt, cntstr := l[0], l[1:]
		dir := string(dirbyt)
		cnt, err := strconv.ParseInt(cntstr, 10, 64)
		if err != nil {
			panic(err)
		}
		for range cnt {
			switch dir {
			case "L":
				position--
				if position < min {
					position = max
				}
			case "R":
				position++
				if position > max {
					position = min
				}
			}
		}
		if position == 0 {
			count++
		}
	}
	return int64(count)
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}
