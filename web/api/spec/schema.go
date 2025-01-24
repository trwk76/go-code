package spec

type (
	NamedSchemas      map[string]Schema
	NamedSchemaOrRefs map[string]ItemOrRef[Schema]
	SchemaOrRef       = ItemOrRef[Schema]

	Schema struct {
		Title                string              `json:"title,omitempty" yaml:"title,omitempty"`
		Description          string              `json:"description,omitempty" yaml:"description,omitempty"`
		Deprecated           bool                `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
		AllOf                []ItemOrRef[Schema] `json:"allOf,omitempty" yaml:"allOf,omitempty"`
		OneOf                []ItemOrRef[Schema] `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
		AnyOf                []ItemOrRef[Schema] `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
		Type                 Type                `json:"type,omitempty" yaml:"type,omitempty"`
		Format               Format              `json:"format,omitempty" yaml:"format,omitempty"`
		Const                any                 `json:"const,omitempty" yaml:"const,omitempty"`
		Enum                 []any               `json:"enum,omitempty" yaml:"enum,omitempty"`
		Minimum              any                 `json:"minimum,omitempty" yaml:"minimum,omitempty"`
		ExclusiveMinimum     any                 `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
		Maximum              any                 `json:"maximum,omitempty" yaml:"maximum,omitempty"`
		ExclusiveMaximum     any                 `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
		MultipleOf           any                 `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
		MinLength            uint64              `json:"minLength,omitempty" yaml:"minLength,omitempty"`
		MaxLength            uint64              `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
		Pattern              string              `json:"pattern,omitempty" yaml:"pattern,omitempty"`
		Items                *ItemOrRef[Schema]  `json:"items,omitempty" yaml:"items,omitempty"`
		MinItems             uint64              `json:"minItems,omitempty" yaml:"minItems,omitempty"`
		MaxItems             uint64              `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
		UniqueItems          bool                `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
		Properties           NamedSchemaOrRefs   `json:"properties,omitempty" yaml:"properties,omitempty"`
		PatternProperties    NamedSchemaOrRefs   `json:"patternProperties,omitempty" yaml:"patternProperties,omitempty"`
		Required             []string            `json:"required,omitempty" yaml:"required,omitempty"`
		AdditionalProperties *ItemOrRef[Schema]  `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
		Discriminator        *Discriminator      `json:"discriminator,omitempty" yaml:"discriminator,omitempty"`
		XML                  *XML                `json:"xml,omitempty" yaml:"xml,omitempty"`
		ExternalDocs         *ExternalDoc        `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
		Examples             NamedExampleOrRefs  `json:"examples,omitempty" yaml:"examples,omitempty"`
	}

	Type   string
	Format string

	Discriminator struct {
		PropertyName string            `json:"propertyName" yaml:"propertyName"`
		Mapping      map[string]string `json:"mapping,omitempty" yaml:"mapping,omitempty"`
	}

	XML struct {
		Name      string `json:"name,omitempty" yaml:"name,omitempty"`
		Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
		Prefix    string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
		Attribute bool   `json:"attribute,omitempty" yaml:"attribute,omitempty"`
		Wrapped   bool   `json:"wrapped,omitempty" yaml:"wrapped,omitempty"`
	}

	NamedExampleOrRefs map[string]ExampleOrRef
	ExampleOrRef       = ItemOrRef[Example]

	Example struct {
		Summary       string `json:"summary,omitempty" yaml:"summary,omitempty"`
		Description   string `json:"description,omitempty" yaml:"description,omitempty"`
		Value         any    `json:"value,omitempty" yaml:"value,omitempty"`
		ExternalValue string `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
	}
)

const (
	TypeNone    Type = ""
	TypeNull    Type = "null"
	TypeBoolean Type = "boolean"
	TypeInteger Type = "integer"
	TypeNumber  Type = "number"
	TypeString  Type = "string"
	TypeArray   Type = "array"
	TypeObject  Type = "object"
)

const (
	FormatNone   Format = ""
	FormatInt32  Format = "int32"
	FormatInt64  Format = "int64"
	FormatFloat  Format = "float"
	FormatDouble Format = "double"
)
