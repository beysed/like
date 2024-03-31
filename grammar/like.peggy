{
  package grammar
  import (
    "strconv"
  )
}

file = 
  _s:line* eof { 
    return arrayify[Expression](_s,), nil
  }

eol = ('\r'? '\n') / '\r';

eoa = eol / eof

empty_line 
  = (eol / ws+ eoa / ws* comment eoa) { return MakeLiteral(""), nil }

line = template / empty_line / (ws* _c:code comment? eoa { return _c, nil })

comment = "#" (any !eoa)* any? { return string(c.text), nil }
code = operator / directive / call 

directive = "// " directive:(include) { return directive, nil }
include = "include" ws+ fn:literal { return Include {FileName: fn.(Literal)}, nil }

quotedString = singleQuotedString / doubleQuotedString
singleQuotedString = "'" _s:("\\'" / [^'])* "'" { return MakeParsedString("'", castify(_s, "")), nil }
doubleQuotedString = '"' _s:("\\\"" / [^"])* '"'{ return MakeParsedString("\"", castify(_s, "")), nil }

expression_list = _e:expression _r:(ws* _e:expression { return _e, nil} )* {
    return ExpressionList(listFrom[Expression](_e, _r)), nil
}

call = _i:store ws* "(" ws* _m:expression_list? ws* ")" {
  var t ExpressionList 

  if _m != nil {
    t = _m.(ExpressionList)
  } else {
    t = ExpressionList {}
  }

  return Call {
    Store: _i.(StoreAccess),
    Arguments: t }, nil
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
template_end = eol dblquot
template_body = _a:( _a:any !template_end { return MakeLiteral(_a.(string)), nil})*  { return _a, nil }
template = 
  dblquot ws+ _i:identifier _a:argument_decl ws* eol 
  _b:( reference/_a:any !dblquot { return MakeLiteral(_a.(string)), nil } )* 
  template_end
{
  e := arrayify[Expression](_b)
  return Assign {
    Store: StoreAccess { Reference: MakeLiteral(_i.(string)) },
    Value: Lambda {
      Arguments : _a.(IdentifierList),
      Body: Expressions(e)}},nil }

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
  invoke /
  reference /
  literal /
  (_l:pipe_send pipe _r:pipe_receive) /
  ( '(' _e: expression ')' { return _e, nil} )

block= '{' ws* comment* eol _l:line* '}' comment* (eol / eof) { 
  return Block(arrayify[Expression](_l)) ,nil 
}

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

super_expression = _s:expression+ _r:(_w:ws+ _l:expression+ { 
      return Expressions{ 
        MakeLiteral(" "), 
        Expressions(arrayify[Expression](_l)) } , nil})* 
    { 
      return Expressions(listFrom[Expression](Expressions(arrayify[Expression](_s)), _r)), nil 
    }

lambda = _a:argument_decl ws* 
  _e:( block / operator / super_expression ) 
  {
    return Lambda { 
      Arguments : _a.(IdentifierList), 
      Body: _e.(Expression) }, nil
  }

assign = _i:store ws* eq ws* _e:(lambda / expression)
{
  return Assign{
    Store: _i.(StoreAccess),
    Value: _e.(Expression) }, nil
}

write = quot ws* _e:super_expression? { 
  return Write { Expressions : []Expression{ castify(_e, (Expression)(MakeLiteral(""))) } }, nil }

invoke = amper ws* _e:super_expression { 
  return Invoke{Expressions: _e.(Expressions)}, nil }

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
sym = [^{}()#=&|$<>~ \n\r\t]

ws = ' ' / '\t'
dash = '-'
unds = '_'
slsh = '/'
bslh = '\\'
quot = '`'
dblsign = sign sign
sign = '$'
pipe = '|'
gt = '>'
lt = '<'
eq = '='
amper = '&'
tilda = '~'
pipe_in = sign lt // $<
any = . { return string(c.text), nil }
eof = !.




