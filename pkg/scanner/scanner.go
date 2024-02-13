package scanner

import (
	"errors"
	"fmt"
)

type (
	Scanner struct {
		source string
		tokens []Token

		start   int
		current int
		line    int
	}
)

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	errs := []error{}
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			errs = append(errs, err)
		}
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", "", s.line))

	return s.tokens, errors.Join(errs...)
}

func (s *Scanner) scanToken() error {
	var err error
	r := s.advance()
	switch r {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		s.line += 1
	case '"':
		err = s.string()
	default:
		if isDigit(r) {
			s.number()
		} else if isAlpha(r) {
			s.identifier()
		} else {
			err = s.error("unexpected character " + string(r))
		}
	}

	return err
}

func (s *Scanner) advance() rune {
	s.current += 1
	return rune(s.source[s.current-1])
}

func (s *Scanner) match(expected rune) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != byte(expected) {
		return false
	}

	s.current++

	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\x00'
	}

	return rune(s.source[s.current])
}

func (s *Scanner) peekNext() rune {
	if s.current+1 >= len(s.source) {
		return '\x00'
	}

	return rune(s.source[s.current+1])
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenLiteral(tokenType, "")
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal string) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line += 1
		}
		s.advance()
	}

	if s.isAtEnd() {
		return s.error("unterminated string")
	}

	s.advance()

	value := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(STRING, value)

	return nil
}

func (s *Scanner) error(message string) error {
	return fmt.Errorf("[line %v] Error: %v", s.line, message)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}

	s.addTokenLiteral(NUMBER, s.source[s.start:s.current])
}

func isAlpha(r rune) bool {
	lowerCase := r >= 'a' && r <= 'z'
	upperCase := r >= 'A' && r <= 'Z'
	underscore := r == '_'

	return lowerCase || upperCase || underscore
}

func isAlphaNumeric(r rune) bool {
	return isAlpha(r) || isDigit(r)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = IDENTIFIER
	}

	s.addToken(tokenType)
}
