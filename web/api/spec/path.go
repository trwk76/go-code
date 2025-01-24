package spec

type (
	Paths               map[string]*PathItem
	NamedPathItemOrRefs map[string]PathItemOfRef
	PathItemOfRef       = ItemOrRef[PathItem]

	PathItem struct {
		Summary     string          `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description string          `json:"description,omitempty" yaml:"description,omitempty"`
		GET         *Operation      `json:"get,omitempty" yaml:"get,omitempty"`
		PUT         *Operation      `json:"put,omitempty" yaml:"put,omitempty"`
		POST        *Operation      `json:"post,omitempty" yaml:"post,omitempty"`
		DELETE      *Operation      `json:"delete,omitempty" yaml:"delete,omitempty"`
		OPTIONS     *Operation      `json:"options,omitempty" yaml:"options,omitempty"`
		HEAD        *Operation      `json:"head,omitempty" yaml:"head,omitempty"`
		PATCH       *Operation      `json:"patch,omitempty" yaml:"patch,omitempty"`
		TRACE       *Operation      `json:"trace,omitempty" yaml:"trace,omitempty"`
		Servers     []Server        `json:"servers,omitempty" yaml:"servers,omitempty"`
		Parameters  ParameterOrRefs `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	}

	Operation struct {
		Tags         []string             `json:"tags,omitempty" yaml:"tags,omitempty"`
		Summary      string               `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description  string               `json:"description,omitempty" yaml:"description,omitempty"`
		ExternalDocs *ExternalDoc         `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
		OperationID  string               `json:"operationId,omitempty" yaml:"operationId,omitempty"`
		Parameters   ParameterOrRefs      `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		RequestBody  *RequestBodyOrRef    `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
		Responses    Responses            `json:"responses" yaml:"responses"`
		Callbacks    NamedCallbackOrRefs  `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
		Deprecated   bool                 `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		Security     SecurityRequirements `json:"security,omitempty" yaml:"security,omitempty"`
		Servers      []Server             `json:"servers,omitempty" yaml:"servers,omitempty"`
	}

	NamedParameterOrRefs map[string]ParameterOrRef
	ParameterOrRefs      []ParameterOrRef
	ParameterOrRef       = ItemOrRef[Parameter]

	Parameter struct {
		Name            string             `json:"name" yaml:"name"`
		In              ParameterIn        `json:"in" yaml:"in"`
		Description     string             `json:"description,omitempty" yaml:"description,omitempty"`
		Required        bool               `json:"required,omitempty" yaml:"required,omitempty"`
		Deprecated      bool               `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		AllowEmptyValue bool               `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
		Schema          *SchemaOrRef       `json:"schema,omitempty" yaml:"schema,omitempty"`
		Examples        NamedExampleOrRefs `json:"examples,omitempty" yaml:"examples,omitempty"`
	}

	ParameterIn string

	NamedHeaderOrRefs map[string]HeaderOrRef
	HeaderOrRef       = ItemOrRef[Header]

	Header struct {
		Description     string             `json:"description,omitempty" yaml:"description,omitempty"`
		Required        bool               `json:"required,omitempty" yaml:"required,omitempty"`
		Deprecated      bool               `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		AllowEmptyValue bool               `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
		Schema          *Schema            `json:"schema,omitempty" yaml:"schema,omitempty"`
		Examples        NamedExampleOrRefs `json:"examples,omitempty" yaml:"examples,omitempty"`
	}

	MediaTypes map[string]*MediaType

	MediaType struct {
		Schema   *SchemaOrRef       `json:"schema,omitempty" yaml:"schema,omitempty"`
		Examples NamedExampleOrRefs `json:"examples,omitempty" yaml:"examples,omitempty"`
	}

	NamedRequestBodyOrRefs map[string]RequestBodyOrRef
	RequestBodyOrRef       = ItemOrRef[RequestBody]

	RequestBody struct {
		Description string     `json:"description,omitempty" yaml:"description,omitempty"`
		Content     MediaTypes `json:"content,omitempty" yaml:"content,omitempty"`
		Required    bool       `json:"required,omitempty" yaml:"required,omitempty"`
	}

	Responses           map[string]ResponseOrRef
	NamedResponseOrRefs map[string]ResponseOrRef
	ResponseOrRef       = ItemOrRef[Response]

	Response struct {
		Description string            `json:"description" yaml:"description"`
		Headers     NamedHeaderOrRefs `json:"headers,omitempty" yaml:"headers,omitempty"`
		Content     MediaTypes        `json:"content,omitempty" yaml:"content,omitempty"`
		Links       NamedLinkOrRefs   `json:"links,omitempty" yaml:"links,omitempty"`
	}

	NamedCallbackOrRefs map[string]CallbackOrRef
	CallbackOrRef       = ItemOrRef[Callback]
	Callback            map[string]PathItemOfRef

	NamedLinkOrRefs map[string]LinkOrRef
	LinkOrRef       = ItemOrRef[Link]

	Link struct {
		OperationRef string         `json:"operationRef,omitempty" yaml:"operationRef,omitempty"`
		OperationID  string         `json:"operationId,omitempty" yaml:"operationId,omitempty"`
		Parameters   map[string]any `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		RequestBody  any            `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
		Description  string         `json:"description,omitempty" yaml:"description,omitempty"`
		Servers      *Server        `json:"server,omitempty" yaml:"server,omitempty"`
	}
)

const (
	ParameterCookie ParameterIn = "cookie"
	ParameterHeader ParameterIn = "header"
	ParameterPath   ParameterIn = "path"
	ParameterQuery  ParameterIn = "query"
)
