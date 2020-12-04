package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadIdentities(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)))

	m := LoadIdentities(input)

	assert.Equal(t, 4, len(m))
	assert.Equal(t, "860033327", m[0].PassportID)
	assert.Equal(t, "028048884", m[1].PassportID)
	assert.Equal(t, "760753108", m[2].PassportID)
	assert.Equal(t, "166559648", m[3].PassportID)
}

func TestCanValidadeIdentities(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)))

	m := LoadIdentities(input)

	assert.Equal(t, 4, len(m))
	assert.Equal(t, "860033327", m[0].PassportID)
	assert.True(t, m[0].Validate())
	assert.Equal(t, "028048884", m[1].PassportID)
	assert.False(t, m[1].Validate())
	assert.Equal(t, "760753108", m[2].PassportID)
	assert.True(t, m[2].Validate())
	assert.Equal(t, "166559648", m[3].PassportID)
	assert.False(t, m[3].Validate())
}
