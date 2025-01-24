package api

import (
	"fmt"
	"strconv"

	golang "github.com/trwk76/gocode/go"
	"github.com/trwk76/gocode/web/api/spec"
)

type (
	Operation struct {
		ID          string
		Summary     string
		Description string
		Parameters  []Parameter
		RequestBody RequestBody
		Responses   ResponseMap
		Deprecated  bool
		Security    spec.SecurityRequirements
		Tags        []string
	}

	ResponseMap struct {
		Codes   map[int]Response
		Default Response
	}
)

func (o *Operation) build(ctx buildContext, method string, acceptBody bool, body *golang.BlockStmt) *spec.Operation {
	if o == nil {
		return nil
	}

	if o.RequestBody != nil && !acceptBody {
		panic(fmt.Errorf("http method '%s' does not accept a request body", method))
	}

	params := make(spec.ParameterOrRefs, 0, len(ctx.params)+len(o.Parameters))

	for _, item := range ctx.params {
		params = append(params, item.paramSpec())
	}

	for _, item := range o.Parameters {
		s := item.paramImpl().spec()

		if s.In == spec.ParameterPath {
			panic(fmt.Errorf("path parameter '%s' must be defined in the path", s.Name))
		}

		params = append(params, item.paramSpec())
	}

	res := &spec.Operation{
		OperationID: o.ID,
		Summary:     o.Summary,
		Description: o.Description,
		Parameters:  params,
		Responses:   o.Responses.spec(),
		Deprecated:  o.Deprecated,
		Security:    o.Security,
	}

	if o.RequestBody != nil {
		s := o.RequestBody.reqBodySpec()
		res.RequestBody = &s
	}

	*body = append(*body, golang.ExprStmt{
		Expr: golang.CallExpr{
			Func: golang.MemberExpr{
				Value: golang.Symbol{
					ID: golang.ID("m"),
				},
				ID:    golang.ID("HandleFunc"),
			},
			Args: golang.Exprs{
				golang.StringExpr(fmt.Sprintf("%s %s", method, ctx.path)),
				golang.Symbol{
					ID: golang.ID(o.ID),
				},
			},
		},
	})

	return res
}

func (r ResponseMap) spec() spec.Responses {
	res := make(spec.Responses)

	for code, item := range r.Codes {
		res[strconv.Itoa(code)] = item.respSpec()
	}

	if r.Default != nil {
		res["default"] = r.Default.respSpec()
	}

	return res
}
