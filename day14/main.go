package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	MaskMatcher   = regexp.MustCompile("mask = (?P<bitmask>[X10]*)")
	MemoryMatcher = regexp.MustCompile("mem\\[(?P<addr>[0-9]*)\\] = (?P<value>[0-9]*)")
)

type MaskFunction func(uint64) uint64
type AddressFunction func(uint64) []uint64

func GenMaskFunction(mask string) MaskFunction {
	return func(value uint64) uint64 {
		for i := 0; i < len(mask); i++ {
			n := len(mask) - 1 - i
			v := mask[n]
			switch v {
			case 'X':
			case '0':
				value = value &^ (1 << i)
			case '1':
				value = value | (1 << i)
			}
		}
		return value
	}
}

func GenAddressesFunction(mask string) AddressFunction {
	return func(value uint64) []uint64 {
		addresses := []uint64{}
		totalX := 1 << strings.Count(mask, "X")
		for i := 0; i < len(mask); i++ {
			n := len(mask) - 1 - i
			v := mask[n]
			switch v {
			case 'X':
			case '0':
			case '1':
				value = value | (1 << i)
			}
		}
		for variation := 0; variation < totalX; variation++ {
			nv := value
			xCount := 0
			for j := 0; j < len(mask); j++ {
				n := len(mask) - 1 - j
				m := mask[n]
				if m == 'X' {
					bit := variation & (1 << xCount)
					if bit == 0 {
						nv = nv &^ (1 << j)
					} else {
						nv = nv | (1 << j)
					}
					xCount++
				}
			}
			addresses = append(addresses, nv)
		}
		return addresses
	}
}

type Address struct {
	Addr  uint64
	Value uint64
}

type Mem struct {
	Addresses []*Address
}

func NewMem() *Mem {
	return &Mem{
		Addresses: []*Address{},
	}
}

func (m *Mem) ProcessMemory(input io.ReadCloser, multiaddr bool) {
	s := bufio.NewScanner(input)
	defer input.Close()
	var fMask MaskFunction
	var aMask AddressFunction
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		parts := MaskMatcher.FindStringSubmatch(txt)
		if len(parts) == 2 {
			fMask = GenMaskFunction(parts[1])
			aMask = GenAddressesFunction(parts[1])
			continue
		}
		parts = MemoryMatcher.FindStringSubmatch(txt)
		if len(parts) != 3 {
			continue
		}
		addr, _ := strconv.ParseUint(parts[1], 10, 64)
		value, _ := strconv.ParseUint(parts[2], 10, 64)
		if multiaddr {
			addrs := aMask(addr)
			for _, a := range addrs {
				m.AddOrUpdate(a, value)
			}
		} else {
			value = fMask(value)
			m.AddOrUpdate(addr, value)
		}
	}
}

func (m *Mem) AddOrUpdate(addr, value uint64) {
	for _, a := range m.Addresses {
		if a.Addr == addr {
			a.Value = value
			return
		}
	}
	m.Addresses = append(m.Addresses, &Address{Addr: addr, Value: value})
}

func (m *Mem) Sum() uint64 {
	sum := uint64(0)
	for _, a := range m.Addresses {
		sum += a.Value
	}
	return sum
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	m := NewMem()
	m.ProcessMemory(f, true)

	fmt.Printf("Result: %d\n", m.Sum())
}
