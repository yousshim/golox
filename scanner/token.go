package scanner

import "fmt"

type Token struct {
	Type   TokenType
	Lexeme any
	Line   int
}

func (t Token) String() string {
	switch t.Lexeme.(type) {
	case string:
		if t.Lexeme == "" {
			return fmt.Sprintf("%s %d", t.Type, t.Line)
		}
		return fmt.Sprintf("%s \"%s\" %d", t.Type, t.Lexeme, t.Line)
	case float64:
		return fmt.Sprintf("%s %f %d", t.Type, t.Lexeme, t.Line)
	}
	panic("Lox only supports string and float64 lexemes")
}
