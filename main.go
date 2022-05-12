package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MEM_SIZE uint32 = 30000
)

func main() {
	commands := readCommands()
	tokens, loopMap := generateTokens(commands)
	runCommands(tokens, loopMap)
}

func readCommands() []byte {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic("Couldn't open file!")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []byte
	// Note: This will ignore newlines,
	// as the scanner throws them away
	for scanner.Scan() {
		for _, v := range []byte(scanner.Text()) {
			if !isToken(v) {
				if v == 'X' {
					return arr
				}
				continue
			}
			arr = append(arr, v)
		}
	}
	return arr
}

// Given an array of commands, return a corresponding array of tokens
func generateTokens(commands []byte) ([]Token, map[int]int) {
	loopStack := StackNew()
	loopMap := make(map[int]int)
	var tokens []Token
	for i, v := range commands {
		token := stringToToken(v)
		tokens = append(tokens, token)
		if token == LOOP_START {
			loopStack.Push(i)
		} else if token == LOOP_END {
			start := loopStack.Pop()
			end := i
			loopMap[start] = end
			loopMap[end] = start
		}
	}
	return tokens, loopMap
}

func runCommands(tokens []Token, loopMap map[int]int) {
	reader := bufio.NewReader(os.Stdin)

	instruction := 0
	memPointer := 0
	var memory [MEM_SIZE]int

	for instruction < len(tokens) {
		switch tokens[instruction] {
		case MOVE_RIGHT:
			memPointer++
		case MOVE_LEFT:
			memPointer--
		case INC:
			memory[memPointer]++
		case DEC:
			memory[memPointer]--
		case OUTPUT:
			fmt.Print(string(memory[memPointer]))
		case INPUT:
			char, _, err := reader.ReadRune()
			if err != nil {
				panic("Couldn't read rune!")
			}
			memory[memPointer] = int(char)
		case LOOP_START:
			if memory[memPointer] == 0 {
				instruction = loopMap[instruction]
			}
		case LOOP_END:
			if memory[memPointer] != 0 {
				instruction = loopMap[instruction]
			}
		default:
			fmt.Printf("instruction %d, length %d\n", instruction, len(tokens))
			errorString := fmt.Sprintf("Unimplemented instruction %d", int(tokens[instruction]))
			panic(errorString)
		}
		instruction++
	}
}
