package lexer

import (
	"fmt"
)

// Mewl Lexer Object
type MwlLexer struct {
	// Raw source code
	Source string
	// Length of source code string
	slen int
	// Start position of current token being read
	start int
	// Current position of the current token being read
	cur int
	// line number position
	line int
	// token output
	Tokens []MwlToken
}

// Create new mewl lexer
func NewMewlLexer(src string) *MwlLexer {
	return &MwlLexer{
		Source: src,
		slen:   len(src),
		start:  0,
		cur:    0,
		line:   1,
		Tokens: make([]MwlToken, 0),
	}
}

// Check if the lexer is at the end of source
func (ml *MwlLexer) atEnd() bool {
	return ml.cur >= ml.slen
}

// Advance the lexer position and return the eaten character
func (ml *MwlLexer) advance() byte {
	c := ml.Source[ml.cur]
	ml.cur++
	return c
}

// Peek into the Source string 'n'th postion
// n = 0 means return the next character
func (ml *MwlLexer) npeek(n int) byte {
	if ml.atEnd() || n < 0 { //todo: safety
		return 0
	}
	if ml.cur+n >= ml.slen {
		return 0
	}
	return ml.Source[ml.cur+n]
}

// Peek into the next character
func (ml *MwlLexer) peek() byte {
	return ml.npeek(0)
}

// Peek into the character after the next character
func (ml *MwlLexer) peekNext() byte {
	return ml.npeek(1)
}

// Peek into two character after the next character
func (ml *MwlLexer) peekNextNext() byte {
	return ml.npeek(2)
}

// If the next character is `a` advance and return true
// otherwise return false
func (ml *MwlLexer) match(a byte) bool {
	if ml.atEnd() {
		return false
	}

	if ml.peek() != a {
		return false
	}

	ml.advance()
	return true
}

// Run lexer on the Source and fillup the Tokens array
func (ml *MwlLexer) ScanTokens() {
	for !ml.atEnd() {
		ml.start = ml.cur
		ml.scanToken()
	}

	for item, value := range ml.Tokens {
		fmt.Printf("%d|%s\n", item, value)
	}
}

func (ml *MwlLexer) addToken(ttype MwlTokenType) {
	tok := MwlToken{
		TokType: ttype,
		Lexeme:  ml.Source[ml.start:ml.cur],
		Pos:     ml.start,
		Len:     ml.cur - ml.start,
	}

	ml.Tokens = append(ml.Tokens, tok)
}

func (ml *MwlLexer) eatMews() {
	for ml.peek() == 'm' &&
		ml.peekNext() == 'e' &&
		ml.peekNextNext() == 'w' {

		ml.advance()
		ml.advance()
		ml.advance()
	}
}

func (ml *MwlLexer) scanRawMews() {
	// first 'm' has been eaten
	ml.advance() // eat 'e'
	ml.advance() // eat 'w'

	ml.eatMews()

	if ml.peek() == '.' {
		ml.advance()

		if ml.peek() != 'm' ||
			ml.peekNext() != 'e' ||
			ml.peekNextNext() != 'w' {

			fmt.Printf("Invalid mew decimal number\n")
			return
		}

		ml.advance()
		ml.advance()
		ml.advance()

		ml.eatMews()
	}

	ml.addToken(MwlTokNumber)
}

func (ml *MwlLexer) scanToken() {
	ch := ml.advance()

	switch ch {
	case '[':
		ml.addToken(MwlTokStarter)
	case '(':
		ml.addToken(MwlTokStarter)
	case ']':
		ml.addToken(MwlTokEnder)
	case ')':
		ml.addToken(MwlTokEnder)
	case '+':
		ml.addToken(MwlTokPlus)
	case '-':
		ml.addToken(MwlTokMinus)
	case '/':
		if ml.match('/') {
			for ml.peek() != '\n' && !ml.atEnd() {
				ml.advance()
			}
		} else {
			ml.addToken(MwlTokSlash)
		}
	case 'm':
		if ml.peek() == 'e' && ml.peekNext() == 'w' {
			ml.scanRawMews()
		} else {
			fmt.Printf("Error Token '%c'\n", ch)
		}
	case ' ':
		break
	case '\t':
		break
	case '\r':
		break
	case '\n':
		ml.line++
	default:
		fmt.Printf("Unknown Token '%c'\n", ch)
	}
}
