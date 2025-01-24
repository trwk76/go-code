package api

import (
	"net/http"

	golang "github.com/trwk76/gocode/go"
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

func (p Path) build(ctx buildContext, dest spec.Paths, body *golang.BlockStmt) {
	ctx.tags = mergeTags(ctx.tags, p.Tags)

	p.Named.build(ctx, dest, body)
	p.Param.build(ctx, dest, body)

	if p.Summary != "" || p.Description != "" || p.GET != nil || p.POST != nil || p.PUT != nil || p.DELETE != nil || p.OPTIONS != nil || p.HEAD != nil || p.PATCH != nil || p.TRACE != nil {
		dest[ctx.path] = &spec.PathItem{
			Summary:     p.Summary,
			Description: p.Description,
			GET:         p.GET.build(ctx, http.MethodGet, false, body),
			POST:        p.POST.build(ctx, http.MethodPost, true, body),
			PUT:         p.PUT.build(ctx, http.MethodPut, true, body),
			DELETE:      p.DELETE.build(ctx, http.MethodDelete, false, body),
			OPTIONS:     p.OPTIONS.build(ctx, http.MethodOptions, false, body),
			HEAD:        p.HEAD.build(ctx, http.MethodHead, false, body),
			PATCH:       p.PATCH.build(ctx, http.MethodPatch, true, body),
			TRACE:       p.TRACE.build(ctx, http.MethodTrace, false, body),
		}
	}
}

func (p *ParamPath) build(ctx buildContext, dest spec.Paths, body *golang.BlockStmt) {
	if p == nil {
		return
	}

	p.Path.build(ctx.paramChild(p.Param), dest, body)
}

func (p NamedPaths) build(ctx buildContext, dest spec.Paths, body *golang.BlockStmt) {
	for name, item := range p {
		item.build(ctx.namedChild(name), dest, body)
	}
}
