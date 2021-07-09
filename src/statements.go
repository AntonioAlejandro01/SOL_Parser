package parser

type Program ServiceStatement

type ServiceStatement struct {
	Bases          []*BaseStatement
	Options        map[string]string
	Before         map[string]string
	ErrorsHandlers map[string]string
}

type BaseStatement struct {
	BasePath  string
	Endpoints []*EndpointStatement
}

type EndpointStatement struct {
	Endpoint string
	Methods  map[string]*MethodStatement
}

type MethodStatement struct {
	Handler string
	Body    *BodyStatement
	Params  map[string]string
	Headers map[string]string
}

type BodyStatement struct {
}
