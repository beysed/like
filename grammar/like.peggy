{
  package grammar
  import (
    "strconv"
    "strings"
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
singleQuotedString = "'" ("\\'" / [^'])* "'" { return MakeParsedString("'", string(c.text[1:len(c.text)-1])), nil }
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return MakeParsedString("\"", string(c.text[1:len(c.text)-1])), nil }

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

member = _m:(identifier / s:quotedString { return unquote(s), nil }) { 
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

pointer = '^' 
  _m: (
    store / 
    ('(' _m:store ')' { return _m, nil })) 
  { 
    return Pointer { Expression : _m.(Expression) }, nil 
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

condition 
  = '?' ws+ _c:expression ws+ _t:body _f:(eol '%' ws+ _f:body {return _f, nil})? { 
    return MakeCondition(_c.(Expression), _t.(Expression), castify[Expression](_f, nil)), nil
  }
  / '?' ws+ _c:expression ws* eol '%' ws+ _f:body {
    return MakeCondition(_c.(Expression), nil, _f.(Expression)), nil
  } 

each 
  = '@' ws+ _e:expression ws+ _b:body { return MakeEach(_e.(Expression), _b.(Expression)), nil }

operator =
  each /
  condition /
  parse /
  assign /
  write /
  writeln /
  invoke;

list 
  = '[' ws* _e:expression_list ws* ']' { return _e, nil }

expression =
  each /
  condition /
  "!" ws* e:expression { return MakeNot(e.(Expression)), nil } /
  _l:expression ws* "==" ws* _r:expression { return MakeEqual(_l.(Expression), _r.(Expression)), nil} /
  _l:expression ws* "!=" ws* _r:expression { return MakeNot(MakeEqual(_l.(Expression), _r.(Expression))), nil} /
  call  /
  invoke /
  reference /
  literal /
  list /
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

body =  block / operator / super_expression

lambda = _a:argument_decl ws* 
  _e:body 
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

writeln = quot ws* _e:super_expression { 
  return WriteLn { Expression : _e.(Expressions) }, nil }

write = tilda ws* _e:super_expression { 
  return Write { Expression : _e.(Expressions) }, nil }

invoke = amper ws* _e:super_expression { 
  return Invoke{Expressions: _e.(Expressions)}, nil }

identifier = (unds / let) (dig / unds / let)* { 
  return string(c.text), nil }

literal_list = l:literal r:(ws+ t:literal { 
  return t, nil} )* { 
    return (LiteralList)(listFrom[Literal](l, r)), nil 
  }

literal = t:( l:lit+ { return convert(l), nil }/ s:quotedString { return unquote(s), nil})
{
  return MakeLiteral(t.(string)), nil
}

// escaped= '\\' s:sym { return string(s.([]byte)), nil }
lit = ("\\" _a:[@!={}()#=&|$<>~'"?% \][.]  { return string(_a.([]byte)), nil }) / (sym / let / dig) { return string(c.text), nil}
//lit =  '\\"' { return "\"", nil } / (sym / let / dig) { return string(c.text), nil}
dig = [0-9]
let = [a-zA-Z]
sym = [^@!={}()#=&|$<>~'"?% \n\r\t\][.] 

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




