package lexer

type MwlLexer struct {
	Source string
	slen   int
	start  int
	cur    int
	line   int
}

func NewMewlLexer(src string) *MwlLexer {
	return &MwlLexer{
		Source: src,
		slen:   len(src),
		start:  0,
		cur:    0,
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

func (ml *MwlLexer) peek() byte {
	if ml.atEnd() {
		return 0
	}
	return ml.Source[ml.cur]
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
