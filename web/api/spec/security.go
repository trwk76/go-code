package spec

type (
	SecurityRequirements []SecurityRequirement
	SecurityRequirement  map[string][]string

	NamedSecuritySchemeOrRefs map[string]SecuritySchemeOrRef
	SecuritySchemeOrRef       = ItemOrRef[SecurityScheme]

	SecurityScheme struct {
		Type             SecurityType `json:"type" yaml:"type"`
		Description      string       `json:"description,omitempty" yaml:"description,omitempty"`
		Name             string       `json:"name,omitempty" yaml:"name,omitempty"`
		In               SecurityIn   `json:"in,omitempty" yaml:"in,omitempty"`
		Scheme           string       `json:"scheme,omitempty" yaml:"scheme,omitempty"`
		BearerFormat     string       `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`
		Flows            *OAuthFlows  `json:"flows,omitempty" yaml:"flows,omitempty"`
		OpenIDConnectURL string       `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`
	}

	SecurityType string
	SecurityIn   string

	OAuthFlows struct {
		Implicit          *OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`
		Password          *OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`
		ClientCredentials *OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
		AuthorizationCode *OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`
	}

	OAuthFlow struct {
		AuthorizationURL string            `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`
		TokenURL         string            `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`
		RefreshURL       string            `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`
		Scopes           map[string]string `json:"scopes" yaml:"scopes"`
	}
)

const (
	SecurityAPIKey        SecurityType = "apiKey"
	SecurityHTTP          SecurityType = "http"
	SecurityMutualTLS     SecurityType = "mutualTLS"
	SecurityOAuth2        SecurityType = "oauth2"
	SecurityOpenIDConnect SecurityType = "openIdConnect"
)

const (
	SecurityInCookie SecurityIn = "cookie"
	SecurityInHeader SecurityIn = "header"
	SecurityInQuery  SecurityIn = "query"
)

var AnonymousSecurityRequirements SecurityRequirements = SecurityRequirements{{}}
