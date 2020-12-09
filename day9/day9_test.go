package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadMessage(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)))

	msg := LoadMessage(input)

	assert.Equal(t, 25, msg[3])
	assert.Equal(t, 576, msg[len(msg)-1])
}

func TestCanFindWrongNumber(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)))

	msg := LoadMessage(input)
	assert.Equal(t, 25, msg[3])
	assert.Equal(t, 576, msg[len(msg)-1])

	wrong := FindWrongNumber(msg, 5)
	assert.Equal(t, 127, wrong)
}

func TestCanFindContiguousSum(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`)))

	msg := LoadMessage(input)
	r := FindContiguousSum(msg, 127)

	assert.Equal(t, []int{15, 25, 47, 40}, r)

	s := SumMinMax(r)
	assert.Equal(t, 62, s)
}

func TestCombine(t *testing.T) {
	c := Combine([]int{1, 2, 3, 4}, 2)

	assert.Equal(t, [][]int{
		{1, 2},
		{1, 3},
		{1, 4},
		{2, 3},
		{2, 4},
		{3, 4},
	}, c)
}

func TestSumCombinations(t *testing.T) {
	c := SumCombination([]int{1, 2, 3, 4}, 2)

	assert.Equal(t, []int{3, 4, 5, 6, 7}, c)
}
