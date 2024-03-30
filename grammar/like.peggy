{
  package grammar
  import (
    "strconv"
  )
}

file = line*;

eol = ('\r'? '\n') / '\r';

line = _l:(comment { return nil, nil} / _c:code comment? { return _c, nil}) (eol / eof) {{ return _l, nil}}

comment = "#" (any !eol)* { return string(c.text), nil }
code = directive / expression / template

directive = "// " directive:(include) { return directive, nil }
include = "include" ws+ fn:literal { return Include {FileName: fn.(Literal)}, nil }

quotedString = singleQuotedString / doubleQuotedString{ return string(c.text), nil }
singleQuotedString = "'" ("\\'" / [^'])* "'" { return string(c.text), nil }
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return string(c.text), nil }

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

pipe_receive = invoke / call
pipe_send = invoke / call / reference / literal

expression =
  parse /
  assign /
  call /
  reference /
  literal_list /
  write /
  invoke /
  (_l:pipe_send pipe _r:pipe_receive) /
  ( '(' _e: expression ')' { return _e, nil} )

block= '{' ws* comment* eol line* '}' comment* eol / eof
identifier_list= identifier (ws+ identifier)
argument_decl = '(' ws* identifier_list ws* ')' 
lambda = argument_decl ws* _e:(block / expression)

assign = _i:member_list ws* eq ws* _e:(expression / lambda)
{
  return Assign{
    Identifier: _i.(MemberList),
    Value: _e.(Expression) }, nil
}

write = quot ws* _e:expression { return Write { Expression : _e.(Expression) }, nil }
invoke = amper _e:expression { return Invoke { Expression : _e.(Expression) }, nil }

identifier = (unds / let) (dig / unds / let)* { return string(c.text), nil }

literal_list = l:literal r:(ws+ t:literal { return t, nil} )* 
  { 
    return (LiteralList)(listFrom[Literal](l, r)), nil 
  }

literal = t:( lit+ { return string(c.text), nil }/ s:quotedString { return s, nil}) 
  { 
    return MakeLiteral(t.(string)), nil 
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
lt = '<'
eq = '='
amper = '&'
tilda = '~'
pipe_in = sign lt // $<
any = .
eof = !.




