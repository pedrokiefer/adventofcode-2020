package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileInputToPasswords(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
`)))

	list := InputToPasswords(input)

	assert.Equal(t, []PasswordInput{
		{
			Policy: PasswordPolicy{
				Low:  1,
				High: 3,
				Char: "a",
			},
			Password: "abcde",
		},
		{
			Policy: PasswordPolicy{
				Low:  1,
				High: 3,
				Char: "b",
			},
			Password: "cdefg",
		},
		{
			Policy: PasswordPolicy{
				Low:  2,
				High: 9,
				Char: "c",
			},
			Password: "ccccccccc",
		},
	}, list)
}

func TestCanCheckPassword(t *testing.T) {
	p := PasswordInput{
		Policy: PasswordPolicy{
			Low:  1,
			High: 3,
			Char: "a",
		},
		Password: "abcde",
	}

	result := p.Policy.CheckPassword(p.Password)

	assert.True(t, result)
}

func TestCanCountValidPasswords(t *testing.T) {
	input := []PasswordInput{
		{
			Policy: PasswordPolicy{
				Low:  1,
				High: 3,
				Char: "a",
			},
			Password: "abcde",
		},
		{
			Policy: PasswordPolicy{
				Low:  1,
				High: 3,
				Char: "b",
			},
			Password: "cdefg",
		},
		{
			Policy: PasswordPolicy{
				Low:  2,
				High: 9,
				Char: "c",
			},
			Password: "ccccccccc",
		},
	}

	result := CountValidPasswords(input)

	assert.Equal(t, 2, result)
}
