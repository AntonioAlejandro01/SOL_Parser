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
		handler: "handlerText",
		body:    &BodyStatement{},
		params:  paramsExpected,
		headers: headersExpected,
	}
	baseExpected := &BaseStatement{
		basePath: "random",
		endpoints: []*EndpointStatement{
			{
				endpoint: "text",
				methods:  methodExpected,
			},
		},
	}

	expectedProgram := Program(ServiceStatement{
		bases:          []*BaseStatement{baseExpected},
		options:        make(map[string]string),
		before:         beforeExpected,
		errorsHandlers: errorsHandlersExpected,
	})

	parser := NewParser(source)

	program := parser.ParseProgram()

	if len(expectedProgram.bases) != len(program.bases) {
		t.Errorf("Expected bases are %d but got %d", len(expectedProgram.bases), len(program.bases))
	}

	if len(expectedProgram.before) != len(program.before) {
		t.Errorf("Expected before are %d but got %d", len(expectedProgram.bases), len(program.bases))
	}

}
