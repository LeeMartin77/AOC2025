package solution

import (
	"strconv"
	"strings"
)

func ComputeSolutionOne(data []byte) int64 {
	ingredients := false
	unspoilt := map[int64]bool{}
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
			for strint <= eint {
				unspoilt[strint] = true
				strint++
			}
		} else {
			strint, _ := strconv.ParseInt(ln, 10, 64)
			if unspoilt[strint] {
				unspoilting = append(unspoilting, strint)
			}
		}
	}
	return int64(len(unspoilting))
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}
