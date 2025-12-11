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

func path_prereq(devices map[string][]string, position string, visited []string, target string) int {
	if position == target {
		if slices.Contains(visited, "fft") && slices.Contains(visited, "dac") {
			return 1
		}
		return 0
	}
	culm := 0
	for _, dest := range devices[position] {
		if slices.Contains(visited, dest) {
			//loop
			continue
		}
		culm += path_prereq(devices, dest, append(visited, position), target)
	}
	return culm
}

func ComputeSolutionTwo(data []byte) int64 {
	devices := map[string][]string{}
	for _, ln := range strings.Split(string(data), "\n") {
		prs := strings.Split(ln, ":")
		id := prs[0]
		dests := strings.Split(strings.Trim(prs[1], " "), " ")
		devices[id] = dests
	}
	totalpaths := path_prereq(devices, "svr", []string{}, "out")
	return int64(totalpaths)
}
