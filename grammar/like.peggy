{
  package grammar

  type Include struct {
    fileName string
  }
}

file = line*;

ws = ' ' / '\t';
eol = ('\r'? '\n') / '\r';

line = (comment / code:code comment? { return code, nil}) (eol / eof);

comment = "#" (any !eol)* { return string(c.text), nil };
code = directive;

directive = "// " directive:(include / context / source) { return directive, nil };
include = "include" ws+ fn:stringParam { return Include(fileName=fn), nil };
context = "context" "s"?;
source = "source" "s"?;

stringParam = quotedString / unquotedString { return string(c.text), nil };
quotedString = singleQuotedString / doubleQuotedString{ return string(c.text), nil };
singleQuotedString = "'" ("\\'" / [^'])* "'" { return string(c.text), nil };
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return string(c.text), nil };
unquotedString = [^ \t\r\n]* { return string(c.text), nil };

out = "| "

any = .
eof = !.




