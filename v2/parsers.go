package v2

import parsec "github.com/bricef/goparsec/v2"

// Parser function parses input text encapsulated by Scanner, higher
// order parsers are constructed using combinators.
type Parser func(Scanner) (Node, Scanner)

type Regex string

type Identifier string

type ShortCircuitCallback func(Node, parsec.Scanner) Node

// ParseNothing is a parser that will reject all input as invalid.
// It will not create a node for any input and return the scanner unchanged.
func ParseNothing(s Scanner) (Node, Scanner) {
	return nil, s
}

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

func Token(name Identifier, pattern Regex) Parser {
	if len(pattern) == 0 {
		return ParseNothing
	}
	if pattern[0] != '^' {
		pattern = "^" + pattern
	}
	return func(s Scanner) (Node, Scanner) {
		if s.Endof() {
			return nil, s
		}
		news := s.Clone()
		news.SkipWS()
		cursor := news.GetCursor()
		if tok, _ := news.Match(string(pattern)); tok != nil {
			return NewTerminal(name, string(tok), Cursor(cursor)), news
		}
		return nil, s
	}
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
