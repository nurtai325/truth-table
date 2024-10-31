package scanner

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var (
	ErrEOF          = errors.New("end of source")
	ErrInvalidRune  = errors.New("invalid rune")
	ErrInvalidToken = errors.New("invalid token")
)

type Scanner struct {
	src    []byte
	cursor int
}

func (s *Scanner) SetSrc(src []byte) {
	s.src = src
	s.cursor = 0
}

func (s *Scanner) Scan() (Token, string, error) {
	r, err := s.NextRune()
	if err != nil {
		if errors.Is(err, ErrEOF) {
			return EOF, "", nil
		} else {
			return ILLEGAL, "", err
		}
	}

	if r >= 97 && r <= 122 {
		return VAR, string(r), nil
	}

	switch r {
	case '(':
		return LPAREN, "", nil
	case ')':
		return RPAREN, "", nil
	case '!':
		return NOT, "", nil
	case '&':
		tok, err := s.scanOper(AND)
		return tok, "", err
	case '|':
		tok, err := s.scanOper(OR)
		return tok, "", err
	case '-':
		tok, err := s.scanOper(IMPLICATION)
		return tok, "", err
	case '<':
		tok, err := s.scanOper(IF_AND_ONLY_IF)
		return tok, "", err
	default:
		return ILLEGAL, "", ErrInvalidToken
	}
}

func (s *Scanner) NextRune() (rune, error) {
	if s.cursor >= len(s.src) {
		return utf8.RuneError, ErrEOF
	}
	r, size := utf8.DecodeRune(s.src[s.cursor:])
	if r == utf8.RuneError {
		return utf8.RuneError, fmt.Errorf("pos: %d, src: %v: %w", s.cursor, s.src, ErrInvalidRune)
	}
	s.cursor += size
	if r == ' ' || r == '\n' {
		r, err := s.NextRune()
		return r, err
	}

	return r, nil
}

func (s *Scanner) scanOper(op Token) (Token, error) {
	r := operInitals[op]
	literal := tokens[op]
	scanned := string(r)
	for i, e := range literal {
		if i == 0 {
			continue
		}

		r, nextRuneErr := s.NextRune()
		scanned += string(r)
		if nextRuneErr != nil || r != e {
			err := fmt.Errorf("expected operator: %s, got %s", literal, scanned)
			if errors.Is(nextRuneErr, ErrEOF) {
				return EOF, errors.Join(nextRuneErr, err, ErrInvalidToken)
			} else {
				return ILLEGAL, errors.Join(nextRuneErr, err, ErrInvalidToken)
			}
		}
	}
	return op, nil
}
