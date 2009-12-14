package graph

import "container/list"

type Graph struct {
    Vertices *list.List;
    Edges *list.List;
}

func New()(*Graph) {
    g := new(Graph);
    g.Edges = list.New();
    g.Vertices = list.New();
    return g;
}

func (g *Graph) EdgeCount() (int) {
    return g.Edges.Len();
}

func (g *Graph) VertexCount() (int) {
    return g.Vertices.Len();
}

func (g *Graph) AddVertex() (*Vertex) {
    v := newVertex();
    g.Vertices.PushFront(v);
    return v;
}

func (g *Graph) ConnectVertices(v1, v2 *Vertex) (*Edge) {
    edge := new(Edge);
    edge.endpoint1 = v1;
    edge.endpoint2 = v2;
    g.Edges.PushFront(edge);
    v1.Edges.PushFront(edge);
    v2.Edges.PushFront(edge);
    return edge;
}

