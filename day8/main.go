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

type OpCode string

const (
	Accumulate  = OpCode("acc")
	Jump        = OpCode("jmp")
	NoOperation = OpCode("nop")
)

type Instruction struct {
	OpCode OpCode
	Value  int
}

func ToOpCode(input string) OpCode {
	switch input {
	case "acc":
		return Accumulate
	case "jmp":
		return Jump
	case "nop":
		return NoOperation
	}
	return NoOperation
}

func ParseInstruction(input string) Instruction {
	parts := strings.Split(input, " ")
	v, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf("Unable to parse: %v", err)
	}

	op := ToOpCode(parts[0])

	return Instruction{
		OpCode: op,
		Value:  v,
	}
}

func LoadInstructions(input io.ReadCloser) []Instruction {
	insts := []Instruction{}
	s := bufio.NewScanner(input)
	defer input.Close()

	for s.Scan() {
		txt := s.Text()
		txt = strings.TrimSpace(txt)
		insts = append(insts, ParseInstruction(txt))
	}

	return insts
}

type VirtualMachine struct {
	Ip             int  // InstructionPointer
	Accumulator    int  // Accumulator
	Running        bool // Running flag
	Halted         bool
	ExecutionOrder []int
}

func NewVirtualMachine() *VirtualMachine {
	return &VirtualMachine{
		Ip:             0,
		Accumulator:    0,
		Running:        false,
		Halted:         false,
		ExecutionOrder: []int{},
	}
}

func (vm *VirtualMachine) Reset() {
	vm.Ip = 0
	vm.Accumulator = 0
	vm.Running = false
	vm.Halted = false
}

func (vm *VirtualMachine) Execute(insts []Instruction, haltOnLoop bool) {
	vm.Running = true
	executed := map[int]bool{}
	for vm.Running {
		if v, ok := executed[vm.Ip]; ok && v && haltOnLoop {
			vm.Halted = true
			break
		}
		if vm.Ip > len(insts)-1 {
			vm.Running = false
			break
		}
		inst := insts[vm.Ip]
		vm.ExecutionOrder = append(vm.ExecutionOrder, vm.Ip)
		executed[vm.Ip] = true
		switch inst.OpCode {
		case Accumulate:
			vm.Accumulator += inst.Value
			vm.Ip++
		case Jump:
			vm.Ip += inst.Value
		case NoOperation:
			vm.Ip++
		}
	}
}

func (vm *VirtualMachine) PrintAccumulator() {
	fmt.Printf("Acc: %d\n", vm.Accumulator)
}

func MutateInstruction(insts []Instruction, order []int) []Instruction {
	vm := NewVirtualMachine()
	for x := range order {
		x = order[len(order)-1-x]
		if insts[x].OpCode == Jump {
			mutate := insts
			mutate[x].OpCode = NoOperation
			vm.Reset()
			vm.Execute(mutate, true)
			if vm.Running == false {
				return mutate
			}
		}
	}
	return insts
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	insts := LoadInstructions(f)
	vm := NewVirtualMachine()
	vm.Execute(insts, true)
	vm.PrintAccumulator()

	newInsts := MutateInstruction(insts, vm.ExecutionOrder)
	vm.Reset()
	vm.Execute(newInsts, false)
	vm.PrintAccumulator()
}
