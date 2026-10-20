package lexer

import (
	"fmt"
)

// Mewl Token Type
type MwlTokenType int

const (
	// Mewl Expression Starter '['
	MwlTokStarter MwlTokenType = iota
	// Mewl Expression Ender ']'
	MwlTokEnder
	// Mewl Token '+'
	MwlTokPlus
	// Mewl Token '-'
	MwlTokMinus
	// Mewl Number 'mewmewmew...'
	MwlTokNumber
	// Mewl Token '*'
	MwlTokAstr
	// Mewl Token '/'
	MwlTokSlash
	// Mewl End of File 'EOF'
	MwlTokEOF
)

func (tt MwlTokenType) String() string {
	switch tt {
	case MwlTokStarter:
		return "TokStarter"
	case MwlTokEnder:
		return "TokEnder"
	case MwlTokPlus:
		return "TokPlus"
	case MwlTokMinus:
		return "TokMinus"
	case MwlTokNumber:
		return "TokNumber"
	case MwlTokEOF:
		return "TokEOF"
	default:
		return "TokUnknown"
	}
}

// Mewl Token Structure
// Outputted by Lexer
type MwlToken struct {
	// Type of the token
	TokType MwlTokenType
	// Actual raw lexeme of the token
	Lexeme string
	// Global position
	Pos int
	// Lexeme Length
	Len int
}

func NewDummyToken() MwlToken {
	return MwlToken{
		TokType: MwlTokEOF,
		Lexeme:  "",
	}
}

func (t MwlToken) String() string {
	return fmt.Sprintf("MwlToken(%d,%d)(%s|'%s')", t.Pos, t.Len, t.TokType, t.Lexeme)
}
