package parser

import (
	lexer "github.com/AntonioAlejandro01/SOL_Lexer/src"
)

// Parser - Parser
type Parser struct {
	lexer        lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
}

func (p *Parser) parseProgram() Program {
	rootStatement := ServiceStatement{}

	rootStatement.bases = p.parseBases()

	return Program(rootStatement)
}

func (p *Parser) parseBases() []*BaseStatement {
	return []*BaseStatement{}
}

// NewParser - Create a new Parser
func NewParser(program string) *Parser {
	return &Parser{
		lexer: lexer.NewLexer(program),
	}
}
