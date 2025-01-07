package models

type Components struct {
	Schemas         map[string]interface{} `yaml:"schemas,omitempty"`
	Responses       map[string]interface{} `yaml:"responses,omitempty"`
	SecuritySchemes map[string]interface{} `yaml:"securitySchemes,omitempty"`
}

type OpenAPI struct {
	OpenAPI      string                   `yaml:"openapi"`
	Info         Info                     `yaml:"info,omitempty"`
	Security     []map[string]interface{} `yaml:"security,omitempty"`
	Tags         []map[string]interface{} `yaml:"tags,omitempty"`
	Paths        map[string]interface{}   `yaml:"paths,omitempty"`
	ExternalDocs map[string]interface{}   `yaml:"externalDocs,omitempty"`
	Servers      []Server                 `yaml:"servers,omitempty"`
	Components   Components               `yaml:"components,omitempty"`
	Definitions  map[string]interface{}   `yaml:"definitions,omitempty"`
}
