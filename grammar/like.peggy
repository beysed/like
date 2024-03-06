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

value_prefix = "$" { return string(c.text), nil }
value = _v:value_prefix? _i:(identifier/integer/quotedString) { 
  // value
  return Value{ prefix: castify(_v, ""), identifier: _i.(string)}, nil 
}

array_access = _a:(value) _i:('[' _i:(value) ']' { return _i, nil } )* { 
  // array_access
  return ArrayAccess{value: _a.(Value), indexes : valueArray(_i)}, nil 
}

assign = _i:identifier ws* "=" ws* _f:array_access _r:('.' _a:array_access { return _a, nil })* 
{ 
  
  return Assign{ 
    identifier: castify(_i, ""), 
    source: append([]ArrayAccess{_f.(ArrayAccess)}, arrayify[ArrayAccess](_r)...) }, nil 
}

expression = assign
any = .
eof = !.




