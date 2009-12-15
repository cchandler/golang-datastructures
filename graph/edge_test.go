package graph

import "testing"

func TestEdgeId(t *testing.T) {
    g := New();
    v1 := g.AddVertex();
    v2 := g.AddVertex();
    
    edge := g.ConnectVertices(v1,v2);

    if edge.Identifier() == "" {
        t.Errorf("Edge does not have a proper identifier");
    }
}

func TestEdgeIdSize(t *testing.T) {
    g := New();
    v1 := g.AddVertex();
    v2 := g.AddVertex();
    
    edge := g.ConnectVertices(v1,v2);

    if len(edge.IdentifierAsBytes()) != 16 {
        t.Errorf("Edge identifier not 16 bytes: %d", len(edge.IdentifierAsBytes()));
    }
}