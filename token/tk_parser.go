package token

import (
	"errors"
	"fmt"
)

type Contexter interface {
	GetVariable(name string) (string, error)
	GetConst(name string) (string, error)
	Calculate(alg string, params []interface{}) (string, error)
}

type TkParser interface {
	GetToken(format string) (string, error)
}

type tkParserImpl struct {
	ctx   Contexter
	yy    *yyParserImpl
	lexer *lexImpl
}

func NewParser(c Contexter) TkParser {
	return &tkParserImpl{
		ctx:   c,
		yy:    &yyParserImpl{},
		lexer: newLexer(),
	}
}

func (p *tkParserImpl) GetToken(format string) (result string, err error) {
	defer func() {
		if pn := recover(); pn != nil {
			err = errors.New(fmt.Sprintf("PANIC!!! %v", pn))
		}
	}()

	p.lexer.SetSource(format)
	n := p.yy.Parse(p.lexer)
	if n != 0 {
		err = errors.New("yyParser return none zero value:" + string(n))
		return
	}

	fmt.Println("========", p.yy)
	//result = p.yy.lval.str
	result = p.yy.stack[1].str
	return
}

func TestParser() {
	src := " $DEF"
	fmt.Println("src:", src)
	parser := NewParser(nil)
	tk, err := parser.GetToken(src)
	fmt.Println("|", tk, err)
}
