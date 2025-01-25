package api

import (
	"net/http"

	"github.com/trwk76/gocode/web/api/spec"
)

type (
	Path struct {
		Summary     string
		Description string
		GET         *Operation
		POST        *Operation
		PUT         *Operation
		DELETE      *Operation
		OPTIONS     *Operation
		HEAD        *Operation
		PATCH       *Operation
		TRACE       *Operation
		Tags        []string
		Named       NamedPaths
		Param       *ParamPath
	}

	ParamPath struct {
		Path
		Param Parameter
	}

	NamedPaths map[string]Path
)

func (p Path) build(ctx buildContext, dest spec.Paths) {
	ctx.tags = mergeTags(ctx.tags, p.Tags)

	p.Named.build(ctx, dest)
	p.Param.build(ctx, dest)

	if p.Summary != "" || p.Description != "" || p.GET != nil || p.POST != nil || p.PUT != nil || p.DELETE != nil || p.OPTIONS != nil || p.HEAD != nil || p.PATCH != nil || p.TRACE != nil {
		dest[ctx.path] = &spec.PathItem{
			Summary:     p.Summary,
			Description: p.Description,
			GET:         p.GET.build(ctx, http.MethodGet, false),
			POST:        p.POST.build(ctx, http.MethodPost, true),
			PUT:         p.PUT.build(ctx, http.MethodPut, true),
			DELETE:      p.DELETE.build(ctx, http.MethodDelete, false),
			OPTIONS:     p.OPTIONS.build(ctx, http.MethodOptions, false),
			HEAD:        p.HEAD.build(ctx, http.MethodHead, false),
			PATCH:       p.PATCH.build(ctx, http.MethodPatch, true),
			TRACE:       p.TRACE.build(ctx, http.MethodTrace, false),
		}
	}
}

func (p *ParamPath) build(ctx buildContext, dest spec.Paths) {
	if p == nil {
		return
	}

	p.Path.build(ctx.paramChild(p.Param), dest)
}

func (p NamedPaths) build(ctx buildContext, dest spec.Paths) {
	for name, item := range p {
		item.build(ctx.namedChild(name), dest)
	}
}
