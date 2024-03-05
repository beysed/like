{
  package grammar;
}

file = line*;

ws = ' ' / '\t';
eol = ('\r'? '\n') / '\r';

line = (comment / code comment?) (eol / eof);

comment = "#" (any !eol)* { return string(c.text), nil };
code = out / command;

command = "// " (include / context / source);
include = "include" ws+ fileName;
context = "context" "s"?;
source = "source" "s"?;

fileName=stringParam
stringParam = quotedString / unquotedString;
quotedString = singleQuotedString / doubleQuotedString;
singleQuotedString = "'" !"'"* "'";
doubleQuotedString = '"' !'"'* '"';
unquotedString = !ws+ / !eol;

out = ">> "

any = .
eof = !.




