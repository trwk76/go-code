package spec

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func ParseJSON(raw []byte) (*OpenAPI, error) {
	var res OpenAPI

	if err := json.Unmarshal(raw, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func ParseYAML(raw []byte) (*OpenAPI, error) {
	var res OpenAPI

	if err := yaml.Unmarshal(raw, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (o OpenAPI) JSON() []byte {
	res, _ := json.Marshal(o)
	return res
}

func (o OpenAPI) YAML() []byte {
	res, _ := yaml.Marshal(o)
	return res
}

type (
	OpenAPI struct {
		OpenAPI           string               `json:"openapi" yaml:"openapi"`
		Info              Info                 `json:"info" yaml:"info"`
		JSONSchemaDialect string               `json:"jsonSchemaDialect,omitempty" yaml:"jsonSchemaDialect,omitempty"`
		Servers           []Server             `json:"servers,omitempty" yaml:"servers,omitempty"`
		Paths             Paths                `json:"paths" yaml:"paths"`
		Webhooks          NamedPathItemOrRefs  `json:"webhooks,omitempty" yaml:"webhooks,omitempty"`
		Components        *Components          `json:"components,omitempty" yaml:"components,omitempty"`
		Security          SecurityRequirements `json:"security,omitempty" yaml:"security,omitempty"`
		Tags              []Tag                `json:"tags,omitempty" yaml:"tags,omitempty"`
		ExternalDocs      *ExternalDoc         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	}

	Info struct {
		Title          string   `json:"title" yaml:"title"`
		Version        string   `json:"version" yaml:"version"`
		Summary        string   `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description    string   `json:"description,omitempty" yaml:"description,omitempty"`
		TermsOfService string   `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`
		Contact        *Contact `json:"contact,omitempty" yaml:"contact,omitempty"`
		License        *License `json:"license,omitempty" yaml:"license,omitempty"`
	}

	Contact struct {
		Name  string `json:"name,omitempty" yaml:"name,omitempty"`
		URL   string `json:"url,omitempty" yaml:"url,omitempty"`
		Email string `json:"email,omitempty" yaml:"email,omitempty"`
	}

	License struct {
		Name       string `json:"name" yaml:"name"`
		Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
		URL        string `json:"url,omitempty" yaml:"url,omitempty"`
	}

	Server struct {
		URL         string          `json:"url" yaml:"url"`
		Description string          `json:"description,omitempty" yaml:"description,omitempty"`
		Variables   ServerVariables `json:"variables,omitempty" yaml:"variables,omitempty"`
	}

	ServerVariables map[string]ServerVariable

	ServerVariable struct {
		Default     string   `json:"default" yaml:"default"`
		Description string   `json:"description,omitempty" yaml:"description,omitempty"`
		Enum        []string `json:"enum,omitempty" yaml:"enum,omitempty"`
	}

	Components struct {
		Schemas         NamedSchemas              `json:"schemas,omitempty" yaml:"schemas,omitempty"`
		Responses       NamedResponseOrRefs       `json:"responses,omitempty" yaml:"responses,omitempty"`
		Parameters      NamedParameterOrRefs      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		Examples        NamedExampleOrRefs        `json:"examples,omitempty" yaml:"examples,omitempty"`
		RequestBodies   NamedRequestBodyOrRefs    `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`
		Headers         NamedHeaderOrRefs         `json:"headers,omitempty" yaml:"headers,omitempty"`
		SecuritySchemes NamedSecuritySchemeOrRefs `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`
		Links           NamedLinkOrRefs           `json:"links,omitempty" yaml:"links,omitempty"`
		Callbacks       NamedCallbackOrRefs       `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
		PathItems       NamedPathItemOrRefs       `json:"pathItems,omitempty" yaml:"pathItems,omitempty"`
	}

	Tag struct {
		Name         string       `json:"name" yaml:"name"`
		Description  string       `json:"description,omitempty" yaml:"description,omitempty"`
		ExternalDocs *ExternalDoc `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	}

	ExternalDoc struct {
		URL         string `json:"url" yaml:"url"`
		Description string `json:"description,omitempty" yaml:"description,omitempty"`
	}
)

const Version string = "3.1.0"
