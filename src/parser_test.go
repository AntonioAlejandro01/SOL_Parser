package parser

import (
	"testing"
)

func TestParseProgramOk(t *testing.T) {

	source := "service : {base random:{text: { GET as handlerText : {body : {} params: {size as int,} headers: {Authorization as token,Cookie as cookie,}}}} options:{} before:{* as middlewareAuth,text as middlewareAuthText,} errorsHandlers : {Error as  handlerError, CustomError as handleError,}}"

	errorsHandlersExpected := make(map[string]string)
	errorsHandlersExpected["Error"] = "handleError"
	errorsHandlersExpected["CustomError"] = "handlerError"

	beforeExpected := make(map[string]string)
	beforeExpected["*"] = "middlewareAuth"
	beforeExpected["text"] = "middlewareAuthText"

	paramsExpected := make(map[string]string)
	paramsExpected["size"] = "int"

	headersExpected := make(map[string]string)
	headersExpected["Authorization"] = "token"
	headersExpected["Cookie"] = "cookie"

	methodExpected := make(map[string]*MethodStatement)
	methodExpected["GET"] = &MethodStatement{
		Handler: "handlerText",
		Body:    &BodyStatement{},
		Params:  paramsExpected,
		Headers: headersExpected,
	}
	baseExpected := &BaseStatement{
		BasePath: "random",
		Endpoints: []*EndpointStatement{
			{
				Endpoint: "text",
				Methods:  methodExpected,
			},
		},
	}

	expectedProgram := Program(ServiceStatement{
		Bases:          []*BaseStatement{baseExpected},
		Options:        make(map[string]string),
		Before:         beforeExpected,
		ErrorsHandlers: errorsHandlersExpected,
	})

	parser := NewParser(source)

	program := parser.ParseProgram()

	if len(expectedProgram.Bases) != len(program.Bases) {
		t.Errorf("Expected bases are %d but got %d", len(expectedProgram.Bases), len(program.Bases))
	}

	if len(expectedProgram.Before) != len(program.Before) {
		t.Errorf("Expected before are %d but got %d", len(expectedProgram.Bases), len(program.Bases))
	}

}
