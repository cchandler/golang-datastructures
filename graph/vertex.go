package graph

import "rand"

type Vertex struct {
    edgeMap map[string]*Edge;
    id int64;
    Value interface{};
}

func newVertex() (*Vertex) {
    v := new(Vertex);
    v.id = rand.Int63();
    v.edgeMap = make(map[string]*Edge);
    return v;
}

func (v *Vertex) edgeiterate(c chan<- *Edge) () {
    for _, val := range v.edgeMap {
        c <- val;
    }
    close(c);
}

func (v *Vertex) EdgeIter() (chan *Edge) {
    c := make(chan *Edge);
    go v.edgeiterate(c);
    return c;
}

func (v *Vertex) registerEdge(e *Edge) (){
    v.edgeMap[e.Identifier()] = e;
}

func (v *Vertex) removeEdge(e *Edge) () {
    v.edgeMap[e.Identifier()] = e, false;
}

func (v *Vertex) Identifier() (int64) {
    return v.id;
}

func (v *Vertex) EdgeCount() (int) {
    return len(v.edgeMap);
}