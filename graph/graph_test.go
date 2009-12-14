package graph

import "testing"

func TestGraphCreation(t *testing.T) {
    g := New();
    
    if g.EdgeCount() != 0 {
        t.Errorf("Started with edges");
    }
    
    if g.VertexCount() != 0 {
        t.Errorf("Started with vertices");
    }
}

func TestAddVertex(t *testing.T) {
    g := New();
    
    g.AddVertex();
    
    if g.VertexCount() != 1 {
        t.Errorf("Didn't add vertex");
    }
}

func TestJoinVerticesWithEdge(t *testing.T) {
    g := New();
    
    v1 := g.AddVertex();
    v2 := g.AddVertex();
    edge := g.ConnectVertices(v1,v2);
    
    if g.Edges.Len() != 1 {
        t.Errorf("Didn't correctly add edge to edge count");
    }
    
    v3, v4 := edge.Endpoints();
    if v3 != v1 || v4 != v2 {
        t.Errorf("Didn't correctly set endpoints");
    }
    
    if v1.EdgeCount() != 1 || v2.EdgeCount() != 1 {
        t.Errorf("Didn't correctly add edge reference to vertex")
    }
}

func TestAddEdgeWeight(t *testing.T) {
    g := New();
    
    v1 := g.AddVertex();
    v2 := g.AddVertex();
    edge := g.ConnectVertices(v1,v2);
    
    edge.Weight = 3;
    
    if edge.Weight != 3 {
        t.Errorf("Didn't correctly set edge weight");
    }
}