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

var bagsMatcher = regexp.MustCompile(`^(?P<n>[0-9]+)\ (?P<color>.*)\ bag(?:s)?[\.]?$`)

type BagDAG struct {
	Nodes []*Node
}

type Node struct {
	Name  string
	Nodes map[string]*Pair
}

type Pair struct {
	Count int
	Node  *Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Nodes: map[string]*Pair{},
	}
}

func CountBags(node *Node) int {
	total := 0
	for _, n := range node.Nodes {
		total += n.Count
		total += n.Count * CountBags(n.Node)
	}
	return total
}

func findNode(node *Node, name string) *Node {
	for _, n := range node.Nodes {
		if n.Node.Name == name {
			return n.Node
		}
		f := findNode(n.Node, name)
		if f != nil {
			return f
		}
	}
	return nil
}

func (d *BagDAG) FindNode(name string) *Node {
	for _, n := range d.Nodes {
		if n.Name == name {
			return n
		}
		f := findNode(n, name)
		if f != nil {
			return f
		}
	}
	return nil
}

func countPathsTo(node *Node, name string, path string, paths *[]string) {
	for _, n := range node.Nodes {
		p := fmt.Sprintf("%s,%s", path, n.Node.Name)
		if n.Node.Name == name {
			*paths = append(*paths, p)
			continue
		}
		countPathsTo(n.Node, name, p, paths)
	}
}

func (d *BagDAG) CountPathsTo(name string) int {
	paths := []string{}
	for _, n := range d.Nodes {
		if n.Name == name {
			paths = append(paths, n.Name)
		}
		countPathsTo(n, name, n.Name, &paths)
	}
	countMap := map[string]bool{}
	for _, v := range paths {
		ps := strings.Split(v, ",")
		for _, p := range ps {
			countMap[p] = true
		}
	}
	return len(countMap) - 1
}

func printNode(n *Node, level int) {
	for _, n := range n.Nodes {
		fmt.Printf("%s- %s: %d\n", strings.Repeat("\t", level), n.Node.Name, n.Count)
		printNode(n.Node, level+1)
	}
}

func (d *BagDAG) Print() {
	for _, n := range d.Nodes {
		fmt.Printf("- %s\n", n.Name)
		printNode(n, 1)
	}
}

func (d *BagDAG) ParseLine(input string) {
	x := strings.SplitN(input, " bags contain ", 2)
	name := x[0]

	shouldAppend := false
	n := d.FindNode(name)
	if n == nil {
		shouldAppend = true
		n = NewNode(name)
	}

	contains := x[1]
	for _, v := range strings.Split(contains, ",") {
		v = strings.TrimSpace(v)
		parts := bagsMatcher.FindStringSubmatch(v)
		if len(parts) == 0 {
			continue
		}
		count, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		nn := d.FindNode(parts[2])
		if nn == nil {
			nn = NewNode(parts[2])
		}
		n.Nodes[parts[2]] = &Pair{Count: count, Node: nn}
	}

	if shouldAppend {
		d.Nodes = append(d.Nodes, n)
	}
}

func LoadBagsDAG(input io.ReadCloser) *BagDAG {
	bags := &BagDAG{
		Nodes: []*Node{},
	}
	s := bufio.NewScanner(input)
	defer input.Close()

	for s.Scan() {
		bags.ParseLine(s.Text())
	}

	return bags
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	bags := LoadBagsDAG(f)
	n := bags.CountPathsTo("shiny gold")
	fmt.Printf("Bags: %v\n", n)
	total := CountBags(bags.FindNode("shiny gold"))
	fmt.Printf("Bags Inside Shiny Gold: %d\n", total)
}
