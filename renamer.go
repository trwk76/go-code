package code

type (
	// IDRenamer is an interface that defines a way to rename a code symbol from one identifier to another.
	IDRenamer interface {
		Rename(id string) string
	}

	// IDRecaser is a Renamer that changes the case of identifiers.
	IDRecaser struct {
		Transform IDTransformer
	}

	// IDMapper is a Renamer that maps identifiers to other identifiers and uses a potential NotFound
	// IDTransformer in case the identifier is not in the map. NotFound is allowed to do anything:
	// - return the same identifier
	// - return a different identifier
	// - panic if the identifier is not found
	IDMapper struct {
		Map map[string]string
		// NotFound is called whenever the symbol is not found in the dictionary.
		NotFound IDTransformer
	}
)

func (c IDRecaser) Rename(id string) string {
	if c.Transform == nil {
		return id
	}

	return c.Transform(id)
}

func (c IDMapper) Rename(id string) string {
	if c.Map == nil {
		if match, ok := c.Map[id]; ok {
			return match
		}
	}

	if c.NotFound == nil {
		return id
	}

	return c.NotFound(id)
}

var (
	_ IDRenamer = IDRecaser{}
)
