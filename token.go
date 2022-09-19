package sqlast

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	EQUALS   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"
	PERCENT  = "%"
	BANG     = "!"

	PLACEHOLDER_DOLLER = "$"
	PLACEHOLDER_AT     = "@"

	LT = "<"
	GT = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	SELECT  = "SELECT"
	FROM    = "FROM"
	INNER   = "INNER"
	OUTER   = "OUTER"
	CROSS   = "CROSS"
	JOIN    = "JOIN"
	ON      = "ON"
	WHERE   = "WHERE"
	AND     = "AND"
	OR      = "OR"
	BETWEEN = "BETWEEN"
	IN      = "IN"
	GROUP   = "GROUP"
	BY      = "BY"
	HAVING  = "HAVING"
	ORDER   = "ORDER"
	LIMIT   = "LIMIT"
	WITH    = "WITH"
)

// keywords はSQLのキーワードを表現します
// PostgreSQL 8.3.7 https://www.postgresql.jp/document/8.3/html/sql-keywords-appendix.html
var keywords = map[string]TokenType{
	"select":  SELECT,
	"from":    FROM,
	"inner":   INNER,
	"outer":   OUTER,
	"cross":   CROSS,
	"join":    JOIN,
	"on":      ON,
	"where":   WHERE,
	"and":     AND,
	"or":      OR,
	"between": BETWEEN,
	"in":      IN,
	"group":   GROUP,
	"by":      BY,
	"having":  HAVING,
	"order":   ORDER,
	"limit":   LIMIT,
	"with":    WITH,
}

func LookupIdent(ident string) TokenType {
	key := strings.ToLower(ident)
	if tok, ok := keywords[key]; ok {
		return tok
	}
	return IDENT
}

var placeHolderPrefixKeys = map[string]TokenType{
	"$": PLACEHOLDER_DOLLER,
	"@": PLACEHOLDER_AT,
}

func LookupPlaceHolder(placeholder string) TokenType {
	if tok, ok := placeHolderPrefixKeys[placeholder]; ok {
		return tok
	}
	return IDENT
}
