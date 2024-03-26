{
  package grammar
  import (
    "strconv"
  )
}

file = line*;

eol = ('\r'? '\n') / '\r';

line = (comment / code:code comment? { return code, nil}) (eol / eof)

comment = "#" (any !eol)* { return string(c.text), nil }
code = directive / expression

directive = "// " directive:(include) { return directive, nil }
include = "include" ws+ fn:literal { return Include {FileName: fn.(Literal)}, nil }

quotedString = singleQuotedString / doubleQuotedString{ return string(c.text), nil }
singleQuotedString = "'" ("\\'" / [^'])* "'" { return string(c.text), nil }
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return string(c.text), nil }

// integer = [0-9]+ { return string(c.text), nil }

expression_list = expression (ws* expression)*;
call = _i:identifier "(" _m:expression_list? ")"

index = '[' _e:expression ']' { return _e, nil };

member = _m:(identifier / quotedString) _e:index* { 
  return Member{ 
    Identifier: castify(_m, ""), 
    Indexes: arrayify[Expression](_e) }, nil
};

member_list = _m:member _r:('.' _r:member)* {
  return (MemberList)(listFrom[Member](_m, _r)), nil
};

reference = sign 
  _m: (
    _m:(member_list / expression) { return _m, nil } / 
    ('(' _m:reference ')' { return _m, nil })) 
  { 
    return _m, nil 
  }

dblquot = quot quot
template = dblquot ws* _i:identifier ws* _a:argument_decl comment* eol !dblquot dblquot
parse = identifier ws* eq tilda expression

expression =
  parse /
  assign /
  call /
  reference /
  literal /
  ( '(' _e: expression ')' { return _e, nil} )

block= '{' ws* comment* eol line* '}' comment* eol / eof
identifier_list= identifier (ws+ identifier)
argument_decl = '(' ws* identifier_list ws* ')' 
lambda = argument_decl ws* _e:(block / expression)

assign = _i:identifier ws* eq ws* _e:(expression / lambda)
{
  return Assign{
    Identifier: castify(_i, ""),
    Value: _e.(Expression) }, nil
}

writeln = gt ws* expression
write = pipe ws* expression
invoke = amper expression

identifier = (unds / let) (dig / unds / let)* { return string(c.text), nil }

literal_list = l:literal r:(ws+ t:literal { return t, nil} )* 
  { 
    return (LiteralList)(listFrom[Literal](l, r)), nil 
  }

literal = t:( lit+ { return string(c.text), nil }/ s:quotedString { return s, nil}) 
  { 
    return LiteralMake(t.(string)), nil 
  }

lit = sym / let / dig
dig = [0-9]
let = [a-zA-Z]
sym = dash / unds / slsh / bslh

ws = ' ' / '\t'
dash = '-'
unds = '_'
slsh = '/'
bslh = '\\'
quot = '`'
sign = '$'
pipe = '|'
gt = '>'
eq = '='
amper = '&'
tilda = '~'

any = .
eof = !.




