{
  package grammar
  import "strconv"
}

file = line*;

ws = ' ' / '\t';
eol = ('\r'? '\n') / '\r';

line = (comment / code:code comment? { return code, nil}) (eol / eof);

comment = "#" (any !eol)* { return string(c.text), nil };
code = directive;

directive = "// " directive:(include / context / source) { return directive, nil };
include = "include" ws+ fn:stringParam { return Include {fileName: string(fn.(string))}, nil };
context = "context" "s"?;
source = "source" "s"?;

stringParam = quotedString / unquotedString { return string(c.text), nil };
quotedString = singleQuotedString / doubleQuotedString{ return string(c.text), nil };
singleQuotedString = "'" ("\\'" / [^'])* "'" { return string(c.text), nil };
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return string(c.text), nil };
unquotedString = [^ \t\r\n#]* { return string(c.text), nil };

out = "| "

identifier = identifier_start identifier_symbol* { return string(c.text), nil }
identifier_symbol = [A-Za-z0-9_]
identifier_start = [A-Za-z_]

integer = [0-9]+ { return string(c.text), nil }

value = _i:(identifier/integer/quotedString) { 
  // value
  return Value{ value: _i.(string) }, nil 
}

index = '[' value ']'
memeber_list = identifier index* ('.' memeber_list)*;
reference = '$' ( ( '(' memeber_list ')' ) / memeber_list );

call_argument_list = _f:reference _r:(ws+ _a:reference { return _a, nil })* { 
  return append([]IndexedAccess{_f.(IndexedAccess)}, arrayify[IndexedAccess](_r)...), nil 
}

call = _i:identifier "(" _m:call_argument_list? ")"

assign = _i:identifier ws* "=" ws* _m:memeber_list 
{ 
  return Assign{ 
    identifier: castify(_i, ""), 
    source: _m.([]IndexedAccess), }, nil 
}

any = .
eof = !.




