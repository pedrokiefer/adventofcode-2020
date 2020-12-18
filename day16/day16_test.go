package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanLoadTicketsPuzzle(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)))

	tickets := LoadTicketsPuzzle(input)

	assert.Equal(t, []Interval{{6, 11}, {33, 44}}, tickets.Fields["row"])
	assert.Equal(t, Ticket{7, 1, 14}, tickets.MyTicket)
	assert.Equal(t, []Ticket{
		{7, 3, 47},
		{40, 4, 50},
		{55, 2, 20},
		{38, 6, 12},
	}, tickets.NearbyTickets)

	assert.Equal(t, 71, tickets.ValidateNearbyTickets())
	assert.Equal(t, []Ticket{
		{7, 1, 14},
		{7, 3, 47},
	}, tickets.ValidTickets)
}

func TestCanFindRowMeaning(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`)))

	tickets := LoadTicketsPuzzle(input)

	assert.Equal(t, []Interval{{0, 1}, {4, 19}}, tickets.Fields["class"])
	assert.Equal(t, Ticket{11, 12, 13}, tickets.MyTicket)
	assert.Equal(t, []Ticket{
		{3, 9, 18},
		{15, 1, 5},
		{5, 14, 9},
	}, tickets.NearbyTickets)

	assert.Equal(t, 0, tickets.ValidateNearbyTickets())
	assert.Equal(t, []Ticket{
		{11, 12, 13},
		{3, 9, 18},
		{15, 1, 5},
		{5, 14, 9},
	}, tickets.ValidTickets)

	tickets.FindColumnsName()
	assert.Equal(t, []string{"row", "class", "seat"}, tickets.ColNames)
}
