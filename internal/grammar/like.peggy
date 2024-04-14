{
  package grammar
  import "github.com/beysed/like/internal/grammar/common"
}

bom = "\xEF\xBB\xBF" / "\xFE\xFF"
file =
  bom? _s:line* eof {
    return arrayify[Expression](_s,), nil
  }

eol = ('\r'? '\n') / '\r';

eoa = eol / eof

empty_line
  = (eol / ws+ eoa / ws* comment eoa) { return MakeLiteral(""), nil }

code_line =ws* _c:code comment? eoa { return _c, nil }
line = template / code_line / empty_line

comment = "#" (any !eoa)* any? { return string(c.text), nil }
code = expression / directive

directive = "#" directive:(include) { return directive, nil }
include = "include" ws+ fn:expression { return Include {FileName: fn.(Expression)}, nil }

quotedString = singleQuotedString / doubleQuotedString
singleQuotedString = "'" ("\\'" / [^'])* "'" { return MakeParsedString("'", string(c.text[1:len(c.text)-1])), nil }
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return MakeParsedString("\"", string(c.text[1:len(c.text)-1])), nil }

expression_list = _e:expression _r:(ws* _e:expression { return _e, nil} )* {
    return ExpressionList(listFrom[Expression](_e, _r)), nil
}

named_expression = _i:(_i:identifier ws* ':' ws* {return _i, nil})? _e:expression {
  var k string
  if _i == nil {
    k = ""
  } else {
    k = _i.(string)
  }

  return NamedExpression { Key: k, Value: _e.(Expression)}, nil
}

named_expression_list = _e:named_expression _r:(ws* _e:named_expression { return _e, nil} )* {
    return NamedExpressionList(listFrom[NamedExpression](_e, _r)), nil
}

call = _i:reference ws* "(" ws* _m:named_expression_list? ws* ")" {
  var t NamedExpressionList

  if _m != nil {
    t = _m.(NamedExpressionList)
  } else {
    t = NamedExpressionList {}
  }

  return Call {
    Reference: _i.(Reference),
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

string_content
  = _a:("$$" { return MakeLiteral("$"), nil}
  / call
  / reference
  / ws  { return MakeLiteral(string(c.text)), nil }
  /_a:[^ \t$]+ { return MakeLiteral(string(c.text)), nil})* {
    return Expressions(arrayify[Expression](_a)), nil
}

dblquot = quot quot
template =
  dblquot ws+ _i:identifier _a:argument_decl? ws* eol
  _s:(_s:any !"``" { return _s, nil} )*
  eol dblquot
{
  b := strings.Join(arrayify[string](_s), "")
  s := StoreAccess { Reference: MakeLiteral(_i.(string)) }
  i := castify[IdentifierList](_a)

  return MakeTemplate(s, i, b),nil
}

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
  / '?' ws+ _c:expression ws+ _t:body _f:(ws+ '%' ws+ _f:body {return _f, nil})? {
    return MakeCondition(_c.(Expression), _t.(Expression), castify[Expression](_f)), nil
  }

each
  = '@' ws+ _e:expression ws+ _b:body { return MakeEach(_e.(Expression), _b.(Expression)), nil }

list
  = '[' ws* _e:expression_list? ws* ']' {
  if (_e == nil) {
    return Constant { Value : common.List{} }, nil
  }

  return _e, nil
}

expression =
  write /
  writeln /
  assign /
  assign_add /
  parse_data /
  format_data /
  each /
  condition /
  "!" ws* e:expression { return MakeNot(e.(Expression)), nil } /
  _l:expression ws* "+" ws* _r:expression { return MakeAdd(_l.(Expression), _r.(Expression)), nil} /
  _l:expression ws* "==" ws* _r:expression { return MakeEqual(_l.(Expression), _r.(Expression)), nil} /
  _l:expression ws* "!=" ws* _r:expression { return MakeNot(MakeEqual(_l.(Expression), _r.(Expression))), nil} /
  _l:expression ws* "|" ws* _r:expression { return MakePiped[Pipe](_l.(Expression), _r.(Expression)), nil } /
  _l:expression ws* ">>" ws* _r:expression { return MakePiped[PipeAppend](_l.(Expression), _r.(Expression)), nil } /
  _l:expression ws* ">" ws* _r:expression { return MakePiped[PipeOut](_l.(Expression), _r.(Expression)), nil } /
  _l:expression ws* "<" ws* _r:expression { return MakePiped[PipeIn](_r.(Expression), _l.(Expression)), nil } /
  call  /
  invoke /
  reference /
  literal /
  quotedString /
  list /
  lambda /
  "{}" { return MakeConstant(common.Store{}), nil} /
  ( '(' _e: expression ')' { return _e, nil} )

block= '{' eol _l:line* ws* '}' {
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
  = _s:expression+ _r:(_w:ws+ _l:expression+ {
      return Expressions{
        MakeLiteral(" "),
        Expressions(arrayify[Expression](_l)) } , nil})*
    {
      return Expressions(listFrom[Expression](Expressions(arrayify[Expression](_s)), _r)), nil
    }

body = super_expression / block

lambda = _a:argument_decl ws*
  _e:body
  {
    return Lambda {
      Arguments : _a.(IdentifierList),
      Body: _e.(Expression) }, nil
  }

assign = _i:store ws* eq ws* _e:expression
{
  return Assign{
    Store: _i.(StoreAccess),
    Value: _e.(Expression) }, nil
}

assign_add = _i:store ws* "+=" ws* _e:expression
{
  return Assign{
    Store: _i.(StoreAccess),
    Value: MakeAdd(Reference{ Expression: _i.(StoreAccess)}, _e.(Expression)) }, nil
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
  = ("\\" _a:[@!={}()#=&|$<>~'"?%:`+ \][.\\] { return string(_a.([]byte)), nil })
  / (sym / let / dig)+ { return string(c.text), nil}
  / "\\" _s:sym+ { return string(c.text[1:]), nil }
  / "\\" { return string(c.text), nil }

dig = [0-9]
let = [a-zA-Z]
sym = [^\\@!={}()#=&|$<>~'"?%:`+ \n\r\t\][.]

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




