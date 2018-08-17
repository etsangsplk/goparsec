package v2

import parsec "github.com/bricef/goparsec"

// Parser function parses input text encapsulated by Scanner, higher
// order parsers are constructed using combinators.
type Parser func(Scanner) (Node, Scanner)

type Regex string

type Identifier string

type ShortCircuitCallback func(Node, parsec.Scanner) Node

func And(name Identifier, parsers ...interface{}) Parser {
	return nil
}

func Or(name Identifier, parsers ...interface{}) Parser {
	return nil
}

func Kleen(name Identifier, parser interface{}) Parser {
	return nil
}

func Many(name Identifier, parser interface{}) Parser {
	return nil
}

func ManyUntil(name Identifier, itemParser interface{}, stopParser interface{}) Parser {
	return nil
}

func Maybe(name Identifier, parser interface{}) Parser {
	return nil
}

func ShortCircuit(parser Parser, callback ShortCircuitCallback) Parser {
	return nil
}

func Char(name Identifier) Parser {
	return nil
}

func float(name Identifier) Parser {
	return nil
}

func Hex(name Identifier) Parser {
	return nil
}

func Int(name Identifier) Parser {
	return nil
}

func Oct(name Identifier) Parser {
	return nil
}

func String(name Identifier) Parser {
	return nil
}

func Ident(name Identifier) Parser {
	return nil
}

func Atom(name Identifier, atom string) Parser {
	return nil
}

func AtomExact(name Identifier, atom string) Parser {
	return nil
}

func Token(name Identifier, r Regex) Parser {
	return nil
}

func TokenExact(name Identifier, r Regex) Parser {
	return nil
}

func OrToken(name Identifier, r Regex) Parser {
	return nil
}

func End() Parser {
	return nil
}

func NotEnd() Parser {
	return nil
}
