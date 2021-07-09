package parser

import (
	"fmt"

	lexer "github.com/AntonioAlejandro01/SOL_Lexer/src"
)

// NewParser - Create a new Parser
func NewParser(program string) *Parser {
	parser := &Parser{lexer: lexer.NewLexer(program)}

	parser.peekToken = parser.lexer.NextToken()

	return parser
}

// Parser - Parser
type Parser struct {
	lexer        lexer.Lexer
	currentToken *lexer.Token
	peekToken    *lexer.Token
}

func (p *Parser) ParseProgram() Program {
	rootStatement := ServiceStatement{}
	bases := []*BaseStatement{}
	if !p.expectedToken(lexer.SERVICE) {
		panic("The Service have to start with keyword service")
	}
	if !p.expectedToken(lexer.COLON) {
		panic("Bad service structure")
	}
	if !p.expectedToken(lexer.LBRACE) {
		panic("Bad service structure")
	}

	for p.expectedToken(lexer.BASE) {
		bases = append(bases, p.parseBaseStatement())
	}
	rootStatement.Bases = bases

	if p.expectedToken(lexer.SERVICEOPTIONS) {
		rootStatement.Options = p.parseStatementMap(false)
	}

	if p.expectedToken(lexer.BEFORE) {
		rootStatement.Before = p.parseStatementMap(true)
	}

	if p.expectedToken(lexer.ERRORSHANDLERS) {
		rootStatement.ErrorsHandlers = p.parseStatementMap(false)
	}

	if !p.expectedToken(lexer.RBRACE) {
		panic("Service Block maybe not closed")
	}

	if !p.expectedToken(lexer.EOF) {
		panic("Malformed structure for Service")
	}

	return Program(rootStatement)
}

func (p *Parser) parseStatementMap(isBeforeStatement bool) map[string]string {
	if !p.expectedToken(lexer.COLON) {
		panic("Expected colon after Options declaration")
	}
	if !p.expectedToken(lexer.LBRACE) {
		panic("Expected Left curly Braces to start block of options")
	}
	var key string
	var value string
	mapOptions := make(map[string]string)

	for !p.expectedToken(lexer.RBRACE) {

		if p.expectedToken(lexer.EOF) {
			panic("Malformed file")
		}

		if !p.expectedToken(lexer.IDENT) && isBeforeStatement && !p.expectedToken(lexer.ALL) {
			panic("Expected Key for options value")
		}
		key = p.currentToken.Literal
		if !p.expectedToken(lexer.AS) {
			panic("Expected Key word as for associated value to options key")
		}

		if !p.expectedToken(lexer.IDENT) {
			panic(fmt.Sprintf("Expected Value for options key %s", key))
		}

		value = p.currentToken.Literal
		mapOptions[key] = value
		if !p.expectedToken(lexer.COMMA) {
			panic("Expected comma after each key value pair")
		}
	}
	return mapOptions
}

func (p *Parser) parseBaseStatement() *BaseStatement {
	baseStatement := &BaseStatement{}
	isCustomPath := false
	// Starting base path
	if p.expectedToken(lexer.IDENT) {
		isCustomPath = true
		baseStatement.BasePath = p.currentToken.Literal
	} else if p.expectedToken(lexer.COLON) {
		baseStatement.BasePath = "/"
	} else {
		panic("Bad Structure in BASE Statement. Remain colon before BASE statement")
	}

	if isCustomPath && !p.expectedToken(lexer.COLON) {
		panic("Bad Structure in BASE Statement. Remain colon before custom base path")
	}

	if !p.expectedToken(lexer.LBRACE) {
		panic("Bad Structure in BASE statement. Remain Left curly Braces before colon")
	}
	// Starting Endpoints

	for !p.expectedToken(lexer.RBRACE) {
		baseStatement.Endpoints = append(baseStatement.Endpoints, p.parseEndpointStatement())
		if p.expectedToken(lexer.EOF) {
			panic("Malformed BASE")
		}
	}

	return baseStatement
}

func (p *Parser) parseEndpointStatement() *EndpointStatement {
	if !p.expectedToken(lexer.IDENT) {
		panic("Endpoint needed path")
	}
	endpointStatement := &EndpointStatement{
		Endpoint: p.currentToken.Literal,
		Methods:  make(map[string]*MethodStatement),
	}
	if !p.expectedToken(lexer.COLON) {
		panic("EXPECTED COLON")
	}

	if !p.expectedToken(lexer.LBRACE) {
		panic("EXPECTED Left Brace")
	}

	for p.expectedToken(lexer.GET) ||
		p.expectedToken(lexer.POST) ||
		p.expectedToken(lexer.PUT) ||
		p.expectedToken(lexer.PATCH) ||
		p.expectedToken(lexer.OPTIONS) ||
		p.expectedToken(lexer.HEAD) ||
		p.expectedToken(lexer.DELETE) {

		method := p.currentToken.Literal
		endpointStatement.Methods[method] = p.parseMethodStatement()

	}

	if !p.expectedToken(lexer.RBRACE) {
		panic("Endpoint declaration not closed")
	}

	return endpointStatement
}

func (p *Parser) parseMethodStatement() *MethodStatement {
	method := &MethodStatement{}
	if !p.expectedToken(lexer.AS) {
		panic("Expected Keyword as after Method Declaration")
	}

	if !p.expectedToken(lexer.IDENT) {
		panic("Expected Handler to be used with this method")
	}

	method.Handler = p.currentToken.Literal

	if !p.expectedToken(lexer.COLON) {
		panic("Expected Colon after handlerDeclaration")
	}

	if !p.expectedToken(lexer.LBRACE) {
		panic("Expected Left curly braces to start block method declaration")
	}

	if !p.expectedToken(lexer.BODY) {
		panic("Missing Body declaration")
	}

	method.Body = p.parseBodyStatement()

	if !p.expectedToken(lexer.PARAMS) {
		panic("Missing Params declaration")
	}

	method.Params = p.parseStatementMap(false)

	if !p.expectedToken(lexer.HEADERS) {
		panic("Missing Headers Declaration")
	}

	method.Headers = p.parseStatementMap(false)

	if !p.expectedToken(lexer.RBRACE) {
		panic("Method declaration not closed")
	}

	return method
}

func (p *Parser) parseBodyStatement() *BodyStatement {
	if !p.expectedToken(lexer.COLON) {
		panic("After Body need colon")
	}
	if !p.expectedToken(lexer.LBRACE) {
		panic("Block start with left brace")
	}
	if !p.expectedToken(lexer.RBRACE) {
		panic("Block start with rigth brace")
	}
	return &BodyStatement{}
}

func (p *Parser) advanceToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) expectedToken(tt lexer.TokenType) bool {
	if tt == p.peekToken.TokenType {
		p.advanceToken()
		return true
	}
	return false
}
