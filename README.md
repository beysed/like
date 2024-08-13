# Like | Template Engine

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=beysed_like&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=beysed_like)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=beysed_like&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=beysed_like)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=beysed_like&metric=coverage)](https://sonarcloud.io/summary/new_code?id=beysed_like)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=beysed_like&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=beysed_like)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=beysed_like&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=beysed_like)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=beysed_like&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=beysed_like)

[![Build](https://github.com/beysed/like/actions/workflows/build.yml/badge.svg)](https://github.com/beysed/like/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/beysed/like.svg)](https://pkg.go.dev/github.com/beysed/like)
[![Go Report Card](https://goreportcard.com/badge/github.com/beysed/like?04112024)](https://goreportcard.com/report/github.com/beysed/like)

- [Like | Template Engine](#like--template-engine)
  - [Motivation / Get started](#motivation--get-started)
  - [Hello World](#hello-world)
  - [Data types](#data-types)
  - [String interpolation](#string-interpolation)
  - [Control Symbols](#control-symbols)
  - [Output operators](#output-operators)
  - [References and variables](#references-and-variables)
    - [Assigments](#assigments)
  - [Operators and code blocks](#operators-and-code-blocks)
  - [Loops](#loops)
  - [Condition operator](#condition-operator)
  - [Lambdas](#lambdas)
  - [Templates](#templates)
  - [Builtin functions](#builtin-functions)
  - [Closures](#closures)
  - [Parse operator](#parse-operator)
  - [Format opertor](#format-opertor)
  - [Supported parsers formatters](#supported-parsers-formatters)
  - [Arguments](#arguments)
  - [Environmnet variables](#environmnet-variables)
  - [Directives](#directives)
    - [Include](#include)

## Motivation / Get started

During setup a [k8s](https://kubernetes.io/) cluster I start thinking about a language that can make esier to handle tons of mostly identical parts of text and can be/have abilities to be template/meta language. Language as simple as powerful and without extra expressions, so here we are...

## Hello World

Following example writes 'Hello World' to console
```
` Hello World
```

`` ` `` is output operator with new line ending, ``Hello World`` string expression that should be output

## Data types

The language supports the following data types:

- String expression

Any expression that does not contain [control symbols](#control-symbols)

- Quoted strings

The language supports single and double quoted strings, all the difference that inside single quoted string you shoud escape single qoute and double quote for double quoted strings

- Objects

Objects(or key value storage) can be defined following way:

```
a = { my_prop:a other_prop: 'Hello' }
```

- Arrays

Arrays can be defined following way:
```
a = [1 a 'asd' {a: 'object_property'} ]
@ $a `$_v
```

## String interpolation

Inside strings it is possible to use references so
```
a='Hello World'
` He said $a
` 'He said $a'
` "He said $a"
```
Will output: ``He said Hello World`` 3 times

## Control Symbols

- `` ` ``  `~` - [output operators](#output-operators)
- `$` - [reference](#references-and-variables)
- `=` - [assigment operator](#references-and-variables)
- `@` - [loop](#loops) operator
- `?` `%`- if/else [#condition operator](#condition-operator)
- `{}` - code block(see [lambdas](#lambdas) and [templates](#templates)
- `()` - parentheses
- `#` - comment
- `=` - [assigment](#assigments)
- `!` - not operator
- `&` - execute system command
- `|` - pipe operator
- `<>` - less and greater than
- `'` `"` - string quotes

- `:>` `:<`- [parse](#parse-operator) and [format](#format-operator) operators
- `+` - add operator
- `\` - back slash(used for escape control characters)

Any control symbol can be escaped using ``\``

## Output operators

- `` ~ `` - outputs expression without new-line ending
- `` ` `` - outputs expression with `` \n `` new-line ending

## References and variables

### Assigments
```
a = Hello World
a = 'Hello World'
a = "Hello world"
# a = 'Hello World'

a = [Hello world]
# a is array with two elements

a = {Hello: world}
# a is object with property Hello
```

The variable ``a`` can be referenced by ``$`` operator so
```
` $a
```

will output value of ``$a``

## Operators and code blocks

For loops and conditions every operator can be introduced as code-block e.g.

```
a = [Hello world]
@ $a ` $_v
# same as
@ $a {
    ` $_v
}
```

## Loops

`@` the loop operator has following syntax

```
@ <reference-to-array-or-object> <operator>
```

inside a loop there are two predefined variables are assigned for every iteration:
` ``$_k``(key) and ``$_v``(value)

```
a = [Hello world]
@ $a ` "$_k: $_v"
```

will output
```
0: Hello
1: world
```

## Condition operator

Common form of conditional opertors looks following

```
? <reference> <true-operator> % <else-operator>
```

also there is a ternary form for conditions

```
<reference> ? <true-operator> : <else-operator>
```

```
a = Hello
? $a ` World

# will output World
```

Condition operator consider empty(empty string, array, or object) as false and vice versa

## Lambdas

Functions or lambdas can be defined and invoked following way:

```
# without arguments
a = () ` Hello World
$a()

# with agruments
a = (p1 p2 p3) {
    ` $p1 ? $p2 % $p3
}

$a([] TRUE FALSE)
# will output FALSE
```

Lambda returns a value of last evaluated expression

## Templates

Templates can be considered as lambdas but with specific format e.g. temlate consists only with string template

```
`` template_name(p1 p2)
This template text, where
p1 = $p1 and p2 = $p2
``

` $template_name(a b)
```

will output
```
This template text, where
p1 = a and p2 = b
```

## Builtin functions
- joinPath - accepts multiple arguments which will be joined into single path
- resolvePath - acts same as [include](#include) and returns full path
- len - returns length of an expression
- error - raises error and stop further execution
- eval - accept string and executes it as Like program

## Closures

In Like it is possible to create functions on the fly using closures for example
```
a = (b) {
    () {
        ` it was $b
    }
}

# create new lambda with closure to $b
b=$a(Hello)

# executes new lambda
$b()
```

will output
```
it was Hello
```

## Parse operator

`:<` - parse operator it can be useful when you need JSON or other supported by the operator values

for example
```
# we have JSON file with following content
# a.json
# {
#   "environment" : "dev"
# }
#
# reading file
(& cat "file.json") | $f
a = :< json $f
~ $a.environment

# outputs: dev
```

## Format opertor

Format operator turns objects into specified [format](#supported-parsers-formatters)

```
a = { env: dev }
~ :> json $a

# outputs
env: dev
```

## Supported parsers formatters

- json
- yaml
- env

## Arguments

Arguments passed via command line are accessible via predefined variable $args

```
# like file.like one two three
@ $_args ` $_v
```

## Environmnet variables

Environmnet variables are accessible via predefined variable `_env`

```
` $_env[TEMP]
```

also Like try to get file with name ``.env`` from current directory and add/replace environment variables from it

## Directives

### Include

```
#include './some.like'
```

Includes and evaluates the specified file, prefix `` ./ `` means that path spcified from the location of the file(not current directory)