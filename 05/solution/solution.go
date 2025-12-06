package solution

import (
	"strconv"
	"strings"
)

type unspoiltranges struct {
	ranges [][]int64
}

func (u *unspoiltranges) IsUnspoilt(p int64) bool {
	for _, r := range u.ranges {
		if p >= r[0] && p <= r[1] {
			return true
		}
	}
	return false
}

func ComputeSolutionOne(data []byte) int64 {
	ingredients := false
	unspoilt := unspoiltranges{}
	unspoilting := []int64{}
	for _, ln := range strings.Split(string(data), "\n") {
		if ln == "" {
			ingredients = true
			continue
		}
		if !ingredients {
			str, end := strings.Split(ln, "-")[0], strings.Split(ln, "-")[1]
			strint, _ := strconv.ParseInt(str, 10, 64)
			eint, _ := strconv.ParseInt(end, 10, 64)
			unspoilt.ranges = append(unspoilt.ranges, []int64{strint, eint})
		} else {
			strint, _ := strconv.ParseInt(ln, 10, 64)
			if unspoilt.IsUnspoilt(strint) {
				unspoilting = append(unspoilting, strint)
			}
		}
	}
	return int64(len(unspoilting))
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}
