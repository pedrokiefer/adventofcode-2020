package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadInitialMap(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`)))

	m := LoadPuzzle(input)
	m.GenerateRule0()

	assert.Equal(t, "4 1 5", m.Rules[0])
	assert.Equal(t, []string{"aaaabb", "aaabab", "abbabb", "abbbab", "aabaab", "aabbbb", "abaaab", "ababbb"}, m.Rule0)
	assert.Equal(t, 2, m.MatchRule0())
}
