package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParseLine(t *testing.T) {
	bags := BagDAG{
		Nodes: []*Node{},
	}
	bags.ParseLine("dark orange bags contain 3 bright white bags, 4 muted yellow bags.")

	assert.Equal(t, "dark orange", bags.Nodes[0].Name)
	assert.Equal(t, 3, bags.Nodes[0].Nodes["bright white"].Count)
	assert.Equal(t, 4, bags.Nodes[0].Nodes["muted yellow"].Count)
}

func TestCanLoadBagsDAG(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`light red bags contain 1 bright white bag, 2 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)))

	bags := LoadBagsDAG(input)

	bags.Print()
	n := bags.CountPathsTo("shiny gold")
	total := CountBags(bags.FindNode("shiny gold"))

	assert.Equal(t, 4, n)
	assert.Equal(t, 32, total)
}

func TestCanLoadBagsDAG2(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`)))

	bags := LoadBagsDAG(input)

	n := bags.FindNode("shiny gold")
	total := CountBags(n)

	assert.Equal(t, 126, total)
}
