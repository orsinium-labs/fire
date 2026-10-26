package fire

import "net/http"

type Registry map[Type]interface {
	extend(Type, []Operation) []Operation
}

func (r Registry) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	ops := []Operation{}
	for t, op := range r {
		op.extend(t, ops)
	}
	for _, op := range ops {
		op.Register(mux)
	}
	return mux
}

type Operations[T any] struct {
	// Read the current state of the resource.
	Read ReadHandler
	// Read the state of a specific version of the resource.
	VRead VReadHandler
	// Update an existing resource by its id (or create it if it is new).
	Update UpdateHandler
	// Update an existing resource by posting a set of changes to it.
	Patch PatchHandler
	// Delete a resource.
	Delete DeleteHandler
	// Create a new resource with a server assigned id.
	Create CreateHandler
	// Search the resource type based on some filter criteria.
	Search SearchHandler
}

func (ops Operations[T]) extend(t Type, res []Operation) []Operation {
	if ops.Read != nil {
		res = append(res, Operation{
			Verb:        VerbGet,
			Path:        "/" + string(t) + "/{id}",
			Interaction: InteractionRead,
			Type:        t,
			Handler:     ops.Read,
		})
	}

	if ops.VRead != nil {
		res = append(res, Operation{
			Verb:        VerbGet,
			Path:        "/" + string(t) + "/{id}/_history/{vid}",
			Interaction: InteractionVRead,
			Type:        t,
			Handler:     ops.VRead,
		})
	}

	if ops.Update != nil {
		res = append(res, Operation{
			Verb:        VerbPut,
			Path:        "/" + string(t) + "/{id}",
			Interaction: InteractionUpdate,
			Type:        t,
			Handler:     ops.Update,
		})
	}

	if ops.Patch != nil {
		res = append(res, Operation{
			Verb:        VerbPatch,
			Path:        "/" + string(t) + "/{id}",
			Interaction: InteractionPatch,
			Type:        t,
			Handler:     ops.Patch,
		})
	}

	if ops.Delete != nil {
		res = append(res, Operation{
			Verb:        VerbDelete,
			Path:        "/" + string(t) + "/{id}",
			Interaction: InteractionDelete,
			Type:        t,
			Handler:     ops.Delete,
		})
	}

	if ops.Create != nil {
		res = append(res, Operation{
			Verb:        VerbPost,
			Path:        "/" + string(t),
			Interaction: InteractionCreate,
			Type:        t,
			Handler:     ops.Create,
		})
	}

	if ops.Search != nil {
		res = append(res, Operation{
			Verb:        VerbPost,
			Path:        "/" + string(t) + "/_search",
			Interaction: InteractionSearch,
			Type:        t,
			Handler:     ops.Search,
		})
	}
	return res
}
