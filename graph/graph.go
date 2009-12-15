package graph

// import "fmt"

type Graph struct {
    vertexMap map[int64] *Vertex;
    edgeMap map[string] *Edge;
}

func New()(*Graph) {
    g := new(Graph);
    g.edgeMap = make(map[string] *Edge);
    g.vertexMap = make(map[int64] *Vertex);
    return g;
}

func (g *Graph) EdgeCount() (int) {
    return len(g.edgeMap);
}

func (g *Graph) VertexCount() (int) {
    return len(g.vertexMap);
}

func (g *Graph) AddVertex() (*Vertex) {
    v := newVertex();
    g.vertexMap[v.Identifier()] = v;
    return v;
}

func (g *Graph) GetVertex(id int64) (*Vertex) {
    vertex, present := g.vertexMap[id];
    if present {
        return vertex;
    }
    return nil;
}

func (g *Graph) RemoveVertex(v *Vertex) () {
    if !g.VertexExists(v) {
        return;
    }
    
    iteration_channel := v.EdgeIter();
    
    for !closed(iteration_channel) {
        e := <-iteration_channel;
        if e == nil { continue }
        e.removeSelf();
        g.edgeMap[e.Identifier()] = e, false;
    }
    
    g.vertexMap[v.Identifier()] = v, false;
}

func (g *Graph) RemoveEdge(e *Edge) () {
    if !g.EdgeExists(e) {
        return;
    }
    
    e.removeSelf();
    g.edgeMap[e.Identifier()] = e, false;
}

func (g *Graph) GetEdge(id string) (*Edge) {
    edge, present := g.edgeMap[id];
    if present {
        return edge;
    }
    return nil;
}

func (g *Graph) ConnectVertices(v1, v2 *Vertex) (edge *Edge) {
    edge = newEdge(v1,v2);
    
    if g.EdgeExists(edge) {
        //Return the edge we already have
        edge = g.edgeMap[edge.Identifier()];
    }
    else
    {
        g.edgeMap[edge.Identifier()] = edge;
        v1.registerEdge(edge);
        v2.registerEdge(edge);
    }
    
    return edge;
}

func (g *Graph) VertexExists(v *Vertex) (bool) {
    _, present := g.vertexMap[v.Identifier()];
    return present;
}

func (g *Graph) EdgeExists(edge *Edge) (bool) {
    _, present := g.edgeMap[edge.Identifier()];
    return present;
}