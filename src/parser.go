package parser

import (
	"errors"

	lexer "github.com/AntonioAlejandro01/SOL_Lexer"
)

// Parser - Parser
type Parser struct {
	lexer        lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
}

func (p *Parser) parseProgram() Program {
	program := &Program{
		statements: []Statement{},
	}

	for p.currentToken.TokenType != lexer.EOF {
		statement, err := p.parseStatement()
		if err == nil {
			program.statements = append(program.statements, statement)
		}
	}

	return *program
}

func (p *Parser) advanceToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expectedToken(tokentype lexer.TokenType) bool {
	if p.peekToken.TokenType == tokentype {
		p.advanceToken()
		return true
	}
	return false
}

func (p *Parser) parseStatement() (astNode, error) {
	// TODO: fix types
	if p.currentToken.TokenType == lexer.LET {
		statement, err := p.parseLetStatement()
		if err == nil {
			return statement, nil
		}
	}
	return Statement{}, errors.New("Statement no indetified")
}

func (p *Parser) parseLetStatement() (LetStatement, error) {
	letStatement := LetStatement{statement: Statement(p.currentToken)}

	if !p.expectedToken(lexer.IDENT) {
		return LetStatement{}, errors.New("")

	}

	letStatement.name = Identifier{Statement: Statement(p.currentToken), value: p.currentToken.Literal}

	if !p.expectedToken(lexer.ASSING) {
		return LetStatement{}, errors.New("")
	}

	for p.currentToken.TokenType != lexer.SEMICOLON {
		p.advanceToken()
	}

	return letStatement, nil

}

// NewParser - Create a new Parser
func NewParser(lexer lexer.Lexer) Parser {
	return Parser{
		lexer: lexer,
	}
}
