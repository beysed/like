{
  package grammar
  import (
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

code_line =ws* _c:code comment? eoa { return _c, nil }
line = template / code_line / empty_line

comment = "#" ws+ (any !eoa)* any? { return string(c.text), nil }
code = operator / directive / call

directive = "#" directive:(include) { return directive, nil }
include = "include" ws+ fn:expression { return Include {FileName: fn.(Expression)}, nil }

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

member = _m:(identifier / s:quotedString) {
  return Member{
    Identifier: castify[string](_m) }, nil
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

string_content
  = ("$$" { return MakeLiteral("$"), nil}
  / reference
  /_a:any { return MakeLiteral(_a.(string)), nil })*

dblquot = quot quot
template_end = eol dblquot
template_body = _a:( _a:any !template_end { return MakeLiteral(_a.(string)), nil})*  { return _a, nil }
template =
  dblquot ws+ _i:identifier _a:argument_decl? ws* eol
  _b:( reference/_a:any !dblquot { return MakeLiteral(_a.(string)), nil } )*
  template_end
{
  e := arrayify[Expression](_b)
  return Assign {
    Store: StoreAccess { Reference: MakeLiteral(_i.(string)) },
    Value: Lambda {
      Arguments : castify[IdentifierList](_a),
      Body: Expressions(e)}},nil }

parse_data
  = colon lt ws* _f:literal ws+ _d:expression {
    return ParseData {
      Format: _f.(Expression),
      Data: _d.(Expression) }, nil
  }

format_data
  = colon gt ws* _f:literal ws+ _d:expression {
    return FormatData {
      Format: _f.(Expression),
      Data: _d.(Expression) }, nil
  }

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
  =  _c:expression ws+ '?' ws+ _t:expression _f:(ws+ '%' ws+ _f:expression { return _f, nil})? {
    return MakeCondition(_c.(Expression), _t.(Expression), castify[Expression](_f)), nil
  }
  / '?' ws+ _c:expression ws+ _t:body _f:(eol '%' ws+ _f:body {return _f, nil})? {
    return MakeCondition(_c.(Expression), _t.(Expression), castify[Expression](_f)), nil
  }
  / '?' ws+ _c:expression ws* eol '%' ws+ _f:body {
    return MakeCondition(_c.(Expression), nil, _f.(Expression)), nil
  }


each
  = '@' ws+ _e:expression ws+ _b:body { return MakeEach(_e.(Expression), _b.(Expression)), nil }

operator =
  reference /
  each /
  condition /
  assign /
  write /
  writeln /
  invoke;

list
  = '[' ws* _e:expression_list ws* ']' { return _e, nil }

expression =
  parse_data /
  format_data /
  each /
  condition /
  "!" ws* e:expression { return MakeNot(e.(Expression)), nil } /
  _l:expression ws* "==" ws* _r:expression { return MakeEqual(_l.(Expression), _r.(Expression)), nil} /
  _l:expression ws* "!=" ws* _r:expression { return MakeNot(MakeEqual(_l.(Expression), _r.(Expression))), nil} /
  call  /
  invoke /
  reference /
  literal /
  quotedString /
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

super_expression
  = _s:super_expression+ ws+ { return Expressions{ Expressions(arrayify[Expression](_s)), MakeLiteral(" ") } , nil }
  / _s:expression+ _r:(_w:ws+ _l:expression+ {
      return Expressions{
        MakeLiteral(" "),
        Expressions(arrayify[Expression](_l)) } , nil})*
    {
      return Expressions(listFrom[Expression](Expressions(arrayify[Expression](_s)), _r)), nil
    }

body =  block / code_line

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

literal
  = a:lit+ { return MakeLiteral(convert(a)), nil }
  / s:quotedString

lit
  = ("\\" _a:[@!={}()#=&|$<>~'"?%: \][.\\] { return string(_a.([]byte)), nil })
  / (sym / let / dig) { return string(c.text), nil}
  / "\\" _s:sym+ { return string(c.text[1:]), nil }
  / "\\" { return string(c.text), nil }

dig = [0-9]
let = [a-zA-Z]
sym = [^\\@!={}()#=&|$<>~'"?%: \n\r\t\][.]

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
colon = ':'
pipe_in = sign lt // $<
any = . { return string(c.text), nil }
eof = !.




