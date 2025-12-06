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

func (u *unspoiltranges) TotalUnspoiltIds() int64 {
	// consolidate ranges
	consolidated_ranges := [][]int64{}
outer:
	for _, r := range u.ranges {
		for i, cr := range consolidated_ranges {
			consolidated := false
			if r[1] <= cr[1] && r[1] >= cr[0] && r[0] < cr[0] {
				// extend start of range
				consolidated_ranges[i][0] = r[0]
				consolidated = true
			}

			if r[0] >= cr[0] && r[0] <= cr[1] && r[1] > cr[1] {
				// extend end of range
				consolidated_ranges[i][1] = r[1]
				consolidated = true
			}
			if consolidated {
				continue outer
			}
		}
		consolidated_ranges = append(consolidated_ranges, r)
	}
	total := int64(0)
	for _, cr := range consolidated_ranges {
		total += (cr[1] - cr[0])
	}
	return total
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
	unspoilt := unspoiltranges{}
	for _, ln := range strings.Split(string(data), "\n") {
		if ln == "" {
			break
		}
		str, end := strings.Split(ln, "-")[0], strings.Split(ln, "-")[1]
		strint, _ := strconv.ParseInt(str, 10, 64)
		eint, _ := strconv.ParseInt(end, 10, 64)
		unspoilt.ranges = append(unspoilt.ranges, []int64{strint, eint})

	}
	return unspoilt.TotalUnspoiltIds()
}
