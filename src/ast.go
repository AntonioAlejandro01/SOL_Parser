package ast

import (
	"fmt"
	"strings"

	"github.com/AntonioAlejandro01/SOL_Lexer"
)

// AstNode -  node ast
type astNode interface {
	tokenLiteral() string
	ToString() string
	myType() string
}

// Statement statement
type Statement lexer.Token

func (s Statement) myType() string {
	return "statement"
}

func (s Statement) tokenLiteral() string {
	return lexer.Token(s).Literal
}

// ToString - ToString
func (s Statement) ToString() string {
	return lexer.Token(s).ToString()
}

// Expression expression
type Expression struct {
	Token lexer.Token
}

func (e Expression) tokenLiteral() string {
	return e.Token.Literal
}

// ToString - toString
func (e Expression) ToString() string {
	return e.Token.ToString()
}

// Program node
type Program struct {
	statements []Statement
}

func (p Program) tokenLiteral() string {
	if len(p.statements) > 0 {
		return lexer.Token(p.statements[0]).Literal
	}
	return ""
}

func (p Program) toString() string {
	builder := strings.Builder{}

	for _, statement := range p.statements {
		builder.WriteString(statement.ToString())
	}
	return builder.String()
}

// LetStatement - Let statement
type LetStatement struct {
	statement Statement
	name      Identifier
	value     Expression
}

func (l LetStatement) myType() string {
	return "letstatement"
}
func (l LetStatement) tokenLiteral() string {
	return fmt.Sprintf("%+v", l)
}

// ToString - toString
func (l LetStatement) ToString() string {
	return fmt.Sprintf("%s %s = %s;", l.statement.tokenLiteral(), l.name.ToString(), l.value.ToString())
}

// Identifier - identifier
type Identifier struct {
	Statement
	value string
}

// ToString - toString
func (i Identifier) ToString() string {
	return i.value
}
