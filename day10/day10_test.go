package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDifferences(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`16
10
15
5
1
11
7
19
6
12
4`)))

	list := InputToIntList(input)
	diffs := CountDifferences(list)

	sorted := SortJolts(list)
	branchs := CountBranchsRecursive(sorted, 0)
	b2 := CountBranchs(sorted)

	assert.Equal(t, []int{7, 0, 5}, diffs)
	assert.Equal(t, 8, branchs)
	assert.Equal(t, 8, b2)
}

func TestCountDifferencesBig(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`)))

	list := InputToIntList(input)
	diffs := CountDifferences(list)

	sorted := SortJolts(list)
	branchs := CountBranchsRecursive(sorted, 0)
	b2 := CountBranchs(sorted)

	assert.Equal(t, []int{22, 0, 10}, diffs)
	assert.Equal(t, 19208, branchs)
	assert.Equal(t, 19208, b2)
}
