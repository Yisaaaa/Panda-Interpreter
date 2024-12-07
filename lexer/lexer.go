package lexer

import (
	"panda/token"
)

type Lexer struct {
	input   string
	pos     int  // current position
	readPos int  // next read position
	char    byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readNextChar()
	return l
}

// readNextChar reads the next character in our lexer's input.
// It only supports ASCII characters.
func (l *Lexer) readNextChar() {
	if l.readPos >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhiteSpace()

	switch l.char {
	case '=':
		if l.peekReadPosChar() == '=' {
			tok = l.makeTwoCharToken(token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '!':
		if l.peekReadPosChar() == '=' {
			tok = l.makeTwoCharToken(token.NOT_EQ)
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier(isLetter)
			tok.Type = token.LookUpIdent(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Literal = l.readIdentifier(isDigit)
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readNextChar()
	return tok
}

func (l *Lexer) readIdentifier(identifier func(char byte) bool) string {
	startPos := l.pos
	for identifier(l.char) {
		l.readNextChar()
	}
	return l.input[startPos:l.pos]
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{tokenType, string(char)}
}

func (l *Lexer) eatWhiteSpace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readNextChar()
	}
}

func (l *Lexer) peekReadPosChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

func (l *Lexer) makeTwoCharToken(tokenType token.TokenType) token.Token {
	var tok token.Token

	char := l.char
	l.readNextChar()
	tok.Literal = string(char) + string(l.char)
	tok.Type = tokenType
	return tok
}
