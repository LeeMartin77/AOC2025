package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/05/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `3-5
10-14
16-20
12-18

1
5
8
11
17
32`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(3), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}
