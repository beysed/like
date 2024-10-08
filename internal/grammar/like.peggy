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

named_expression_list = _e:named_expression _r:( (ws / eol)+ _e:named_expression { return _e, nil} )* {
    return NamedExpressionList(listFrom[NamedExpression](_e, _r)), nil
}

call = _i:reference ws* "(" (ws/eol)* _m:named_expression_list? (ws/eol)* ")" {
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

index
    = '[' _e:expression ']' { return _e, nil }
    / '.' _s:store { return _s, nil };

member = _m:(identifier / s:quotedString) {
  return Member{
    Identifier: castify[string](_m) }, nil
};

reference = sign
  _a: store { return Reference { Expression : _a.(Expression) }, nil }

reference_paren = sign '(' _m:expression ')' { return Reference { Expression : _m.(Expression) }, nil }

reference_block = "${" _e: expression '}' { return _e, nil }

string_content
  = _q:(_b:escaped { return MakeLiteral(_b.(string)), nil}
  / "\\b" { return MakeLiteral("\b"), nil}
  / "\\f" { return MakeLiteral("\f"), nil}
  / "\\n" { return MakeLiteral("\n"), nil}
  / "\\r" { return MakeLiteral("\r"), nil}
  / "\\t" { return MakeLiteral("\t"), nil}
  / call
  / reference
  / reference_paren
  / reference_block
  / ws { return MakeLiteral(string(c.text)), nil }
  /_a:[^ \t\\$]+ { return MakeLiteral(string(c.text)), nil})* eof {
    return Expressions(arrayify[Expression](_q)), nil
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
  _e:identifier {
    return StoreAccess {
      Reference: MakeLiteral(_e.(string)) }, nil} /
  _e:quotedString {
    return StoreAccess {
      Reference: _e.(Expression) }, nil} /
  _e:reference _i:index? {
    var idx Expression
    if _i != nil {
      idx = _i.(Expression)
    }
    return StoreAccess {
      Reference: _e.(Expression),
      Index: idx }, nil}

condition
  =  _c:piped ws+ '?' ws+ _t:piped _f:(ws+ '%' ws+ _f:piped { return _f, nil})? {
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

object
  = "{" ws* _n:named_expression_list? ws* "}" {
    if _n == nil {
      return MakeExpressionStore(NamedExpressionList{})
    }
    return MakeExpressionStore(_n.(NamedExpressionList))
  } /
  ( '(' _e: expression ')' { return _e, nil} );

output =
  writeErr /
  writeErrLn /
  write /
  writeln;

piped =
  not /
  equal /
  not_equal /
  output /
  invoke /
  parse_data /
  format_data /
  condition/
  call  /
  reference /
  reference_paren /
  quotedString /
  list /
  object /
  literal;

not
  = "!" ws* e:piped { return MakeNot(e.(Expression)), nil };

equal
    = _l:piped ws* "==" ws* _r:piped { return MakeEqual(_l.(Expression), _r.(Expression)), nil};

not_equal
    = _l:piped ws* "!=" ws* _r:piped { return MakeNot(MakeEqual(_l.(Expression), _r.(Expression))), nil};

add
    = _l:piped ws* "+" ws* _r:piped { return MakeAdd(_l.(Expression), _r.(Expression)), nil};

pipe
    = _l:pipe ws* "|" ws* _r:piped ws* _e:piped? { return MakePiped[Pipe](_l.(Expression), _r.(Expression), castify[Expression](_e)), nil }
    / piped;

expression =
  assign /
  assign_add /
  parse_data /
  format_data /
  each /
  condition /
  not /
  add /
  equal /
  not_equal /
  lambda /
  _l:piped ws* ">>" ws* _r:piped ws* _e:piped? { return MakePiped[PipeAppend](_l.(Expression), _r.(Expression), castify[Expression](_e)), nil } /
  _l:piped ws* ">" ws* _r:piped ws* _e:piped? { return MakePiped[PipeOut](_l.(Expression), _r.(Expression), castify[Expression](_e)), nil } /
  _l:piped ws* "<" ws* _r:piped { return MakePiped[PipeIn](_r.(Expression), _l.(Expression), nil), nil } /
  pipe /
  invoke /
  output /
  call /
  reference /
  reference_paren /
  quotedString /
  list /
  object /
  literal;

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

super_piped_expression
  = _s:piped+ _r:(_w:ws+ _l:piped+ {
      return Expressions{
        MakeLiteral(" "),
        Expressions(arrayify[Expression](_l)) } , nil})*
    {
      return Expressions(listFrom[Expression](Expressions(arrayify[Expression](_s)), _r)), nil
    }

super_expression
  = _s:expression+ _r:(_w:ws+ _l:expression+ {
      return Expressions{
        MakeLiteral(" "),
        Expressions(arrayify[Expression](_l)) } , nil})*
    {
      return Expressions(listFrom[Expression](Expressions(arrayify[Expression](_s)), _r)), nil
    }

body = lambda / super_expression / block

lambda = _a:argument_decl ws*
  _e:body
  {
    return Lambda {
      Arguments : _a.(IdentifierList),
      Body: _e.(Expression) }, nil
  }

assign = _i:store ws* eq ws* _e:super_expression
{
  return Assign{
    Store: _i.(StoreAccess),
    Value: _e.(Expression) }, nil
}

assign_add = _i:store ws* "+=" ws* _e:super_expression
{
  return Assign{
    Store: _i.(StoreAccess),
    Value: MakeAdd(Reference{ Expression: _i.(StoreAccess)}, MakeUnwrap(_e.(Expression))) }, nil
}

writeln = quot ws* _e:super_piped_expression {
  return WriteLn { Expression : _e.(Expressions) }, nil }

write = tilda ws* _e:super_piped_expression {
  return Write { Expression : _e.(Expressions) }, nil }

writeErrLn = quot '*' ws* _e:super_piped_expression {
  return WriteLn { Expression : _e.(Expressions), Error: true }, nil }

writeErr = tilda '*' ws* _e:super_piped_expression {
  return Write { Expression : _e.(Expressions), Error: true }, nil }

invoke = amper ws* _e:super_piped_expression {
  return Invoke{Expressions: _e.(Expressions)}, nil }

identifier = (unds / let) (dig / unds / let)* {
  return string(c.text), nil }

literal
  = a:lit+ { return MakeLiteral(convert(a)), nil }
  / s:quotedString

escaped
  = '\\' _a:[@!={}()#=&|$<>~'"?%:`+ \][.] { return string(_a.([]byte)), nil }

lit
  = escaped
  / (sym / let / dig)+ { return string(c.text), nil}
  / '\\' { return "\\", nil }

dig = [0-9]
let = [a-zA-Z]
sym = [^\\@!={}()#=&|$<>~'"?%:`+ \n\r\t\][.]

ws = ' ' / '\t'
dash = '-'
unds = '_'
quot = '`'
dblsign = sign sign
sign = '$'
gt = '>'
lt = '<'
eq = '='
amper = '&'
tilda = '~'
colon = ':'
any = . { return string(c.text), nil }
eof = !.




