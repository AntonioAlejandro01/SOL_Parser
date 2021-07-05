package parser

type Program ServiceStatement

type ServiceStatement struct {
	bases          []*BaseStatement
	options        map[string]string
	before         map[string]string
	errorsHandlers map[string]string
}

type BaseStatement struct {
	basePath  string
	endpoints []*EndpointStatement
}

type EndpointStatement struct {
	endpoint string
	methods  map[string]*MethodStatement
}

type MethodStatement struct {
	handler string
	body    *BodyStatement
	params  map[string]string
	headers map[string]string
}

type BodyStatement struct {
}
