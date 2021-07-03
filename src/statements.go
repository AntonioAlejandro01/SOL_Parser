package parser

type ServiceStatement struct {
	bases []*BaseStatement
}

type BaseStatement struct {
	basePath  string
	endpoints []EndpointStatement
}

type EndpointStatement struct {
	endpoint string
	methods  map[string]MethodStatement
}

type MethodStatement struct {
	handler string
	body    BodyStatement
	params  ParamsStatement
	headers HeaderStatement
}

type BodyStatement struct {
}

type ParamsStatement struct {
	name string
	typ  string
}

type HeaderStatement struct {
	name  string
	alias string
}
