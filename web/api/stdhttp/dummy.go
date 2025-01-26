package stdhttp

import "github.com/trwk76/go-code/web/api"

func (gen *Generator) Parameter(key string, impl *api.ParameterImpl)     {}
func (gen *Generator) RequestBody(key string, impl *api.RequestBodyImpl) {}
func (gen *Generator) Response(key string, impl *api.ResponseImpl)       {}
