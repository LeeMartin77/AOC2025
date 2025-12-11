package solution

import (
	"slices"
	"strings"
)

func path(devices map[string][]string, position string, visited []string, target string) int {
	if position == target {
		return 1
	}
	culm := 0
	for _, dest := range devices[position] {
		if slices.Contains(visited, dest) {
			//loop
			continue
		}
		culm += path(devices, dest, append(visited, position), target)
	}
	return culm
}

func ComputeSolutionOne(data []byte) int64 {
	devices := map[string][]string{}
	for _, ln := range strings.Split(string(data), "\n") {
		prs := strings.Split(ln, ":")
		id := prs[0]
		dests := strings.Split(strings.Trim(prs[1], " "), " ")
		devices[id] = dests
	}
	totalpaths := path(devices, "you", []string{}, "out")
	return int64(totalpaths)
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}
