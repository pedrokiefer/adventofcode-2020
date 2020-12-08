package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanParseInstruction(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Instruction
	}{
		{name: "nop", input: "nop +0", expected: Instruction{OpCode: NoOperation, Value: 0}},
		{name: "acc", input: "acc +1", expected: Instruction{OpCode: Accumulate, Value: 1}},
		{name: "jmp", input: "jmp -3", expected: Instruction{OpCode: Jump, Value: -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := ParseInstruction(tt.input)
			assert.Equal(t, tt.expected, inst)
		})
	}
}

func TestCanLoadInstruction(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)))

	insts := LoadInstructions(input)

	assert.Equal(t, []Instruction{
		{OpCode: NoOperation, Value: 0},
		{OpCode: Accumulate, Value: 1},
		{OpCode: Jump, Value: 4},
		{OpCode: Accumulate, Value: 3},
		{OpCode: Jump, Value: -3},
		{OpCode: Accumulate, Value: -99},
		{OpCode: Accumulate, Value: 1},
		{OpCode: Jump, Value: -4},
		{OpCode: Accumulate, Value: 6},
	}, insts)
}

func TestCanExecuteAndHalt(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)))

	insts := LoadInstructions(input)
	vm := NewVirtualMachine()
	vm.Execute(insts, true)

	assert.Equal(t, 5, vm.Accumulator)
	assert.Equal(t, []int{0, 1, 2, 6, 7, 3, 4}, vm.ExecutionOrder)
}

func TestCanMutateInstructions(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`)))

	insts := LoadInstructions(input)
	vm := NewVirtualMachine()
	vm.Execute(insts, true)

	assert.Equal(t, 5, vm.Accumulator)
	assert.Equal(t, []int{0, 1, 2, 6, 7, 3, 4}, vm.ExecutionOrder)

	newInsts := MutateInstruction(insts, vm.ExecutionOrder)
	vm.Reset()
	vm.Execute(newInsts, false)
	assert.Equal(t, 8, vm.Accumulator)
}
