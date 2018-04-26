package rfc5425

import (
	"strconv"
)

// Token represents a lexical token
type Token struct {
	typ TokenType
	lit []byte
}

// TokenType represents a lexical token type
type TokenType int

// Tokens
const (
	ILLEGAL TokenType = iota
	EOF
	WS
	MSGLEN
	SYSLOGMSG
)

// String outputs the string representation of the receiving Token
func (t Token) String() string {
	switch t.typ {
	case WS, EOF:
		return "Token{" + t.typ.String() + "}"
	default:
		return "Token{" + t.typ.String() + ", " + string(t.lit) + "}"
	}
}

const tokentypename = "ILLEGALEOFWSMSGLENSYSLOGMSG"

var tokentypeindex = [...]uint8{0, 7, 10, 12, 18, 27}

// String outputs the string representation of the receiving TokenType
func (i TokenType) String() string {
	if i < 0 || i >= TokenType(len(tokentypeindex)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return tokentypename[tokentypeindex[i]:tokentypeindex[i+1]]
}
