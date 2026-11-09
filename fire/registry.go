package fire

import "net/http"

type Registry map[Type]interface {
	extend(Type, []Operation) []Operation
}

func (r Registry) Operations() []Operation {
	ops := []Operation{}
	for t, op := range r {
		op.extend(t, ops)
	}
	return ops
}

func (r Registry) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	for _, op := range r.Operations() {
		op.Register(mux)
	}
	return mux
}

type Operations[R Resource] struct {
	// Read the current state of the resource.
	Read ReadHandler[R]
	// Read the state of a specific version of the resource.
	VRead VReadHandler[R]
	// Update an existing resource by its id (or create it if it is new).
	Update UpdateHandler[R]
	// Update an existing resource by posting a set of changes to it.
	Patch PatchHandler[R]
	// Delete a resource.
	Delete DeleteHandler[R]
	// Create a new resource with a server assigned id.
	Create CreateHandler[R]
	// Search the resource type based on some filter criteria.
	Search SearchHandler[R]

	// Additional operations on the resource instance.
	//
	// https://www.hl7.org/fhir/operations.html
	InstanceExtensions map[string]InstanceExtHandler[R]
	TypeExtensions     map[string]TypeExtHandler[R]
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

	for name, h := range ops.InstanceExtensions {
		res = append(res, Operation{
			Verb:        VerbPost,
			Path:        "/" + string(t) + "/{id}/$" + name,
			Interaction: InteractionExecute,
			Type:        t,
			Handler:     h,
		})
	}

	for name, h := range ops.TypeExtensions {
		res = append(res, Operation{
			Verb:        VerbPost,
			Path:        "/" + string(t) + "/$" + name,
			Interaction: InteractionExecute,
			Type:        t,
			Handler:     h,
		})
	}

	return res
}
