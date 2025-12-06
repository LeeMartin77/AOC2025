package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/03/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `987654321111111
811111111111119
234234234234278
818181911112111`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(357), res)
}

func xTestPhaseTwo(t *testing.T) {
	teststring := ``
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(0), res)
}
