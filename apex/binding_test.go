package apex_test

import (
	"context"
	"testing"

	"github.com/octoberswimmer/go-tree-sitter-sfapex/apex"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/stretchr/testify/assert"
)

var apexCode = `
public class MyClass {
	public void myMethod() {
		System.debug(a);
	 }
}
`

var output = `(parser_output (class_declaration (modifiers (modifier)) name: (identifier) body: (class_body (method_declaration (modifiers (modifier)) type: (void_type) name: (identifier) parameters: (formal_parameters) body: (block (expression_statement (method_invocation object: (identifier) name: (identifier) arguments: (argument_list (identifier)))))))))`

func TestGrammar(t *testing.T) {
	n, err := sitter.ParseCtx(context.Background(), []byte(apexCode), apex.GetLanguage())
	assert.Nil(t, err)
	assert.Equal(t, output, n.String())
}
