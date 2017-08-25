
%{

package token

import (
)

%}

%union {
    str string
}
%type <str> result part 
%token <str> _STR _WS
%start result

%%

result:
        part
|       result '+' part
            {
                debugOut("result + part|", $1, $3)
                $$ = $1 + $3
            }

part:
        part _WS
            {
                debugOut("part _WS|", $1, $2)
                $$ = $1
            }
|       _WS part
            {
                debugOut("_WS part|", $1, $2)
                $$ = $2
            }
|       '\'' _STR '\''
            {
                debugOut("' _STR '|", $2)
                $$ = $2
            }
|       '$' _STR
            {
                debugOut("$ _STR|", $2)
                $$ = $2
            }

