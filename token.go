package main

type Token byte

const (
	UNDEFINED Token = iota
	MOVE_RIGHT
	MOVE_LEFT
	INC
	DEC
	OUTPUT
	INPUT
	LOOP_START
	LOOP_END
)

func isToken(b byte) bool {
	switch b {
	case '>':
		fallthrough
	case '<':
		fallthrough
	case '+':
		fallthrough
	case '-':
		fallthrough
	case '.':
		fallthrough
	case ',':
		fallthrough
	case '[':
		fallthrough
	case ']':
		return true
	}
	return false
}

func (t Token) String() string {
	switch t {
	case MOVE_RIGHT:
		return ">"
	case MOVE_LEFT:
		return "<"
	case INC:
		return "+"
	case DEC:
		return "-"
	case OUTPUT:
		return "."
	case INPUT:
		return ","
	case LOOP_START:
		return "["
	case LOOP_END:
		return "]"
	}
	return "Undefined"
}

func stringToToken(b byte) Token {
	switch b {
	case '>':
		return MOVE_RIGHT
	case '<':
		return MOVE_LEFT
	case '+':
		return INC
	case '-':
		return DEC
	case '.':
		return OUTPUT
	case ',':
		return INPUT
	case '[':
		return LOOP_START
	case ']':
		return LOOP_END
	}
	return UNDEFINED
}
