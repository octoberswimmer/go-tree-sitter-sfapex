package soql_test

import (
	"context"
	"testing"

	"github.com/octoberswimmer/go-tree-sitter-sfapex/soql"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/stretchr/testify/assert"
)

var soqlQuery = `
SELECT
	Id,
	Name,
	(SELECT
		Id,
		Subject
	FROM
		Tasks
	)
FROM
	Account
`

var output = `(source_file (soql_query_body (select_clause (field_identifier (identifier)) (field_identifier (identifier)) (subquery (soql_query_body (select_clause (field_identifier (identifier)) (field_identifier (identifier))) (from_clause (storage_identifier (identifier)))))) (from_clause (storage_identifier (identifier)))))`

func TestGrammar(t *testing.T) {
	n, err := sitter.ParseCtx(context.Background(), []byte(soqlQuery), soql.GetLanguage())
	assert.Nil(t, err)
	assert.Equal(t, output, n.String())
}
