package parser

import (
	lexer "github.com/AntonioAlejandro01/SOL_Lexer/src"
)

// Parser - Parser
type Parser struct {
	lexer        lexer.Lexer
	currentToken *lexer.Token
	peekToken    *lexer.Token
}

func (p *Parser) ParseProgram() Program {
	rootStatement := ServiceStatement{}
	bases := []*BaseStatement{}
	for p.expectedToken(lexer.BASE) {
		bases = append(bases, p.parseBaseStatement())
	}
	rootStatement.bases = bases

	if p.expectedToken(lexer.SERVICEOPTIONS) {
		rootStatement.options = p.parseOptionsStatement()
	}

	if p.expectedToken(lexer.BEFORE) {
		rootStatement.before = p.parseBeforeStatement()
	}

	if p.expectedToken(lexer.ERRORSHANDLERS) {
		rootStatement.errorsHandlers = p.parseErrorsHandlersStatement()
	}

	if !p.expectedToken(lexer.EOF) {
		panic("Malformed structure for Service")
	}

	return Program(rootStatement)
}

func (p *Parser) parseOptionsStatement() *OptionsStatement {
	//TODO: Method
	return &OptionsStatement{}
}

func (p *Parser) parseBeforeStatement() *BeforeStatement {
	//TODO: Method
	return &BeforeStatement{}
}

func (p *Parser) parseErrorsHandlersStatement() *ErrorsHandlersStatement {
	//TODO: Method
	return &ErrorsHandlersStatement{}
}

func (p *Parser) parseBaseStatement() *BaseStatement {
	baseStatement := &BaseStatement{}
	isCustomPath := false
	// Starting base path
	if p.expectedToken(lexer.IDENT) {
		isCustomPath = true
		baseStatement.basePath = p.currentToken.Literal
	} else if p.expectedToken(lexer.COLON) {
		baseStatement.basePath = "/"
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
		baseStatement.endpoints = append(baseStatement.endpoints, p.parseEndpointStatement())
		if p.expectedToken(lexer.EOF) {
			panic("Malformed BASE")
		}
	}

	return baseStatement
}

func (p *Parser) parseEndpointStatement() *EndpointStatement {
	endpointStatement := &EndpointStatement{
		endpoint: p.currentToken.Literal,
		methods:  make(map[string]*MethodStatement),
	}

	for p.expectedToken(lexer.GET) ||
		p.expectedToken(lexer.POST) ||
		p.expectedToken(lexer.PUT) ||
		p.expectedToken(lexer.PATCH) ||
		p.expectedToken(lexer.OPTIONS) ||
		p.expectedToken(lexer.HEAD) ||
		p.expectedToken(lexer.DELETE) {

		endpointStatement.methods[p.currentToken.Literal] = p.parseMethodStatement()

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

	method.handler = p.currentToken.Literal

	if !p.expectedToken(lexer.COLON) {
		panic("Expected Colon after handlerDeclaration")
	}

	if !p.expectedToken(lexer.LBRACE) {
		panic("Expected Left curly braces to start block method declaration")
	}

	if !p.expectedToken(lexer.BODY) {
		panic("Missing Body declaration")
	}

	method.body = p.parseBodyStatement()

	if !p.expectedToken(lexer.PARAMS) {
		panic("Missing Params declaration")
	}

	method.params = p.parseParamsStatement()

	if !p.expectedToken(lexer.HEADERS) {
		panic("Missing Headers Declaration")
	}

	method.headers = p.parseHeadersStatement()

	if !p.expectedToken(lexer.RBRACE) {
		panic("Method declaration not closed")
	}

	return method
}

func (p *Parser) parseHeadersStatement() *HeaderStatement {
	//TODO: Method
	return &HeaderStatement{}
}

func (p *Parser) parseParamsStatement() *ParamsStatement {
	//TODO: Method
	return &ParamsStatement{}
}

func (p *Parser) parseBodyStatement() *BodyStatement {
	//TODO: Method
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

// NewParser - Create a new Parser
func NewParser(program string) *Parser {
	return &Parser{
		lexer: lexer.NewLexer(program),
	}
}
