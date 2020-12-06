package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCountUniqueDeclarations(t *testing.T) {
	input := `abcx
abcy
abcz`

	n := CountUniqueDeclarations(input, false)

	assert.Equal(t, 6, n)
}

func TestCanLoadDeclarations(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`abc

a
b
c

ab
ac

a
a
a
a

b`)))

	decs := LoadDeclarations(input, false)

	assert.Equal(t, 11, decs)
}

func TestCanLoadDeclarationsGlobal(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`abc

a
b
c

ab
ac

a
a
a
a

b`)))

	decs := LoadDeclarations(input, true)

	assert.Equal(t, 6, decs)
}
