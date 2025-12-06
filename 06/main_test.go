package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/06/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(4277556), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}
