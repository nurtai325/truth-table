package scanner

func OperAction(a, b bool, tok Token) bool {
	action := operActions[tok]
	return action(a, b)
}

func IsValidToken(t Token) bool {
	return t > BEGIN && t < END && t != operBegin && t != operEnd
}

func IsOperator(t Token) bool {
	return t > operBegin && t < operEnd
}

func IsParentheses(t Token) bool {
	return t == LPAREN || t == RPAREN
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

type operFunc func(a, b bool) bool

var operActions = []operFunc{
	AND: func(a, b bool) bool {
		return a && b
	},
	OR: func(a, b bool) bool {
		return a || b
	},
	IMPLICATION: func(a, b bool) bool {
		return !a || b
	},
	IF_AND_ONLY_IF: func(a, b bool) bool {
		return a == b
	},
}

func Operate(a, b bool, oper Token) bool {
	f := operActions[oper]
	return f(a, b)
}
