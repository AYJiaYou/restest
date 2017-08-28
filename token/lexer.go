package token

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
)

const _EOF = 0

type lexImpl struct {
	r   *strings.Reader
	ctx Contexter
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isLetter(r rune) bool {
	return ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') ||
		('0' <= r && r <= '9') ||
		r == '_' || r == ':' || r == '\\' || r == '\n'
}

func isSpecial(r rune) bool {
	return r == '+' ||
		r == '\'' ||
		r == '$' ||
		r == '(' ||
		r == ')' ||
		r == ',' ||
		r == _EOF
}

func newLexer() *lexImpl {
	return &lexImpl{}
}

func (l *lexImpl) Error(s string) {
	panic(errors.New(s))
}

func (l *lexImpl) Lex(yylval *yySymType) int {
	r := l.readRune()

	if isWhitespace(r) {
		l.unreadRune()
		return l.scanWhitespace(yylval)
	} else if isLetter(r) {
		l.unreadRune()
		return l.scanString(yylval)
	} else if isSpecial(r) {
		return int(r)
	}

	panic(errors.New("invalid char:" + string(r)))
}

func (l *lexImpl) SetSource(s string) {
	l.r = strings.NewReader(s)
}

func (l *lexImpl) SetContexter(ctx Contexter) {
	l.ctx = ctx
}

func (l *lexImpl) readRune() rune {
	r, _, err := l.r.ReadRune()
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
		return _EOF
	}
	return r
}

func (l *lexImpl) unreadRune() {
	err := l.r.UnreadRune()
	if err != nil {
		panic(err)
	}
}

func (l *lexImpl) scanWhitespace(yylval *yySymType) int {
	var buf bytes.Buffer
	buf.WriteRune(l.readRune())

	for {
		if r := l.readRune(); r == _EOF {
			break
		} else if !isWhitespace(r) {
			l.unreadRune()
			break
		} else {
			buf.WriteRune(r)
		}
	}
	return _WS
}

func (l *lexImpl) scanString(yylval *yySymType) int {
	var buf bytes.Buffer
	buf.WriteRune(l.readRune())

	for {
		if r := l.readRune(); r == _EOF {
			break
		} else if isSpecial(r) || (!isLetter(r) && !_YY_IsString) {
			l.unreadRune()
			break
		} else {
			buf.WriteRune(r)
		}
	}
	yylval.str = buf.String()
	return _STR
}

func TestLexer() {
	src := "'TSign ' + $SerialNumber + ':' + urlsafe_base64(hmac_sha1(%REQ_PATH + '\n' + %REQ_BODY, $SecretKey))"
	fmt.Println("src:", src)
	lex := newLexer()
	lex.SetSource(src)
	for {
		sym := yySymType{}
		n := lex.Lex(&sym)
		if n == _EOF {
			break
		}
		fmt.Println(n, sym.str)
	}
}
