package scanner

func IsValidToken(t Token) bool {
	return t > BEGIN && t < END && t != operBegin && t != operEnd
}

func IsOperator(t Token) bool {
	return t > operBegin && t < operEnd
}

type Token int

func (t Token) String() string {
	return tokens[t]
}

const (
	ILLEGAL Token = iota
	BEGIN

	EOF
	VAR // a-z

	operBegin
	AND            // &&
	OR             // ||
	NOT            // !
	IMPLICATION    // ->
	IF_AND_ONLY_IF // <=>
	LPAREN         // (
	RPAREN         // )
	operEnd

	END
)

var tokens = [...]string{
	ILLEGAL:        "ILLEGAL",
	EOF:            "EOF",
	VAR:            "VAR",
	AND:            "&&",
	OR:             "||",
	NOT:            "!",
	IMPLICATION:    "->",
	IF_AND_ONLY_IF: "<=>",
	LPAREN:         "(",
	RPAREN:         ")",
}

var operInitals = [...]rune{
	AND:            '&',
	OR:             '|',
	IMPLICATION:    '-',
	IF_AND_ONLY_IF: '<',
}
