package gogu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_Substr(t *testing.T) {
	assert := assert.New(t)

	str1 := Substr("abcdef", 0, 0)
	str2 := Substr("abcdef", -1, 0)
	str3 := Substr("abcdef", 7, 7)
	str4 := Substr("abcdef", 0, 20)
	str5 := Substr("abcdef", 5, 10)
	str6 := Substr("abcdef", 0, -1)
	str7 := Substr("abcdef", 2, -1)
	str8 := Substr("abcdef", 4, -4)
	str9 := Substr("abcdef", -3, -1)
	str10 := Substr("abcdef", 1, 3)
	str11 := Substr("abcdef", 0, 4)
	str12 := Substr("abcdef", 0, 8)
	str13 := Substr("abcdef", -1, 1)
	str14 := Substr("abcdef", -2, -8)
	str15 := Substr("abcdef", -4, 1)
	str16 := Substr("abcdef", -10, -10)

	assert.Empty(str1)
	assert.Empty(str2)
	assert.Empty(str3)
	assert.Equal("abcdef", str4)
	assert.Equal("f", str5)
	assert.Equal("abcde", str6)
	assert.Equal("cde", str7)
	assert.Empty(str8)
	assert.Equal("de", str9)
	assert.Equal("bcd", str10)
	assert.Equal("abcd", str11)
	assert.Equal("abcdef", str12)
	assert.Equal("f", str13)
	assert.Equal("", str14)
	assert.Equal("c", str15)
	assert.Equal("", str16)
}

func TestString_Capitalize(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Abc", Capitalize("Abc"))
	assert.Equal("Abc", Capitalize("abc"))
	assert.Equal("Abc", Capitalize("abC"))
	assert.Equal("Abc", Capitalize("aBC"))
	assert.Equal("Abc", Capitalize("ABC"))
	assert.Equal("Abø", Capitalize("aBø"))
	assert.Equal("", Capitalize(""))
}

func TestString_CamelCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("fooBar", CamelCase("Foo Bar"))
	assert.Equal("fooBar", CamelCase("--foo-Bar--"))
	assert.Equal("fooBar", CamelCase("__foo-_Bar__"))
	assert.Equal("fooBar", CamelCase("__FOO BAR__"))
	assert.Equal("fooBar", CamelCase(" FOO BAR "))
	assert.Equal("fooBar", CamelCase("&FOO&baR "))
	assert.Equal("fooBar", CamelCase("&&foo&&bar__"))
}

func TestString_SnakeCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("foo", SnakeCase("Foo"))
	assert.Equal("foo_bar", SnakeCase("Foo Bar"))
	assert.Equal("foo_bar", SnakeCase("Foo_Bar"))
	assert.Equal("foo_bar", SnakeCase("fooBar"))
	assert.Equal("foo_bar", SnakeCase("FooBar"))
	assert.Equal("foo_bar", SnakeCase("foo__Bar"))
	assert.Equal("foo_bar", SnakeCase("foo_&Bar"))
	assert.Equal("foo_bar", SnakeCase("&Foo_&Bar"))
	assert.Equal("foo_bar", SnakeCase(" Foo_Bar "))
}

func TestString_SplitAtIndex(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"", "abcdef"}, SplitAtIndex("abcdef", -1))
	assert.Equal([]string{"a", "bcdef"}, SplitAtIndex("abcdef", 0))
	assert.Equal([]string{"ab", "cdef"}, SplitAtIndex("abcdef", 1))
	assert.Equal([]string{"abc", "def"}, SplitAtIndex("abcdef", 2))
	assert.Equal([]string{"abcdef", ""}, SplitAtIndex("abcdef", 5))
	assert.Equal([]string{"abcdef", ""}, SplitAtIndex("abcdef", 6))
}

func TestString_Pad(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", PadLeft("abc", 2, "."))
	assert.Equal("abc", PadLeft("abc", 3, "."))
	assert.Equal("_abc", PadLeft("abc", 4, "_"))
	assert.Equal(".....abc", PadLeft("abc", 8, "..."))
	assert.Equal("...abc", PadLeft("abc", 6, "........"))
	assert.Equal("_-_abc", PadLeft("abc", 6, "_-"))
	assert.Equal("_-|_abc", PadLeft("abc", 7, "_-|"))

	assert.Equal("abc", PadRight("abc", 2, "."))
	assert.Equal("abc", PadRight("abc", 3, "."))
	assert.Equal("abc_", PadRight("abc", 4, "_"))
	assert.Equal("abc.....", PadRight("abc", 8, "..."))
	assert.Equal("abc...", PadRight("abc", 6, "........"))
	assert.Equal("abc_-_", PadRight("abc", 6, "_-"))
	assert.Equal("abc_-|_", PadRight("abc", 7, "_-|"))
}

func TestString_ReverseStr(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("fedcba", ReverseStr("abcdef"))
	assert.Equal("FEDCBA", ReverseStr("ABCDEF"))
	assert.Equal("654321", ReverseStr("123456"))
}

func TestString_Wrap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", Wrap("abc", ""))
	assert.Equal("'abc'", Wrap("abc", "'"))
	assert.Equal("*abc*", Wrap("abc", "*"))
	assert.Equal(`\abc\`, Wrap("abc", `\`))
	assert.Equal(`|abc|`, Wrap("abc", `|`))
}

func TestString_Unwrap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", Unwrap("'abc'", "'"))
	assert.Equal("abc", Unwrap("*abc*", "*"))
	assert.Equal("a*bc", Unwrap("*a*bc*", "*"))
	assert.Equal("abc", Unwrap("''abc''", "''"))
	assert.Equal("abc", Unwrap("\"abc\"", "\""))
	assert.Equal("'", Unwrap("'''", "'"))
}

func TestString_WrapAllRune(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", WrapAllRune("abc", ""))
	assert.Equal("'a''b''c'", WrapAllRune("abc", "'"))
	assert.Equal("*a**b**c*", WrapAllRune("abc", "*"))
	assert.Equal("-a--b--c-", WrapAllRune("abc", "-"))
}
