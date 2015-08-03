package tokeq

// wget http://golang.org/ref/spec
var benchTestHTML = `
<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#375EAB">

  <title>The Go Programming Language Specification - The Go Programming Language</title>

<link type="text/css" rel="stylesheet" href="/lib/godoc/style.css">

<link rel="search" type="application/opensearchdescription+xml" title="godoc" href="/opensearch.xml" />

<link rel="stylesheet" href="/lib/godoc/jquery.treeview.css">
<script type="text/javascript">window.initFuncs = [];</script>
<script type="text/javascript">
var _gaq = _gaq || [];
_gaq.push(["_setAccount", "UA-11222381-2"]);
_gaq.push(["b._setAccount", "UA-49880327-6"]);
window.trackPageview = function() {
  _gaq.push(["_trackPageview", location.pathname+location.hash]);
  _gaq.push(["b._trackPageview", location.pathname+location.hash]);
};
window.trackPageview();
window.trackEvent = function(category, action, opt_label, opt_value, opt_noninteraction) {
  _gaq.push(["_trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
  _gaq.push(["b._trackEvent", category, action, opt_label, opt_value, opt_noninteraction]);
};
</script>
</head>
<body>

<div id='lowframe' style="position: fixed; bottom: 0; left: 0; height: 0; width: 100%; border-top: thin solid grey; background-color: white; overflow: auto;">
...
</div><!-- #lowframe -->

<div id="topbar" class="wide"><div class="container">
<div class="top-heading" id="heading-wide"><a href="/">The Go Programming Language</a></div>
<div class="top-heading" id="heading-narrow"><a href="/">Go</a></div>
<a href="#" id="menu-button"><span id="menu-button-arrow">&#9661;</span></a>
<form method="GET" action="/search">
<div id="menu">
<a href="/doc/">Documents</a>
<a href="/pkg/">Packages</a>
<a href="/project/">The Project</a>
<a href="/help/">Help</a>
<a href="/blog/">Blog</a>

<a id="playgroundButton" href="http://play.golang.org/" title="Show Go Playground">Play</a>

<input type="text" id="search" name="q" class="inactive" value="Search" placeholder="Search">
</div>
</form>

</div></div>


<div id="playground" class="play">
	<div class="input"><textarea class="code">package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}</textarea></div>
	<div class="output"></div>
	<div class="buttons">
		<a class="run" title="Run this code [shift-enter]">Run</a>
		<a class="fmt" title="Format this code">Format</a>
		<a class="share" title="Share this code">Share</a>
	</div>
</div>


<div id="page" class="wide">
<div class="container">


  <h1>The Go Programming Language Specification</h1>


  <h2>Version of November 11, 2014</h2>



<div id="nav"></div>




<!--
TODO
[ ] need language about function/method calls and parameter passing rules
[ ] last paragraph of #Assignments (constant promotion) should be elsewhere
    and mention assignment to empty interface.
[ ] need to say something about "scope" of selectors?
[ ] clarify what a field name is in struct declarations
    (struct{T} vs struct {T T} vs struct {t T})
[ ] need explicit language about the result type of operations
[ ] should probably write something about evaluation order of statements even
	though obvious
[ ] in Selectors section, clarify what receiver value is passed in method invocations
-->


<h2 id="Introduction">Introduction</h2>

<p>
This is a reference manual for the Go programming language. For
more information and other documents, see <a href="/">golang.org</a>.
</p>

<p>
Go is a general-purpose language designed with systems programming
in mind. It is strongly typed and garbage-collected and has explicit
support for concurrent programming.  Programs are constructed from
<i>packages</i>, whose properties allow efficient management of
dependencies. The existing implementations use a traditional
compile/link model to generate executable binaries.
</p>

<p>
The grammar is compact and regular, allowing for easy analysis by
automatic tools such as integrated development environments.
</p>

<h2 id="Notation">Notation</h2>
<p>
The syntax is specified using Extended Backus-Naur Form (EBNF):
</p>

<pre class="grammar">
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
</pre>

<p>
Productions are expressions constructed from terms and the following
operators, in increasing precedence:
</p>
<pre class="grammar">
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
</pre>

<p>
Lower-case production names are used to identify lexical tokens.
Non-terminals are in CamelCase. Lexical tokens are enclosed in
double quotes <code>""</code> or back quotes <code></code>.
</p>

<p>
The form <code>a … b</code> represents the set of characters from
<code>a</code> through <code>b</code> as alternatives. The horizontal
ellipsis <code>…</code> is also used elsewhere in the spec to informally denote various
enumerations or code snippets that are not further specified. The character <code>…</code>
(as opposed to the three characters <code>...</code>) is not a token of the Go
language.
</p>

<h2 id="Source_code_representation">Source code representation</h2>

<p>
Source code is Unicode text encoded in
<a href="http://en.wikipedia.org/wiki/UTF-8">UTF-8</a>. The text is not
canonicalized, so a single accented code point is distinct from the
same character constructed from combining an accent and a letter;
those are treated as two code points.  For simplicity, this document
will use the unqualified term <i>character</i> to refer to a Unicode code point
in the source text.
</p>
<p>
Each code point is distinct; for instance, upper and lower case letters
are different characters.
</p>
<p>
Implementation restriction: For compatibility with other tools, a
compiler may disallow the NUL character (U+0000) in the source text.
</p>
<p>
Implementation restriction: For compatibility with other tools, a
compiler may ignore a UTF-8-encoded byte order mark
(U+FEFF) if it is the first Unicode code point in the source text.
A byte order mark may be disallowed anywhere else in the source.
</p>

<h3 id="Characters">Characters</h3>

<p>
The following terms are used to denote specific Unicode character classes:
</p>
<pre class="ebnf">
<a id="newline">newline</a>        = /* the Unicode code point U+000A */ .
<a id="unicode_char">unicode_char</a>   = /* an arbitrary Unicode code point except newline */ .
<a id="unicode_letter">unicode_letter</a> = /* a Unicode code point classified as "Letter" */ .
<a id="unicode_digit">unicode_digit</a>  = /* a Unicode code point classified as "Decimal Digit" */ .
</pre>

<p>
In <a href="http://www.unicode.org/versions/Unicode6.3.0/">The Unicode Standard 6.3</a>,
Section 4.5 "General Category"
defines a set of character categories.  Go treats
those characters in category Lu, Ll, Lt, Lm, or Lo as Unicode letters,
and those in category Nd as Unicode digits.
</p>

<h3 id="Letters_and_digits">Letters and digits</h3>

<p>
The underscore character <code>_</code> (U+005F) is considered a letter.
</p>
<pre class="ebnf">
<a id="letter">letter</a>        = <a href="#unicode_letter" class="noline">unicode_letter</a> | "_" .
<a id="decimal_digit">decimal_digit</a> = "0" … "9" .
<a id="octal_digit">octal_digit</a>   = "0" … "7" .
<a id="hex_digit">hex_digit</a>     = "0" … "9" | "A" … "F" | "a" … "f" .
</pre>

<h2 id="Lexical_elements">Lexical elements</h2>

<h3 id="Comments">Comments</h3>

<p>
There are two forms of comments:
</p>

<ol>
<li>
<i>Line comments</i> start with the character sequence <code>//</code>
and stop at the end of the line. A line comment acts like a newline.
</li>
<li>
<i>General comments</i> start with the character sequence <code>/*</code>
and continue through the character sequence <code>*/</code>. A general
comment containing one or more newlines acts like a newline, otherwise it acts
like a space.
</li>
</ol>

<p>
Comments do not nest.
</p>


<h3 id="Tokens">Tokens</h3>

<p>
Tokens form the vocabulary of the Go language.
There are four classes: <i>identifiers</i>, <i>keywords</i>, <i>operators
and delimiters</i>, and <i>literals</i>.  <i>White space</i>, formed from
spaces (U+0020), horizontal tabs (U+0009),
carriage returns (U+000D), and newlines (U+000A),
is ignored except as it separates tokens
that would otherwise combine into a single token. Also, a newline or end of file
may trigger the insertion of a <a href="#Semicolons">semicolon</a>.
While breaking the input into tokens,
the next token is the longest sequence of characters that form a
valid token.
</p>

<h3 id="Semicolons">Semicolons</h3>

<p>
The formal grammar uses semicolons <code>";"</code> as terminators in
a number of productions. Go programs may omit most of these semicolons
using the following two rules:
</p>

<ol>
<li>
<p>
When the input is broken into tokens, a semicolon is automatically inserted
into the token stream at the end of a non-blank line if the line's final
token is
</p>
<ul>
	<li>an
	    <a href="#Identifiers">identifier</a>
	</li>

	<li>an
	    <a href="#Integer_literals">integer</a>,
	    <a href="#Floating-point_literals">floating-point</a>,
	    <a href="#Imaginary_literals">imaginary</a>,
	    <a href="#Rune_literals">rune</a>, or
	    <a href="#String_literals">string</a> literal
	</li>

	<li>one of the <a href="#Keywords">keywords</a>
	    <code>break</code>,
	    <code>continue</code>,
	    <code>fallthrough</code>, or
	    <code>return</code>
	</li>

	<li>one of the <a href="#Operators_and_Delimiters">operators and delimiters</a>
	    <code>++</code>,
	    <code>--</code>,
	    <code>)</code>,
	    <code>]</code>, or
	    <code>}</code>
	</li>
</ul>
</li>

<li>
To allow complex statements to occupy a single line, a semicolon
may be omitted before a closing <code>")"</code> or <code>"}"</code>.
</li>
</ol>

<p>
To reflect idiomatic use, code examples in this document elide semicolons
using these rules.
</p>


<h3 id="Identifiers">Identifiers</h3>

<p>
Identifiers name program entities such as variables and types.
An identifier is a sequence of one or more letters and digits.
The first character in an identifier must be a letter.
</p>
<pre class="ebnf">
<a id="identifier">identifier</a> = <a href="#letter" class="noline">letter</a> { <a href="#letter" class="noline">letter</a> | <a href="#unicode_digit" class="noline">unicode_digit</a> } .
</pre>
<pre>
a
_x9
ThisVariableIsExported
αβ
</pre>

<p>
Some identifiers are <a href="#Predeclared_identifiers">predeclared</a>.
</p>


<h3 id="Keywords">Keywords</h3>

<p>
The following keywords are reserved and may not be used as identifiers.
</p>
<pre class="grammar">
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
</pre>

<h3 id="Operators_and_Delimiters">Operators and Delimiters</h3>

<p>
The following character sequences represent <a href="#Operators">operators</a>, delimiters, and other special tokens:
</p>
<pre class="grammar">
+    &amp;     +=    &amp;=     &amp;&amp;    ==    !=    (    )
-    |     -=    |=     ||    &lt;     &lt;=    [    ]
*    ^     *=    ^=     &lt;-    &gt;     &gt;=    {    }
/    &lt;&lt;    /=    &lt;&lt;=    ++    =     :=    ,    ;
%    &gt;&gt;    %=    &gt;&gt;=    --    !     ...   .    :
     &amp;^          &amp;^=
</pre>

<h3 id="Integer_literals">Integer literals</h3>

<p>
An integer literal is a sequence of digits representing an
<a href="#Constants">integer constant</a>.
An optional prefix sets a non-decimal base: <code>0</code> for octal, <code>0x</code> or
<code>0X</code> for hexadecimal.  In hexadecimal literals, letters
<code>a-f</code> and <code>A-F</code> represent values 10 through 15.
</p>
<pre class="ebnf">
<a id="int_lit">int_lit</a>     = <a href="#decimal_lit" class="noline">decimal_lit</a> | <a href="#octal_lit" class="noline">octal_lit</a> | <a href="#hex_lit" class="noline">hex_lit</a> .
<a id="decimal_lit">decimal_lit</a> = ( "1" … "9" ) { <a href="#decimal_digit" class="noline">decimal_digit</a> } .
<a id="octal_lit">octal_lit</a>   = "0" { <a href="#octal_digit" class="noline">octal_digit</a> } .
<a id="hex_lit">hex_lit</a>     = "0" ( "x" | "X" ) <a href="#hex_digit" class="noline">hex_digit</a> { <a href="#hex_digit" class="noline">hex_digit</a> } .
</pre>

<pre>
42
0600
0xBadFace
170141183460469231731687303715884105727
</pre>

<h3 id="Floating-point_literals">Floating-point literals</h3>
<p>
A floating-point literal is a decimal representation of a
<a href="#Constants">floating-point constant</a>.
It has an integer part, a decimal point, a fractional part,
and an exponent part.  The integer and fractional part comprise
decimal digits; the exponent part is an <code>e</code> or <code>E</code>
followed by an optionally signed decimal exponent.  One of the
integer part or the fractional part may be elided; one of the decimal
point or the exponent may be elided.
</p>
<pre class="ebnf">
<a id="float_lit">float_lit</a> = <a href="#decimals" class="noline">decimals</a> "." [ <a href="#decimals" class="noline">decimals</a> ] [ <a href="#exponent" class="noline">exponent</a> ] |
            <a href="#decimals" class="noline">decimals</a> <a href="#exponent" class="noline">exponent</a> |
            "." <a href="#decimals" class="noline">decimals</a> [ <a href="#exponent" class="noline">exponent</a> ] .
<a id="decimals">decimals</a>  = <a href="#decimal_digit" class="noline">decimal_digit</a> { <a href="#decimal_digit" class="noline">decimal_digit</a> } .
<a id="exponent">exponent</a>  = ( "e" | "E" ) [ "+" | "-" ] <a href="#decimals" class="noline">decimals</a> .
</pre>

<pre>
0.
72.40
072.40  // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
</pre>

<h3 id="Imaginary_literals">Imaginary literals</h3>
<p>
An imaginary literal is a decimal representation of the imaginary part of a
<a href="#Constants">complex constant</a>.
It consists of a
<a href="#Floating-point_literals">floating-point literal</a>
or decimal integer followed
by the lower-case letter <code>i</code>.
</p>
<pre class="ebnf">
<a id="imaginary_lit">imaginary_lit</a> = (<a href="#decimals" class="noline">decimals</a> | <a href="#float_lit" class="noline">float_lit</a>) "i" .
</pre>

<pre>
0i
011i  // == 11i
0.i
2.71828i
1.e+0i
6.67428e-11i
1E6i
.25i
.12345E+5i
</pre>


<h3 id="Rune_literals">Rune literals</h3>

<p>
A rune literal represents a <a href="#Constants">rune constant</a>,
an integer value identifying a Unicode code point.
A rune literal is expressed as one or more characters enclosed in single quotes.
Within the quotes, any character may appear except single
quote and newline. A single quoted character represents the Unicode value
of the character itself,
while multi-character sequences beginning with a backslash encode
values in various formats.
</p>
<p>
The simplest form represents the single character within the quotes;
since Go source text is Unicode characters encoded in UTF-8, multiple
UTF-8-encoded bytes may represent a single integer value.  For
instance, the literal <code>'a'</code> holds a single byte representing
a literal <code>a</code>, Unicode U+0061, value <code>0x61</code>, while
<code>'ä'</code> holds two bytes (<code>0xc3</code> <code>0xa4</code>) representing
a literal <code>a</code>-dieresis, U+00E4, value <code>0xe4</code>.
</p>
<p>
Several backslash escapes allow arbitrary values to be encoded as
ASCII text.  There are four ways to represent the integer value
as a numeric constant: <code>\x</code> followed by exactly two hexadecimal
digits; <code>\u</code> followed by exactly four hexadecimal digits;
<code>\U</code> followed by exactly eight hexadecimal digits, and a
plain backslash <code>\</code> followed by exactly three octal digits.
In each case the value of the literal is the value represented by
the digits in the corresponding base.
</p>
<p>
Although these representations all result in an integer, they have
different valid ranges.  Octal escapes must represent a value between
0 and 255 inclusive.  Hexadecimal escapes satisfy this condition
by construction. The escapes <code>\u</code> and <code>\U</code>
represent Unicode code points so within them some values are illegal,
in particular those above <code>0x10FFFF</code> and surrogate halves.
</p>
<p>
After a backslash, certain single-character escapes represent special values:
</p>
<pre class="grammar">
\a   U+0007 alert or bell
\b   U+0008 backspace
\f   U+000C form feed
\n   U+000A line feed or newline
\r   U+000D carriage return
\t   U+0009 horizontal tab
\v   U+000b vertical tab
\\   U+005c backslash
\'   U+0027 single quote  (valid escape only within rune literals)
\"   U+0022 double quote  (valid escape only within string literals)
</pre>
<p>
All other sequences starting with a backslash are illegal inside rune literals.
</p>
<pre class="ebnf">
<a id="rune_lit">rune_lit</a>         = "'" ( <a href="#unicode_value" class="noline">unicode_value</a> | <a href="#byte_value" class="noline">byte_value</a> ) "'" .
<a id="unicode_value">unicode_value</a>    = <a href="#unicode_char" class="noline">unicode_char</a> | <a href="#little_u_value" class="noline">little_u_value</a> | <a href="#big_u_value" class="noline">big_u_value</a> | <a href="#escaped_char" class="noline">escaped_char</a> .
<a id="byte_value">byte_value</a>       = <a href="#octal_byte_value" class="noline">octal_byte_value</a> | <a href="#hex_byte_value" class="noline">hex_byte_value</a> .
<a id="octal_byte_value">octal_byte_value</a> = '\' <a href="#octal_digit" class="noline">octal_digit</a> <a href="#octal_digit" class="noline">octal_digit</a> <a href="#octal_digit" class="noline">octal_digit</a> .
<a id="hex_byte_value">hex_byte_value</a>   = '\' "x" <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> .
<a id="little_u_value">little_u_value</a>   = '\' "u" <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> .
<a id="big_u_value">big_u_value</a>      = '\' "U" <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a>
                           <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> <a href="#hex_digit" class="noline">hex_digit</a> .
<a id="escaped_char">escaped_char</a>     = '\' ( "a" | "b" | "f" | "n" | "r" | "t" | "v" | '\' | "'" | '"' ) .
</pre>

<pre>
'a'
'ä'
'本'
'\t'
'\000'
'\007'
'\377'
'\x07'
'\xff'
'\u12e4'
'\U00101234'
'aa'         // illegal: too many characters
'\xa'        // illegal: too few hexadecimal digits
'\0'         // illegal: too few octal digits
'\uDFFF'     // illegal: surrogate half
'\U00110000' // illegal: invalid Unicode code point
</pre>


<h3 id="String_literals">String literals</h3>

<p>
A string literal represents a <a href="#Constants">string constant</a>
obtained from concatenating a sequence of characters. There are two forms:
raw string literals and interpreted string literals.
</p>
<p>
Raw string literals are character sequences between back quotes
<code>''</code>.  Within the quotes, any character is legal except
back quote. The value of a raw string literal is the
string composed of the uninterpreted (implicitly UTF-8-encoded) characters
between the quotes;
in particular, backslashes have no special meaning and the string may
contain newlines.
Carriage return characters ('\r') inside raw string literals
are discarded from the raw string value.
</p>
<p>
Interpreted string literals are character sequences between double
quotes <code>&quot;&quot;</code>. The text between the quotes,
which may not contain newlines, forms the
value of the literal, with backslash escapes interpreted as they
are in <a href="#Rune_literals">rune literals</a> (except that <code>\'</code> is illegal and
<code>\"</code> is legal), with the same restrictions.
The three-digit octal (<code>\</code><i>nnn</i>)
and two-digit hexadecimal (<code>\x</code><i>nn</i>) escapes represent individual
<i>bytes</i> of the resulting string; all other escapes represent
the (possibly multi-byte) UTF-8 encoding of individual <i>characters</i>.
Thus inside a string literal <code>\377</code> and <code>\xFF</code> represent
a single byte of value <code>0xFF</code>=255, while <code>ÿ</code>,
<code>\u00FF</code>, <code>\U000000FF</code> and <code>\xc3\xbf</code> represent
the two bytes <code>0xc3</code> <code>0xbf</code> of the UTF-8 encoding of character
U+00FF.
</p>

<pre class="ebnf">
<a id="string_lit">string_lit</a>             = <a href="#raw_string_lit" class="noline">raw_string_lit</a> | <a href="#interpreted_string_lit" class="noline">interpreted_string_lit</a> .
<a id="raw_string_lit">raw_string_lit</a>         = "'" { <a href="#unicode_char" class="noline">unicode_char</a> | <a href="#newline" class="noline">newline</a> } "'" .
<a id="interpreted_string_lit">interpreted_string_lit</a> = '"' { <a href="#unicode_value" class="noline">unicode_value</a> | <a href="#byte_value" class="noline">byte_value</a> } '"' .
</pre>

<pre>
'abc'  // same as "abc"
'\n
\n'    // same as "\\n\n\\n"
"\n"
""
"Hello, world!\n"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
"\uD800"       // illegal: surrogate half
"\U00110000"   // illegal: invalid Unicode code point
</pre>

<p>
These examples all represent the same string:
</p>

<pre>
"日本語"                                 // UTF-8 input text
'日本語'                                 // UTF-8 input text as a raw literal
"\u65e5\u672c\u8a9e"                    // the explicit Unicode code points
"\U000065e5\U0000672c\U00008a9e"        // the explicit Unicode code points
"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e"  // the explicit UTF-8 bytes
</pre>

<p>
If the source code represents a character as two code points, such as
a combining form involving an accent and a letter, the result will be
an error if placed in a rune literal (it is not a single code
point), and will appear as two code points if placed in a string
literal.
</p>


<h2 id="Constants">Constants</h2>

<p>There are <i>boolean constants</i>,
<i>rune constants</i>,
<i>integer constants</i>,
<i>floating-point constants</i>, <i>complex constants</i>,
and <i>string constants</i>. Rune, integer, floating-point,
and complex constants are
collectively called <i>numeric constants</i>.
</p>

<p>
A constant value is represented by a
<a href="#Rune_literals">rune</a>,
<a href="#Integer_literals">integer</a>,
<a href="#Floating-point_literals">floating-point</a>,
<a href="#Imaginary_literals">imaginary</a>,
or
<a href="#String_literals">string</a> literal,
an identifier denoting a constant,
a <a href="#Constant_expressions">constant expression</a>,
a <a href="#Conversions">conversion</a> with a result that is a constant, or
the result value of some built-in functions such as
<code>unsafe.Sizeof</code> applied to any value,
<code>cap</code> or <code>len</code> applied to
<a href="#Length_and_capacity">some expressions</a>,
<code>real</code> and <code>imag</code> applied to a complex constant
and <code>complex</code> applied to numeric constants.
The boolean truth values are represented by the predeclared constants
<code>true</code> and <code>false</code>. The predeclared identifier
<a href="#Iota">iota</a> denotes an integer constant.
</p>

<p>
In general, complex constants are a form of
<a href="#Constant_expressions">constant expression</a>
and are discussed in that section.
</p>

<p>
Numeric constants represent values of arbitrary precision and do not overflow.
</p>

<p>
Constants may be <a href="#Types">typed</a> or <i>untyped</i>.
Literal constants, <code>true</code>, <code>false</code>, <code>iota</code>,
and certain <a href="#Constant_expressions">constant expressions</a>
containing only untyped constant operands are untyped.
</p>

<p>
A constant may be given a type explicitly by a <a href="#Constant_declarations">constant declaration</a>
or <a href="#Conversions">conversion</a>, or implicitly when used in a
<a href="#Variable_declarations">variable declaration</a> or an
<a href="#Assignments">assignment</a> or as an
operand in an <a href="#Expressions">expression</a>.
It is an error if the constant value
cannot be represented as a value of the respective type.
For instance, <code>3.0</code> can be given any integer or any
floating-point type, while <code>2147483648.0</code> (equal to <code>1&lt;&lt;31</code>)
can be given the types <code>float32</code>, <code>float64</code>, or <code>uint32</code> but
not <code>int32</code> or <code>string</code>.
</p>

<p>
An untyped constant has a <i>default type</i> which is the type to which the
constant is implicitly converted in contexts where a typed value is required,
for instance, in a <a href="#Short_variable_declarations">short variable declaration</a>
such as <code>i := 0</code> where there is no explicit type.
The default type of an untyped constant is <code>bool</code>, <code>rune</code>,
<code>int</code>, <code>float64</code>, <code>complex128</code> or <code>string</code>
respectively, depending on whether it is a boolean, rune, integer, floating-point,
complex, or string constant.
</p>

<p>
There are no constants denoting the IEEE-754 infinity and not-a-number values,
but the <a href="/pkg/math/"><code>math</code> package</a>'s
<a href="/pkg/math/#Inf">Inf</a>,
<a href="/pkg/math/#NaN">NaN</a>,
<a href="/pkg/math/#IsInf">IsInf</a>, and
<a href="/pkg/math/#IsNaN">IsNaN</a>
functions return and test for those values at run time.
</p>

<p>
Implementation restriction: Although numeric constants have arbitrary
precision in the language, a compiler may implement them using an
internal representation with limited precision.  That said, every
implementation must:
</p>
<ul>
	<li>Represent integer constants with at least 256 bits.</li>

	<li>Represent floating-point constants, including the parts of
	    a complex constant, with a mantissa of at least 256 bits
	    and a signed exponent of at least 32 bits.</li>

	<li>Give an error if unable to represent an integer constant
	    precisely.</li>

	<li>Give an error if unable to represent a floating-point or
	    complex constant due to overflow.</li>

	<li>Round to the nearest representable constant if unable to
	    represent a floating-point or complex constant due to limits
	    on precision.</li>
</ul>
<p>
These requirements apply both to literal constants and to the result
of evaluating <a href="#Constant_expressions">constant
expressions</a>.
</p>

<h2 id="Variables">Variables</h2>

<p>
A variable is a storage location for holding a <i>value</i>.
The set of permissible values is determined by the
variable's <i><a href="#Types">type</a></i>.
</p>

<p>
A <a href="#Variable_declarations">variable declaration</a>
or, for function parameters and results, the signature
of a <a href="#Function_declarations">function declaration</a>
or <a href="#Function_literals">function literal</a> reserves
storage for a named variable.

Calling the built-in function <a href="#Allocation"><code>new</code></a>
or taking the address of a <a href="#Composite_literals">composite literal</a>
allocates storage for a variable at run time.
Such an anonymous variable is referred to via a (possibly implicit)
<a href="#Address_operators">pointer indirection</a>.
</p>

<p>
<i>Structured</i> variables of <a href="#Array_types">array</a>, <a href="#Slice_types">slice</a>,
and <a href="#Struct_types">struct</a> types have elements and fields that may
be <a href="#Address_operators">addressed</a> individually. Each such element
acts like a variable.
</p>

<p>
The <i>static type</i> (or just <i>type</i>) of a variable is the	
type given in its declaration, the type provided in the
<code>new</code> call or composite literal, or the type of
an element of a structured variable.
Variables of interface type also have a distinct <i>dynamic type</i>,
which is the concrete type of the value assigned to the variable at run time
(unless the value is the predeclared identifier <code>nil</code>,
which has no type).
The dynamic type may vary during execution but values stored in interface
variables are always <a href="#Assignability">assignable</a>
to the static type of the variable.	
</p>	

<pre>
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
</pre>

<p>
A variable's value is retrieved by referring to the variable in an
<a href="#Expressions">expression</a>; it is the most recent value
<a href="#Assignments">assigned</a> to the variable.
If a variable has not yet been assigned a value, its value is the
<a href="#The_zero_value">zero value</a> for its type.
</p>


<h2 id="Types">Types</h2>

<p>
A type determines the set of values and operations specific to values of that
type. Types may be <i>named</i> or <i>unnamed</i>. Named types are specified
by a (possibly <a href="#Qualified_identifiers">qualified</a>)
<a href="#Type_declarations"><i>type name</i></a>; unnamed types are specified
using a <i>type literal</i>, which composes a new type from existing types.
</p>

<pre class="ebnf">
<a id="Type">Type</a>      = <a href="#TypeName" class="noline">TypeName</a> | <a href="#TypeLit" class="noline">TypeLit</a> | "(" <a href="#Type" class="noline">Type</a> ")" .
<a id="TypeName">TypeName</a>  = <a href="#identifier" class="noline">identifier</a> | <a href="#QualifiedIdent" class="noline">QualifiedIdent</a> .
<a id="TypeLit">TypeLit</a>   = <a href="#ArrayType" class="noline">ArrayType</a> | <a href="#StructType" class="noline">StructType</a> | <a href="#PointerType" class="noline">PointerType</a> | <a href="#FunctionType" class="noline">FunctionType</a> | <a href="#InterfaceType" class="noline">InterfaceType</a> |
	    <a href="#SliceType" class="noline">SliceType</a> | <a href="#MapType" class="noline">MapType</a> | <a href="#ChannelType" class="noline">ChannelType</a> .
</pre>

<p>
Named instances of the boolean, numeric, and string types are
<a href="#Predeclared_identifiers">predeclared</a>.
<i>Composite types</i>&mdash;array, struct, pointer, function,
interface, slice, map, and channel types&mdash;may be constructed using
type literals.
</p>

<p>
Each type <code>T</code> has an <i>underlying type</i>: If <code>T</code>
is one of the predeclared boolean, numeric, or string types, or a type literal,
the corresponding underlying
type is <code>T</code> itself. Otherwise, <code>T</code>'s underlying type
is the underlying type of the type to which <code>T</code> refers in its
<a href="#Type_declarations">type declaration</a>.
</p>

<pre>
   type T1 string
   type T2 T1
   type T3 []T1
   type T4 T3
</pre>

<p>
The underlying type of <code>string</code>, <code>T1</code>, and <code>T2</code>
is <code>string</code>. The underlying type of <code>[]T1</code>, <code>T3</code>,
and <code>T4</code> is <code>[]T1</code>.
</p>

<h3 id="Method_sets">Method sets</h3>
<p>
A type may have a <i>method set</i> associated with it.
The method set of an <a href="#Interface_types">interface type</a> is its interface.
The method set of any other type <code>T</code> consists of all
<a href="#Method_declarations">methods</a> declared with receiver type <code>T</code>.
The method set of the corresponding <a href="#Pointer_types">pointer type</a> <code>*T</code>
is the set of all methods declared with receiver <code>*T</code> or <code>T</code>
(that is, it also contains the method set of <code>T</code>).
Further rules apply to structs containing anonymous fields, as described
in the section on <a href="#Struct_types">struct types</a>.
Any other type has an empty method set.
In a method set, each method must have a
<a href="#Uniqueness_of_identifiers">unique</a>
non-<a href="#Blank_identifier">blank</a> <a href="#MethodName">method name</a>.
</p>

<p>
The method set of a type determines the interfaces that the
type <a href="#Interface_types">implements</a>
and the methods that can be <a href="#Calls">called</a>
using a receiver of that type.
</p>

<h3 id="Boolean_types">Boolean types</h3>

<p>
A <i>boolean type</i> represents the set of Boolean truth values
denoted by the predeclared constants <code>true</code>
and <code>false</code>. The predeclared boolean type is <code>bool</code>.
</p>

<h3 id="Numeric_types">Numeric types</h3>

<p>
A <i>numeric type</i> represents sets of integer or floating-point values.
The predeclared architecture-independent numeric types are:
</p>

<pre class="grammar">
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
</pre>

<p>
The value of an <i>n</i>-bit integer is <i>n</i> bits wide and represented using
<a href="http://en.wikipedia.org/wiki/Two's_complement">two's complement arithmetic</a>.
</p>

<p>
There is also a set of predeclared numeric types with implementation-specific sizes:
</p>

<pre class="grammar">
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
</pre>

<p>
To avoid portability issues all numeric types are distinct except
<code>byte</code>, which is an alias for <code>uint8</code>, and
<code>rune</code>, which is an alias for <code>int32</code>.
Conversions
are required when different numeric types are mixed in an expression
or assignment. For instance, <code>int32</code> and <code>int</code>
are not the same type even though they may have the same size on a
particular architecture.


<h3 id="String_types">String types</h3>

<p>
A <i>string type</i> represents the set of string values.
A string value is a (possibly empty) sequence of bytes.
Strings are immutable: once created,
it is impossible to change the contents of a string.
The predeclared string type is <code>string</code>.
</p>

<p>
The length of a string <code>s</code> (its size in bytes) can be discovered using
the built-in function <a href="#Length_and_capacity"><code>len</code></a>.
The length is a compile-time constant if the string is a constant.
A string's bytes can be accessed by integer <a href="#Index_expressions">indices</a>
0 through <code>len(s)-1</code>.
It is illegal to take the address of such an element; if
<code>s[i]</code> is the <code>i</code>'th byte of a
string, <code>&amp;s[i]</code> is invalid.
</p>


<h3 id="Array_types">Array types</h3>

<p>
An array is a numbered sequence of elements of a single
type, called the element type.
The number of elements is called the length and is never
negative.
</p>

<pre class="ebnf">
<a id="ArrayType">ArrayType</a>   = "[" <a href="#ArrayLength" class="noline">ArrayLength</a> "]" <a href="#ElementType" class="noline">ElementType</a> .
<a id="ArrayLength">ArrayLength</a> = <a href="#Expression" class="noline">Expression</a> .
<a id="ElementType">ElementType</a> = <a href="#Type" class="noline">Type</a> .
</pre>

<p>
The length is part of the array's type; it must evaluate to a
non-negative <a href="#Constants">constant</a> representable by a value
of type <code>int</code>.
The length of array <code>a</code> can be discovered
using the built-in function <a href="#Length_and_capacity"><code>len</code></a>.
The elements can be addressed by integer <a href="#Index_expressions">indices</a>
0 through <code>len(a)-1</code>.
Array types are always one-dimensional but may be composed to form
multi-dimensional types.
</p>

<pre>
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
</pre>

<h3 id="Slice_types">Slice types</h3>

<p>
A slice is a descriptor for a contiguous segment of an <i>underlying array</i> and
provides access to a numbered sequence of elements from that array.
A slice type denotes the set of all slices of arrays of its element type.
The value of an uninitialized slice is <code>nil</code>.
</p>

<pre class="ebnf">
<a id="SliceType">SliceType</a> = "[" "]" <a href="#ElementType" class="noline">ElementType</a> .
</pre>

<p>
Like arrays, slices are indexable and have a length.  The length of a
slice <code>s</code> can be discovered by the built-in function
<a href="#Length_and_capacity"><code>len</code></a>; unlike with arrays it may change during
execution.  The elements can be addressed by integer <a href="#Index_expressions">indices</a>
0 through <code>len(s)-1</code>.  The slice index of a
given element may be less than the index of the same element in the
underlying array.
</p>
<p>
A slice, once initialized, is always associated with an underlying
array that holds its elements.  A slice therefore shares storage
with its array and with other slices of the same array; by contrast,
distinct arrays always represent distinct storage.
</p>
<p>
The array underlying a slice may extend past the end of the slice.
The <i>capacity</i> is a measure of that extent: it is the sum of
the length of the slice and the length of the array beyond the slice;
a slice of length up to that capacity can be created by
<a href="#Slice_expressions"><i>slicing</i></a> a new one from the original slice.
The capacity of a slice <code>a</code> can be discovered using the
built-in function <a href="#Length_and_capacity"><code>cap(a)</code></a>.
</p>

<p>
A new, initialized slice value for a given element type <code>T</code> is
made using the built-in function
<a href="#Making_slices_maps_and_channels"><code>make</code></a>,
which takes a slice type
and parameters specifying the length and optionally the capacity.
A slice created with <code>make</code> always allocates a new, hidden array
to which the returned slice value refers. That is, executing
</p>

<pre>
make([]T, length, capacity)
</pre>

<p>
produces the same slice as allocating an array and <a href="#Slice_expressions">slicing</a>
it, so these two expressions are equivalent:
</p>

<pre>
make([]int, 50, 100)
new([100]int)[0:50]
</pre>

<p>
Like arrays, slices are always one-dimensional but may be composed to construct
higher-dimensional objects.
With arrays of arrays, the inner arrays are, by construction, always the same length;
however with slices of slices (or arrays of slices), the inner lengths may vary dynamically.
Moreover, the inner slices must be initialized individually.
</p>

<h3 id="Struct_types">Struct types</h3>

<p>
A struct is a sequence of named elements, called fields, each of which has a
name and a type. Field names may be specified explicitly (IdentifierList) or
implicitly (AnonymousField).
Within a struct, non-<a href="#Blank_identifier">blank</a> field names must
be <a href="#Uniqueness_of_identifiers">unique</a>.
</p>

<pre class="ebnf">
<a id="StructType">StructType</a>     = "struct" "{" { <a href="#FieldDecl" class="noline">FieldDecl</a> ";" } "}" .
<a id="FieldDecl">FieldDecl</a>      = (<a href="#IdentifierList" class="noline">IdentifierList</a> <a href="#Type" class="noline">Type</a> | <a href="#AnonymousField" class="noline">AnonymousField</a>) [ <a href="#Tag" class="noline">Tag</a> ] .
<a id="AnonymousField">AnonymousField</a> = [ "*" ] <a href="#TypeName" class="noline">TypeName</a> .
<a id="Tag">Tag</a>            = <a href="#string_lit" class="noline">string_lit</a> .
</pre>

<pre>
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
</pre>

<p>
A field declared with a type but no explicit field name is an <i>anonymous field</i>,
also called an <i>embedded</i> field or an embedding of the type in the struct.
An embedded type must be specified as
a type name <code>T</code> or as a pointer to a non-interface type name <code>*T</code>,
and <code>T</code> itself may not be
a pointer type. The unqualified type name acts as the field name.
</p>

<pre>
// A struct with four anonymous fields of type T1, *T2, P.T3 and *P.T4
struct {
	T1        // field name is T1
	*T2       // field name is T2
	P.T3      // field name is T3
	*P.T4     // field name is T4
	x, y int  // field names are x and y
}
</pre>

<p>
The following declaration is illegal because field names must be unique
in a struct type:
</p>

<pre>
struct {
	T     // conflicts with anonymous field *T and *P.T
	*T    // conflicts with anonymous field T and *P.T
	*P.T  // conflicts with anonymous field T and *T
}
</pre>

<p>
A field or <a href="#Method_declarations">method</a> <code>f</code> of an
anonymous field in a struct <code>x</code> is called <i>promoted</i> if
<code>x.f</code> is a legal <a href="#Selectors">selector</a> that denotes
that field or method <code>f</code>.
</p>

<p>
Promoted fields act like ordinary fields
of a struct except that they cannot be used as field names in
<a href="#Composite_literals">composite literals</a> of the struct.
</p>

<p>
Given a struct type <code>S</code> and a type named <code>T</code>,
promoted methods are included in the method set of the struct as follows:
</p>
<ul>
	<li>
	If <code>S</code> contains an anonymous field <code>T</code>,
	the <a href="#Method_sets">method sets</a> of <code>S</code>
	and <code>*S</code> both include promoted methods with receiver
	<code>T</code>. The method set of <code>*S</code> also
	includes promoted methods with receiver <code>*T</code>.
	</li>

	<li>
	If <code>S</code> contains an anonymous field <code>*T</code>,
	the method sets of <code>S</code> and <code>*S</code> both
	include promoted methods with receiver <code>T</code> or
	<code>*T</code>.
	</li>
</ul>

<p>
A field declaration may be followed by an optional string literal <i>tag</i>,
which becomes an attribute for all the fields in the corresponding
field declaration. The tags are made
visible through a <a href="/pkg/reflect/#StructTag">reflection interface</a>
and take part in <a href="#Type_identity">type identity</a> for structs
but are otherwise ignored.
</p>

<pre>
// A struct corresponding to the TimeStamp protocol buffer.
// The tag strings define the protocol buffer field numbers.
struct {
	microsec  uint64 "field 1"
	serverIP6 uint64 "field 2"
	process   string "field 3"
}
</pre>

<h3 id="Pointer_types">Pointer types</h3>

<p>
A pointer type denotes the set of all pointers to <a href="#Variables">variables</a> of a given
type, called the <i>base type</i> of the pointer.
The value of an uninitialized pointer is <code>nil</code>.
</p>

<pre class="ebnf">
<a id="PointerType">PointerType</a> = "*" <a href="#BaseType" class="noline">BaseType</a> .
<a id="BaseType">BaseType</a>    = <a href="#Type" class="noline">Type</a> .
</pre>

<pre>
*Point
*[4]int
</pre>

<h3 id="Function_types">Function types</h3>

<p>
A function type denotes the set of all functions with the same parameter
and result types. The value of an uninitialized variable of function type
is <code>nil</code>.
</p>

<pre class="ebnf">
<a id="FunctionType">FunctionType</a>   = "func" <a href="#Signature" class="noline">Signature</a> .
<a id="Signature">Signature</a>      = <a href="#Parameters" class="noline">Parameters</a> [ <a href="#Result" class="noline">Result</a> ] .
<a id="Result">Result</a>         = <a href="#Parameters" class="noline">Parameters</a> | <a href="#Type" class="noline">Type</a> .
<a id="Parameters">Parameters</a>     = "(" [ <a href="#ParameterList" class="noline">ParameterList</a> [ "," ] ] ")" .
<a id="ParameterList">ParameterList</a>  = <a href="#ParameterDecl" class="noline">ParameterDecl</a> { "," <a href="#ParameterDecl" class="noline">ParameterDecl</a> } .
<a id="ParameterDecl">ParameterDecl</a>  = [ <a href="#IdentifierList" class="noline">IdentifierList</a> ] [ "..." ] <a href="#Type" class="noline">Type</a> .
</pre>

<p>
Within a list of parameters or results, the names (IdentifierList)
must either all be present or all be absent. If present, each name
stands for one item (parameter or result) of the specified type and
all non-<a href="#Blank_identifier">blank</a> names in the signature
must be <a href="#Uniqueness_of_identifiers">unique</a>.
If absent, each type stands for one item of that type.
Parameter and result
lists are always parenthesized except that if there is exactly
one unnamed result it may be written as an unparenthesized type.
</p>

<p>
The final parameter in a function signature may have
a type prefixed with <code>...</code>.
A function with such a parameter is called <i>variadic</i> and
may be invoked with zero or more arguments for that parameter.
</p>

<pre>
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
</pre>


<h3 id="Interface_types">Interface types</h3>

<p>
An interface type specifies a <a href="#Method_sets">method set</a> called its <i>interface</i>.
A variable of interface type can store a value of any type with a method set
that is any superset of the interface. Such a type is said to
<i>implement the interface</i>.
The value of an uninitialized variable of interface type is <code>nil</code>.
</p>

<pre class="ebnf">
<a id="InterfaceType">InterfaceType</a>      = "interface" "{" { <a href="#MethodSpec" class="noline">MethodSpec</a> ";" } "}" .
<a id="MethodSpec">MethodSpec</a>         = <a href="#MethodName" class="noline">MethodName</a> <a href="#Signature" class="noline">Signature</a> | <a href="#InterfaceTypeName" class="noline">InterfaceTypeName</a> .
<a id="MethodName">MethodName</a>         = <a href="#identifier" class="noline">identifier</a> .
<a id="InterfaceTypeName">InterfaceTypeName</a>  = <a href="#TypeName" class="noline">TypeName</a> .
</pre>

<p>
As with all method sets, in an interface type, each method must have a
<a href="#Uniqueness_of_identifiers">unique</a>
non-<a href="#Blank_identifier">blank</a> name.
</p>

<pre>
// A simple File interface
interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
	Close()
}
</pre>

<p>
More than one type may implement an interface.
For instance, if two types <code>S1</code> and <code>S2</code>
have the method set
</p>

<pre>
func (p T) Read(b Buffer) bool { return … }
func (p T) Write(b Buffer) bool { return … }
func (p T) Close() { … }
</pre>

<p>
(where <code>T</code> stands for either <code>S1</code> or <code>S2</code>)
then the <code>File</code> interface is implemented by both <code>S1</code> and
<code>S2</code>, regardless of what other methods
<code>S1</code> and <code>S2</code> may have or share.
</p>

<p>
A type implements any interface comprising any subset of its methods
and may therefore implement several distinct interfaces. For
instance, all types implement the <i>empty interface</i>:
</p>

<pre>
interface{}
</pre>

<p>
Similarly, consider this interface specification,
which appears within a <a href="#Type_declarations">type declaration</a>
to define an interface called <code>Locker</code>:
</p>

<pre>
type Locker interface {
	Lock()
	Unlock()
}
</pre>

<p>
If <code>S1</code> and <code>S2</code> also implement
</p>

<pre>
func (p T) Lock() { … }
func (p T) Unlock() { … }
</pre>

<p>
they implement the <code>Locker</code> interface as well
as the <code>File</code> interface.
</p>

<p>
An interface <code>T</code> may use a (possibly qualified) interface type
name <code>E</code> in place of a method specification. This is called
<i>embedding</i> interface <code>E</code> in <code>T</code>; it adds
all (exported and non-exported) methods of <code>E</code> to the interface
<code>T</code>.
</p>

<pre>
type ReadWriter interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}

type File interface {
	ReadWriter  // same as adding the methods of ReadWriter
	Locker      // same as adding the methods of Locker
	Close()
}

type LockedFile interface {
	Locker
	File        // illegal: Lock, Unlock not unique
	Lock()      // illegal: Lock not unique
}
</pre>

<p>
An interface type <code>T</code> may not embed itself
or any interface type that embeds <code>T</code>, recursively.
</p>

<pre>
// illegal: Bad cannot embed itself
type Bad interface {
	Bad
}

// illegal: Bad1 cannot embed itself using Bad2
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}
</pre>

<h3 id="Map_types">Map types</h3>

<p>
A map is an unordered group of elements of one type, called the
element type, indexed by a set of unique <i>keys</i> of another type,
called the key type.
The value of an uninitialized map is <code>nil</code>.
</p>

<pre class="ebnf">
<a id="MapType">MapType</a>     = "map" "[" <a href="#KeyType" class="noline">KeyType</a> "]" <a href="#ElementType" class="noline">ElementType</a> .
<a id="KeyType">KeyType</a>     = <a href="#Type" class="noline">Type</a> .
</pre>

<p>
The <a href="#Comparison_operators">comparison operators</a>
<code>==</code> and <code>!=</code> must be fully defined
for operands of the key type; thus the key type must not be a function, map, or
slice.
If the key type is an interface type, these
comparison operators must be defined for the dynamic key values;
failure will cause a <a href="#Run_time_panics">run-time panic</a>.

</p>

<pre>
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
</pre>

<p>
The number of map elements is called its length.
For a map <code>m</code>, it can be discovered using the
built-in function <a href="#Length_and_capacity"><code>len</code></a>
and may change during execution. Elements may be added during execution
using <a href="#Assignments">assignments</a> and retrieved with
<a href="#Index_expressions">index expressions</a>; they may be removed with the
<a href="#Deletion_of_map_elements"><code>delete</code></a> built-in function.
</p>
<p>
A new, empty map value is made using the built-in
function <a href="#Making_slices_maps_and_channels"><code>make</code></a>,
which takes the map type and an optional capacity hint as arguments:
</p>

<pre>
make(map[string]int)
make(map[string]int, 100)
</pre>

<p>
The initial capacity does not bound its size:
maps grow to accommodate the number of items
stored in them, with the exception of <code>nil</code> maps.
A <code>nil</code> map is equivalent to an empty map except that no elements
may be added.

<h3 id="Channel_types">Channel types</h3>

<p>
A channel provides a mechanism for
<a href="#Go_statements">concurrently executing functions</a>
to communicate by
<a href="#Send_statements">sending</a> and
<a href="#Receive_operator">receiving</a>
values of a specified element type.
The value of an uninitialized channel is <code>nil</code>.
</p>

<pre class="ebnf">
<a id="ChannelType">ChannelType</a> = ( "chan" | "chan" "&lt;-" | "&lt;-" "chan" ) <a href="#ElementType" class="noline">ElementType</a> .
</pre>

<p>
The optional <code>&lt;-</code> operator specifies the channel <i>direction</i>,
<i>send</i> or <i>receive</i>. If no direction is given, the channel is
<i>bidirectional</i>.
A channel may be constrained only to send or only to receive by
<a href="#Conversions">conversion</a> or <a href="#Assignments">assignment</a>.
</p>

<pre>
chan T          // can be used to send and receive values of type T
chan&lt;- float64  // can only be used to send float64s
&lt;-chan int      // can only be used to receive ints
</pre>

<p>
The <code>&lt;-</code> operator associates with the leftmost <code>chan</code>
possible:
</p>

<pre>
chan&lt;- chan int    // same as chan&lt;- (chan int)
chan&lt;- &lt;-chan int  // same as chan&lt;- (&lt;-chan int)
&lt;-chan &lt;-chan int  // same as &lt;-chan (&lt;-chan int)
chan (&lt;-chan int)
</pre>

<p>
A new, initialized channel
value can be made using the built-in function
<a href="#Making_slices_maps_and_channels"><code>make</code></a>,
which takes the channel type and an optional <i>capacity</i> as arguments:
</p>

<pre>
make(chan int, 100)
</pre>

<p>
The capacity, in number of elements, sets the size of the buffer in the channel.
If the capacity is zero or absent, the channel is unbuffered and communication
succeeds only when both a sender and receiver are ready. Otherwise, the channel
is buffered and communication succeeds without blocking if the buffer
is not full (sends) or not empty (receives).
A <code>nil</code> channel is never ready for communication.
</p>

<p>
A channel may be closed with the built-in function
<a href="#Close"><code>close</code></a>.
The multi-valued assignment form of the
<a href="#Receive_operator">receive operator</a>
reports whether a received value was sent before
the channel was closed.
</p>

<p>
A single channel may be used in
<a href="#Send_statements">send statements</a>,
<a href="#Receive_operator">receive operations</a>,
and calls to the built-in functions
<a href="#Length_and_capacity"><code>cap</code></a> and
<a href="#Length_and_capacity"><code>len</code></a>
by any number of goroutines without further synchronization.
Channels act as first-in-first-out queues.
For example, if one goroutine sends values on a channel
and a second goroutine receives them, the values are
received in the order sent.
</p>

<h2 id="Properties_of_types_and_values">Properties of types and values</h2>

<h3 id="Type_identity">Type identity</h3>

<p>
Two types are either <i>identical</i> or <i>different</i>.
</p>

<p>
Two <a href="#Types">named types</a> are identical if their type names originate in the same
<a href="#Type_declarations">TypeSpec</a>.
A named and an <a href="#Types">unnamed type</a> are always different. Two unnamed types are identical
if the corresponding type literals are identical, that is, if they have the same
literal structure and corresponding components have identical types. In detail:
</p>

<ul>
	<li>Two array types are identical if they have identical element types and
	    the same array length.</li>

	<li>Two slice types are identical if they have identical element types.</li>

	<li>Two struct types are identical if they have the same sequence of fields,
	    and if corresponding fields have the same names, and identical types,
	    and identical tags.
	    Two anonymous fields are considered to have the same name. Lower-case field
	    names from different packages are always different.</li>

	<li>Two pointer types are identical if they have identical base types.</li>

	<li>Two function types are identical if they have the same number of parameters
	    and result values, corresponding parameter and result types are
	    identical, and either both functions are variadic or neither is.
	    Parameter and result names are not required to match.</li>

	<li>Two interface types are identical if they have the same set of methods
	    with the same names and identical function types. Lower-case method names from
	    different packages are always different. The order of the methods is irrelevant.</li>

	<li>Two map types are identical if they have identical key and value types.</li>

	<li>Two channel types are identical if they have identical value types and
	    the same direction.</li>
</ul>

<p>
Given the declarations
</p>

<pre>
type (
	T0 []string
	T1 []string
	T2 struct{ a, b int }
	T3 struct{ a, c int }
	T4 func(int, float64) *T0
	T5 func(x int, y float64) *[]string
)
</pre>

<p>
these types are identical:
</p>

<pre>
T0 and T0
[]int and []int
struct{ a, b *T5 } and struct{ a, b *T5 }
func(x int, y float64) *[]string and func(int, float64) (result *[]string)
</pre>

<p>
<code>T0</code> and <code>T1</code> are different because they are named types
with distinct declarations; <code>func(int, float64) *T0</code> and
<code>func(x int, y float64) *[]string</code> are different because <code>T0</code>
is different from <code>[]string</code>.
</p>


<h3 id="Assignability">Assignability</h3>

<p>
A value <code>x</code> is <i>assignable</i> to a <a href="#Variables">variable</a> of type <code>T</code>
("<code>x</code> is assignable to <code>T</code>") in any of these cases:
</p>

<ul>
<li>
<code>x</code>'s type is identical to <code>T</code>.
</li>
<li>
<code>x</code>'s type <code>V</code> and <code>T</code> have identical
<a href="#Types">underlying types</a> and at least one of <code>V</code>
or <code>T</code> is not a <a href="#Types">named type</a>.
</li>
<li>
<code>T</code> is an interface type and
<code>x</code> <a href="#Interface_types">implements</a> <code>T</code>.
</li>
<li>
<code>x</code> is a bidirectional channel value, <code>T</code> is a channel type,
<code>x</code>'s type <code>V</code> and <code>T</code> have identical element types,
and at least one of <code>V</code> or <code>T</code> is not a named type.
</li>
<li>
<code>x</code> is the predeclared identifier <code>nil</code> and <code>T</code>
is a pointer, function, slice, map, channel, or interface type.
</li>
<li>
<code>x</code> is an untyped <a href="#Constants">constant</a> representable
by a value of type <code>T</code>.
</li>
</ul>


<h2 id="Blocks">Blocks</h2>

<p>
A <i>block</i> is a possibly empty sequence of declarations and statements
within matching brace brackets.
</p>

<pre class="ebnf">
<a id="Block">Block</a> = "{" <a href="#StatementList" class="noline">StatementList</a> "}" .
<a id="StatementList">StatementList</a> = { <a href="#Statement" class="noline">Statement</a> ";" } .
</pre>

<p>
In addition to explicit blocks in the source code, there are implicit blocks:
</p>

<ol>
	<li>The <i>universe block</i> encompasses all Go source text.</li>

	<li>Each <a href="#Packages">package</a> has a <i>package block</i> containing all
	    Go source text for that package.</li>

	<li>Each file has a <i>file block</i> containing all Go source text
	    in that file.</li>

	<li>Each <a href="#If_statements">"if"</a>,
	    <a href="#For_statements">"for"</a>, and
	    <a href="#Switch_statements">"switch"</a>
	    statement is considered to be in its own implicit block.</li>

	<li>Each clause in a <a href="#Switch_statements">"switch"</a>
	    or <a href="#Select_statements">"select"</a> statement
	    acts as an implicit block.</li>
</ol>

<p>
Blocks nest and influence <a href="#Declarations_and_scope">scoping</a>.
</p>


<h2 id="Declarations_and_scope">Declarations and scope</h2>

<p>
A <i>declaration</i> binds a non-<a href="#Blank_identifier">blank</a> identifier to a
<a href="#Constant_declarations">constant</a>,
<a href="#Type_declarations">type</a>,
<a href="#Variable_declarations">variable</a>,
<a href="#Function_declarations">function</a>,
<a href="#Labeled_statements">label</a>, or
<a href="#Import_declarations">package</a>.
Every identifier in a program must be declared.
No identifier may be declared twice in the same block, and
no identifier may be declared in both the file and package block.
</p>

<p>
The <a href="#Blank_identifier">blank identifier</a> may be used like any other identifier
in a declaration, but it does not introduce a binding and thus is not declared.
In the package block, the identifier <code>init</code> may only be used for
<a href="#Package_initialization"><code>init</code> function</a> declarations,
and like the blank identifier it does not introduce a new binding.
</p>

<pre class="ebnf">
<a id="Declaration">Declaration</a>   = <a href="#ConstDecl" class="noline">ConstDecl</a> | <a href="#TypeDecl" class="noline">TypeDecl</a> | <a href="#VarDecl" class="noline">VarDecl</a> .
<a id="TopLevelDecl">TopLevelDecl</a>  = <a href="#Declaration" class="noline">Declaration</a> | <a href="#FunctionDecl" class="noline">FunctionDecl</a> | <a href="#MethodDecl" class="noline">MethodDecl</a> .
</pre>

<p>
The <i>scope</i> of a declared identifier is the extent of source text in which
the identifier denotes the specified constant, type, variable, function, label, or package.
</p>

<p>
Go is lexically scoped using <a href="#Blocks">blocks</a>:
</p>

<ol>
	<li>The scope of a <a href="#Predeclared_identifiers">predeclared identifier</a> is the universe block.</li>

	<li>The scope of an identifier denoting a constant, type, variable,
	    or function (but not method) declared at top level (outside any
	    function) is the package block.</li>

	<li>The scope of the package name of an imported package is the file block
	    of the file containing the import declaration.</li>

	<li>The scope of an identifier denoting a method receiver, function parameter,
	    or result variable is the function body.</li>

	<li>The scope of a constant or variable identifier declared
	    inside a function begins at the end of the ConstSpec or VarSpec
	    (ShortVarDecl for short variable declarations)
	    and ends at the end of the innermost containing block.</li>

	<li>The scope of a type identifier declared inside a function
	    begins at the identifier in the TypeSpec
	    and ends at the end of the innermost containing block.</li>
</ol>

<p>
An identifier declared in a block may be redeclared in an inner block.
While the identifier of the inner declaration is in scope, it denotes
the entity declared by the inner declaration.
</p>

<p>
The <a href="#Package_clause">package clause</a> is not a declaration; the package name
does not appear in any scope. Its purpose is to identify the files belonging
to the same <a href="#Packages">package</a> and to specify the default package name for import
declarations.
</p>


<h3 id="Label_scopes">Label scopes</h3>

<p>
Labels are declared by <a href="#Labeled_statements">labeled statements</a> and are
used in the <a href="#Break_statements">"break"</a>,
<a href="#Continue_statements">"continue"</a>, and
<a href="#Goto_statements">"goto"</a> statements.
It is illegal to define a label that is never used.
In contrast to other identifiers, labels are not block scoped and do
not conflict with identifiers that are not labels. The scope of a label
is the body of the function in which it is declared and excludes
the body of any nested function.
</p>


<h3 id="Blank_identifier">Blank identifier</h3>

<p>
The <i>blank identifier</i> is represented by the underscore character <code>_</code>.
It serves as an anonymous placeholder instead of a regular (non-blank)
identifier and has special meaning in <a href="#Declarations_and_scope">declarations</a>,
as an <a href="#Operands">operand</a>, and in <a href="#Assignments">assignments</a>.
</p>


<h3 id="Predeclared_identifiers">Predeclared identifiers</h3>

<p>
The following identifiers are implicitly declared in the
<a href="#Blocks">universe block</a>:
</p>
<pre class="grammar">
Types:
	bool byte complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap close complex copy delete imag len
	make new panic print println real recover
</pre>


<h3 id="Exported_identifiers">Exported identifiers</h3>

<p>
An identifier may be <i>exported</i> to permit access to it from another package.
An identifier is exported if both:
</p>
<ol>
	<li>the first character of the identifier's name is a Unicode upper case
	letter (Unicode class "Lu"); and</li>
	<li>the identifier is declared in the <a href="#Blocks">package block</a>
	or it is a <a href="#Struct_types">field name</a> or
	<a href="#MethodName">method name</a>.</li>
</ol>
<p>
All other identifiers are not exported.
</p>


<h3 id="Uniqueness_of_identifiers">Uniqueness of identifiers</h3>

<p>
Given a set of identifiers, an identifier is called <i>unique</i> if it is
<i>different</i> from every other in the set.
Two identifiers are different if they are spelled differently, or if they
appear in different <a href="#Packages">packages</a> and are not
<a href="#Exported_identifiers">exported</a>. Otherwise, they are the same.
</p>

<h3 id="Constant_declarations">Constant declarations</h3>

<p>
A constant declaration binds a list of identifiers (the names of
the constants) to the values of a list of <a href="#Constant_expressions">constant expressions</a>.
The number of identifiers must be equal
to the number of expressions, and the <i>n</i>th identifier on
the left is bound to the value of the <i>n</i>th expression on the
right.
</p>

<pre class="ebnf">
<a id="ConstDecl">ConstDecl</a>      = "const" ( <a href="#ConstSpec" class="noline">ConstSpec</a> | "(" { <a href="#ConstSpec" class="noline">ConstSpec</a> ";" } ")" ) .
<a id="ConstSpec">ConstSpec</a>      = <a href="#IdentifierList" class="noline">IdentifierList</a> [ [ <a href="#Type" class="noline">Type</a> ] "=" <a href="#ExpressionList" class="noline">ExpressionList</a> ] .

<a id="IdentifierList">IdentifierList</a> = <a href="#identifier" class="noline">identifier</a> { "," <a href="#identifier" class="noline">identifier</a> } .
<a id="ExpressionList">ExpressionList</a> = <a href="#Expression" class="noline">Expression</a> { "," <a href="#Expression" class="noline">Expression</a> } .
</pre>

<p>
If the type is present, all constants take the type specified, and
the expressions must be <a href="#Assignability">assignable</a> to that type.
If the type is omitted, the constants take the
individual types of the corresponding expressions.
If the expression values are untyped <a href="#Constants">constants</a>,
the declared constants remain untyped and the constant identifiers
denote the constant values. For instance, if the expression is a
floating-point literal, the constant identifier denotes a floating-point
constant, even if the literal's fractional part is zero.
</p>

<pre>
const Pi float64 = 3.14159265358979323846
const zero = 0.0         // untyped floating-point constant
const (
	size int64 = 1024
	eof        = -1  // untyped integer constant
)
const a, b, c = 3, 4, "foo"  // a = 3, b = 4, c = "foo", untyped integer and string constants
const u, v float32 = 0, 3    // u = 0.0, v = 3.0
</pre>

<p>
Within a parenthesized <code>const</code> declaration list the
expression list may be omitted from any but the first declaration.
Such an empty list is equivalent to the textual substitution of the
first preceding non-empty expression list and its type if any.
Omitting the list of expressions is therefore equivalent to
repeating the previous list.  The number of identifiers must be equal
to the number of expressions in the previous list.
Together with the <a href="#Iota"><code>iota</code> constant generator</a>
this mechanism permits light-weight declaration of sequential values:
</p>

<pre>
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays  // this constant is not exported
)
</pre>


<h3 id="Iota">Iota</h3>

<p>
Within a <a href="#Constant_declarations">constant declaration</a>, the predeclared identifier
<code>iota</code> represents successive untyped integer <a href="#Constants">
constants</a>. It is reset to 0 whenever the reserved word <code>const</code>
appears in the source and increments after each <a href="#ConstSpec">ConstSpec</a>.
It can be used to construct a set of related constants:
</p>

<pre>
const (  // iota is reset to 0
	c0 = iota  // c0 == 0
	c1 = iota  // c1 == 1
	c2 = iota  // c2 == 2
)

const (
	a = 1 &lt;&lt; iota  // a == 1 (iota has been reset)
	b = 1 &lt;&lt; iota  // b == 2
	c = 1 &lt;&lt; iota  // c == 4
)

const (
	u         = iota * 42  // u == 0     (untyped integer constant)
	v float64 = iota * 42  // v == 42.0  (float64 constant)
	w         = iota * 42  // w == 84    (untyped integer constant)
)

const x = iota  // x == 0 (iota has been reset)
const y = iota  // y == 0 (iota has been reset)
</pre>

<p>
Within an ExpressionList, the value of each <code>iota</code> is the same because
it is only incremented after each ConstSpec:
</p>

<pre>
const (
	bit0, mask0 = 1 &lt;&lt; iota, 1&lt;&lt;iota - 1  // bit0 == 1, mask0 == 0
	bit1, mask1                           // bit1 == 2, mask1 == 1
	_, _                                  // skips iota == 2
	bit3, mask3                           // bit3 == 8, mask3 == 7
)
</pre>

<p>
This last example exploits the implicit repetition of the
last non-empty expression list.
</p>


<h3 id="Type_declarations">Type declarations</h3>

<p>
A type declaration binds an identifier, the <i>type name</i>, to a new type
that has the same <a href="#Types">underlying type</a> as an existing type,
and operations defined for the existing type are also defined for the new type.
The new type is <a href="#Type_identity">different</a> from the existing type.
</p>

<pre class="ebnf">
<a id="TypeDecl">TypeDecl</a>     = "type" ( <a href="#TypeSpec" class="noline">TypeSpec</a> | "(" { <a href="#TypeSpec" class="noline">TypeSpec</a> ";" } ")" ) .
<a id="TypeSpec">TypeSpec</a>     = <a href="#identifier" class="noline">identifier</a> <a href="#Type" class="noline">Type</a> .
</pre>

<pre>
type IntArray [16]int

type (
	Point struct{ x, y float64 }
	Polar Point
)

type TreeNode struct {
	left, right *TreeNode
	value *Comparable
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}
</pre>

<p>
The declared type does not inherit any <a href="#Method_declarations">methods</a>
bound to the existing type, but the <a href="#Method_sets">method set</a>
of an interface type or of elements of a composite type remains unchanged:
</p>

<pre>
// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of the <a href="#Pointer_types">base type</a> of PtrMutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its anonymous field Mutex.
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block
</pre>

<p>
A type declaration may be used to define a different boolean, numeric, or string
type and attach methods to it:
</p>

<pre>
type TimeZone int

const (
	EST TimeZone = -(5 + iota)
	CST
	MST
	PST
)

func (tz TimeZone) String() string {
	return fmt.Sprintf("GMT+%dh", tz)
}
</pre>


<h3 id="Variable_declarations">Variable declarations</h3>

<p>
A variable declaration creates one or more variables, binds corresponding
identifiers to them, and gives each a type and an initial value.
</p>

<pre class="ebnf">
<a id="VarDecl">VarDecl</a>     = "var" ( <a href="#VarSpec" class="noline">VarSpec</a> | "(" { <a href="#VarSpec" class="noline">VarSpec</a> ";" } ")" ) .
<a id="VarSpec">VarSpec</a>     = <a href="#IdentifierList" class="noline">IdentifierList</a> ( <a href="#Type" class="noline">Type</a> [ "=" <a href="#ExpressionList" class="noline">ExpressionList</a> ] | "=" <a href="#ExpressionList" class="noline">ExpressionList</a> ) .
</pre>

<pre>
var i int
var U, V, W float64
var k = 0
var x, y float32 = -1, -2
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
var re, im = complexSqrt(-1)
var _, found = entries[name]  // map lookup; only interested in "found"
</pre>

<p>
If a list of expressions is given, the variables are initialized
with the expressions following the rules for <a href="#Assignments">assignments</a>.
Otherwise, each variable is initialized to its <a href="#The_zero_value">zero value</a>.
</p>

<p>
If a type is present, each variable is given that type.
Otherwise, each variable is given the type of the corresponding
initialization value in the assignment.
If that value is an untyped constant, it is first
<a href="#Conversions">converted</a> to its <a href="#Constants">default type</a>;
if it is an untyped boolean value, it is first converted to type <code>bool</code>.
The predeclared value <code>nil</code> cannot be used to initialize a variable
with no explicit type.
</p>

<pre>
var d = math.Sin(0.5)  // d is int64
var i = 42             // i is int
var t, ok = x.(T)      // t is T, ok is bool
var n = nil            // illegal
</pre>

<p>
Implementation restriction: A compiler may make it illegal to declare a variable
inside a <a href="#Function_declarations">function body</a> if the variable is
never used.
</p>

<h3 id="Short_variable_declarations">Short variable declarations</h3>

<p>
A <i>short variable declaration</i> uses the syntax:
</p>

<pre class="ebnf">
<a id="ShortVarDecl">ShortVarDecl</a> = <a href="#IdentifierList" class="noline">IdentifierList</a> ":=" <a href="#ExpressionList" class="noline">ExpressionList</a> .
</pre>

<p>
It is shorthand for a regular <a href="#Variable_declarations">variable declaration</a>
with initializer expressions but no types:
</p>

<pre class="grammar">
"var" IdentifierList = ExpressionList .
</pre>

<pre>
i, j := 0, 10
f := func() int { return 7 }
ch := make(chan int)
r, w := os.Pipe(fd)  // os.Pipe() returns two values
_, y, _ := coord(p)  // coord() returns three values; only interested in y coordinate
</pre>

<p>
Unlike regular variable declarations, a short variable declaration may redeclare variables provided they
were originally declared earlier in the same block with the same type, and at
least one of the non-<a href="#Blank_identifier">blank</a> variables is new.  As a consequence, redeclaration
can only appear in a multi-variable short declaration.
Redeclaration does not introduce a new
variable; it just assigns a new value to the original.
</p>

<pre>
field1, offset := nextField(str, 0)
field2, offset := nextField(str, offset)  // redeclares offset
a, a := 1, 2                              // illegal: double declaration of a or no new variable if a was declared elsewhere
</pre>

<p>
Short variable declarations may appear only inside functions.
In some contexts such as the initializers for
<a href="#If_statements">"if"</a>,
<a href="#For_statements">"for"</a>, or
<a href="#Switch_statements">"switch"</a> statements,
they can be used to declare local temporary variables.
</p>

<h3 id="Function_declarations">Function declarations</h3>

<p>
A function declaration binds an identifier, the <i>function name</i>,
to a function.
</p>

<pre class="ebnf">
<a id="FunctionDecl">FunctionDecl</a> = "func" <a href="#FunctionName" class="noline">FunctionName</a> ( <a href="#Function" class="noline">Function</a> | <a href="#Signature" class="noline">Signature</a> ) .
<a id="FunctionName">FunctionName</a> = <a href="#identifier" class="noline">identifier</a> .
<a id="Function">Function</a>     = <a href="#Signature" class="noline">Signature</a> <a href="#FunctionBody" class="noline">FunctionBody</a> .
<a id="FunctionBody">FunctionBody</a> = <a href="#Block" class="noline">Block</a> .
</pre>

<p>
If the function's <a href="#Function_types">signature</a> declares
result parameters, the function body's statement list must end in
a <a href="#Terminating_statements">terminating statement</a>.
</p>

<pre>
func findMarker(c &lt;-chan int) int {
	for i := range c {
		if x := &lt;-c; isMarker(x) {
			return x
		}
	}
	// invalid: missing return statement.
}
</pre>

<p>
A function declaration may omit the body. Such a declaration provides the
signature for a function implemented outside Go, such as an assembly routine.
</p>

<pre>
func min(x int, y int) int {
	if x &lt; y {
		return x
	}
	return y
}

func flushICache(begin, end uintptr)  // implemented externally
</pre>

<h3 id="Method_declarations">Method declarations</h3>

<p>
A method is a <a href="#Function_declarations">function</a> with a <i>receiver</i>.
A method declaration binds an identifier, the <i>method name</i>, to a method,
and associates the method with the receiver's <i>base type</i>.
</p>

<pre class="ebnf">
<a id="MethodDecl">MethodDecl</a>   = "func" <a href="#Receiver" class="noline">Receiver</a> <a href="#MethodName" class="noline">MethodName</a> ( <a href="#Function" class="noline">Function</a> | <a href="#Signature" class="noline">Signature</a> ) .
<a id="Receiver">Receiver</a>     = <a href="#Parameters" class="noline">Parameters</a> .
</pre>

<p>
The receiver is specified via an extra parameter section preceeding the method
name. That parameter section must declare a single parameter, the receiver.
Its type must be of the form <code>T</code> or <code>*T</code> (possibly using
parentheses) where <code>T</code> is a type name. The type denoted by <code>T</code> is called
the receiver <i>base type</i>; it must not be a pointer or interface type and
it must be declared in the same package as the method.
The method is said to be <i>bound</i> to the base type and the method name
is visible only within selectors for that type.
</p>

<p>
A non-<a href="#Blank_identifier">blank</a> receiver identifier must be
<a href="#Uniqueness_of_identifiers">unique</a> in the method signature.
If the receiver's value is not referenced inside the body of the method,
its identifier may be omitted in the declaration. The same applies in
general to parameters of functions and methods.
</p>

<p>
For a base type, the non-blank names of methods bound to it must be unique.
If the base type is a <a href="#Struct_types">struct type</a>,
the non-blank method and field names must be distinct.
</p>

<p>
Given type <code>Point</code>, the declarations
</p>

<pre>
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
</pre>

<p>
bind the methods <code>Length</code> and <code>Scale</code>,
with receiver type <code>*Point</code>,
to the base type <code>Point</code>.
</p>

<p>
The type of a method is the type of a function with the receiver as first
argument.  For instance, the method <code>Scale</code> has type
</p>

<pre>
func(p *Point, factor float64)
</pre>

<p>
However, a function declared this way is not a method.
</p>


<h2 id="Expressions">Expressions</h2>

<p>
An expression specifies the computation of a value by applying
operators and functions to operands.
</p>

<h3 id="Operands">Operands</h3>

<p>
Operands denote the elementary values in an expression. An operand may be a
literal, a (possibly <a href="#Qualified_identifiers">qualified</a>)
non-<a href="#Blank_identifier">blank</a> identifier denoting a
<a href="#Constant_declarations">constant</a>,
<a href="#Variable_declarations">variable</a>, or
<a href="#Function_declarations">function</a>,
a <a href="#Method_expressions">method expression</a> yielding a function,
or a parenthesized expression.
</p>

<p>
The <a href="#Blank_identifier">blank identifier</a> may appear as an
operand only on the left-hand side of an <a href="#Assignments">assignment</a>.
</p>

<pre class="ebnf">
<a id="Operand">Operand</a>     = <a href="#Literal" class="noline">Literal</a> | <a href="#OperandName" class="noline">OperandName</a> | <a href="#MethodExpr" class="noline">MethodExpr</a> | "(" <a href="#Expression" class="noline">Expression</a> ")" .
<a id="Literal">Literal</a>     = <a href="#BasicLit" class="noline">BasicLit</a> | <a href="#CompositeLit" class="noline">CompositeLit</a> | <a href="#FunctionLit" class="noline">FunctionLit</a> .
<a id="BasicLit">BasicLit</a>    = <a href="#int_lit" class="noline">int_lit</a> | <a href="#float_lit" class="noline">float_lit</a> | <a href="#imaginary_lit" class="noline">imaginary_lit</a> | <a href="#rune_lit" class="noline">rune_lit</a> | <a href="#string_lit" class="noline">string_lit</a> .
<a id="OperandName">OperandName</a> = <a href="#identifier" class="noline">identifier</a> | <a href="#QualifiedIdent" class="noline">QualifiedIdent</a>.
</pre>

<h3 id="Qualified_identifiers">Qualified identifiers</h3>

<p>
A qualified identifier is an identifier qualified with a package name prefix.
Both the package name and the identifier must not be
<a href="#Blank_identifier">blank</a>.
</p>

<pre class="ebnf">
<a id="QualifiedIdent">QualifiedIdent</a> = <a href="#PackageName" class="noline">PackageName</a> "." <a href="#identifier" class="noline">identifier</a> .
</pre>

<p>
A qualified identifier accesses an identifier in a different package, which
must be <a href="#Import_declarations">imported</a>.
The identifier must be <a href="#Exported_identifiers">exported</a> and
declared in the <a href="#Blocks">package block</a> of that package.
</p>

<pre>
math.Sin	// denotes the Sin function in package math
</pre>

<h3 id="Composite_literals">Composite literals</h3>

<p>
Composite literals construct values for structs, arrays, slices, and maps
and create a new value each time they are evaluated.
They consist of the type of the value
followed by a brace-bound list of composite elements. An element may be
a single expression or a key-value pair.
</p>

<pre class="ebnf">
<a id="CompositeLit">CompositeLit</a>  = <a href="#LiteralType" class="noline">LiteralType</a> <a href="#LiteralValue" class="noline">LiteralValue</a> .
<a id="LiteralType">LiteralType</a>   = <a href="#StructType" class="noline">StructType</a> | <a href="#ArrayType" class="noline">ArrayType</a> | "[" "..." "]" <a href="#ElementType" class="noline">ElementType</a> |
                <a href="#SliceType" class="noline">SliceType</a> | <a href="#MapType" class="noline">MapType</a> | <a href="#TypeName" class="noline">TypeName</a> .
<a id="LiteralValue">LiteralValue</a>  = "{" [ <a href="#ElementList" class="noline">ElementList</a> [ "," ] ] "}" .
<a id="ElementList">ElementList</a>   = <a href="#Element" class="noline">Element</a> { "," <a href="#Element" class="noline">Element</a> } .
<a id="Element">Element</a>       = [ <a href="#Key" class="noline">Key</a> ":" ] <a href="#Value" class="noline">Value</a> .
<a id="Key">Key</a>           = <a href="#FieldName" class="noline">FieldName</a> | <a href="#ElementIndex" class="noline">ElementIndex</a> .
<a id="FieldName">FieldName</a>     = <a href="#identifier" class="noline">identifier</a> .
<a id="ElementIndex">ElementIndex</a>  = <a href="#Expression" class="noline">Expression</a> .
<a id="Value">Value</a>         = <a href="#Expression" class="noline">Expression</a> | <a href="#LiteralValue" class="noline">LiteralValue</a> .
</pre>

<p>
The LiteralType must be a struct, array, slice, or map type
(the grammar enforces this constraint except when the type is given
as a TypeName).
The types of the expressions must be <a href="#Assignability">assignable</a>
to the respective field, element, and key types of the LiteralType;
there is no additional conversion.
The key is interpreted as a field name for struct literals,
an index for array and slice literals, and a key for map literals.
For map literals, all elements must have a key. It is an error
to specify multiple elements with the same field name or
constant key value.
</p>

<p>
For struct literals the following rules apply:
</p>
<ul>
	<li>A key must be a field name declared in the LiteralType.
	</li>
	<li>An element list that does not contain any keys must
	    list an element for each struct field in the
	    order in which the fields are declared.
	</li>
	<li>If any element has a key, every element must have a key.
	</li>
	<li>An element list that contains keys does not need to
	    have an element for each struct field. Omitted fields
	    get the zero value for that field.
	</li>
	<li>A literal may omit the element list; such a literal evaluates
	    to the zero value for its type.
	</li>
	<li>It is an error to specify an element for a non-exported
	    field of a struct belonging to a different package.
	</li>
</ul>

<p>
Given the declarations
</p>
<pre>
type Point3D struct { x, y, z float64 }
type Line struct { p, q Point3D }
</pre>

<p>
one may write
</p>

<pre>
origin := Point3D{}                            // zero value for Point3D
line := Line{origin, Point3D{y: -4, z: 12.3}}  // zero value for line.q.x
</pre>

<p>
For array and slice literals the following rules apply:
</p>
<ul>
	<li>Each element has an associated integer index marking
	    its position in the array.
	</li>
	<li>An element with a key uses the key as its index; the
	    key must be a constant integer expression.
	</li>
	<li>An element without a key uses the previous element's index plus one.
	    If the first element has no key, its index is zero.
	</li>
</ul>

<p>
<a href="#Address_operators">Taking the address</a> of a composite literal
generates a pointer to a unique <a href="#Variables">variable</a> initialized
with the literal's value.
</p>
<pre>
var pointer *Point3D = &amp;Point3D{y: 1000}
</pre>

<p>
The length of an array literal is the length specified in the LiteralType.
If fewer elements than the length are provided in the literal, the missing
elements are set to the zero value for the array element type.
It is an error to provide elements with index values outside the index range
of the array. The notation <code>...</code> specifies an array length equal
to the maximum element index plus one.
</p>

<pre>
buffer := [10]string{}             // len(buffer) == 10
intSet := [6]int{1, 2, 3, 5}       // len(intSet) == 6
days := [...]string{"Sat", "Sun"}  // len(days) == 2
</pre>

<p>
A slice literal describes the entire underlying array literal.
Thus, the length and capacity of a slice literal are the maximum
element index plus one. A slice literal has the form
</p>

<pre>
[]T{x1, x2, … xn}
</pre>

<p>
and is shorthand for a slice operation applied to an array:
</p>

<pre>
tmp := [n]T{x1, x2, … xn}
tmp[0 : n]
</pre>

<p>
Within a composite literal of array, slice, or map type <code>T</code>,
elements that are themselves composite literals may elide the respective
literal type if it is identical to the element type of <code>T</code>.
Similarly, elements that are addresses of composite literals may elide
the <code>&amp;T</code> when the element type is <code>*T</code>.
</p>

<pre>
[...]Point{{1.5, -3.5}, {0, 0}}   // same as [...]Point{Point{1.5, -3.5}, Point{0, 0}}
[][]int{{1, 2, 3}, {4, 5}}        // same as [][]int{[]int{1, 2, 3}, []int{4, 5}}

[...]*Point{{1.5, -3.5}, {0, 0}}  // same as [...]*Point{&amp;Point{1.5, -3.5}, &amp;Point{0, 0}}
</pre>

<p>
A parsing ambiguity arises when a composite literal using the
TypeName form of the LiteralType appears as an operand between the
<a href="#Keywords">keyword</a> and the opening brace of the block
of an "if", "for", or "switch" statement, and the composite literal
is not enclosed in parentheses, square brackets, or curly braces.
In this rare case, the opening brace of the literal is erroneously parsed
as the one introducing the block of statements. To resolve the ambiguity,
the composite literal must appear within parentheses.
</p>

<pre>
if x == (T{a,b,c}[i]) { … }
if (x == T{a,b,c}[i]) { … }
</pre>

<p>
Examples of valid array, slice, and map literals:
</p>

<pre>
// list of prime numbers
primes := []int{2, 3, 5, 7, 9, 2147483647}

// vowels[ch] is true if ch is a vowel
vowels := [128]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// the array [10]float32{-1, 0, 0, 0, -0.1, -0.1, 0, 0, 0, -1}
filter := [10]float32{-1, 4: -0.1, -0.1, 9: -1}

// frequencies in Hz for equal-tempered scale (A4 = 440Hz)
noteFrequency := map[string]float32{
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87,
}
</pre>


<h3 id="Function_literals">Function literals</h3>

<p>
A function literal represents an anonymous <a href="#Function_declarations">function</a>.
</p>

<pre class="ebnf">
<a id="FunctionLit">FunctionLit</a> = "func" <a href="#Function" class="noline">Function</a> .
</pre>

<pre>
func(a, b int, z float64) bool { return a*b &lt; int(z) }
</pre>

<p>
A function literal can be assigned to a variable or invoked directly.
</p>

<pre>
f := func(x, y int) int { return x + y }
func(ch chan int) { ch &lt;- ACK }(replyChan)
</pre>

<p>
Function literals are <i>closures</i>: they may refer to variables
defined in a surrounding function. Those variables are then shared between
the surrounding function and the function literal, and they survive as long
as they are accessible.
</p>


<h3 id="Primary_expressions">Primary expressions</h3>

<p>
Primary expressions are the operands for unary and binary expressions.
</p>

<pre class="ebnf">
<a id="PrimaryExpr">PrimaryExpr</a> =
	<a href="#Operand" class="noline">Operand</a> |
	<a href="#Conversion" class="noline">Conversion</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Selector" class="noline">Selector</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Index" class="noline">Index</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Slice" class="noline">Slice</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#TypeAssertion" class="noline">TypeAssertion</a> |
	<a href="#PrimaryExpr" class="noline">PrimaryExpr</a> <a href="#Arguments" class="noline">Arguments</a> .

<a id="Selector">Selector</a>       = "." <a href="#identifier" class="noline">identifier</a> .
<a id="Index">Index</a>          = "[" <a href="#Expression" class="noline">Expression</a> "]" .
<a id="Slice">Slice</a>          = "[" ( [ <a href="#Expression" class="noline">Expression</a> ] ":" [ <a href="#Expression" class="noline">Expression</a> ] ) |
                     ( [ <a href="#Expression" class="noline">Expression</a> ] ":" <a href="#Expression" class="noline">Expression</a> ":" <a href="#Expression" class="noline">Expression</a> )
                 "]" .
<a id="TypeAssertion">TypeAssertion</a>  = "." "(" <a href="#Type" class="noline">Type</a> ")" .
<a id="Arguments">Arguments</a>      = "(" [ ( <a href="#ExpressionList" class="noline">ExpressionList</a> | <a href="#Type" class="noline">Type</a> [ "," <a href="#ExpressionList" class="noline">ExpressionList</a> ] ) [ "..." ] [ "," ] ] ")" .
</pre>


<pre>
x
2
(s + ".txt")
f(3.1415, true)
Point{1, 2}
m["foo"]
s[i : j + 1]
obj.color
f.p[i].x()
</pre>


<h3 id="Selectors">Selectors</h3>

<p>
For a <a href="#Primary_expressions">primary expression</a> <code>x</code>
that is not a <a href="#Package_clause">package name</a>, the
<i>selector expression</i>
</p>

<pre>
x.f
</pre>

<p>
denotes the field or method <code>f</code> of the value <code>x</code>
(or sometimes <code>*x</code>; see below).
The identifier <code>f</code> is called the (field or method) <i>selector</i>;
it must not be the <a href="#Blank_identifier">blank identifier</a>.
The type of the selector expression is the type of <code>f</code>.
If <code>x</code> is a package name, see the section on
<a href="#Qualified_identifiers">qualified identifiers</a>.
</p>

<p>
A selector <code>f</code> may denote a field or method <code>f</code> of
a type <code>T</code>, or it may refer
to a field or method <code>f</code> of a nested
<a href="#Struct_types">anonymous field</a> of <code>T</code>.
The number of anonymous fields traversed
to reach <code>f</code> is called its <i>depth</i> in <code>T</code>.
The depth of a field or method <code>f</code>
declared in <code>T</code> is zero.
The depth of a field or method <code>f</code> declared in
an anonymous field <code>A</code> in <code>T</code> is the
depth of <code>f</code> in <code>A</code> plus one.
</p>

<p>
The following rules apply to selectors:
</p>

<ol>
<li>
For a value <code>x</code> of type <code>T</code> or <code>*T</code>
where <code>T</code> is not a pointer or interface type,
<code>x.f</code> denotes the field or method at the shallowest depth
in <code>T</code> where there
is such an <code>f</code>.
If there is not exactly <a href="#Uniqueness_of_identifiers">one <code>f</code></a>
with shallowest depth, the selector expression is illegal.
</li>

<li>
For a value <code>x</code> of type <code>I</code> where <code>I</code>
is an interface type, <code>x.f</code> denotes the actual method with name
<code>f</code> of the dynamic value of <code>x</code>.
If there is no method with name <code>f</code> in the
<a href="#Method_sets">method set</a> of <code>I</code>, the selector
expression is illegal.
</li>

<li>
As an exception, if the type of <code>x</code> is a named pointer type
and <code>(*x).f</code> is a valid selector expression denoting a field
(but not a method), <code>x.f</code> is shorthand for <code>(*x).f</code>.
</li>

<li>
In all other cases, <code>x.f</code> is illegal.
</li>

<li>
If <code>x</code> is of pointer type and has the value
<code>nil</code> and <code>x.f</code> denotes a struct field,
assigning to or evaluating <code>x.f</code>
causes a <a href="#Run_time_panics">run-time panic</a>.
</li>

<li>
If <code>x</code> is of interface type and has the value
<code>nil</code>, <a href="#Calls">calling</a> or
<a href="#Method_values">evaluating</a> the method <code>x.f</code>
causes a <a href="#Run_time_panics">run-time panic</a>.
</li>
</ol>

<p>
For example, given the declarations:
</p>

<pre>
type T0 struct {
	x int
}

func (*T0) M0()

type T1 struct {
	y int
}

func (T1) M1()

type T2 struct {
	z int
	T1
	*T0
}

func (*T2) M2()

type Q *T2

var t T2     // with t.T0 != nil
var p *T2    // with p != nil and (*p).T0 != nil
var q Q = p
</pre>

<p>
one may write:
</p>

<pre>
t.z          // t.z
t.y          // t.T1.y
t.x          // (*t.TO).x

p.z          // (*p).z
p.y          // (*p).T1.y
p.x          // (*(*p).T0).x

q.x          // (*(*q).T0).x        (*q).x is a valid field selector

p.M2()       // p.M2()              M2 expects *T2 receiver
p.M1()       // ((*p).T1).M1()      M1 expects T1 receiver
p.M0()       // ((&(*p).T0)).M0()   M0 expects *T0 receiver, see section on Calls
</pre>

<p>
but the following is invalid:
</p>

<pre>
q.M0()       // (*q).M0 is valid but not a field selector
</pre>


<h3 id="Method_expressions">Method expressions</h3>

<p>
If <code>M</code> is in the <a href="#Method_sets">method set</a> of type <code>T</code>,
<code>T.M</code> is a function that is callable as a regular function
with the same arguments as <code>M</code> prefixed by an additional
argument that is the receiver of the method.
</p>

<pre class="ebnf">
<a id="MethodExpr">MethodExpr</a>    = <a href="#ReceiverType" class="noline">ReceiverType</a> "." <a href="#MethodName" class="noline">MethodName</a> .
<a id="ReceiverType">ReceiverType</a>  = <a href="#TypeName" class="noline">TypeName</a> | "(" "*" <a href="#TypeName" class="noline">TypeName</a> ")" | "(" <a href="#ReceiverType" class="noline">ReceiverType</a> ")" .
</pre>

<p>
Consider a struct type <code>T</code> with two methods,
<code>Mv</code>, whose receiver is of type <code>T</code>, and
<code>Mp</code>, whose receiver is of type <code>*T</code>.
</p>

<pre>
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
</pre>

<p>
The expression
</p>

<pre>
T.Mv
</pre>

<p>
yields a function equivalent to <code>Mv</code> but
with an explicit receiver as its first argument; it has signature
</p>

<pre>
func(tv T, a int) int
</pre>

<p>
That function may be called normally with an explicit receiver, so
these five invocations are equivalent:
</p>

<pre>
t.Mv(7)
T.Mv(t, 7)
(T).Mv(t, 7)
f1 := T.Mv; f1(t, 7)
f2 := (T).Mv; f2(t, 7)
</pre>

<p>
Similarly, the expression
</p>

<pre>
(*T).Mp
</pre>

<p>
yields a function value representing <code>Mp</code> with signature
</p>

<pre>
func(tp *T, f float32) float32
</pre>

<p>
For a method with a value receiver, one can derive a function
with an explicit pointer receiver, so
</p>

<pre>
(*T).Mv
</pre>

<p>
yields a function value representing <code>Mv</code> with signature
</p>

<pre>
func(tv *T, a int) int
</pre>

<p>
Such a function indirects through the receiver to create a value
to pass as the receiver to the underlying method;
the method does not overwrite the value whose address is passed in
the function call.
</p>

<p>
The final case, a value-receiver function for a pointer-receiver method,
is illegal because pointer-receiver methods are not in the method set
of the value type.
</p>

<p>
Function values derived from methods are called with function call syntax;
the receiver is provided as the first argument to the call.
That is, given <code>f := T.Mv</code>, <code>f</code> is invoked
as <code>f(t, 7)</code> not <code>t.f(7)</code>.
To construct a function that binds the receiver, use a
<a href="#Function_literals">function literal</a> or
<a href="#Method_values">method value</a>.
</p>

<p>
It is legal to derive a function value from a method of an interface type.
The resulting function takes an explicit receiver of that interface type.
</p>

<h3 id="Method_values">Method values</h3>

<p>
If the expression <code>x</code> has static type <code>T</code> and
<code>M</code> is in the <a href="#Method_sets">method set</a> of type <code>T</code>,
<code>x.M</code> is called a <i>method value</i>.
The method value <code>x.M</code> is a function value that is callable
with the same arguments as a method call of <code>x.M</code>.
The expression <code>x</code> is evaluated and saved during the evaluation of the
method value; the saved copy is then used as the receiver in any calls,
which may be executed later.
</p>

<p>
The type <code>T</code> may be an interface or non-interface type.
</p>

<p>
As in the discussion of <a href="#Method_expressions">method expressions</a> above,
consider a struct type <code>T</code> with two methods,
<code>Mv</code>, whose receiver is of type <code>T</code>, and
<code>Mp</code>, whose receiver is of type <code>*T</code>.
</p>

<pre>
type T struct {
	a int
}
func (tv  T) Mv(a int) int         { return 0 }  // value receiver
func (tp *T) Mp(f float32) float32 { return 1 }  // pointer receiver

var t T
var pt *T
func makeT() T
</pre>

<p>
The expression
</p>

<pre>
t.Mv
</pre>

<p>
yields a function value of type
</p>

<pre>
func(int) int
</pre>

<p>
These two invocations are equivalent:
</p>

<pre>
t.Mv(7)
f := t.Mv; f(7)
</pre>

<p>
Similarly, the expression
</p>

<pre>
pt.Mp
</pre>

<p>
yields a function value of type
</p>

<pre>
func(float32) float32
</pre>

<p>
As with <a href="#Selectors">selectors</a>, a reference to a non-interface method with a value receiver
using a pointer will automatically dereference that pointer: <code>pt.Mv</code> is equivalent to <code>(*pt).Mv</code>.
</p>

<p>
As with <a href="#Calls">method calls</a>, a reference to a non-interface method with a pointer receiver
using an addressable value will automatically take the address of that value: <code>t.Mp</code> is equivalent to <code>(&amp;t).Mp</code>.
</p>

<pre>
f := t.Mv; f(7)   // like t.Mv(7)
f := pt.Mp; f(7)  // like pt.Mp(7)
f := pt.Mv; f(7)  // like (*pt).Mv(7)
f := t.Mp; f(7)   // like (&amp;t).Mp(7)
f := makeT().Mp   // invalid: result of makeT() is not addressable
</pre>

<p>
Although the examples above use non-interface types, it is also legal to create a method value
from a value of interface type.
</p>

<pre>
var i interface { M(int) } = myVal
f := i.M; f(7)  // like i.M(7)
</pre>


<h3 id="Index_expressions">Index expressions</h3>

<p>
A primary expression of the form
</p>

<pre>
a[x]
</pre>

<p>
denotes the element of the array, pointer to array, slice, string or map <code>a</code> indexed by <code>x</code>.
The value <code>x</code> is called the <i>index</i> or <i>map key</i>, respectively.
The following rules apply:
</p>

<p>
If <code>a</code> is not a map:
</p>
<ul>
	<li>the index <code>x</code> must be of integer type or untyped;
	    it is <i>in range</i> if <code>0 &lt;= x &lt; len(a)</code>,
	    otherwise it is <i>out of range</i></li>
	<li>a <a href="#Constants">constant</a> index must be non-negative
	    and representable by a value of type <code>int</code>
</ul>

<p>
For <code>a</code> of <a href="#Array_types">array type</a> <code>A</code>:
</p>
<ul>
	<li>a <a href="#Constants">constant</a> index must be in range</li>
	<li>if <code>x</code> is out of range at run time,
	    a <a href="#Run_time_panics">run-time panic</a> occurs</li>
	<li><code>a[x]</code> is the array element at index <code>x</code> and the type of
	    <code>a[x]</code> is the element type of <code>A</code></li>
</ul>

<p>
For <code>a</code> of <a href="#Pointer_types">pointer</a> to array type:
</p>
<ul>
	<li><code>a[x]</code> is shorthand for <code>(*a)[x]</code></li>
</ul>

<p>
For <code>a</code> of <a href="#Slice_types">slice type</a> <code>S</code>:
</p>
<ul>
	<li>if <code>x</code> is out of range at run time,
	    a <a href="#Run_time_panics">run-time panic</a> occurs</li>
	<li><code>a[x]</code> is the slice element at index <code>x</code> and the type of
	    <code>a[x]</code> is the element type of <code>S</code></li>
</ul>

<p>
For <code>a</code> of <a href="#String_types">string type</a>:
</p>
<ul>
	<li>a <a href="#Constants">constant</a> index must be in range
	    if the string <code>a</code> is also constant</li>
	<li>if <code>x</code> is out of range at run time,
	    a <a href="#Run_time_panics">run-time panic</a> occurs</li>
	<li><code>a[x]</code> is the non-constant byte value at index <code>x</code> and the type of
	    <code>a[x]</code> is <code>byte</code></li>
	<li><code>a[x]</code> may not be assigned to</li>
</ul>

<p>
For <code>a</code> of <a href="#Map_types">map type</a> <code>M</code>:
</p>
<ul>
	<li><code>x</code>'s type must be
	    <a href="#Assignability">assignable</a>
	    to the key type of <code>M</code></li>
	<li>if the map contains an entry with key <code>x</code>,
	    <code>a[x]</code> is the map value with key <code>x</code>
	    and the type of <code>a[x]</code> is the value type of <code>M</code></li>
	<li>if the map is <code>nil</code> or does not contain such an entry,
	    <code>a[x]</code> is the <a href="#The_zero_value">zero value</a>
	    for the value type of <code>M</code></li>
</ul>

<p>
Otherwise <code>a[x]</code> is illegal.
</p>

<p>
An index expression on a map <code>a</code> of type <code>map[K]V</code>
used in an <a href="#Assignments">assignment</a> or initialization of the special form
</p>

<pre>
v, ok = a[x]
v, ok := a[x]
var v, ok = a[x]
</pre>

<p>
yields an additional untyped boolean value. The value of <code>ok</code> is
<code>true</code> if the key <code>x</code> is present in the map, and
<code>false</code> otherwise.
</p>

<p>
Assigning to an element of a <code>nil</code> map causes a
<a href="#Run_time_panics">run-time panic</a>.
</p>


<h3 id="Slice_expressions">Slice expressions</h3>

<p>
Slice expressions construct a substring or slice from a string, array, pointer
to array, or slice. There are two variants: a simple form that specifies a low
and high bound, and a full form that also specifies a bound on the capacity.
</p>

<h4>Simple slice expressions</h4>

<p>
For a string, array, pointer to array, or slice <code>a</code>, the primary expression
</p>

<pre>
a[low : high]
</pre>

<p>
constructs a substring or slice. The <i>indices</i> <code>low</code> and
<code>high</code> select which elements of operand <code>a</code> appear
in the result. The result has indices starting at 0 and length equal to
<code>high</code>&nbsp;-&nbsp;<code>low</code>.
After slicing the array <code>a</code>
</p>

<pre>
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4]
</pre>

<p>
the slice <code>s</code> has type <code>[]int</code>, length 3, capacity 4, and elements
</p>

<pre>
s[0] == 2
s[1] == 3
s[2] == 4
</pre>

<p>
For convenience, any of the indices may be omitted. A missing <code>low</code>
index defaults to zero; a missing <code>high</code> index defaults to the length of the
sliced operand:
</p>

<pre>
a[2:]  // same as a[2 : len(a)]
a[:3]  // same as a[0 : 3]
a[:]   // same as a[0 : len(a)]
</pre>

<p>
If <code>a</code> is a pointer to an array, <code>a[low : high]</code> is shorthand for
<code>(*a)[low : high]</code>.
</p>

<p>
For arrays or strings, the indices are <i>in range</i> if
<code>0</code> &lt;= <code>low</code> &lt;= <code>high</code> &lt;= <code>len(a)</code>,
otherwise they are <i>out of range</i>.
For slices, the upper index bound is the slice capacity <code>cap(a)</code> rather than the length.
A <a href="#Constants">constant</a> index must be non-negative and representable by a value of type
<code>int</code>; for arrays or constant strings, constant indices must also be in range.
If both indices are constant, they must satisfy <code>low &lt;= high</code>.
If the indices are out of range at run time, a <a href="#Run_time_panics">run-time panic</a> occurs.
</p>

<p>
Except for <a href="#Constants">untyped strings</a>, if the sliced operand is a string or slice,
the result of the slice operation is a non-constant value of the same type as the operand.
For untyped string operands the result is a non-constant value of type <code>string</code>.
If the sliced operand is an array, it must be <a href="#Address_operators">addressable</a>
and the result of the slice operation is a slice with the same element type as the array.
</p>

<p>
If the sliced operand of a valid slice expression is a <code>nil</code> slice, the result
is a <code>nil</code> slice. Otherwise, the result shares its underlying array with the
operand.
</p>

<h4>Full slice expressions</h4>

<p>
For an array, pointer to array, or slice <code>a</code> (but not a string), the primary expression
</p>

<pre>
a[low : high : max]
</pre>

<p>
constructs a slice of the same type, and with the same length and elements as the simple slice
expression <code>a[low : high]</code>. Additionally, it controls the resulting slice's capacity
by setting it to <code>max - low</code>. Only the first index may be omitted; it defaults to 0.
After slicing the array <code>a</code>
</p>

<pre>
a := [5]int{1, 2, 3, 4, 5}
t := a[1:3:5]
</pre>

<p>
the slice <code>t</code> has type <code>[]int</code>, length 2, capacity 4, and elements
</p>

<pre>
t[0] == 2
t[1] == 3
</pre>

<p>
As for simple slice expressions, if <code>a</code> is a pointer to an array,
<code>a[low : high : max]</code> is shorthand for <code>(*a)[low : high : max]</code>.
If the sliced operand is an array, it must be <a href="#Address_operators">addressable</a>.
</p>

<p>
The indices are <i>in range</i> if <code>0 &lt;= low &lt;= high &lt;= max &lt;= cap(a)</code>,
otherwise they are <i>out of range</i>.
A <a href="#Constants">constant</a> index must be non-negative and representable by a value of type
<code>int</code>; for arrays, constant indices must also be in range.
If multiple indices are constant, the constants that are present must be in range relative to each
other.
If the indices are out of range at run time, a <a href="#Run_time_panics">run-time panic</a> occurs.
</p>

<h3 id="Type_assertions">Type assertions</h3>

<p>
For an expression <code>x</code> of <a href="#Interface_types">interface type</a>
and a type <code>T</code>, the primary expression
</p>

<pre>
x.(T)
</pre>

<p>
asserts that <code>x</code> is not <code>nil</code>
and that the value stored in <code>x</code> is of type <code>T</code>.
The notation <code>x.(T)</code> is called a <i>type assertion</i>.
</p>
<p>
More precisely, if <code>T</code> is not an interface type, <code>x.(T)</code> asserts
that the dynamic type of <code>x</code> is <a href="#Type_identity">identical</a>
to the type <code>T</code>.
In this case, <code>T</code> must <a href="#Method_sets">implement</a> the (interface) type of <code>x</code>;
otherwise the type assertion is invalid since it is not possible for <code>x</code>
to store a value of type <code>T</code>.
If <code>T</code> is an interface type, <code>x.(T)</code> asserts that the dynamic type
of <code>x</code> implements the interface <code>T</code>.
</p>
<p>
If the type assertion holds, the value of the expression is the value
stored in <code>x</code> and its type is <code>T</code>. If the type assertion is false,
a <a href="#Run_time_panics">run-time panic</a> occurs.
In other words, even though the dynamic type of <code>x</code>
is known only at run time, the type of <code>x.(T)</code> is
known to be <code>T</code> in a correct program.
</p>

<pre>
var x interface{} = 7  // x has dynamic type int and value 7
i := x.(int)           // i has type int and value 7

type I interface { m() }
var y I
s := y.(string)        // illegal: string does not implement I (missing method m)
r := y.(io.Reader)     // r has type io.Reader and y must implement both I and io.Reader
</pre>

<p>
A type assertion used in an <a href="#Assignments">assignment</a> or initialization of the special form
</p>

<pre>
v, ok = x.(T)
v, ok := x.(T)
var v, ok = x.(T)
</pre>

<p>
yields an additional untyped boolean value. The value of <code>ok</code> is <code>true</code>
if the assertion holds. Otherwise it is <code>false</code> and the value of <code>v</code> is
the <a href="#The_zero_value">zero value</a> for type <code>T</code>.
No run-time panic occurs in this case.
</p>


<h3 id="Calls">Calls</h3>

<p>
Given an expression <code>f</code> of function type
<code>F</code>,
</p>

<pre>
f(a1, a2, … an)
</pre>

<p>
calls <code>f</code> with arguments <code>a1, a2, … an</code>.
Except for one special case, arguments must be single-valued expressions
<a href="#Assignability">assignable</a> to the parameter types of
<code>F</code> and are evaluated before the function is called.
The type of the expression is the result type
of <code>F</code>.
A method invocation is similar but the method itself
is specified as a selector upon a value of the receiver type for
the method.
</p>

<pre>
math.Atan2(x, y)  // function call
var pt *Point
pt.Scale(3.5)     // method call with receiver pt
</pre>

<p>
In a function call, the function value and arguments are evaluated in
<a href="#Order_of_evaluation">the usual order</a>.
After they are evaluated, the parameters of the call are passed by value to the function
and the called function begins execution.
The return parameters of the function are passed by value
back to the calling function when the function returns.
</p>

<p>
Calling a <code>nil</code> function value
causes a <a href="#Run_time_panics">run-time panic</a>.
</p>

<p>
As a special case, if the return values of a function or method
<code>g</code> are equal in number and individually
assignable to the parameters of another function or method
<code>f</code>, then the call <code>f(g(<i>parameters_of_g</i>))</code>
will invoke <code>f</code> after binding the return values of
<code>g</code> to the parameters of <code>f</code> in order.  The call
of <code>f</code> must contain no parameters other than the call of <code>g</code>,
and <code>g</code> must have at least one return value.
If <code>f</code> has a final <code>...</code> parameter, it is
assigned the return values of <code>g</code> that remain after
assignment of regular parameters.
</p>

<pre>
func Split(s string, pos int) (string, string) {
	return s[0:pos], s[pos:]
}

func Join(s, t string) string {
	return s + t
}

if Join(Split(value, len(value)/2)) != value {
	log.Panic("test fails")
}
</pre>

<p>
A method call <code>x.m()</code> is valid if the <a href="#Method_sets">method set</a>
of (the type of) <code>x</code> contains <code>m</code> and the
argument list can be assigned to the parameter list of <code>m</code>.
If <code>x</code> is <a href="#Address_operators">addressable</a> and <code>&amp;x</code>'s method
set contains <code>m</code>, <code>x.m()</code> is shorthand
for <code>(&amp;x).m()</code>:
</p>

<pre>
var p Point
p.Scale(3.5)
</pre>

<p>
There is no distinct method type and there are no method literals.
</p>

<h3 id="Passing_arguments_to_..._parameters">Passing arguments to <code>...</code> parameters</h3>

<p>
If <code>f</code> is <a href="#Function_types">variadic</a> with a final
parameter <code>p</code> of type <code>...T</code>, then within <code>f</code>
the type of <code>p</code> is equivalent to type <code>[]T</code>.
If <code>f</code> is invoked with no actual arguments for <code>p</code>,
the value passed to <code>p</code> is <code>nil</code>.
Otherwise, the value passed is a new slice
of type <code>[]T</code> with a new underlying array whose successive elements
are the actual arguments, which all must be <a href="#Assignability">assignable</a>
to <code>T</code>. The length and capacity of the slice is therefore
the number of arguments bound to <code>p</code> and may differ for each
call site.
</p>

<p>
Given the function and calls
</p>
<pre>
func Greeting(prefix string, who ...string)
Greeting("nobody")
Greeting("hello:", "Joe", "Anna", "Eileen")
</pre>

<p>
within <code>Greeting</code>, <code>who</code> will have the value
<code>nil</code> in the first call, and
<code>[]string{"Joe", "Anna", "Eileen"}</code> in the second.
</p>

<p>
If the final argument is assignable to a slice type <code>[]T</code>, it may be
passed unchanged as the value for a <code>...T</code> parameter if the argument
is followed by <code>...</code>. In this case no new slice is created.
</p>

<p>
Given the slice <code>s</code> and call
</p>

<pre>
s := []string{"James", "Jasmine"}
Greeting("goodbye:", s...)
</pre>

<p>
within <code>Greeting</code>, <code>who</code> will have the same value as <code>s</code>
with the same underlying array.
</p>


<h3 id="Operators">Operators</h3>

<p>
Operators combine operands into expressions.
</p>

<pre class="ebnf">
<a id="Expression">Expression</a> = <a href="#UnaryExpr" class="noline">UnaryExpr</a> | <a href="#Expression" class="noline">Expression</a> <a href="#binary_op" class="noline">binary_op</a> <a href="#UnaryExpr" class="noline">UnaryExpr</a> .
<a id="UnaryExpr">UnaryExpr</a>  = <a href="#PrimaryExpr" class="noline">PrimaryExpr</a> | <a href="#unary_op" class="noline">unary_op</a> <a href="#UnaryExpr" class="noline">UnaryExpr</a> .

<a id="binary_op">binary_op</a>  = "||" | "&amp;&amp;" | <a href="#rel_op" class="noline">rel_op</a> | <a href="#add_op" class="noline">add_op</a> | <a href="#mul_op" class="noline">mul_op</a> .
<a id="rel_op">rel_op</a>     = "==" | "!=" | "&lt;" | "&lt;=" | ">" | ">=" .
<a id="add_op">add_op</a>     = "+" | "-" | "|" | "^" .
<a id="mul_op">mul_op</a>     = "*" | "/" | "%" | "&lt;&lt;" | "&gt;&gt;" | "&amp;" | "&amp;^" .

<a id="unary_op">unary_op</a>   = "+" | "-" | "!" | "^" | "*" | "&amp;" | "&lt;-" .
</pre>

<p>
Comparisons are discussed <a href="#Comparison_operators">elsewhere</a>.
For other binary operators, the operand types must be <a href="#Type_identity">identical</a>
unless the operation involves shifts or untyped <a href="#Constants">constants</a>.
For operations involving constants only, see the section on
<a href="#Constant_expressions">constant expressions</a>.
</p>

<p>
Except for shift operations, if one operand is an untyped <a href="#Constants">constant</a>
and the other operand is not, the constant is <a href="#Conversions">converted</a>
to the type of the other operand.
</p>

<p>
The right operand in a shift expression must have unsigned integer type
or be an untyped constant that can be converted to unsigned integer type.
If the left operand of a non-constant shift expression is an untyped constant,
the type of the constant is what it would be if the shift expression were
replaced by its left operand alone.
</p>

<pre>
var s uint = 33
var i = 1&lt;&lt;s           // 1 has type int
var j int32 = 1&lt;&lt;s     // 1 has type int32; j == 0
var k = uint64(1&lt;&lt;s)   // 1 has type uint64; k == 1&lt;&lt;33
var m int = 1.0&lt;&lt;s     // 1.0 has type int
var n = 1.0&lt;&lt;s != i    // 1.0 has type int; n == false if ints are 32bits in size
var o = 1&lt;&lt;s == 2&lt;&lt;s   // 1 and 2 have type int; o == true if ints are 32bits in size
var p = 1&lt;&lt;s == 1&lt;&lt;33  // illegal if ints are 32bits in size: 1 has type int, but 1&lt;&lt;33 overflows int
var u = 1.0&lt;&lt;s         // illegal: 1.0 has type float64, cannot shift
var u1 = 1.0&lt;&lt;s != 0   // illegal: 1.0 has type float64, cannot shift
var u2 = 1&lt;&lt;s != 1.0   // illegal: 1 has type float64, cannot shift
var v float32 = 1&lt;&lt;s   // illegal: 1 has type float32, cannot shift
var w int64 = 1.0&lt;&lt;33  // 1.0&lt;&lt;33 is a constant shift expression
</pre>

<h3 id="Operator_precedence">Operator precedence</h3>
<p>
Unary operators have the highest precedence.
As the  <code>++</code> and <code>--</code> operators form
statements, not expressions, they fall
outside the operator hierarchy.
As a consequence, statement <code>*p++</code> is the same as <code>(*p)++</code>.
<p>
There are five precedence levels for binary operators.
Multiplication operators bind strongest, followed by addition
operators, comparison operators, <code>&amp;&amp;</code> (logical AND),
and finally <code>||</code> (logical OR):
</p>

<pre class="grammar">
Precedence    Operator
    5             *  /  %  &lt;&lt;  &gt;&gt;  &amp;  &amp;^
    4             +  -  |  ^
    3             ==  !=  &lt;  &lt;=  &gt;  &gt;=
    2             &amp;&amp;
    1             ||
</pre>

<p>
Binary operators of the same precedence associate from left to right.
For instance, <code>x / y * z</code> is the same as <code>(x / y) * z</code>.
</p>

<pre>
+x
23 + 3*x[i]
x &lt;= f()
^a &gt;&gt; b
f() || g()
x == y+1 &amp;&amp; &lt;-chanPtr &gt; 0
</pre>


<h3 id="Arithmetic_operators">Arithmetic operators</h3>
<p>
Arithmetic operators apply to numeric values and yield a result of the same
type as the first operand. The four standard arithmetic operators (<code>+</code>,
<code>-</code>,  <code>*</code>, <code>/</code>) apply to integer,
floating-point, and complex types; <code>+</code> also applies
to strings. All other arithmetic operators apply to integers only.
</p>

<pre class="grammar">
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers

&amp;    bitwise AND            integers
|    bitwise OR             integers
^    bitwise XOR            integers
&amp;^   bit clear (AND NOT)    integers

&lt;&lt;   left shift             integer &lt;&lt; unsigned integer
&gt;&gt;   right shift            integer &gt;&gt; unsigned integer
</pre>

<p>
Strings can be concatenated using the <code>+</code> operator
or the <code>+=</code> assignment operator:
</p>

<pre>
s := "hi" + string(c)
s += " and good bye"
</pre>

<p>
String addition creates a new string by concatenating the operands.
</p>
<p>
For two integer values <code>x</code> and <code>y</code>, the integer quotient
<code>q = x / y</code> and remainder <code>r = x % y</code> satisfy the following
relationships:
</p>

<pre>
x = q*y + r  and  |r| &lt; |y|
</pre>

<p>
with <code>x / y</code> truncated towards zero
(<a href="http://en.wikipedia.org/wiki/Modulo_operation">"truncated division"</a>).
</p>

<pre>
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
</pre>

<p>
As an exception to this rule, if the dividend <code>x</code> is the most
negative value for the int type of <code>x</code>, the quotient
<code>q = x / -1</code> is equal to <code>x</code> (and <code>r = 0</code>).
</p>

<pre>
			 x, q
int8                     -128
int16                  -32768
int32             -2147483648
int64    -9223372036854775808
</pre>

<p>
If the divisor is a <a href="#Constants">constant</a>, it must not be zero.
If the divisor is zero at run time, a <a href="#Run_time_panics">run-time panic</a> occurs.
If the dividend is non-negative and the divisor is a constant power of 2,
the division may be replaced by a right shift, and computing the remainder may
be replaced by a bitwise AND operation:
</p>

<pre>
 x     x / 4     x % 4     x &gt;&gt; 2     x &amp; 3
 11      2         3         2          3
-11     -2        -3        -3          1
</pre>

<p>
The shift operators shift the left operand by the shift count specified by the
right operand. They implement arithmetic shifts if the left operand is a signed
integer and logical shifts if it is an unsigned integer.
There is no upper limit on the shift count. Shifts behave
as if the left operand is shifted <code>n</code> times by 1 for a shift
count of <code>n</code>.
As a result, <code>x &lt;&lt; 1</code> is the same as <code>x*2</code>
and <code>x &gt;&gt; 1</code> is the same as
<code>x/2</code> but truncated towards negative infinity.
</p>

<p>
For integer operands, the unary operators
<code>+</code>, <code>-</code>, and <code>^</code> are defined as
follows:
</p>

<pre class="grammar">
+x                          is 0 + x
-x    negation              is 0 - x
^x    bitwise complement    is m ^ x  with m = "all bits set to 1" for unsigned x
                                      and  m = -1 for signed x
</pre>

<p>
For floating-point and complex numbers,
<code>+x</code> is the same as <code>x</code>,
while <code>-x</code> is the negation of <code>x</code>.
The result of a floating-point or complex division by zero is not specified beyond the
IEEE-754 standard; whether a <a href="#Run_time_panics">run-time panic</a>
occurs is implementation-specific.
</p>

<h3 id="Integer_overflow">Integer overflow</h3>

<p>
For unsigned integer values, the operations <code>+</code>,
<code>-</code>, <code>*</code>, and <code>&lt;&lt;</code> are
computed modulo 2<sup><i>n</i></sup>, where <i>n</i> is the bit width of
the <a href="#Numeric_types">unsigned integer</a>'s type.
Loosely speaking, these unsigned integer operations
discard high bits upon overflow, and programs may rely on ''wrap around''.
</p>
<p>
For signed integers, the operations <code>+</code>,
<code>-</code>, <code>*</code>, and <code>&lt;&lt;</code> may legally
overflow and the resulting value exists and is deterministically defined
by the signed integer representation, the operation, and its operands.
No exception is raised as a result of overflow. A
compiler may not optimize code under the assumption that overflow does
not occur. For instance, it may not assume that <code>x &lt; x + 1</code> is always true.
</p>


<h3 id="Comparison_operators">Comparison operators</h3>

<p>
Comparison operators compare two operands and yield an untyped boolean value.
</p>

<pre class="grammar">
==    equal
!=    not equal
&lt;     less
&lt;=    less or equal
&gt;     greater
&gt;=    greater or equal
</pre>

<p>
In any comparison, the first operand
must be <a href="#Assignability">assignable</a>
to the type of the second operand, or vice versa.
</p>
<p>
The equality operators <code>==</code> and <code>!=</code> apply
to operands that are <i>comparable</i>.
The ordering operators <code>&lt;</code>, <code>&lt;=</code>, <code>&gt;</code>, and <code>&gt;=</code>
apply to operands that are <i>ordered</i>.
These terms and the result of the comparisons are defined as follows:
</p>

<ul>
	<li>
	Boolean values are comparable.
	Two boolean values are equal if they are either both
	<code>true</code> or both <code>false</code>.
	</li>

	<li>
	Integer values are comparable and ordered, in the usual way.
	</li>

	<li>
	Floating point values are comparable and ordered,
	as defined by the IEEE-754 standard.
	</li>

	<li>
	Complex values are comparable.
	Two complex values <code>u</code> and <code>v</code> are
	equal if both <code>real(u) == real(v)</code> and
	<code>imag(u) == imag(v)</code>.
	</li>

	<li>
	String values are comparable and ordered, lexically byte-wise.
	</li>

	<li>
	Pointer values are comparable.
	Two pointer values are equal if they point to the same variable or if both have value <code>nil</code>.
	Pointers to distinct <a href="#Size_and_alignment_guarantees">zero-size</a> variables may or may not be equal.
	</li>

	<li>
	Channel values are comparable.
	Two channel values are equal if they were created by the same call to
	<a href="#Making_slices_maps_and_channels"><code>make</code></a>
	or if both have value <code>nil</code>.
	</li>

	<li>
	Interface values are comparable.
	Two interface values are equal if they have <a href="#Type_identity">identical</a> dynamic types
	and equal dynamic values or if both have value <code>nil</code>.
	</li>

	<li>
	A value <code>x</code> of non-interface type <code>X</code> and
	a value <code>t</code> of interface type <code>T</code> are comparable when values
	of type <code>X</code> are comparable and
	<code>X</code> implements <code>T</code>.
	They are equal if <code>t</code>'s dynamic type is identical to <code>X</code>
	and <code>t</code>'s dynamic value is equal to <code>x</code>.
	</li>

	<li>
	Struct values are comparable if all their fields are comparable.
	Two struct values are equal if their corresponding
	non-<a href="#Blank_identifier">blank</a> fields are equal.
	</li>

	<li>
	Array values are comparable if values of the array element type are comparable.
	Two array values are equal if their corresponding elements are equal.
	</li>
</ul>

<p>
A comparison of two interface values with identical dynamic types
causes a <a href="#Run_time_panics">run-time panic</a> if values
of that type are not comparable.  This behavior applies not only to direct interface
value comparisons but also when comparing arrays of interface values
or structs with interface-valued fields.
</p>

<p>
Slice, map, and function values are not comparable.
However, as a special case, a slice, map, or function value may
be compared to the predeclared identifier <code>nil</code>.
Comparison of pointer, channel, and interface values to <code>nil</code>
is also allowed and follows from the general rules above.
</p>

<pre>
const c = 3 &lt; 4            // c is the untyped bool constant true

type MyBool bool
var x, y int
var (
	// The result of a comparison is an untyped bool.
	// The usual assignment rules apply.
	b3        = x == y // b3 has type bool
	b4 bool   = x == y // b4 has type bool
	b5 MyBool = x == y // b5 has type MyBool
)
</pre>

<h3 id="Logical_operators">Logical operators</h3>

<p>
Logical operators apply to <a href="#Boolean_types">boolean</a> values
and yield a result of the same type as the operands.
The right operand is evaluated conditionally.
</p>

<pre class="grammar">
&amp;&amp;    conditional AND    p &amp;&amp; q  is  "if p then q else false"
||    conditional OR     p || q  is  "if p then true else q"
!     NOT                !p      is  "not p"
</pre>


<h3 id="Address_operators">Address operators</h3>

<p>
For an operand <code>x</code> of type <code>T</code>, the address operation
<code>&amp;x</code> generates a pointer of type <code>*T</code> to <code>x</code>.
The operand must be <i>addressable</i>,
that is, either a variable, pointer indirection, or slice indexing
operation; or a field selector of an addressable struct operand;
or an array indexing operation of an addressable array.
As an exception to the addressability requirement, <code>x</code> may also be a
(possibly parenthesized)
<a href="#Composite_literals">composite literal</a>.
If the evaluation of <code>x</code> would cause a <a href="#Run_time_panics">run-time panic</a>,
then the evaluation of <code>&amp;x</code> does too.
</p>

<p>
For an operand <code>x</code> of pointer type <code>*T</code>, the pointer
indirection <code>*x</code> denotes the <a href="#Variables">variable</a> of type <code>T</code> pointed
to by <code>x</code>.
If <code>x</code> is <code>nil</code>, an attempt to evaluate <code>*x</code>
will cause a <a href="#Run_time_panics">run-time panic</a>.
</p>

<pre>
&amp;x
&amp;a[f(2)]
&amp;Point{2, 3}
*p
*pf(x)

var x *int = nil
*x   // causes a run-time panic
&amp;*x  // causes a run-time panic
</pre>


<h3 id="Receive_operator">Receive operator</h3>

<p>
For an operand <code>ch</code> of <a href="#Channel_types">channel type</a>,
the value of the receive operation <code>&lt;-ch</code> is the value received
from the channel <code>ch</code>. The channel direction must permit receive operations,
and the type of the receive operation is the element type of the channel.
The expression blocks until a value is available.
Receiving from a <code>nil</code> channel blocks forever.
A receive operation on a <a href="#Close">closed</a> channel can always proceed
immediately, yielding the element type's <a href="#The_zero_value">zero value</a>
after any previously sent values have been received.
</p>

<pre>
v1 := &lt;-ch
v2 = &lt;-ch
f(&lt;-ch)
&lt;-strobe  // wait until clock pulse and discard received value
</pre>

<p>
A receive expression used in an <a href="#Assignments">assignment</a> or initialization of the special form
</p>

<pre>
x, ok = &lt;-ch
x, ok := &lt;-ch
var x, ok = &lt;-ch
</pre>

<p>
yields an additional untyped boolean result reporting whether the
communication succeeded. The value of <code>ok</code> is <code>true</code>
if the value received was delivered by a successful send operation to the
channel, or <code>false</code> if it is a zero value generated because the
channel is closed and empty.
</p>


<h3 id="Conversions">Conversions</h3>

<p>
Conversions are expressions of the form <code>T(x)</code>
where <code>T</code> is a type and <code>x</code> is an expression
that can be converted to type <code>T</code>.
</p>

<pre class="ebnf">
<a id="Conversion">Conversion</a> = <a href="#Type" class="noline">Type</a> "(" <a href="#Expression" class="noline">Expression</a> [ "," ] ")" .
</pre>

<p>
If the type starts with the operator <code>*</code> or <code>&lt;-</code>,
or if the type starts with the keyword <code>func</code>
and has no result list, it must be parenthesized when
necessary to avoid ambiguity:
</p>

<pre>
*Point(p)        // same as *(Point(p))
(*Point)(p)      // p is converted to *Point
&lt;-chan int(c)    // same as &lt;-(chan int(c))
(&lt;-chan int)(c)  // c is converted to &lt;-chan int
func()(x)        // function signature func() x
(func())(x)      // x is converted to func()
(func() int)(x)  // x is converted to func() int
func() int(x)    // x is converted to func() int (unambiguous)
</pre>

<p>
A <a href="#Constants">constant</a> value <code>x</code> can be converted to
type <code>T</code> in any of these cases:
</p>

<ul>
	<li>
	<code>x</code> is representable by a value of type <code>T</code>.
	</li>
	<li>
	<code>x</code> is a floating-point constant,
	<code>T</code> is a floating-point type,
	and <code>x</code> is representable by a value
	of type <code>T</code> after rounding using
	IEEE 754 round-to-even rules.
	The constant <code>T(x)</code> is the rounded value.
	</li>
	<li>
	<code>x</code> is an integer constant and <code>T</code> is a
	<a href="#String_types">string type</a>.
	The <a href="#Conversions_to_and_from_a_string_type">same rule</a>
	as for non-constant <code>x</code> applies in this case.
	</li>
</ul>

<p>
Converting a constant yields a typed constant as result.
</p>

<pre>
uint(iota)               // iota value of type uint
float32(2.718281828)     // 2.718281828 of type float32
complex128(1)            // 1.0 + 0.0i of type complex128
float32(0.49999999)      // 0.5 of type float32
string('x')              // "x" of type string
string(0x266c)           // "♬" of type string
MyString("foo" + "bar")  // "foobar" of type MyString
string([]byte{'a'})      // not a constant: []byte{'a'} is not a constant
(*int)(nil)              // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
int(1.2)                 // illegal: 1.2 cannot be represented as an int
string(65.0)             // illegal: 65.0 is not an integer constant
</pre>

<p>
A non-constant value <code>x</code> can be converted to type <code>T</code>
in any of these cases:
</p>

<ul>
	<li>
	<code>x</code> is <a href="#Assignability">assignable</a>
	to <code>T</code>.
	</li>
	<li>
	<code>x</code>'s type and <code>T</code> have identical
	<a href="#Types">underlying types</a>.
	</li>
	<li>
	<code>x</code>'s type and <code>T</code> are unnamed pointer types
	and their pointer base types have identical underlying types.
	</li>
	<li>
	<code>x</code>'s type and <code>T</code> are both integer or floating
	point types.
	</li>
	<li>
	<code>x</code>'s type and <code>T</code> are both complex types.
	</li>
	<li>
	<code>x</code> is an integer or a slice of bytes or runes
	and <code>T</code> is a string type.
	</li>
	<li>
	<code>x</code> is a string and <code>T</code> is a slice of bytes or runes.
	</li>
</ul>

<p>
Specific rules apply to (non-constant) conversions between numeric types or
to and from a string type.
These conversions may change the representation of <code>x</code>
and incur a run-time cost.
All other conversions only change the type but not the representation
of <code>x</code>.
</p>

<p>
There is no linguistic mechanism to convert between pointers and integers.
The package <a href="#Package_unsafe"><code>unsafe</code></a>
implements this functionality under
restricted circumstances.
</p>

<h4>Conversions between numeric types</h4>

<p>
For the conversion of non-constant numeric values, the following rules apply:
</p>

<ol>
<li>
When converting between integer types, if the value is a signed integer, it is
sign extended to implicit infinite precision; otherwise it is zero extended.
It is then truncated to fit in the result type's size.
For example, if <code>v := uint16(0x10F0)</code>, then <code>uint32(int8(v)) == 0xFFFFFFF0</code>.
The conversion always yields a valid value; there is no indication of overflow.
</li>
<li>
When converting a floating-point number to an integer, the fraction is discarded
(truncation towards zero).
</li>
<li>
When converting an integer or floating-point number to a floating-point type,
or a complex number to another complex type, the result value is rounded
to the precision specified by the destination type.
For instance, the value of a variable <code>x</code> of type <code>float32</code>
may be stored using additional precision beyond that of an IEEE-754 32-bit number,
but float32(x) represents the result of rounding <code>x</code>'s value to
32-bit precision. Similarly, <code>x + 0.1</code> may use more than 32 bits
of precision, but <code>float32(x + 0.1)</code> does not.
</li>
</ol>

<p>
In all non-constant conversions involving floating-point or complex values,
if the result type cannot represent the value the conversion
succeeds but the result value is implementation-dependent.
</p>

<h4 id="Conversions_to_and_from_a_string_type">Conversions to and from a string type</h4>

<ol>
<li>
Converting a signed or unsigned integer value to a string type yields a
string containing the UTF-8 representation of the integer. Values outside
the range of valid Unicode code points are converted to <code>"\uFFFD"</code>.

<pre>
string('a')       // "a"
string(-1)        // "\ufffd" == "\xef\xbf\xbd"
string(0xf8)      // "\u00f8" == "ø" == "\xc3\xb8"
type MyString string
MyString(0x65e5)  // "\u65e5" == "日" == "\xe6\x97\xa5"
</pre>
</li>

<li>
Converting a slice of bytes to a string type yields
a string whose successive bytes are the elements of the slice.

<pre>
string([]byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'})   // "hellø"
string([]byte{})                                     // ""
string([]byte(nil))                                  // ""

type MyBytes []byte
string(MyBytes{'h', 'e', 'l', 'l', '\xc3', '\xb8'})  // "hellø"
</pre>
</li>

<li>
Converting a slice of runes to a string type yields
a string that is the concatenation of the individual rune values
converted to strings.

<pre>
string([]rune{0x767d, 0x9d6c, 0x7fd4})   // "\u767d\u9d6c\u7fd4" == "白鵬翔"
string([]rune{})                         // ""
string([]rune(nil))                      // ""

type MyRunes []rune
string(MyRunes{0x767d, 0x9d6c, 0x7fd4})  // "\u767d\u9d6c\u7fd4" == "白鵬翔"
</pre>
</li>

<li>
Converting a value of a string type to a slice of bytes type
yields a slice whose successive elements are the bytes of the string.

<pre>
[]byte("hellø")   // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
[]byte("")        // []byte{}

MyBytes("hellø")  // []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8'}
</pre>
</li>

<li>
Converting a value of a string type to a slice of runes type
yields a slice containing the individual Unicode code points of the string.

<pre>
[]rune(MyString("白鵬翔"))  // []rune{0x767d, 0x9d6c, 0x7fd4}
[]rune("")                 // []rune{}

MyRunes("白鵬翔")           // []rune{0x767d, 0x9d6c, 0x7fd4}
</pre>
</li>
</ol>


<h3 id="Constant_expressions">Constant expressions</h3>

<p>
Constant expressions may contain only <a href="#Constants">constant</a>
operands and are evaluated at compile time.
</p>

<p>
Untyped boolean, numeric, and string constants may be used as operands
wherever it is legal to use an operand of boolean, numeric, or string type,
respectively.
Except for shift operations, if the operands of a binary operation are
different kinds of untyped constants, the operation and, for non-boolean operations, the result use
the kind that appears later in this list: integer, rune, floating-point, complex.
For example, an untyped integer constant divided by an
untyped complex constant yields an untyped complex constant.
</p>

<p>
A constant <a href="#Comparison_operators">comparison</a> always yields
an untyped boolean constant.  If the left operand of a constant
<a href="#Operators">shift expression</a> is an untyped constant, the
result is an integer constant; otherwise it is a constant of the same
type as the left operand, which must be of
<a href="#Numeric_types">integer type</a>.
Applying all other operators to untyped constants results in an untyped
constant of the same kind (that is, a boolean, integer, floating-point,
complex, or string constant).
</p>

<pre>
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const Π float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)
const d = 1 &lt;&lt; 3.0         // d == 8     (untyped integer constant)
const e = 1.0 &lt;&lt; 3         // e == 8     (untyped integer constant)
const f = int32(1) &lt;&lt; 33   // illegal    (constant 8589934592 overflows int32)
const g = float64(2) &gt;&gt; 1  // illegal    (float64(2) is a typed floating-point constant)
const h = "foo" &gt; "bar"    // h == true  (untyped boolean constant)
const j = true             // j == true  (untyped boolean constant)
const k = 'w' + 1          // k == 'x'   (untyped rune constant)
const l = "hi"             // l == "hi"  (untyped string constant)
const m = string(k)        // m == "x"   (type string)
const Σ = 1 - 0.707i       //            (untyped complex constant)
const Δ = Σ + 2.0e-4       //            (untyped complex constant)
const Φ = iota*1i - 1/1i   //            (untyped complex constant)
</pre>

<p>
Applying the built-in function <code>complex</code> to untyped
integer, rune, or floating-point constants yields
an untyped complex constant.
</p>

<pre>
const ic = complex(0, c)   // ic == 3.75i  (untyped complex constant)
const iΘ = complex(0, Θ)   // iΘ == 1i     (type complex128)
</pre>

<p>
Constant expressions are always evaluated exactly; intermediate values and the
constants themselves may require precision significantly larger than supported
by any predeclared type in the language. The following are legal declarations:
</p>

<pre>
const Huge = 1 &lt;&lt; 100         // Huge == 1267650600228229401496703205376  (untyped integer constant)
const Four int8 = Huge &gt;&gt; 98  // Four == 4                                (type int8)
</pre>

<p>
The divisor of a constant division or remainder operation must not be zero:
</p>

<pre>
3.14 / 0.0   // illegal: division by zero
</pre>

<p>
The values of <i>typed</i> constants must always be accurately representable as values
of the constant type. The following constant expressions are illegal:
</p>

<pre>
uint(-1)     // -1 cannot be represented as a uint
int(3.14)    // 3.14 cannot be represented as an int
int64(Huge)  // 1267650600228229401496703205376 cannot be represented as an int64
Four * 300   // operand 300 cannot be represented as an int8 (type of Four)
Four * 100   // product 400 cannot be represented as an int8 (type of Four)
</pre>

<p>
The mask used by the unary bitwise complement operator <code>^</code> matches
the rule for non-constants: the mask is all 1s for unsigned constants
and -1 for signed and untyped constants.
</p>

<pre>
^1         // untyped integer constant, equal to -2
uint8(^1)  // illegal: same as uint8(-2), -2 cannot be represented as a uint8
^uint8(1)  // typed uint8 constant, same as 0xFF ^ uint8(1) = uint8(0xFE)
int8(^1)   // same as int8(-2)
^int8(1)   // same as -1 ^ int8(1) = -2
</pre>

<p>
Implementation restriction: A compiler may use rounding while
computing untyped floating-point or complex constant expressions; see
the implementation restriction in the section
on <a href="#Constants">constants</a>.  This rounding may cause a
floating-point constant expression to be invalid in an integer
context, even if it would be integral when calculated using infinite
precision.
</p>


<h3 id="Order_of_evaluation">Order of evaluation</h3>

<p>
At package level, <a href="#Package_initialization">initialization dependencies</a>
determine the evaluation order of individual initialization expressions in
<a href="#Variable_declarations">variable declarations</a>.
Otherwise, when evaluating the <a href="#Operands">operands</a> of an
expression, assignment, or
<a href="#Return_statements">return statement</a>,
all function calls, method calls, and
communication operations are evaluated in lexical left-to-right
order.
</p>

<p>
For example, in the (function-local) assignment
</p>
<pre>
y[f()], ok = g(h(), i()+x[j()], &lt;-c), k()
</pre>
<p>
the function calls and communication happen in the order
<code>f()</code>, <code>h()</code>, <code>i()</code>, <code>j()</code>,
<code>&lt;-c</code>, <code>g()</code>, and <code>k()</code>.
However, the order of those events compared to the evaluation
and indexing of <code>x</code> and the evaluation
of <code>y</code> is not specified.
</p>

<pre>
a := 1
f := func() int { a++; return a }
x := []int{a, f()}            // x may be [1, 2] or [2, 2]: evaluation order between a and f() is not specified
m := map[int]int{a: 1, a: 2}  // m may be {2: 1} or {2: 2}: evaluation order between the two map assignments is not specified
n := map[int]int{a: f()}      // n may be {2: 3} or {3: 3}: evaluation order between the key and the value is not specified
</pre>

<p>
At package level, initialization dependencies override the left-to-right rule
for individual initialization expressions, but not for operands within each
expression:
</p>

<pre>
var a, b, c = f() + v(), g(), sqr(u()) + v()

func f() int        { return c }
func g() int        { return a }
func sqr(x int) int { return x*x }

// functions u and v are independent of all other variables and functions
</pre>

<p>
The function calls happen in the order
<code>u()</code>, <code>sqr()</code>, <code>v()</code>,
<code>f()</code>, <code>v()</code>, and <code>g()</code>.
</p>

<p>
Floating-point operations within a single expression are evaluated according to
the associativity of the operators.  Explicit parentheses affect the evaluation
by overriding the default associativity.
In the expression <code>x + (y + z)</code> the addition <code>y + z</code>
is performed before adding <code>x</code>.
</p>

<h2 id="Statements">Statements</h2>

<p>
Statements control execution.
</p>

<pre class="ebnf">
<a id="Statement">Statement</a> =
	<a href="#Declaration" class="noline">Declaration</a> | <a href="#LabeledStmt" class="noline">LabeledStmt</a> | <a href="#SimpleStmt" class="noline">SimpleStmt</a> |
	<a href="#GoStmt" class="noline">GoStmt</a> | <a href="#ReturnStmt" class="noline">ReturnStmt</a> | <a href="#BreakStmt" class="noline">BreakStmt</a> | <a href="#ContinueStmt" class="noline">ContinueStmt</a> | <a href="#GotoStmt" class="noline">GotoStmt</a> |
	<a href="#FallthroughStmt" class="noline">FallthroughStmt</a> | <a href="#Block" class="noline">Block</a> | <a href="#IfStmt" class="noline">IfStmt</a> | <a href="#SwitchStmt" class="noline">SwitchStmt</a> | <a href="#SelectStmt" class="noline">SelectStmt</a> | <a href="#ForStmt" class="noline">ForStmt</a> |
	<a href="#DeferStmt" class="noline">DeferStmt</a> .

<a id="SimpleStmt">SimpleStmt</a> = <a href="#EmptyStmt" class="noline">EmptyStmt</a> | <a href="#ExpressionStmt" class="noline">ExpressionStmt</a> | <a href="#SendStmt" class="noline">SendStmt</a> | <a href="#IncDecStmt" class="noline">IncDecStmt</a> | <a href="#Assignment" class="noline">Assignment</a> | <a href="#ShortVarDecl" class="noline">ShortVarDecl</a> .
</pre>

<h3 id="Terminating_statements">Terminating statements</h3>

<p>
A terminating statement is one of the following:
</p>

<ol>
<li>
	A <a href="#Return_statements">"return"</a> or
    	<a href="#Goto_statements">"goto"</a> statement.
	<!-- ul below only for regular layout -->
	<ul> </ul>
</li>

<li>
	A call to the built-in function
	<a href="#Handling_panics"><code>panic</code></a>.
	<!-- ul below only for regular layout -->
	<ul> </ul>
</li>

<li>
	A <a href="#Blocks">block</a> in which the statement list ends in a terminating statement.
	<!-- ul below only for regular layout -->
	<ul> </ul>
</li>

<li>
	An <a href="#If_statements">"if" statement</a> in which:
	<ul>
	<li>the "else" branch is present, and</li>
	<li>both branches are terminating statements.</li>
	</ul>
</li>

<li>
	A <a href="#For_statements">"for" statement</a> in which:
	<ul>
	<li>there are no "break" statements referring to the "for" statement, and</li>
	<li>the loop condition is absent.</li>
	</ul>
</li>

<li>
	A <a href="#Switch_statements">"switch" statement</a> in which:
	<ul>
	<li>there are no "break" statements referring to the "switch" statement,</li>
	<li>there is a default case, and</li>
	<li>the statement lists in each case, including the default, end in a terminating
	    statement, or a possibly labeled <a href="#Fallthrough_statements">"fallthrough"
	    statement</a>.</li>
	</ul>
</li>

<li>
	A <a href="#Select_statements">"select" statement</a> in which:
	<ul>
	<li>there are no "break" statements referring to the "select" statement, and</li>
	<li>the statement lists in each case, including the default if present,
	    end in a terminating statement.</li>
	</ul>
</li>

<li>
	A <a href="#Labeled_statements">labeled statement</a> labeling
	a terminating statement.
</li>
</ol>

<p>
All other statements are not terminating.
</p>

<p>
A <a href="#Blocks">statement list</a> ends in a terminating statement if the list
is not empty and its final statement is terminating.
</p>


<h3 id="Empty_statements">Empty statements</h3>

<p>
The empty statement does nothing.
</p>

<pre class="ebnf">
<a id="EmptyStmt">EmptyStmt</a> = .
</pre>


<h3 id="Labeled_statements">Labeled statements</h3>

<p>
A labeled statement may be the target of a <code>goto</code>,
<code>break</code> or <code>continue</code> statement.
</p>

<pre class="ebnf">
<a id="LabeledStmt">LabeledStmt</a> = <a href="#Label" class="noline">Label</a> ":" <a href="#Statement" class="noline">Statement</a> .
<a id="Label">Label</a>       = <a href="#identifier" class="noline">identifier</a> .
</pre>

<pre>
Error: log.Panic("error encountered")
</pre>


<h3 id="Expression_statements">Expression statements</h3>

<p>
With the exception of specific built-in functions,
function and method <a href="#Calls">calls</a> and
<a href="#Receive_operator">receive operations</a>
can appear in statement context. Such statements may be parenthesized.
</p>

<pre class="ebnf">
<a id="ExpressionStmt">ExpressionStmt</a> = <a href="#Expression" class="noline">Expression</a> .
</pre>

<p>
The following built-in functions are not permitted in statement context:
</p>

<pre>
append cap complex imag len make new real
unsafe.Alignof unsafe.Offsetof unsafe.Sizeof
</pre>

<pre>
h(x+y)
f.Close()
&lt;-ch
(&lt;-ch)
len("foo")  // illegal if len is the built-in function
</pre>


<h3 id="Send_statements">Send statements</h3>

<p>
A send statement sends a value on a channel.
The channel expression must be of <a href="#Channel_types">channel type</a>,
the channel direction must permit send operations,
and the type of the value to be sent must be <a href="#Assignability">assignable</a>
to the channel's element type.
</p>

<pre class="ebnf">
<a id="SendStmt">SendStmt</a> = <a href="#Channel" class="noline">Channel</a> "&lt;-" <a href="#Expression" class="noline">Expression</a> .
<a id="Channel">Channel</a>  = <a href="#Expression" class="noline">Expression</a> .
</pre>

<p>
Both the channel and the value expression are evaluated before communication
begins. Communication blocks until the send can proceed.
A send on an unbuffered channel can proceed if a receiver is ready.
A send on a buffered channel can proceed if there is room in the buffer.
A send on a closed channel proceeds by causing a <a href="#Run_time_panics">run-time panic</a>.
A send on a <code>nil</code> channel blocks forever.
</p>

<pre>
ch &lt;- 3  // send value 3 to channel ch
</pre>


<h3 id="IncDec_statements">IncDec statements</h3>

<p>
The "++" and "--" statements increment or decrement their operands
by the untyped <a href="#Constants">constant</a> <code>1</code>.
As with an assignment, the operand must be <a href="#Address_operators">addressable</a>
or a map index expression.
</p>

<pre class="ebnf">
<a id="IncDecStmt">IncDecStmt</a> = <a href="#Expression" class="noline">Expression</a> ( "++" | "--" ) .
</pre>

<p>
The following <a href="#Assignments">assignment statements</a> are semantically
equivalent:
</p>

<pre class="grammar">
IncDec statement    Assignment
x++                 x += 1
x--                 x -= 1
</pre>


<h3 id="Assignments">Assignments</h3>

<pre class="ebnf">
<a id="Assignment">Assignment</a> = <a href="#ExpressionList" class="noline">ExpressionList</a> <a href="#assign_op" class="noline">assign_op</a> <a href="#ExpressionList" class="noline">ExpressionList</a> .

<a id="assign_op">assign_op</a> = [ <a href="#add_op" class="noline">add_op</a> | <a href="#mul_op" class="noline">mul_op</a> ] "=" .
</pre>

<p>
Each left-hand side operand must be <a href="#Address_operators">addressable</a>,
a map index expression, or (for <code>=</code> assignments only) the
<a href="#Blank_identifier">blank identifier</a>.
Operands may be parenthesized.
</p>

<pre>
x = 1
*p = f()
a[i] = 23
(k) = &lt;-ch  // same as: k = &lt;-ch
</pre>

<p>
An <i>assignment operation</i> <code>x</code> <i>op</i><code>=</code>
<code>y</code> where <i>op</i> is a binary arithmetic operation is equivalent
to <code>x</code> <code>=</code> <code>x</code> <i>op</i>
<code>y</code> but evaluates <code>x</code>
only once.  The <i>op</i><code>=</code> construct is a single token.
In assignment operations, both the left- and right-hand expression lists
must contain exactly one single-valued expression, and the left-hand
expression must not be the blank identifier.
</p>

<pre>
a[i] &lt;&lt;= 2
i &amp;^= 1&lt;&lt;n
</pre>

<p>
A tuple assignment assigns the individual elements of a multi-valued
operation to a list of variables.  There are two forms.  In the
first, the right hand operand is a single multi-valued expression
such as a function call, a <a href="#Channel_types">channel</a> or
<a href="#Map_types">map</a> operation, or a <a href="#Type_assertions">type assertion</a>.
The number of operands on the left
hand side must match the number of values.  For instance, if
<code>f</code> is a function returning two values,
</p>

<pre>
x, y = f()
</pre>

<p>
assigns the first value to <code>x</code> and the second to <code>y</code>.
In the second form, the number of operands on the left must equal the number
of expressions on the right, each of which must be single-valued, and the
<i>n</i>th expression on the right is assigned to the <i>n</i>th
operand on the left:
</p>

<pre>
one, two, three = '一', '二', '三'
</pre>

<p>
The <a href="#Blank_identifier">blank identifier</a> provides a way to
ignore right-hand side values in an assignment:
</p>

<pre>
_ = x       // evaluate x but ignore it
x, _ = f()  // evaluate f() but ignore second result value
</pre>

<p>
The assignment proceeds in two phases.
First, the operands of <a href="#Index_expressions">index expressions</a>
and <a href="#Address_operators">pointer indirections</a>
(including implicit pointer indirections in <a href="#Selectors">selectors</a>)
on the left and the expressions on the right are all
<a href="#Order_of_evaluation">evaluated in the usual order</a>.
Second, the assignments are carried out in left-to-right order.
</p>

<pre>
a, b = b, a  // exchange a and b

x := []int{1, 2, 3}
i := 0
i, x[i] = 1, 2  // set i = 1, x[0] = 2

i = 0
x[i], i = 2, 1  // set x[0] = 2, i = 1

x[0], x[0] = 1, 2  // set x[0] = 1, then x[0] = 2 (so x[0] == 2 at end)

x[1], x[3] = 4, 5  // set x[1] = 4, then panic setting x[3] = 5.

type Point struct { x, y int }
var p *Point
x[2], p.x = 6, 7  // set x[2] = 6, then panic setting p.x = 7

i = 2
x = []int{3, 5, 7}
for i, x[i] = range x {  // set i, x[2] = 0, x[0]
	break
}
// after this loop, i == 0 and x == []int{3, 5, 3}
</pre>

<p>
In assignments, each value must be <a href="#Assignability">assignable</a>
to the type of the operand to which it is assigned, with the following special cases:
</p>

<ol>
<li>
	Any typed value may be assigned to the blank identifier.
</li>

<li>
	If an untyped constant
	is assigned to a variable of interface type or the blank identifier,
	the constant is first <a href="#Conversions">converted</a> to its
	 <a href="#Constants">default type</a>.
</li>

<li>
	If an untyped boolean value is assigned to a variable of interface type or
	the blank identifier, it is first converted to type <code>bool</code>.
</li>
</ol>

<h3 id="If_statements">If statements</h3>

<p>
"If" statements specify the conditional execution of two branches
according to the value of a boolean expression.  If the expression
evaluates to true, the "if" branch is executed, otherwise, if
present, the "else" branch is executed.
</p>

<pre class="ebnf">
<a id="IfStmt">IfStmt</a> = "if" [ <a href="#SimpleStmt" class="noline">SimpleStmt</a> ";" ] <a href="#Expression" class="noline">Expression</a> <a href="#Block" class="noline">Block</a> [ "else" ( <a href="#IfStmt" class="noline">IfStmt</a> | <a href="#Block" class="noline">Block</a> ) ] .
</pre>

<pre>
if x &gt; max {
	x = max
}
</pre>

<p>
The expression may be preceded by a simple statement, which
executes before the expression is evaluated.
</p>

<pre>
if x := f(); x &lt; y {
	return x
} else if x &gt; z {
	return z
} else {
	return y
}
</pre>


<h3 id="Switch_statements">Switch statements</h3>

<p>
"Switch" statements provide multi-way execution.
An expression or type specifier is compared to the "cases"
inside the "switch" to determine which branch
to execute.
</p>

<pre class="ebnf">
<a id="SwitchStmt">SwitchStmt</a> = <a href="#ExprSwitchStmt" class="noline">ExprSwitchStmt</a> | <a href="#TypeSwitchStmt" class="noline">TypeSwitchStmt</a> .
</pre>

<p>
There are two forms: expression switches and type switches.
In an expression switch, the cases contain expressions that are compared
against the value of the switch expression.
In a type switch, the cases contain types that are compared against the
type of a specially annotated switch expression.
</p>

<h4 id="Expression_switches">Expression switches</h4>

<p>
In an expression switch,
the switch expression is evaluated and
the case expressions, which need not be constants,
are evaluated left-to-right and top-to-bottom; the first one that equals the
switch expression
triggers execution of the statements of the associated case;
the other cases are skipped.
If no case matches and there is a "default" case,
its statements are executed.
There can be at most one default case and it may appear anywhere in the
"switch" statement.
A missing switch expression is equivalent to the boolean value
<code>true</code>.
</p>

<pre class="ebnf">
<a id="ExprSwitchStmt">ExprSwitchStmt</a> = "switch" [ <a href="#SimpleStmt" class="noline">SimpleStmt</a> ";" ] [ <a href="#Expression" class="noline">Expression</a> ] "{" { <a href="#ExprCaseClause" class="noline">ExprCaseClause</a> } "}" .
<a id="ExprCaseClause">ExprCaseClause</a> = <a href="#ExprSwitchCase" class="noline">ExprSwitchCase</a> ":" <a href="#StatementList" class="noline">StatementList</a> .
<a id="ExprSwitchCase">ExprSwitchCase</a> = "case" <a href="#ExpressionList" class="noline">ExpressionList</a> | "default" .
</pre>

<p>
In a case or default clause, the last non-empty statement
may be a (possibly <a href="#Labeled_statements">labeled</a>)
<a href="#Fallthrough_statements">"fallthrough" statement</a> to
indicate that control should flow from the end of this clause to
the first statement of the next clause.
Otherwise control flows to the end of the "switch" statement.
A "fallthrough" statement may appear as the last statement of all
but the last clause of an expression switch.
</p>

<p>
The expression may be preceded by a simple statement, which
executes before the expression is evaluated.
</p>

<pre>
switch tag {
default: s3()
case 0, 1, 2, 3: s1()
case 4, 5, 6, 7: s2()
}

switch x := f(); {  // missing switch expression means "true"
case x &lt; 0: return -x
default: return x
}

switch {
case x &lt; y: f1()
case x &lt; z: f2()
case x == 4: f3()
}
</pre>

<h4 id="Type_switches">Type switches</h4>

<p>
A type switch compares types rather than values. It is otherwise similar
to an expression switch. It is marked by a special switch expression that
has the form of a <a href="#Type_assertions">type assertion</a>
using the reserved word <code>type</code> rather than an actual type:
</p>

<pre>
switch x.(type) {
// cases
}
</pre>

<p>
Cases then match actual types <code>T</code> against the dynamic type of the
expression <code>x</code>. As with type assertions, <code>x</code> must be of
<a href="#Interface_types">interface type</a>, and each non-interface type
<code>T</code> listed in a case must implement the type of <code>x</code>.
</p>

<pre class="ebnf">
<a id="TypeSwitchStmt">TypeSwitchStmt</a>  = "switch" [ <a href="#SimpleStmt" class="noline">SimpleStmt</a> ";" ] <a href="#TypeSwitchGuard" class="noline">TypeSwitchGuard</a> "{" { <a href="#TypeCaseClause" class="noline">TypeCaseClause</a> } "}" .
<a id="TypeSwitchGuard">TypeSwitchGuard</a> = [ <a href="#identifier" class="noline">identifier</a> ":=" ] <a href="#PrimaryExpr" class="noline">PrimaryExpr</a> "." "(" "type" ")" .
<a id="TypeCaseClause">TypeCaseClause</a>  = <a href="#TypeSwitchCase" class="noline">TypeSwitchCase</a> ":" <a href="#StatementList" class="noline">StatementList</a> .
<a id="TypeSwitchCase">TypeSwitchCase</a>  = "case" <a href="#TypeList" class="noline">TypeList</a> | "default" .
<a id="TypeList">TypeList</a>        = <a href="#Type" class="noline">Type</a> { "," <a href="#Type" class="noline">Type</a> } .
</pre>

<p>
The TypeSwitchGuard may include a
<a href="#Short_variable_declarations">short variable declaration</a>.
When that form is used, the variable is declared at the beginning of
the <a href="#Blocks">implicit block</a> in each clause.
In clauses with a case listing exactly one type, the variable
has that type; otherwise, the variable has the type of the expression
in the TypeSwitchGuard.
</p>

<p>
The type in a case may be <a href="#Predeclared_identifiers"><code>nil</code></a>;
that case is used when the expression in the TypeSwitchGuard
is a <code>nil</code> interface value.
</p>

<p>
Given an expression <code>x</code> of type <code>interface{}</code>,
the following type switch:
</p>

<pre>
switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
</pre>

<p>
could be rewritten:
</p>

<pre>
v := x  // x is evaluated exactly once
if v == nil {
	i := v                                 // type of i is type of x (interface{})
	printString("x is nil")
} else if i, isInt := v.(int); isInt {
	printInt(i)                            // type of i is int
} else if i, isFloat64 := v.(float64); isFloat64 {
	printFloat64(i)                        // type of i is float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	printFunction(i)                       // type of i is func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // type of i is type of x (interface{})
		printString("type is bool or string")
	} else {
		i := v                         // type of i is type of x (interface{})
		printString("don't know the type")
	}
}
</pre>

<p>
The type switch guard may be preceded by a simple statement, which
executes before the guard is evaluated.
</p>

<p>
The "fallthrough" statement is not permitted in a type switch.
</p>

<h3 id="For_statements">For statements</h3>

<p>
A "for" statement specifies repeated execution of a block. The iteration is
controlled by a condition, a "for" clause, or a "range" clause.
</p>

<pre class="ebnf">
<a id="ForStmt">ForStmt</a> = "for" [ <a href="#Condition" class="noline">Condition</a> | <a href="#ForClause" class="noline">ForClause</a> | <a href="#RangeClause" class="noline">RangeClause</a> ] <a href="#Block" class="noline">Block</a> .
<a id="Condition">Condition</a> = <a href="#Expression" class="noline">Expression</a> .
</pre>

<p>
In its simplest form, a "for" statement specifies the repeated execution of
a block as long as a boolean condition evaluates to true.
The condition is evaluated before each iteration.
If the condition is absent, it is equivalent to the boolean value
<code>true</code>.
</p>

<pre>
for a &lt; b {
	a *= 2
}
</pre>

<p>
A "for" statement with a ForClause is also controlled by its condition, but
additionally it may specify an <i>init</i>
and a <i>post</i> statement, such as an assignment,
an increment or decrement statement. The init statement may be a
<a href="#Short_variable_declarations">short variable declaration</a>, but the post statement must not.
Variables declared by the init statement are re-used in each iteration.
</p>

<pre class="ebnf">
<a id="ForClause">ForClause</a> = [ <a href="#InitStmt" class="noline">InitStmt</a> ] ";" [ <a href="#Condition" class="noline">Condition</a> ] ";" [ <a href="#PostStmt" class="noline">PostStmt</a> ] .
<a id="InitStmt">InitStmt</a> = <a href="#SimpleStmt" class="noline">SimpleStmt</a> .
<a id="PostStmt">PostStmt</a> = <a href="#SimpleStmt" class="noline">SimpleStmt</a> .
</pre>

<pre>
for i := 0; i &lt; 10; i++ {
	f(i)
}
</pre>

<p>
If non-empty, the init statement is executed once before evaluating the
condition for the first iteration;
the post statement is executed after each execution of the block (and
only if the block was executed).
Any element of the ForClause may be empty but the
<a href="#Semicolons">semicolons</a> are
required unless there is only a condition.
If the condition is absent, it is equivalent to the boolean value
<code>true</code>.
</p>

<pre>
for cond { S() }    is the same as    for ; cond ; { S() }
for      { S() }    is the same as    for true     { S() }
</pre>

<p>
A "for" statement with a "range" clause
iterates through all entries of an array, slice, string or map,
or values received on a channel. For each entry it assigns <i>iteration values</i>
to corresponding <i>iteration variables</i> if present and then executes the block.
</p>

<pre class="ebnf">
<a id="RangeClause">RangeClause</a> = [ <a href="#ExpressionList" class="noline">ExpressionList</a> "=" | <a href="#IdentifierList" class="noline">IdentifierList</a> ":=" ] "range" <a href="#Expression" class="noline">Expression</a> .
</pre>

<p>
The expression on the right in the "range" clause is called the <i>range expression</i>,
which may be an array, pointer to an array, slice, string, map, or channel permitting
<a href="#Receive_operator">receive operations</a>.
As with an assignment, if present the operands on the left must be
<a href="#Address_operators">addressable</a> or map index expressions; they
denote the iteration variables. If the range expression is a channel, at most
one iteration variable is permitted, otherwise there may be up to two.
If the last iteration variable is the <a href="#Blank_identifier">blank identifier</a>,
the range clause is equivalent to the same clause without that identifier.
</p>

<p>
The range expression is evaluated once before beginning the loop,
with one exception: if the range expression is an array or a pointer to an array
and at most one iteration variable is present, only the range expression's
length is evaluated; if that length is constant,
<a href="#Length_and_capacity">by definition</a>
the range expression itself will not be evaluated.
</p>

<p>
Function calls on the left are evaluated once per iteration.
For each iteration, iteration values are produced as follows
if the respective iteration variables are present:
</p>

<pre class="grammar">
Range expression                          1st value          2nd value

array or slice  a  [n]E, *[n]E, or []E    index    i  int    a[i]       E
string          s  string type            index    i  int    see below  rune
map             m  map[K]V                key      k  K      m[k]       V
channel         c  chan E, &lt;-chan E       element  e  E
</pre>

<ol>
<li>
For an array, pointer to array, or slice value <code>a</code>, the index iteration
values are produced in increasing order, starting at element index 0.
If at most one iteration variable is present, the range loop produces
iteration values from 0 up to <code>len(a)-1</code> and does not index into the array
or slice itself. For a <code>nil</code> slice, the number of iterations is 0.
</li>

<li>
For a string value, the "range" clause iterates over the Unicode code points
in the string starting at byte index 0.  On successive iterations, the index value will be the
index of the first byte of successive UTF-8-encoded code points in the string,
and the second value, of type <code>rune</code>, will be the value of
the corresponding code point.  If the iteration encounters an invalid
UTF-8 sequence, the second value will be <code>0xFFFD</code>,
the Unicode replacement character, and the next iteration will advance
a single byte in the string.
</li>

<li>
The iteration order over maps is not specified
and is not guaranteed to be the same from one iteration to the next.
If map entries that have not yet been reached are removed during iteration,
the corresponding iteration values will not be produced. If map entries are
created during iteration, that entry may be produced during the iteration or
may be skipped. The choice may vary for each entry created and from one
iteration to the next.
If the map is <code>nil</code>, the number of iterations is 0.
</li>

<li>
For channels, the iteration values produced are the successive values sent on
the channel until the channel is <a href="#Close">closed</a>. If the channel
is <code>nil</code>, the range expression blocks forever.
</li>
</ol>

<p>
The iteration values are assigned to the respective
iteration variables as in an <a href="#Assignments">assignment statement</a>.
</p>

<p>
The iteration variables may be declared by the "range" clause using a form of
<a href="#Short_variable_declarations">short variable declaration</a>
(<code>:=</code>).
In this case their types are set to the types of the respective iteration values
and their <a href="#Declarations_and_scope">scope</a> is the block of the "for"
statement; they are re-used in each iteration.
If the iteration variables are declared outside the "for" statement,
after execution their values will be those of the last iteration.
</p>

<pre>
var testdata *struct {
	a *[7]int
}
for i, _ := range testdata.a {
	// testdata.a is never evaluated; len(testdata.a) is constant
	// i ranges from 0 to 6
	f(i)
}

var a [10]string
for i, s := range a {
	// type of i is int
	// type of s is string
	// s == a[i]
	g(i, s)
}

var key string
var val interface {}  // value type of m is assignable to val
m := map[string]int{"mon":0, "tue":1, "wed":2, "thu":3, "fri":4, "sat":5, "sun":6}
for key, val = range m {
	h(key, val)
}
// key == last map key encountered in iteration
// val == map[key]

var ch chan Work = producer()
for w := range ch {
	doWork(w)
}

// empty a channel
for range ch {}
</pre>


<h3 id="Go_statements">Go statements</h3>

<p>
A "go" statement starts the execution of a function call
as an independent concurrent thread of control, or <i>goroutine</i>,
within the same address space.
</p>

<pre class="ebnf">
<a id="GoStmt">GoStmt</a> = "go" <a href="#Expression" class="noline">Expression</a> .
</pre>

<p>
The expression must be a function or method call; it cannot be parenthesized.
Calls of built-in functions are restricted as for
<a href="#Expression_statements">expression statements</a>.
</p>

<p>
The function value and parameters are
<a href="#Calls">evaluated as usual</a>
in the calling goroutine, but
unlike with a regular call, program execution does not wait
for the invoked function to complete.
Instead, the function begins executing independently
in a new goroutine.
When the function terminates, its goroutine also terminates.
If the function has any return values, they are discarded when the
function completes.
</p>

<pre>
go Server()
go func(ch chan&lt;- bool) { for { sleep(10); ch &lt;- true; }} (c)
</pre>


<h3 id="Select_statements">Select statements</h3>

<p>
A "select" statement chooses which of a set of possible
<a href="#Send_statements">send</a> or
<a href="#Receive_operator">receive</a>
operations will proceed.
It looks similar to a
<a href="#Switch_statements">"switch"</a> statement but with the
cases all referring to communication operations.
</p>

<pre class="ebnf">
<a id="SelectStmt">SelectStmt</a> = "select" "{" { <a href="#CommClause" class="noline">CommClause</a> } "}" .
<a id="CommClause">CommClause</a> = <a href="#CommCase" class="noline">CommCase</a> ":" <a href="#StatementList" class="noline">StatementList</a> .
<a id="CommCase">CommCase</a>   = "case" ( <a href="#SendStmt" class="noline">SendStmt</a> | <a href="#RecvStmt" class="noline">RecvStmt</a> ) | "default" .
<a id="RecvStmt">RecvStmt</a>   = [ <a href="#ExpressionList" class="noline">ExpressionList</a> "=" | <a href="#IdentifierList" class="noline">IdentifierList</a> ":=" ] <a href="#RecvExpr" class="noline">RecvExpr</a> .
<a id="RecvExpr">RecvExpr</a>   = <a href="#Expression" class="noline">Expression</a> .
</pre>

<p>
A case with a RecvStmt may assign the result of a RecvExpr to one or
two variables, which may be declared using a
<a href="#Short_variable_declarations">short variable declaration</a>.
The RecvExpr must be a (possibly parenthesized) receive operation.
There can be at most one default case and it may appear anywhere
in the list of cases.
</p>

<p>
Execution of a "select" statement proceeds in several steps:
</p>

<ol>
<li>
For all the cases in the statement, the channel operands of receive operations
and the channel and right-hand-side expressions of send statements are
evaluated exactly once, in source order, upon entering the "select" statement.
The result is a set of channels to receive from or send to,
and the corresponding values to send.
Any side effects in that evaluation will occur irrespective of which (if any)
communication operation is selected to proceed.
Expressions on the left-hand side of a RecvStmt with a short variable declaration
or assignment are not yet evaluated.
</li>

<li>
If one or more of the communications can proceed,
a single one that can proceed is chosen via a uniform pseudo-random selection.
Otherwise, if there is a default case, that case is chosen.
If there is no default case, the "select" statement blocks until
at least one of the communications can proceed.
</li>

<li>
Unless the selected case is the default case, the respective communication
operation is executed.
</li>

<li>
If the selected case is a RecvStmt with a short variable declaration or
an assignment, the left-hand side expressions are evaluated and the
received value (or values) are assigned.
</li>

<li>
The statement list of the selected case is executed.
</li>
</ol>

<p>
Since communication on <code>nil</code> channels can never proceed,
a select with only <code>nil</code> channels and no default case blocks forever.
</p>

<pre>
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = &lt;-c1:
	print("received ", i1, " from c1\n")
case c2 &lt;- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (&lt;-c3):  // same as: i3, ok := &lt;-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
case a[f()] = &lt;-c4:
	// same as:
	// case t := &lt;-c4
	//	a[f()] = t
default:
	print("no communication\n")
}

for {  // send random sequence of bits to c
	select {
	case c &lt;- 0:  // note: no statement, no fallthrough, no folding of cases
	case c &lt;- 1:
	}
}

select {}  // block forever
</pre>


<h3 id="Return_statements">Return statements</h3>

<p>
A "return" statement in a function <code>F</code> terminates the execution
of <code>F</code>, and optionally provides one or more result values.
Any functions <a href="#Defer_statements">deferred</a> by <code>F</code>
are executed before <code>F</code> returns to its caller.
</p>

<pre class="ebnf">
<a id="ReturnStmt">ReturnStmt</a> = "return" [ <a href="#ExpressionList" class="noline">ExpressionList</a> ] .
</pre>

<p>
In a function without a result type, a "return" statement must not
specify any result values.
</p>
<pre>
func noResult() {
	return
}
</pre>

<p>
There are three ways to return values from a function with a result
type:
</p>

<ol>
	<li>The return value or values may be explicitly listed
		in the "return" statement. Each expression must be single-valued
		and <a href="#Assignability">assignable</a>
		to the corresponding element of the function's result type.
<pre>
func simpleF() int {
	return 2
}

func complexF1() (re float64, im float64) {
	return -7.0, -4.0
}
</pre>
	</li>
	<li>The expression list in the "return" statement may be a single
		call to a multi-valued function. The effect is as if each value
		returned from that function were assigned to a temporary
		variable with the type of the respective value, followed by a
		"return" statement listing these variables, at which point the
		rules of the previous case apply.
<pre>
func complexF2() (re float64, im float64) {
	return complexF1()
}
</pre>
	</li>
	<li>The expression list may be empty if the function's result
		type specifies names for its <a href="#Function_types">result parameters</a>.
		The result parameters act as ordinary local variables
		and the function may assign values to them as necessary.
		The "return" statement returns the values of these variables.
<pre>
func complexF3() (re float64, im float64) {
	re = 7.0
	im = 4.0
	return
}

func (devnull) Write(p []byte) (n int, _ error) {
	n = len(p)
	return
}
</pre>
	</li>
</ol>

<p>
Regardless of how they are declared, all the result values are initialized to
the <a href="#The_zero_value">zero values</a> for their type upon entry to the
function. A "return" statement that specifies results sets the result parameters before
any deferred functions are executed.
</p>

<p>
Implementation restriction: A compiler may disallow an empty expression list
in a "return" statement if a different entity (constant, type, or variable)
with the same name as a result parameter is in
<a href="#Declarations_and_scope">scope</a> at the place of the return.
</p>

<pre>
func f(n int) (res int, err error) {
	if _, err := f(n-1); err != nil {
		return  // invalid return statement: err is shadowed
	}
	return
}
</pre>

<h3 id="Break_statements">Break statements</h3>

<p>
A "break" statement terminates execution of the innermost
<a href="#For_statements">"for"</a>,
<a href="#Switch_statements">"switch"</a>, or
<a href="#Select_statements">"select"</a> statement
within the same function.
</p>

<pre class="ebnf">
<a id="BreakStmt">BreakStmt</a> = "break" [ <a href="#Label" class="noline">Label</a> ] .
</pre>

<p>
If there is a label, it must be that of an enclosing
"for", "switch", or "select" statement,
and that is the one whose execution terminates.
</p>

<pre>
OuterLoop:
	for i = 0; i &lt; n; i++ {
		for j = 0; j &lt; m; j++ {
			switch a[i][j] {
			case nil:
				state = Error
				break OuterLoop
			case item:
				state = Found
				break OuterLoop
			}
		}
	}
</pre>

<h3 id="Continue_statements">Continue statements</h3>

<p>
A "continue" statement begins the next iteration of the
innermost <a href="#For_statements">"for" loop</a> at its post statement.
The "for" loop must be within the same function.
</p>

<pre class="ebnf">
<a id="ContinueStmt">ContinueStmt</a> = "continue" [ <a href="#Label" class="noline">Label</a> ] .
</pre>

<p>
If there is a label, it must be that of an enclosing
"for" statement, and that is the one whose execution
advances.
</p>

<pre>
RowLoop:
	for y, row := range rows {
		for x, data := range row {
			if data == endOfRow {
				continue RowLoop
			}
			row[x] = data + bias(x, y)
		}
	}
</pre>

<h3 id="Goto_statements">Goto statements</h3>

<p>
A "goto" statement transfers control to the statement with the corresponding label
within the same function.
</p>

<pre class="ebnf">
<a id="GotoStmt">GotoStmt</a> = "goto" <a href="#Label" class="noline">Label</a> .
</pre>

<pre>
goto Error
</pre>

<p>
Executing the "goto" statement must not cause any variables to come into
<a href="#Declarations_and_scope">scope</a> that were not already in scope at the point of the goto.
For instance, this example:
</p>

<pre>
	goto L  // BAD
	v := 3
L:
</pre>

<p>
is erroneous because the jump to label <code>L</code> skips
the creation of <code>v</code>.
</p>

<p>
A "goto" statement outside a <a href="#Blocks">block</a> cannot jump to a label inside that block.
For instance, this example:
</p>

<pre>
if n%2 == 1 {
	goto L1
}
for n &gt; 0 {
	f()
	n--
L1:
	f()
	n--
}
</pre>

<p>
is erroneous because the label <code>L1</code> is inside
the "for" statement's block but the <code>goto</code> is not.
</p>

<h3 id="Fallthrough_statements">Fallthrough statements</h3>

<p>
A "fallthrough" statement transfers control to the first statement of the
next case clause in a <a href="#Expression_switches">expression "switch" statement</a>.
It may be used only as the final non-empty statement in such a clause.
</p>

<pre class="ebnf">
<a id="FallthroughStmt">FallthroughStmt</a> = "fallthrough" .
</pre>


<h3 id="Defer_statements">Defer statements</h3>

<p>
A "defer" statement invokes a function whose execution is deferred
to the moment the surrounding function returns, either because the
surrounding function executed a <a href="#Return_statements">return statement</a>,
reached the end of its <a href="#Function_declarations">function body</a>,
or because the corresponding goroutine is <a href="#Handling_panics">panicking</a>.
</p>

<pre class="ebnf">
<a id="DeferStmt">DeferStmt</a> = "defer" <a href="#Expression" class="noline">Expression</a> .
</pre>

<p>
The expression must be a function or method call; it cannot be parenthesized.
Calls of built-in functions are restricted as for
<a href="#Expression_statements">expression statements</a>.
</p>

<p>
Each time a "defer" statement
executes, the function value and parameters to the call are
<a href="#Calls">evaluated as usual</a>
and saved anew but the actual function is not invoked.
Instead, deferred functions are invoked immediately before
the surrounding function returns, in the reverse order
they were deferred.
If a deferred function value evaluates
to <code>nil</code>, execution <a href="#Handling_panics">panics</a>
when the function is invoked, not when the "defer" statement is executed.
</p>

<p>
For instance, if the deferred function is
a <a href="#Function_literals">function literal</a> and the surrounding
function has <a href="#Function_types">named result parameters</a> that
are in scope within the literal, the deferred function may access and modify
the result parameters before they are returned.
If the deferred function has any return values, they are discarded when
the function completes.
(See also the section on <a href="#Handling_panics">handling panics</a>.)
</p>

<pre>
lock(l)
defer unlock(l)  // unlocking happens before surrounding function returns

// prints 3 2 1 0 before surrounding function returns
for i := 0; i &lt;= 3; i++ {
	defer fmt.Print(i)
}

// f returns 1
func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}
</pre>

<h2 id="Built-in_functions">Built-in functions</h2>

<p>
Built-in functions are
<a href="#Predeclared_identifiers">predeclared</a>.
They are called like any other function but some of them
accept a type instead of an expression as the first argument.
</p>

<p>
The built-in functions do not have standard Go types,
so they can only appear in <a href="#Calls">call expressions</a>;
they cannot be used as function values.
</p>

<h3 id="Close">Close</h3>

<p>
For a channel <code>c</code>, the built-in function <code>close(c)</code>
records that no more values will be sent on the channel.
It is an error if <code>c</code> is a receive-only channel.
Sending to or closing a closed channel causes a <a href="#Run_time_panics">run-time panic</a>.
Closing the nil channel also causes a <a href="#Run_time_panics">run-time panic</a>.
After calling <code>close</code>, and after any previously
sent values have been received, receive operations will return
the zero value for the channel's type without blocking.
The multi-valued <a href="#Receive_operator">receive operation</a>
returns a received value along with an indication of whether the channel is closed.
</p>


<h3 id="Length_and_capacity">Length and capacity</h3>

<p>
The built-in functions <code>len</code> and <code>cap</code> take arguments
of various types and return a result of type <code>int</code>.
The implementation guarantees that the result always fits into an <code>int</code>.
</p>

<pre class="grammar">
Call      Argument type    Result

len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length
          map[K]T          map length (number of defined keys)
          chan T           number of elements queued in channel buffer

cap(s)    [n]T, *[n]T      array length (== n)
          []T              slice capacity
          chan T           channel buffer capacity
</pre>

<p>
The capacity of a slice is the number of elements for which there is
space allocated in the underlying array.
At any time the following relationship holds:
</p>

<pre>
0 &lt;= len(s) &lt;= cap(s)
</pre>

<p>
The length of a <code>nil</code> slice, map or channel is 0.
The capacity of a <code>nil</code> slice or channel is 0.
</p>

<p>
The expression <code>len(s)</code> is <a href="#Constants">constant</a> if
<code>s</code> is a string constant. The expressions <code>len(s)</code> and
<code>cap(s)</code> are constants if the type of <code>s</code> is an array
or pointer to an array and the expression <code>s</code> does not contain
<a href="#Receive_operator">channel receives</a> or (non-constant)
<a href="#Calls">function calls</a>; in this case <code>s</code> is not evaluated.
Otherwise, invocations of <code>len</code> and <code>cap</code> are not
constant and <code>s</code> is evaluated.
</p>

<pre>
const (
	c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
	c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
	c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
	c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
	c5 = len([10]float64{imag(z)})   // invalid: imag(x) is a (non-constant) function call
)
var z complex128
</pre>

<h3 id="Allocation">Allocation</h3>

<p>
The built-in function <code>new</code> takes a type <code>T</code>,
allocates storage for a <a href="#Variables">variable</a> of that type
at run time, and returns a value of type <code>*T</code>
<a href="#Pointer_types">pointing</a> to it.
The variable is initialized as described in the section on
<a href="#The_zero_value">initial values</a>.
</p>

<pre class="grammar">
new(T)
</pre>

<p>
For instance
</p>

<pre>
type S struct { a int; b float64 }
new(S)
</pre>

<p>
allocates storage for a variable of type <code>S</code>,
initializes it (<code>a=0</code>, <code>b=0.0</code>),
and returns a value of type <code>*S</code> containing the address
of the location.
</p>

<h3 id="Making_slices_maps_and_channels">Making slices, maps and channels</h3>

<p>
The built-in function <code>make</code> takes a type <code>T</code>,
which must be a slice, map or channel type,
optionally followed by a type-specific list of expressions.
It returns a value of type <code>T</code> (not <code>*T</code>).
The memory is initialized as described in the section on
<a href="#The_zero_value">initial values</a>.
</p>

<pre class="grammar">
Call             Type T     Result

make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
</pre>


<p>
The size arguments <code>n</code> and <code>m</code> must be of integer type or untyped.
A <a href="#Constants">constant</a> size argument must be non-negative and
representable by a value of type <code>int</code>.
If both <code>n</code> and <code>m</code> are provided and are constant, then
<code>n</code> must be no larger than <code>m</code>.
If <code>n</code> is negative or larger than <code>m</code> at run time,
a <a href="#Run_time_panics">run-time panic</a> occurs.
</p>

<pre>
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1&lt;&lt;63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for 100 elements
</pre>


<h3 id="Appending_and_copying_slices">Appending to and copying slices</h3>

<p>
The built-in functions <code>append</code> and <code>copy</code> assist in
common slice operations.
For both functions, the result is independent of whether the memory referenced
by the arguments overlaps.
</p>

<p>
The <a href="#Function_types">variadic</a> function <code>append</code>
appends zero or more values <code>x</code>
to <code>s</code> of type <code>S</code>, which must be a slice type, and
returns the resulting slice, also of type <code>S</code>.
The values <code>x</code> are passed to a parameter of type <code>...T</code>
where <code>T</code> is the <a href="#Slice_types">element type</a> of
<code>S</code> and the respective
<a href="#Passing_arguments_to_..._parameters">parameter passing rules</a> apply.
As a special case, <code>append</code> also accepts a first argument
assignable to type <code>[]byte</code> with a second argument of
string type followed by <code>...</code>. This form appends the
bytes of the string.
</p>

<pre class="grammar">
append(s S, x ...T) S  // T is the element type of S
</pre>

<p>
If the capacity of <code>s</code> is not large enough to fit the additional
values, <code>append</code> allocates a new, sufficiently large underlying
array that fits both the existing slice elements and the additional values.
Otherwise, <code>append</code> re-uses the underlying array.
</p>

<pre>
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")                                  t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents      b == []byte{'b', 'a', 'r' }
</pre>

<p>
The function <code>copy</code> copies slice elements from
a source <code>src</code> to a destination <code>dst</code> and returns the
number of elements copied.
Both arguments must have <a href="#Type_identity">identical</a> element type <code>T</code> and must be
<a href="#Assignability">assignable</a> to a slice of type <code>[]T</code>.
The number of elements copied is the minimum of
<code>len(src)</code> and <code>len(dst)</code>.
As a special case, <code>copy</code> also accepts a destination argument assignable
to type <code>[]byte</code> with a source argument of a string type.
This form copies the bytes from the string into the byte slice.
</p>

<pre class="grammar">
copy(dst, src []T) int
copy(dst []byte, src string) int
</pre>

<p>
Examples:
</p>

<pre>
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")
</pre>


<h3 id="Deletion_of_map_elements">Deletion of map elements</h3>

<p>
The built-in function <code>delete</code> removes the element with key
<code>k</code> from a <a href="#Map_types">map</a> <code>m</code>. The
type of <code>k</code> must be <a href="#Assignability">assignable</a>
to the key type of <code>m</code>.
</p>

<pre class="grammar">
delete(m, k)  // remove element m[k] from map m
</pre>

<p>
If the map <code>m</code> is <code>nil</code> or the element <code>m[k]</code>
does not exist, <code>delete</code> is a no-op.
</p>


<h3 id="Complex_numbers">Manipulating complex numbers</h3>

<p>
Three functions assemble and disassemble complex numbers.
The built-in function <code>complex</code> constructs a complex
value from a floating-point real and imaginary part, while
<code>real</code> and <code>imag</code>
extract the real and imaginary parts of a complex value.
</p>

<pre class="grammar">
complex(realPart, imaginaryPart floatT) complexT
real(complexT) floatT
imag(complexT) floatT
</pre>

<p>
The type of the arguments and return value correspond.
For <code>complex</code>, the two arguments must be of the same
floating-point type and the return type is the complex type
with the corresponding floating-point constituents:
<code>complex64</code> for <code>float32</code>,
<code>complex128</code> for <code>float64</code>.
The <code>real</code> and <code>imag</code> functions
together form the inverse, so for a complex value <code>z</code>,
<code>z</code> <code>==</code> <code>complex(real(z),</code> <code>imag(z))</code>.
</p>

<p>
If the operands of these functions are all constants, the return
value is a constant.
</p>

<pre>
var a = complex(2, -2)             // complex128
var b = complex(1.0, -1.4)         // complex128
x := float32(math.Cos(math.Pi/2))  // float32
var c64 = complex(5, -x)           // complex64
var im = imag(b)                   // float64
var rl = real(c64)                 // float32
</pre>

<h3 id="Handling_panics">Handling panics</h3>

<p> Two built-in functions, <code>panic</code> and <code>recover</code>,
assist in reporting and handling <a href="#Run_time_panics">run-time panics</a>
and program-defined error conditions.
</p>

<pre class="grammar">
func panic(interface{})
func recover() interface{}
</pre>

<p>
While executing a function <code>F</code>,
an explicit call to <code>panic</code> or a <a href="#Run_time_panics">run-time panic</a>
terminates the execution of <code>F</code>.
Any functions <a href="#Defer_statements">deferred</a> by <code>F</code>
are then executed as usual.
Next, any deferred functions run by <code>F's</code> caller are run,
and so on up to any deferred by the top-level function in the executing goroutine.
At that point, the program is terminated and the error
condition is reported, including the value of the argument to <code>panic</code>.
This termination sequence is called <i>panicking</i>.
</p>

<pre>
panic(42)
panic("unreachable")
panic(Error("cannot parse"))
</pre>

<p>
The <code>recover</code> function allows a program to manage behavior
of a panicking goroutine.
Suppose a function <code>G</code> defers a function <code>D</code> that calls
<code>recover</code> and a panic occurs in a function on the same goroutine in which <code>G</code>
is executing.
When the running of deferred functions reaches <code>D</code>,
the return value of <code>D</code>'s call to <code>recover</code> will be the value passed to the call of <code>panic</code>.
If <code>D</code> returns normally, without starting a new
<code>panic</code>, the panicking sequence stops. In that case,
the state of functions called between <code>G</code> and the call to <code>panic</code>
is discarded, and normal execution resumes.
Any functions deferred by <code>G</code> before <code>D</code> are then run and <code>G</code>'s
execution terminates by returning to its caller.
</p>

<p>
The return value of <code>recover</code> is <code>nil</code> if any of the following conditions holds:
</p>
<ul>
<li>
<code>panic</code>'s argument was <code>nil</code>;
</li>
<li>
the goroutine is not panicking;
</li>
<li>
<code>recover</code> was not called directly by a deferred function.
</li>
</ul>

<p>
The <code>protect</code> function in the example below invokes
the function argument <code>g</code> and protects callers from
run-time panics raised by <code>g</code>.
</p>

<pre>
func protect(g func()) {
	defer func() {
		log.Println("done")  // Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}
</pre>


<h3 id="Bootstrapping">Bootstrapping</h3>

<p>
Current implementations provide several built-in functions useful during
bootstrapping. These functions are documented for completeness but are not
guaranteed to stay in the language. They do not return a result.
</p>

<pre class="grammar">
Function   Behavior

print      prints all arguments; formatting of arguments is implementation-specific
println    like print but prints spaces between arguments and a newline at the end
</pre>


<h2 id="Packages">Packages</h2>

<p>
Go programs are constructed by linking together <i>packages</i>.
A package in turn is constructed from one or more source files
that together declare constants, types, variables and functions
belonging to the package and which are accessible in all files
of the same package. Those elements may be
<a href="#Exported_identifiers">exported</a> and used in another package.
</p>

<h3 id="Source_file_organization">Source file organization</h3>

<p>
Each source file consists of a package clause defining the package
to which it belongs, followed by a possibly empty set of import
declarations that declare packages whose contents it wishes to use,
followed by a possibly empty set of declarations of functions,
types, variables, and constants.
</p>

<pre class="ebnf">
<a id="SourceFile">SourceFile</a>       = <a href="#PackageClause" class="noline">PackageClause</a> ";" { <a href="#ImportDecl" class="noline">ImportDecl</a> ";" } { <a href="#TopLevelDecl" class="noline">TopLevelDecl</a> ";" } .
</pre>

<h3 id="Package_clause">Package clause</h3>

<p>
A package clause begins each source file and defines the package
to which the file belongs.
</p>

<pre class="ebnf">
<a id="PackageClause">PackageClause</a>  = "package" <a href="#PackageName" class="noline">PackageName</a> .
<a id="PackageName">PackageName</a>    = <a href="#identifier" class="noline">identifier</a> .
</pre>

<p>
The PackageName must not be the <a href="#Blank_identifier">blank identifier</a>.
</p>

<pre>
package math
</pre>

<p>
A set of files sharing the same PackageName form the implementation of a package.
An implementation may require that all source files for a package inhabit the same directory.
</p>

<h3 id="Import_declarations">Import declarations</h3>

<p>
An import declaration states that the source file containing the declaration
depends on functionality of the <i>imported</i> package
(<a href="#Program_initialization_and_execution">§Program initialization and execution</a>)
and enables access to <a href="#Exported_identifiers">exported</a> identifiers
of that package.
The import names an identifier (PackageName) to be used for access and an ImportPath
that specifies the package to be imported.
</p>

<pre class="ebnf">
<a id="ImportDecl">ImportDecl</a>       = "import" ( <a href="#ImportSpec" class="noline">ImportSpec</a> | "(" { <a href="#ImportSpec" class="noline">ImportSpec</a> ";" } ")" ) .
<a id="ImportSpec">ImportSpec</a>       = [ "." | <a href="#PackageName" class="noline">PackageName</a> ] <a href="#ImportPath" class="noline">ImportPath</a> .
<a id="ImportPath">ImportPath</a>       = <a href="#string_lit" class="noline">string_lit</a> .
</pre>

<p>
The PackageName is used in <a href="#Qualified_identifiers">qualified identifiers</a>
to access exported identifiers of the package within the importing source file.
It is declared in the <a href="#Blocks">file block</a>.
If the PackageName is omitted, it defaults to the identifier specified in the
<a href="#Package_clause">package clause</a> of the imported package.
If an explicit period (<code>.</code>) appears instead of a name, all the
package's exported identifiers declared in that package's
<a href="#Blocks">package block</a> will be declared in the importing source
file's file block and must be accessed without a qualifier.
</p>

<p>
The interpretation of the ImportPath is implementation-dependent but
it is typically a substring of the full file name of the compiled
package and may be relative to a repository of installed packages.
</p>

<p>
Implementation restriction: A compiler may restrict ImportPaths to
non-empty strings using only characters belonging to
<a href="http://www.unicode.org/versions/Unicode6.3.0/">Unicode's</a>
L, M, N, P, and S general categories (the Graphic characters without
spaces) and may also exclude the characters
<code>!"#$%&amp;'()*,:;&lt;=&gt;?[\]^'{|}</code>
and the Unicode replacement character U+FFFD.
</p>

<p>
Assume we have compiled a package containing the package clause
<code>package math</code>, which exports function <code>Sin</code>, and
installed the compiled package in the file identified by
<code>"lib/math"</code>.
This table illustrates how <code>Sin</code> is accessed in files
that import the package after the
various types of import declaration.
</p>

<pre class="grammar">
Import declaration          Local name of Sin

import   "lib/math"         math.Sin
import m "lib/math"         m.Sin
import . "lib/math"         Sin
</pre>

<p>
An import declaration declares a dependency relation between
the importing and imported package.
It is illegal for a package to import itself, directly or indirectly,
or to directly import a package without
referring to any of its exported identifiers. To import a package solely for
its side-effects (initialization), use the <a href="#Blank_identifier">blank</a>
identifier as explicit package name:
</p>

<pre>
import _ "lib/math"
</pre>


<h3 id="An_example_package">An example package</h3>

<p>
Here is a complete Go package that implements a concurrent prime sieve.
</p>

<pre>
package main

import "fmt"

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan&lt;- int) {
	for i := 2; ; i++ {
		ch &lt;- i  // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src &lt;-chan int, dst chan&lt;- int, prime int) {
	for i := range src {  // Loop over values received from 'src'.
		if i%prime != 0 {
			dst &lt;- i  // Send 'i' to channel 'dst'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int)  // Create a new channel.
	go generate(ch)       // Start generate() as a subprocess.
	for {
		prime := &lt;-ch
		fmt.Print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
</pre>

<h2 id="Program_initialization_and_execution">Program initialization and execution</h2>

<h3 id="The_zero_value">The zero value</h3>
<p>
When storage is allocated for a <a href="#Variables">variable</a>,
either through a declaration or a call of <code>new</code>, or when
a new value is created, either through a composite literal or a call
of <code>make</code>,
and no explicit initialization is provided, the variable or value is
given a default value.  Each element of such a variable or value is
set to the <i>zero value</i> for its type: <code>false</code> for booleans,
<code>0</code> for integers, <code>0.0</code> for floats, <code>""</code>
for strings, and <code>nil</code> for pointers, functions, interfaces, slices, channels, and maps.
This initialization is done recursively, so for instance each element of an
array of structs will have its fields zeroed if no value is specified.
</p>
<p>
These two simple declarations are equivalent:
</p>

<pre>
var i int
var i int = 0
</pre>

<p>
After
</p>

<pre>
type T struct { i int; f float64; next *T }
t := new(T)
</pre>

<p>
the following holds:
</p>

<pre>
t.i == 0
t.f == 0.0
t.next == nil
</pre>

<p>
The same would also be true after
</p>

<pre>
var t T
</pre>

<h3 id="Package_initialization">Package initialization</h3>

<p>
Within a package, package-level variables are initialized in
<i>declaration order</i> but after any of the variables
they <i>depend</i> on.
</p>

<p>
More precisely, a package-level variable is considered <i>ready for
initialization</i> if it is not yet initialized and either has
no <a href="#Variable_declarations">initialization expression</a> or
its initialization expression has no dependencies on uninitialized variables.
Initialization proceeds by repeatedly initializing the next package-level
variable that is earliest in declaration order and ready for initialization,
until there are no variables ready for initialization.
</p>

<p>
If any variables are still uninitialized when this
process ends, those variables are part of one or more initialization cycles,
and the program is not valid.
</p>

<p>
The declaration order of variables declared in multiple files is determined
by the order in which the files are presented to the compiler: Variables
declared in the first file are declared before any of the variables declared
in the second file, and so on.
</p>

<p>
Dependency analysis does not rely on the actual values of the
variables, only on lexical <i>references</i> to them in the source,
analyzed transitively. For instance, if a variable <code>x</code>'s
initialization expression refers to a function whose body refers to
variable <code>y</code> then <code>x</code> depends on <code>y</code>.
Specifically:
</p>

<ul>
<li>
A reference to a variable or function is an identifier denoting that
variable or function.
</li>

<li>
A reference to a method <code>m</code> is a
<a href="#Method_values">method value</a> or
<a href="#Method_expressions">method expression</a> of the form
<code>t.m</code>, where the (static) type of <code>t</code> is
not an interface type, and the method <code>m</code> is in the
<a href="#Method_sets">method set</a> of <code>t</code>.
It is immaterial whether the resulting function value
<code>t.m</code> is invoked.
</li>

<li>
A variable, function, or method <code>x</code> depends on a variable
<code>y</code> if <code>x</code>'s initialization expression or body
(for functions and methods) contains a reference to <code>y</code>
or to a function or method that depends on <code>y</code>.
</li>
</ul>

<p>
Dependency analysis is performed per package; only references referring
to variables, functions, and methods declared in the current package
are considered.
</p>

<p>
For example, given the declarations
</p>

<pre>
var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}
</pre>

<p>
the initialization order is <code>d</code>, <code>b</code>, <code>c</code>, <code>a</code>.
</p>

<p>
Variables may also be initialized using functions named <code>init</code>
declared in the package block, with no arguments and no result parameters.
</p>

<pre>
func init() { … }
</pre>

<p>
Multiple such functions may be defined, even within a single
source file. The <code>init</code> identifier is not
<a href="#Declarations_and_scope">declared</a> and thus
<code>init</code> functions cannot be referred to from anywhere
in a program.
</p>

<p>
A package with no imports is initialized by assigning initial values
to all its package-level variables followed by calling all <code>init</code>
functions in the order they appear in the source, possibly in multiple files,
as presented to the compiler.
If a package has imports, the imported packages are initialized
before initializing the package itself. If multiple packages import
a package, the imported package will be initialized only once.
The importing of packages, by construction, guarantees that there
can be no cyclic initialization dependencies.
</p>

<p>
Package initialization&mdash;variable initialization and the invocation of
<code>init</code> functions&mdash;happens in a single goroutine,
sequentially, one package at a time.
An <code>init</code> function may launch other goroutines, which can run
concurrently with the initialization code. However, initialization
always sequences
the <code>init</code> functions: it will not invoke the next one
until the previous one has returned.
</p>

<p>
To ensure reproducible initialization behavior, build systems are encouraged
to present multiple files belonging to the same package in lexical file name
order to a compiler.
</p>


<h3 id="Program_execution">Program execution</h3>
<p>
A complete program is created by linking a single, unimported package
called the <i>main package</i> with all the packages it imports, transitively.
The main package must
have package name <code>main</code> and
declare a function <code>main</code> that takes no
arguments and returns no value.
</p>

<pre>
func main() { … }
</pre>

<p>
Program execution begins by initializing the main package and then
invoking the function <code>main</code>.
When that function invocation returns, the program exits.
It does not wait for other (non-<code>main</code>) goroutines to complete.
</p>

<h2 id="Errors">Errors</h2>

<p>
The predeclared type <code>error</code> is defined as
</p>

<pre>
type error interface {
	Error() string
}
</pre>

<p>
It is the conventional interface for representing an error condition,
with the nil value representing no error.
For instance, a function to read data from a file might be defined:
</p>

<pre>
func Read(f *File, b []byte) (n int, err error)
</pre>

<h2 id="Run_time_panics">Run-time panics</h2>

<p>
Execution errors such as attempting to index an array out
of bounds trigger a <i>run-time panic</i> equivalent to a call of
the built-in function <a href="#Handling_panics"><code>panic</code></a>
with a value of the implementation-defined interface type <code>runtime.Error</code>.
That type satisfies the predeclared interface type
<a href="#Errors"><code>error</code></a>.
The exact error values that
represent distinct run-time error conditions are unspecified.
</p>

<pre>
package runtime

type Error interface {
	error
	// and perhaps other methods
}
</pre>

<h2 id="System_considerations">System considerations</h2>

<h3 id="Package_unsafe">Package <code>unsafe</code></h3>

<p>
The built-in package <code>unsafe</code>, known to the compiler,
provides facilities for low-level programming including operations
that violate the type system. A package using <code>unsafe</code>
must be vetted manually for type safety and may not be portable.
The package provides the following interface:
</p>

<pre class="grammar">
package unsafe

type ArbitraryType int  // shorthand for an arbitrary Go type; it is not a real type
type Pointer *ArbitraryType

func Alignof(variable ArbitraryType) uintptr
func Offsetof(selector ArbitraryType) uintptr
func Sizeof(variable ArbitraryType) uintptr
</pre>

<p>
A <code>Pointer</code> is a <a href="#Pointer_types">pointer type</a> but a <code>Pointer</code>
value may not be <a href="#Address_operators">dereferenced</a>.
Any pointer or value of <a href="#Types">underlying type</a> <code>uintptr</code> can be converted to
a <code>Pointer</code> type and vice versa.
The effect of converting between <code>Pointer</code> and <code>uintptr</code> is implementation-defined.
</p>

<pre>
var f float64
bits = *(*uint64)(unsafe.Pointer(&amp;f))

type ptr unsafe.Pointer
bits = *(*uint64)(ptr(&amp;f))

var p ptr = nil
</pre>

<p>
The functions <code>Alignof</code> and <code>Sizeof</code> take an expression <code>x</code>
of any type and return the alignment or size, respectively, of a hypothetical variable <code>v</code>
as if <code>v</code> was declared via <code>var v = x</code>.
</p>
<p>
The function <code>Offsetof</code> takes a (possibly parenthesized) <a href="#Selectors">selector</a>
<code>s.f</code>, denoting a field <code>f</code> of the struct denoted by <code>s</code>
or <code>*s</code>, and returns the field offset in bytes relative to the struct's address.
If <code>f</code> is an <a href="#Struct_types">embedded field</a>, it must be reachable
without pointer indirections through fields of the struct.
For a struct <code>s</code> with field <code>f</code>:
</p>

<pre>
uintptr(unsafe.Pointer(&amp;s)) + unsafe.Offsetof(s.f) == uintptr(unsafe.Pointer(&amp;s.f))
</pre>

<p>
Computer architectures may require memory addresses to be <i>aligned</i>;
that is, for addresses of a variable to be a multiple of a factor,
the variable's type's <i>alignment</i>.  The function <code>Alignof</code>
takes an expression denoting a variable of any type and returns the
alignment of the (type of the) variable in bytes.  For a variable
<code>x</code>:
</p>

<pre>
uintptr(unsafe.Pointer(&amp;x)) % unsafe.Alignof(x) == 0
</pre>

<p>
Calls to <code>Alignof</code>, <code>Offsetof</code>, and
<code>Sizeof</code> are compile-time constant expressions of type <code>uintptr</code>.
</p>

<h3 id="Size_and_alignment_guarantees">Size and alignment guarantees</h3>

<p>
For the <a href="#Numeric_types">numeric types</a>, the following sizes are guaranteed:
</p>

<pre class="grammar">
type                                 size in bytes

byte, uint8, int8                     1
uint16, int16                         2
uint32, int32, float32                4
uint64, int64, float64, complex64     8
complex128                           16
</pre>

<p>
The following minimal alignment properties are guaranteed:
</p>
<ol>
<li>For a variable <code>x</code> of any type: <code>unsafe.Alignof(x)</code> is at least 1.
</li>

<li>For a variable <code>x</code> of struct type: <code>unsafe.Alignof(x)</code> is the largest of
   all the values <code>unsafe.Alignof(x.f)</code> for each field <code>f</code> of <code>x</code>, but at least 1.
</li>

<li>For a variable <code>x</code> of array type: <code>unsafe.Alignof(x)</code> is the same as
   <code>unsafe.Alignof(x[0])</code>, but at least 1.
</li>
</ol>

<p>
A struct or array type has size zero if it contains no fields (or elements, respectively) that have a size greater than zero. Two distinct zero-size variables may have the same address in memory.
</p>


<div id="footer">
Build version go1.4.2.<br>
Except as <a href="https://developers.google.com/site-policies#restrictions">noted</a>,
the content of this page is licensed under the
Creative Commons Attribution 3.0 License,
and code is licensed under a <a href="/LICENSE">BSD license</a>.<br>
<a href="/doc/tos.html">Terms of Service</a> | 
<a href="http://www.google.com/intl/en/policies/privacy/">Privacy Policy</a>
</div>

</div><!-- .container -->
</div><!-- #page -->

<!-- TODO(adonovan): load these from <head> using "defer" attribute? -->
<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.8.2/jquery.min.js"></script>
<script type="text/javascript" src="/lib/godoc/jquery.treeview.js"></script>
<script type="text/javascript" src="/lib/godoc/jquery.treeview.edit.js"></script>


<script type="text/javascript" src="/lib/godoc/playground.js"></script>

<script type="text/javascript" src="/lib/godoc/godocs.js"></script>

<script type="text/javascript">
(function() {
  var ga = document.createElement("script"); ga.type = "text/javascript"; ga.async = true;
  ga.src = ("https:" == document.location.protocol ? "https://ssl" : "http://www") + ".google-analytics.com/ga.js";
  var s = document.getElementsByTagName("script")[0]; s.parentNode.insertBefore(ga, s);
})();
</script>
</body>
</html>
`
