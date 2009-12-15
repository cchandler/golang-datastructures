package graph

import "testing"

func TestVertexId(t *testing.T) {
    g := New();
    v := g.AddVertex();
    if v.Identifier() == 0 {
        t.Errorf("Identifier on vertex was nil");
    }
}