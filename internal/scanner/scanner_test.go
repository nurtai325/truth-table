package scanner_test

import (
	"errors"
	"testing"
	"unicode/utf8"

	"github.com/nurtai325/truth-table/internal/scanner"
)

type scannerTestCase struct {
	Tok scanner.Token
	Lit string
	Err error
}

var scannerCases = []scannerTestCase{
	{Tok: scanner.VAR, Lit: "a", Err: nil},
	{Tok: scanner.AND, Lit: "", Err: nil},
	{Tok: scanner.VAR, Lit: "b", Err: nil},
	{Tok: scanner.OR, Lit: "", Err: nil},
	{Tok: scanner.LPAREN, Lit: "", Err: nil},
	{Tok: scanner.VAR, Lit: "a", Err: nil},
	{Tok: scanner.IMPLICATION, Lit: "", Err: nil},
	{Tok: scanner.VAR, Lit: "b", Err: nil},
	{Tok: scanner.RPAREN, Lit: "", Err: nil},
	{Tok: scanner.IF_AND_ONLY_IF, Lit: "", Err: nil},
	{Tok: scanner.VAR, Lit: "c", Err: nil},
}

func TestScannerNormal(t *testing.T) {
	normalCase := "a&&b||(a->b)<=>c"
	var s scanner.Scanner
	s.SetSrc([]byte(normalCase))
	for {
		tok, lit, err := s.Scan()
		if tok == scanner.EOF {
			break
		}
		if !scanner.IsValidToken(tok) {
			t.Fatalf("invalid token: %v %s", tok, lit)
		}
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
	}

}

func TestScannerErr(t *testing.T) {
	var s scanner.Scanner
	errCase := "1234567890-_=+~`@#$%^&*{}|;:,.<>?/ABCD"
	s.SetSrc([]byte(errCase))
	for {
		tok, lit, err := s.Scan()
		if tok == scanner.EOF {
			break
		}
		t.Log(tok, lit, err)
		if scanner.IsValidToken(tok) {
			t.Fatalf("valid token: %v %s", tok, lit)
		}
		if err == nil {
			t.Fatalf("expected err: %v, got %v", scanner.ErrInvalidToken, err)
		}
	}
}

func TestScannerFull(t *testing.T) {
	var s scanner.Scanner
	fullCheckedCase := "a&&b||(a->b)<=>c"
	s.SetSrc([]byte(fullCheckedCase))
	for _, testCase := range scannerCases {
		tok, lit, err := s.Scan()
		if tok != testCase.Tok {
			t.Fatalf("unexpected token: %v, expected: %v", tok, testCase.Tok)
		}
		if lit != testCase.Lit {
			t.Fatalf("unexpected literal: %s, expected: %s", lit, testCase.Lit)
		}
		if err != testCase.Err {
			t.Fatalf("unexpected error: %v, expected: %v", err, testCase.Err)
		}
	}

}

var nextRuneCases = []string{
	"",
	"Hello, World!",
	"1234567890",
	"!@#$%^&*()",
	"This is a normal string.",
	"   Leading and trailing spaces   ",
	"A very long string that is intended to test the limits of the slice length. It goes on and on and on without any real endpoint!",
	"Newline\ncharacter",
	"Tab\tcharacter",
	"😀 Unicode Emoji",
	"Special chars: ñ, ö, ü",
	"Another line\nwith multiple\nnewlines.",
}

func TestNextRune(t *testing.T) {
	for _, testCase := range nextRuneCases {
		var s scanner.Scanner
		s.SetSrc([]byte(testCase))
		t.Logf("source: %s", testCase)

		for _, testRune := range testCase {
			if testRune == ' ' || testRune == '\n'  {
				continue
			}
			r, err := s.NextRune()
			t.Logf("%c %c", testRune, r)
			if err != nil {
				t.Fatalf("err: %v", err)
			}
			if testRune != r {
				t.Fatalf("asserting %c == %c", testRune, r)
			}
		}
		r, err := s.NextRune()
		if r != utf8.RuneError {
			t.Fatalf("expected RuneError, got: %c", r)
		}
		if !errors.Is(err, scanner.ErrEOF) {
			t.Fatalf("expected EOF, got: %v, %c", err, r)
		}
	}
}
