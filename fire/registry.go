package fire

import "net/http"

type Registry struct {
	ops []Operation
}

func NewRegistry() *Registry {
	return &Registry{}
}

func (r *Registry) Mux() *http.ServeMux {
	mux := http.NewServeMux()
	for _, op := range r.ops {
		op.Register(mux)
	}
	return mux
}

// Read the current state of the resource.
func (r *Registry) Read(t Type, h Handler) {
	r.ops = append(r.ops, Operation{
		Verb:        VerbGet,
		Path:        "/" + string(t) + "/{id}",
		Interaction: InteractionRead,
		Type:        t,
		Handler:     h,
	})
}

// Read the state of a specific version of the resource.
func (r *Registry) VRead(t Type, h Handler) {
	r.ops = append(r.ops, Operation{
		Verb:        VerbGet,
		Path:        "/" + string(t) + "/{id}/_history/{vid}",
		Interaction: InteractionVRead,
		Type:        t,
		Handler:     h,
	})
}

// Update an existing resource by its id (or create it if it is new).
func (r *Registry) Update(t Type, h Handler) {
	r.ops = append(r.ops, Operation{
		Verb:        VerbPut,
		Path:        "/" + string(t) + "/{id}",
		Interaction: InteractionUpdate,
		Type:        t,
		Handler:     h,
	})
}

// Update an existing resource by posting a set of changes to it.
func (r *Registry) Patch(t Type, h Handler) {
	r.ops = append(r.ops, Operation{
		Verb:        VerbPatch,
		Path:        "/" + string(t) + "/{id}",
		Interaction: InteractionPatch,
		Type:        t,
		Handler:     h,
	})
}

// Delete a resource.
func (r *Registry) Delete(t Type, h Handler) {
	r.ops = append(r.ops, Operation{
		Verb:        VerbDelete,
		Path:        "/" + string(t) + "/{id}",
		Interaction: InteractionDelete,
		Type:        t,
		Handler:     h,
	})
}

// Create a new resource with a server assigned id.
func (r *Registry) Create(t Type, h Handler) {
	r.ops = append(r.ops, Operation{
		Verb:        VerbPost,
		Path:        "/" + string(t),
		Interaction: InteractionCreate,
		Type:        t,
		Handler:     h,
	})
}

// Search the resource type based on some filter criteria.
func (r *Registry) Search(t Type, h Handler) {
	r.ops = append(r.ops, Operation{
		Verb:        VerbPost,
		Path:        "/" + string(t) + "/_search",
		Interaction: InteractionSearch,
		Type:        t,
		Handler:     h,
	})
}
