package sqlast

import "testing"

// 	input := `SELECT USER_ID, NAME, ITEM_ID FROM USER WHERE ID = $1 INNER JOIN USER_ITEM ON USER.USER_ID = USER_ITEM.USER_ID ORDER BY USER_ID, ITEM_ID;`

func TestNextToken(t *testing.T) {
	input := `select user_id, name from user where id = $1 order by user_id, item_id;`
	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{SELECT, "select"},
		{IDENT, "user_id"},
		{COMMA, ","},
		{IDENT, "name"},
		{FROM, "from"},
		{IDENT, "user"},
		{WHERE, "where"},
		{IDENT, "id"},
		{EQUALS, "="},
		{PLACEHOLDER_DOLLER, "$1"},
		{ORDER, "order"},
		{BY, "by"},
		{IDENT, "user_id"},
		{COMMA, ","},
		{IDENT, "item_id"},
		{SEMICOLON, ";"},
	}

	l := New(input)

	for _, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("%s want:%v, got=%v", tok.Literal, tt.expectedType, tok.Type)
		}
	}
}
