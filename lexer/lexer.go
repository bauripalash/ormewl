package lexer

import (
	"fmt"
)

type MwlLexer struct {
	Source string
	slen   int
	start  int
	cur    int
	line   int
	Tokens []MwlToken
}

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

func (ml *MwlLexer) atEnd() bool {
	return ml.cur >= ml.slen
}

func (ml *MwlLexer) advance() byte {
	c := ml.Source[ml.cur]
	ml.cur++
	return c
}

func (ml *MwlLexer) npeek(n int) byte {
	if ml.atEnd() || n < 0 { //todo: safety
		return 0
	}
	return ml.Source[ml.cur+n]
}

func (ml *MwlLexer) peek() byte {
	return ml.npeek(0)
}

func (ml *MwlLexer) peekNext() byte {
	return ml.npeek(1)
}

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

func (ml *MwlLexer) eatWhitespaces() {
	for !ml.atEnd() {
		c := ml.peek()
		switch c {
		case ' ', '\t', '\r':
			ml.advance()
		case '\n':
			ml.line++
			ml.advance()
		default:
			return
		}
	}
}

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
		Pos:     ml.cur,
		Len:     ml.cur - ml.start,
	}

	ml.Tokens = append(ml.Tokens, tok)
}

func (ml *MwlLexer) scanToken() {
	ch := ml.advance()

	switch ch {
	case '[':
		ml.addToken(MwlTokStarter)
	case ']':
		ml.addToken(MwlTokEnder)
	case ' ':
		break
	case '\t':
		break
	case '\r':
		break
	case '\n':
		ml.line++
	default:
		fmt.Printf("Token '%c'", ch)
	}
}
