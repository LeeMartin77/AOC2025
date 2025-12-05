package solution

import (
	"fmt"
	"strconv"
	"strings"
)

func ComputeSolutionOne(data []byte) int64 {
	datastr := string(data)
	chnks := strings.Split(datastr, ",")
	ranges := [][]int64{}
	for _, chnk := range chnks {
		ends := strings.Split(chnk, "-")
		sint, _ := strconv.ParseInt(ends[0], 10, 64)
		eint, _ := strconv.ParseInt(ends[1], 10, 64)

		ranges = append(ranges, []int64{sint, eint})
	}
	invalidids := []int64{}
	for _, rng := range ranges {
		for i := range rng[1] - rng[0] + 1 {
			vl := rng[0] + i
			stid := fmt.Sprintf("%d", vl)
			if len(stid)%2 == 0 {
				fhlf := stid[:len(stid)/2]
				shlf := stid[len(stid)/2:]
				if fhlf == shlf {
					invalidids = append(invalidids, vl)
				}
			}
		}
	}
	sum := int64(0)
	for _, vl := range invalidids {
		sum += vl
	}
	return sum
}

func ComputeSolutionTwo(data []byte) int64 {
	datastr := string(data)
	chnks := strings.Split(datastr, ",")
	ranges := [][]int64{}
	for _, chnk := range chnks {
		ends := strings.Split(chnk, "-")
		sint, _ := strconv.ParseInt(ends[0], 10, 64)
		eint, _ := strconv.ParseInt(ends[1], 10, 64)

		ranges = append(ranges, []int64{sint, eint})
	}
	invalidids := []int64{}
	for _, rng := range ranges {
		for i := range rng[1] - rng[0] + 1 {
			vl := rng[0] + i
			stid := fmt.Sprintf("%d", vl)
			modulos := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

		outer:
			for _, m := range modulos {

				if len(stid)%m == 0 {
					s, e := 0, len(stid)/m
					gp := e
					prt := stid[s:e]
					for s < len(stid) {
						if stid[s:e] != prt {
							continue outer
						}
						s += gp
						e += gp
					}

					invalidids = append(invalidids, vl)
					break outer
				}
			}
		}
	}
	sum := int64(0)
	for _, vl := range invalidids {
		sum += vl
	}
	return sum
}
