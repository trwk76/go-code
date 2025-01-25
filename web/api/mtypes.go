package api

import "github.com/trwk76/gocode/web/api/spec"

const (
	MediaTypeJSON string = "application/json"
)

type (
	MediaTypes map[string]MediaType

	MediaType struct {
		Schema Schema
	}
)

func (m MediaTypes) spec() spec.MediaTypes {
	res := make(spec.MediaTypes)

	for key, item := range m {
		s := item.spec()
		res[key] = &s
	}

	return res
}

func (m MediaType) spec() spec.MediaType {
	res := spec.MediaType{}

	if m.Schema != nil {
		s := m.Schema.spec()
		res.Schema = &s
	}

	return res
}
