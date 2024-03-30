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
code = directive / operator / template

directive = "// " directive:(include) { return directive, nil }
include = "include" ws+ fn:literal { return Include {FileName: fn.(Literal)}, nil }

quotedString = singleQuotedString / doubleQuotedString{ return string(c.text), nil }
singleQuotedString = "'" ("\\'" / [^'])* "'" { return string(c.text), nil }
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return string(c.text), nil }

expression_list = _e:expression _r:(ws* _e:expression { return _e, nil} )* {
    return ExpressionList { 
    Expressions: listFrom[Expression](_e, _r) }, nil
}

call = _i:store ws* "(" ws* _m:expression_list? ws* ")" {
  return Call {
    Store: _i.(StoreAccess),
    Arguments: _m.(ExpressionList) }, nil
}

index = '[' _e:expression ']' { return _e, nil };

member = _m:(identifier / quotedString) { 
  return Member{ 
    Identifier: castify(_m, "") }, nil
};

reference = sign 
  _m: (
    store / 
    ('(' _m:store ')' { return _m, nil })) 
  { 
    return Reference { Expression : _m.(Expression) }, nil 
  }

dblquot = quot quot
template = dblquot ws* _i:identifier ws* _a:argument_decl comment* eol !dblquot dblquot
parse = store ws* eq tilda ws* expression

pipe_receive = invoke / call
pipe_send = invoke / call / reference / literal

store = 
  (_e:store _i:index { 
    return StoreAccess { 
      Reference: _e.(Expression), 
      Index: _i.(Expression)}, nil }) /
  (_e:store '.' _i:store { 
    return StoreAccess { 
      Reference: _e.(Expression), 
      Index: _i.(Expression)}, nil }) /
  _e:literal { 
    return StoreAccess { 
      Reference: _e.(Expression) }, nil}

operator =
  parse /
  assign /
  write /
  invoke;

expression =
  call  /
  reference /
  literal /
  (_l:pipe_send pipe _r:pipe_receive) /
  ( '(' _e: expression ')' { return _e, nil} )

block= '{' ws* comment* eol _l:line* '}' comment* eol / eof

identifier_list= _i:identifier _r:(ws+ _i: identifier { return _i, nil } )* { 
  return IdentifierList { 
    Identifiers: listFrom[string](_i, _r) }, nil
}

argument_decl = '(' ws* _i:identifier_list? ws* ')' { 
  if _i == nil {
    return IdentifierList{}, nil
  }

  return _i, nil 
}
lambda = _a:argument_decl ws* _e:(block / operator / _e:expression* { 
  return Expressions(arrayify[Expression](_e)), nil }) {
    return Lambda { Arguments : _a.(IdentifierList), Body: _e.(Expression) }, nil
}

assign = _i:store ws* eq ws* _e:(lambda /expression)
{
  return Assign{
    Store: _i.(StoreAccess),
    Value: _e.(Expression) }, nil
}

write = quot ws* _e:expression* { 
  return Write { Expressions : arrayify[Expression](_e) }, nil }

invoke = amper _e:expression+ { 
  return Invoke { Expressions : arrayify[Expression](_e) }, nil }

identifier = (unds / let) (dig / unds / let)* { 
  return string(c.text), nil }

literal_list = l:literal r:(ws+ t:literal { 
  return t, nil} )* { 
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




