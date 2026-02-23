package swagger

type OpenAPI struct {
	OpenAPI string                 `json:"openapi" yaml:"openapi"`
	Servers []struct{ URL string } `json:"servers" yaml:"servers"`
	Paths   map[string]PathItem    `json:"paths" yaml:"paths"`
}

type PathItem struct {
	Get     *Operation `json:"get,omitempty" yaml:"get,omitempty"`
	Post    *Operation `json:"post,omitempty" yaml:"post,omitempty"`
	Put     *Operation `json:"put,omitempty" yaml:"put,omitempty"`
	Delete  *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`
	Patch   *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`
	Head    *Operation `json:"head,omitempty" yaml:"head,omitempty"`
	Options *Operation `json:"options,omitempty" yaml:"options,omitempty"`
	Trace   *Operation `json:"trace,omitempty" yaml:"trace,omitempty"`
}

type Operation struct {
	OperationID string       `json:"operationId" yaml:"operationId"`
	Parameters  []Parameter  `json:"parameters" yaml:"parameters"`
	RequestBody *RequestBody `json:"requestBody" yaml:"requestBody"`
}

type Parameter struct {
	Name   string                 `json:"name" yaml:"name"`
	In     string                 `json:"in" yaml:"in"` // path, query, header
	Schema map[string]interface{} `json:"schema" yaml:"schema"`
}

type RequestBody struct {
	Content map[string]struct {
		Schema  map[string]interface{} `json:"schema" yaml:"schema"`
		Example interface{}            `json:"example" yaml:"example"` // пример запроса из спецификации OpenAPI 3
	} `json:"content" yaml:"content"`
}
