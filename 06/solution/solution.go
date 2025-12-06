package solution

import (
	"strconv"
	"strings"
)

func ComputeSolutionOne(data []byte) int64 {
	strin := string(data)
	column_indexes := []int{}

	lines := strings.Split(strin, "\n")
outer:
	for i, c := range lines[0] {
		if c == ' ' {
			for _, ln := range lines {
				if ln[i] != ' ' {
					continue outer
				}
			}
			column_indexes = append(column_indexes, i)
		}
	}
	column_indexes = append(column_indexes, len(lines[0]))
	columns := [][]string{}
	prev := 0
	for i, c := range column_indexes {
		columns = append(columns, []string{})
		for _, line := range lines {
			chunk := line[prev:c]
			trimmed := strings.Trim(chunk, " ")
			columns[i] = append(columns[i], trimmed)
		}
		prev = c
	}
	total := int64(0)
	for _, col := range columns {
		operator := col[len(col)-1]
		numbersstr := col[:len(col)-1]
		numbers := []int64{}
		for _, nums := range numbersstr {
			n, _ := strconv.ParseInt(nums, 10, 64)
			numbers = append(numbers, n)
		}
		intotal := int64(0)
		if operator == "*" {
			for _, n := range numbers {
				if intotal == 0 {
					intotal = n
				} else {
					intotal = intotal * n
				}
			}
		}
		if operator == "+" {
			for _, n := range numbers {
				intotal += n
			}
		}
		total += intotal
	}
	return total
}

func ComputeSolutionTwo(data []byte) int64 {
	panic("unimplemented")
}
