package ts

import (
	"slices"
	"unicode"
)

func IsIdentifier(s string) bool {
	if len(s) < 1 {
		return false
	}

	for idx, r := range s {
		if idx == 0 {
			if !isIDStart(r) {
				return false
			}
		} else {
			if !isIDChar(r) {
				return false
			}
		}
	}

	return true
}

func IsKeyword(s string) bool {
	return slices.Contains(Keywords, Keyword(s))
}

const (
	KwAny         Keyword = "any"
	KwAs          Keyword = "as"
	KwBoolean     Keyword = "boolean"
	KwBreak       Keyword = "break"
	KwCase        Keyword = "case"
	KwCatch       Keyword = "catch"
	KwClass       Keyword = "class"
	KwConst       Keyword = "const"
	KwConstructor Keyword = "constructor"
	KwContinue    Keyword = "continue"
	KwDebugger    Keyword = "debugger"
	KwDeclare     Keyword = "declare"
	KwDefault     Keyword = "default"
	KwDelete      Keyword = "delete"
	KwDo          Keyword = "do"
	KwElse        Keyword = "else"
	KwEnum        Keyword = "enum"
	KwExport      Keyword = "export"
	KwExtends     Keyword = "extends"
	KwFalse       Keyword = "false"
	KwFinally     Keyword = "finally"
	KwFor         Keyword = "for"
	KwFrom        Keyword = "from"
	KwFunction    Keyword = "function"
	KwGet         Keyword = "get"
	KwIf          Keyword = "if"
	KwImplements  Keyword = "implements"
	KwImport      Keyword = "import"
	KwIn          Keyword = "in"
	KwInstanceof  Keyword = "instanceof"
	KwInterface   Keyword = "interface"
	KwLet         Keyword = "let"
	KwModule      Keyword = "module"
	KwNew         Keyword = "new"
	KwNull        Keyword = "null"
	KwNumber      Keyword = "number"
	KwOf          Keyword = "of"
	KwPackage     Keyword = "package"
	KwPrivate     Keyword = "private"
	KwProtected   Keyword = "protected"
	KwPublic      Keyword = "public"
	KwRequire     Keyword = "require"
	KwReturn      Keyword = "return"
	KwSet         Keyword = "set"
	KwStatic      Keyword = "static"
	KwString      Keyword = "string"
	KwSuper       Keyword = "super"
	KwSwitch      Keyword = "switch"
	KwSymbol      Keyword = "symbol"
	KwThis        Keyword = "this"
	KwThrow       Keyword = "throw"
	KwTrue        Keyword = "true"
	KwTry         Keyword = "try"
	KwType        Keyword = "type"
	KwTypeof      Keyword = "typeof"
	KwVar         Keyword = "var"
	KwVoid        Keyword = "void"
	KwWhile       Keyword = "while"
	KwWith        Keyword = "with"
	KwYield       Keyword = "yield"
)

var Keywords []Keyword = []Keyword{
	KwAny,
	KwAs,
	KwBoolean,
	KwBreak,
	KwCase,
	KwCatch,
	KwClass,
	KwConst,
	KwConstructor,
	KwContinue,
	KwDebugger,
	KwDeclare,
	KwDefault,
	KwDelete,
	KwDo,
	KwElse,
	KwEnum,
	KwExport,
	KwExtends,
	KwFalse,
	KwFinally,
	KwFor,
	KwFrom,
	KwFunction,
	KwGet,
	KwIf,
	KwImplements,
	KwImport,
	KwIn,
	KwInstanceof,
	KwInterface,
	KwLet,
	KwModule,
	KwNew,
	KwNull,
	KwNumber,
	KwOf,
	KwPackage,
	KwPrivate,
	KwProtected,
	KwPublic,
	KwRequire,
	KwReturn,
	KwSet,
	KwStatic,
	KwString,
	KwSuper,
	KwSwitch,
	KwSymbol,
	KwThis,
	KwThrow,
	KwTrue,
	KwTry,
	KwType,
	KwTypeof,
	KwVar,
	KwVoid,
	KwWhile,
	KwWith,
	KwYield,
}

func isIDStart(r rune) bool {
	return unicode.IsLetter(r) || r == '_' || r == '$'
}

func isIDChar(r rune) bool {
	return isIDStart(r) || (r >= '0' && r <= '9')
}

type (
	ID      string
	Keyword string
)
