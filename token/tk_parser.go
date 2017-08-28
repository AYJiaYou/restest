package token

import (
	"errors"
	"fmt"
)

type TkParser interface {
	SetFormat(format string)
	SetContexter(ctx Contexter)
	GetToken() (string, error)
}

type tkParserImpl struct {
	yy    *yyParserImpl
	lexer *lexImpl
}

func NewParser() TkParser {
	return &tkParserImpl{
		yy:    &yyParserImpl{},
		lexer: newLexer(),
	}
}

func (p *tkParserImpl) SetFormat(format string) {
	p.lexer.SetSource(format)
}

func (p *tkParserImpl) SetContexter(ctx Contexter) {
	p.lexer.SetContexter(ctx)
}

func (p *tkParserImpl) GetToken() (result string, err error) {
	defer func() {
		if pn := recover(); pn != nil {
			err = errors.New(fmt.Sprintf("PANIC!!! %v", pn))
		}
	}()

	n := p.yy.Parse(p.lexer)
	if n != 0 {
		err = errors.New("yyParser return none zero value:" + string(n))
		return
	}

	result = p.yy.stack[1].str
	return
}

func TestParser() {
	//src := "'TSign ' + $SerialNumber + ':' + urlsafe_base64(hmac_sha1($ReqPath + '\\n' + $ReqBody, $SecretKey))"
	//src := "$ReqPath + '\\n' + $RegBody"
	src := "'TSign ' + $SerialNumber + ':' + urlsafe_base64(hmac_sha1($ReqPath + '\\n' + $ReqBody, $SecretKey))"
	fmt.Println("src:", src)
	parser := NewParser()
	parser.SetFormat(src)
	parser.SetContexter(newTestContexter())
	tk, err := parser.GetToken()
	fmt.Println("("+tk+")", err)
}
