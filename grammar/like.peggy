{
  package grammar
  import "strconv"
}

file = line*;

eol = ('\r'? '\n') / '\r';

line = (comment / code:code comment? { return code, nil}) (eol / eof);

comment = "#" (any !eol)* { return string(c.text), nil };
code = directive;

directive = "// " directive:(include) { return directive, nil };
include = "include" ws+ fn:literal { return Include {FileName: string(fn.(string))}, nil };

quotedString = singleQuotedString / doubleQuotedString{ return string(c.text), nil };
singleQuotedString = "'" ("\\'" / [^'])* "'" { return string(c.text), nil };
doubleQuotedString = '"' ("\\\"" / [^"])* '"'{ return string(c.text), nil };

// integer = [0-9]+ { return string(c.text), nil }

expression_list = expression (ws* expression)*;
call = _i:identifier "(" _m:expression_list? ")";

index = '[' _e:expression ']' { return _e, nil };

member = _m:(identifier / quotedString) _e:index* { return Member{ Identifier: castify(_m, ""), Indexes: arrayify[Expression](_e) }, nil} ;
member_list = _m:member _r:('.' _r:member)* {
  // here
  return (MemberList)(list_from[Member](_m, _r)), nil
};

expression =
  _v: literal { return Value { V: castify(_v, "") }, nil } /
  _m: member_list { return _m, nil }/
  call /
  ( '(' _e: expression ')' { return _e, nil} ) ;

assign = _i:identifier ws* "=" ws* _e:expression
{
  return Assign{
    Identifier: castify(_i, ""),
    Value: _e.(Expression) }, nil
}

output = gt ws* expression
output_line = pipe ws* expression

identifier = (unds / let) (dig / unds / let)* { return string(c.text), nil }

literal_list = l:literal r:(ws+ t:literal { return t, nil} )* { return list_from[string](l, r), nil }

literal = t:( lit+ { return string(c.text), nil }/ s:quotedString { return s, nil}) { return t, nil }

lit = sym / let / dig
dig = [0-9]
let = [a-zA-Z]
sym = dash / unds /slsh / bslh / quot / sign

ws = ' ' / '\t';
dash = '-'
unds = '_'
slsh = '/'
bslh = '\\'
quot = '`'
sign = '$'
pipe = '|'
gt = '>'

any = .
eof = !.




