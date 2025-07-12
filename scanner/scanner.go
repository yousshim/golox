package scanner

import (
	"fmt"
	"strconv"
)

var singleCharTokens map[byte]TokenType = map[byte]TokenType{
	'(': LEFT_PAREN,
	')': RIGHT_PAREN,
	'{': LEFT_BRACE,
	'}': RIGHT_BRACE,
	',': COMMA,
	'.': DOT,
	'-': MINUS,
	'+': PLUS,
	';': SEMICOLON,
	'*': STAR,
}

var keywords map[string]TokenType = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func peek(script string, offset int) byte {
	if offset < len(script) {
		return script[offset]
	}
	return 0
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func Scan(script string) ([]Token, error) {
	var tokens []Token
	start := 0
	current := 0
	line := 1

	for current < len(script) {
		start = current
		c := script[current]
		current++
		if _, ok := singleCharTokens[c]; ok {
			t := Token{Type: singleCharTokens[c], Lexeme: "", Line: line}
			tokens = append(tokens, t)
			continue
		}
		switch c {
		case '!':
			if peek(script, current) == '=' {
				current++
				tokens = append(tokens, Token{Type: BANG_EQUAL, Lexeme: "", Line: line})
			} else {
				tokens = append(tokens, Token{Type: BANG, Lexeme: "", Line: line})
			}
			break
		case '=':
			if peek(script, current) == '=' {
				current++
				tokens = append(tokens, Token{Type: EQUAL_EQUAL, Lexeme: "", Line: line})
			} else {
				tokens = append(tokens, Token{Type: EQUAL, Lexeme: "", Line: line})
			}
			break
		case '<':
			if peek(script, current) == '=' {
				current++
				tokens = append(tokens, Token{Type: LESS_EQUAL, Lexeme: "", Line: line})
			} else {
				tokens = append(tokens, Token{Type: LESS, Lexeme: "", Line: line})
			}
			break
		case '>':
			if peek(script, current) == '=' {
				current++
				tokens = append(tokens, Token{Type: GREATER_EQUAL, Lexeme: "", Line: line})
			} else {
				tokens = append(tokens, Token{Type: GREATER, Lexeme: "", Line: line})
			}
			break
		case '/':
			if peek(script, current) == '/' {
				current++
				for peek(script, current) != '\n' {
					current++
				}
			} else {
				tokens = append(tokens, Token{Type: SLASH, Lexeme: "", Line: line})
			}
			break
		case '"':
			for peek(script, current) != '"' {
				current++
				if peek(script, current) == '\n' {
					line++
				}
			}
			if current > len(script) {
				report(line, "Unexpected end of string", string(c))
			}
			current++
			tokens = append(tokens, Token{Type: STRING, Lexeme: string(script[start+1 : current-1]), Line: line})
			break
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			for isDigit(peek(script, current)) {
				current++
			}
			if peek(script, current) == '.' && isDigit(peek(script, current+1)) {
				current++
				for isDigit(peek(script, current)) {
					current++
				}
			}
			n, _ := strconv.ParseFloat(string(script[start:current]), 64)
			tokens = append(tokens, Token{Type: NUMBER, Lexeme: n, Line: line})
			break
		case ' ', '\t', '\r':
			break
		case '\n':
			line++
			break
		default:
			if isAlpha(c) {
				for isAlpha(peek(script, current)) || isDigit(peek(script, current)) {
					current++
				}
				val := string(script[start:current])
				if keyword, ok := keywords[val]; ok {
					tokens = append(tokens, Token{Type: keyword, Lexeme: "", Line: line})
					break
				} else {
					tokens = append(tokens, Token{Type: IDENTIFIER, Lexeme: val, Line: line})
				}
				break
			}
			report(line, "Unexpected character", string(c))
			break
		}
	}

	tokens = append(tokens, Token{Type: EOF, Lexeme: "", Line: line})
	return tokens, nil
}

func report(line int, where string, msg string) error {
	return fmt.Errorf("[line %d Error %s: %s\n", line, where, msg)
}
