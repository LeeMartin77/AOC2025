package solution

import (
	"strconv"
	"strings"
)

func ComputeSolutionOne(data []byte) int64 {
	banks := strings.Split(string(data), "\n")
	totaljoltage := int64(0)

	for _, bank := range banks {
		maxjoltage := int64(0)
		for i, s := range bank {
			if i == len(bank)-1 {
				break
			}
			for _, ss := range bank[i+1:] {

				v, _ := strconv.ParseInt(string(s)+string(ss), 10, 64)

				if v > maxjoltage {
					maxjoltage = v
				}
			}
		}
		totaljoltage += maxjoltage
	}
	return totaljoltage
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}
