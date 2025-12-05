package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/00/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(3), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(6), res)
}
