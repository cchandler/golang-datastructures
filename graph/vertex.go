package graph

import "container/list"

type Vertex struct {
    Edges *list.List;
    Value interface{};
}

func newVertex() (*Vertex) {
    v := new(Vertex);
    v.Edges = list.New();
    return v;
}

func (v *Vertex) EdgeCount() (int) {
    return v.Edges.Len();
}