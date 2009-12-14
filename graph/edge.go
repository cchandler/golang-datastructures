package graph

type Edge struct {
    endpoint1 *Vertex;
    endpoint2 *Vertex;
    Weight int;
}

func (e *Edge) Endpoints() (v1,v2 *Vertex) {
    return e.endpoint1, e.endpoint2;
}