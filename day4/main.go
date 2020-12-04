package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Identity struct
type Identity struct {
	BirthYear      string // byr
	IssueYear      string // iyr
	ExpirationYear string // eyr
	Height         string // hgt
	HairColor      string // hcl
	EyeColor       string // ecl
	PassportID     string // pid
	CountryID      string // cid
}

var ValidEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

func StringInArray(input string, array []string) bool {
	for _, v := range array {
		if input == v {
			return true
		}
	}
	return false
}

func (i *Identity) Validate() bool {
	if i.BirthYear == "" || i.IssueYear == "" || i.ExpirationYear == "" || i.Height == "" || i.HairColor == "" || i.EyeColor == "" || i.PassportID == "" {
		return false
	}

	if len(i.BirthYear) != 4 {
		return false
	}
	v, err := strconv.Atoi(i.BirthYear)
	if err != nil {
		return false
	}
	if v < 1920 || v > 2002 {
		return false
	}

	if len(i.IssueYear) != 4 {
		return false
	}
	v, err = strconv.Atoi(i.IssueYear)
	if err != nil {
		return false
	}
	if v < 2010 || v > 2020 {
		return false
	}

	if len(i.ExpirationYear) != 4 {
		return false
	}
	v, err = strconv.Atoi(i.ExpirationYear)
	if err != nil {
		return false
	}
	if v < 2020 || v > 2030 {
		return false
	}

	if strings.Contains(i.Height, "cm") {
		s := strings.ReplaceAll(i.Height, "cm", "")
		v, err = strconv.Atoi(s)
		if err != nil {
			return false
		}
		if v < 150 || v > 193 {
			return false
		}
	} else if strings.Contains(i.Height, "in") {
		s := strings.ReplaceAll(i.Height, "in", "")
		v, err = strconv.Atoi(s)
		if err != nil {
			return false
		}
		if v < 59 || v > 76 {
			return false
		}
	} else {
		return false
	}

	if strings.HasPrefix(i.HairColor, "#") {
		s := string(i.HairColor[1:])
		_, err = strconv.ParseInt(s, 16, 64)
		if err != nil {
			return false
		}
	} else {
		return false
	}

	if !StringInArray(i.EyeColor, ValidEyeColors) {
		return false
	}

	if len(i.PassportID) != 9 {
		return false
	}

	return true
}

func LineToIdentity(buf string) *Identity {
	i := Identity{}
	for _, kv := range strings.Split(buf, " ") {
		parts := strings.Split(kv, ":")
		switch parts[0] {
		case "byr":
			i.BirthYear = parts[1]
		case "iyr":
			i.IssueYear = parts[1]
		case "eyr":
			i.ExpirationYear = parts[1]
		case "hgt":
			i.Height = parts[1]
		case "hcl":
			i.HairColor = parts[1]
		case "ecl":
			i.EyeColor = parts[1]
		case "pid":
			i.PassportID = parts[1]
		case "cid":
			i.CountryID = parts[1]
		}
	}
	return &i
}

func LoadIdentities(input io.ReadCloser) []*Identity {
	identities := []*Identity{}
	s := bufio.NewScanner(input)
	defer input.Close()

	buf := ""
	for s.Scan() {
		txt := s.Text()
		if txt == "" {
			identities = append(identities, LineToIdentity(buf))
			buf = ""
			continue
		}
		txt = strings.TrimSpace(txt)
		buf = buf + " " + txt
	}
	if buf != "" {
		identities = append(identities, LineToIdentity(buf))
	}
	return identities
}

func CountValid(identities []*Identity) int {
	count := 0
	for _, i := range identities {
		if i.Validate() {
			count++
		}
	}
	return count
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	m := LoadIdentities(f)
	c := CountValid(m)

	fmt.Printf("Result: %d\n", c)
}
