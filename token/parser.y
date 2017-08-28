
%{

package token

import (
)

var (
    _YY_IsString = false
)

%}

%union {
    str string
    arr []interface{}
}
%type <str> result part 
%type <arr> params
%token <str> _STR _WS
%start result

%%

result:
        part
|       result plus part
            {
                debugOut("result + part|", $1, $3)
                $$ = $1 + $3
            }

plus:
        '+'
|       '+' _WS

part:
        part _WS
            {
                debugOut("part _WS|", $1, $2)
                $$ = $1
            }
|       '\'' 
            { 
                _YY_IsString = true 
            }
        _STR '\''
            {
                _YY_IsString = false
                debugOut("' _STR '|", $3)
                $$ = $3
            }
|       '$' _STR
            {
                debugOut("$ _STR|", $2)
                lex, _ := yylex.(*lexImpl)
                str, err := lex.ctx.GetVariable($2)
                if err != nil {
                    panic(err)
                }
                $$ = str
            }
|       _STR '(' params ')'
            {
                debugOut("_STR(params)", $1, $3)
                lex, _ := yylex.(*lexImpl)
                str, err := lex.ctx.Calculate($1, $3)
                if err != nil {
                    panic(err)
                }
                $$ = str
            }

sep:
    ','
|   sep _WS

params:
        /* empty */
            {
                $$ = nil
            }
|       result
            {
                $$ = []interface{}{$1}
            }
|       params sep result
            {
                $$ = append($1, $3)
            }

