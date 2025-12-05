package main_test

import (
	"testing"

	"github.com/LeeMartin77/AOC2024/02/solution"
	"github.com/stretchr/testify/assert"
)

func TestPhaseOne(t *testing.T) {
	teststring := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	res := solution.ComputeSolutionOne([]byte(teststring))
	assert.Equal(t, int64(1227775554), res)
}

func TestPhaseTwo(t *testing.T) {
	teststring := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	res := solution.ComputeSolutionTwo([]byte(teststring))
	assert.Equal(t, int64(4174379265), res)
}
