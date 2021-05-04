package ast

import (
	"reflect"
	"testing"

	"github.com/AntonioAlejandro01/SOL_Lexer"
)

func TestParseProgram(t *testing.T) {
	source := "variable x = 5;"
	lexer := lexer.NewLexer(source)

	parser := NewParser(lexer)

	program := parser.parseProgram()

	if reflect.TypeOf(program).Name() != "Program" {
		t.Errorf("Expected that program was Program struct, but got %s", reflect.TypeOf(program).Name())
	}
}

func TestLetStatement(t *testing.T) {
	source := "variable x = 5;\nvariable y = 10;\nvariable foo = 20;"
	lexer := lexer.NewLexer(source)
	parser := NewParser(lexer)

	program := parser.parseProgram()

	if len(program.statements) != 3 {
		t.Errorf("Expected number of statement was 3, but got %d", len(program.statements))
	}

	for _, statement := range program.statements {
		if statement.tokenLiteral() != "variable" {
			t.Errorf("Expected tokenliteral was variable, but got %s", statement.tokenLiteral())
		}
		if reflect.TypeOf(statement).Name() != "LetStatement" {
			t.Errorf("Expected type of statement was LetStament, but got %s", reflect.TypeOf(statement).Name())
		}
	}
}
