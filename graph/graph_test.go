package graph

import "testing"
import "disjoint"

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

func TestRemoveVertex(t *testing.T) {
    g := New();
    
    v := g.AddVertex();
    g.RemoveVertex(v);
    
    if g.VertexCount() != 0 {
        t.Errorf("Didn't remove vertex from count");
    }
}

func TestRemoveEdgeViaRemoveVertex(t *testing.T) {
    g := New();
    
    v1 := g.AddVertex();
    v2 := g.AddVertex();
    g.ConnectVertices(v1,v2);
    g.RemoveVertex(v1);
    
    if g.EdgeCount() != 0 {
        t.Errorf("Didn't remove edge from count");
    }
    
    if v2.EdgeCount() != 0 {
        t.Errorf("Didn't remove edge from second vertex after edge removal")
    }
}

func TestRemoveEdge(t *testing.T) {
    g := New();

   v1 := g.AddVertex();
   v2 := g.AddVertex();
   edge := g.ConnectVertices(v1,v2);
   
   g.RemoveEdge(edge);
   
   if g.EdgeCount() != 0 {
       t.Errorf("Didn't remove edge from count");
   }
   
   if v1.EdgeCount() != 0 || v2.EdgeCount() != 0 {
       t.Errorf("Didn't correctly remove edge from vertices");
   }
}

func TestLookupVertex(t *testing.T) {
    g := New();
    v := g.AddVertex();
    
    if g.GetVertex(123) != nil {
        t.Errorf("Found a nonsense vertex");
    }
    
    if g.GetVertex(v.Identifier()) == nil {
        t.Errorf("Didn't find a proper vertex");
    }
}

func TestJoinVerticesWithEdge(t *testing.T) {
    g := New();
    
    v1 := g.AddVertex();
    v2 := g.AddVertex();
    edge := g.ConnectVertices(v1,v2);
    
    
    if g.EdgeCount() != 1 {
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

func TestEdgeExistence(t *testing.T) {
    g := New();

    v1 := g.AddVertex();
    v2 := g.AddVertex();
    edge := g.ConnectVertices(v1,v2);
    
    if g.EdgeExists(edge) != true {
        t.Errorf("Failed to correctly determine if an edge exists in the graph")
    }
}

func TestEdgeLookup(t *testing.T) {
    g := New();

    v1 := g.AddVertex();
    v2 := g.AddVertex();
    edge := g.ConnectVertices(v1,v2);
    
    if g.GetEdge(edge.Identifier()) == nil {
        t.Errorf("Failed to correctly lookup an edge");
    }
}

func TestSingleEdgeBetweenVertices(t *testing.T) {
    g := New();
    
    v1 := g.AddVertex();
    v2 := g.AddVertex();
    edge := g.ConnectVertices(v1,v2);
    edge.Weight = 2;
    
    //Sneaking a second test that makes sure params are unordered
    edge1 := g.ConnectVertices(v2,v1);
    
    if edge1.Weight != 2 {
        t.Errorf("Didn't correctly return the same edge: %d", edge1.Weight);
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

func TestDfsTraversal(t *testing.T) {
    g := New();
    
    vertices := new([100]*Vertex);
    
    for i := 0 ; i < 100 ; i++ {
        vertices[i] = g.AddVertex();
    }
    
    for i := 0 ; i < 99 ; i++ {
        v1 := vertices[i];
        v2 := vertices[i + 1];
        g.ConnectVertices(v1,v2);
    }
    
    var i int32 = 0;
    visiter := func(v *Vertex)() {
        i++;
    };
    
    g.Dfs(vertices[0],visiter);
    
    if i != 99 {
        t.Errorf("Did not visit all vertices: %i",i);
    }
}

func TestFindConnectedComponents(t *testing.T) {
    g := New();
    
    vertices := new([100]*Vertex);
    
    for i := 0 ; i < 100 ; i++ {
        vertices[i] = g.AddVertex();
    }
    
    for i := 0 ; i < 99 ; i++ {
        v1 := vertices[i];
        v2 := vertices[i + 1];
        g.ConnectVertices(v1,v2);
    }
    
    result := g.FindConnectedComponents();
    
    for i := 0 ; i < 99 ; i++ {
        v1 := vertices[i];
        v2 := vertices[i + 1];
        
        if disjoint.Find(result[v1.Identifier()]) != disjoint.Find(result[v2.Identifier()]) {
            t.Errorf("Two nodes were not in the same connected component despite being connected");
        }
    }
}

func TestFindConnectedComponentsDifferentComponents(t *testing.T) {
    g := New();
    
    vertices := new([100]*Vertex);
    
    for i := 0 ; i < 100 ; i++ {
        vertices[i] = g.AddVertex();
    }
    
    for i := 0 ; i < 49 ; i++ {
        v1 := vertices[i];
        v2 := vertices[i + 1];
        g.ConnectVertices(v1,v2);
    }
    
    for i := 51 ; i < 99 ; i++ {
        v1 := vertices[i];
        v2 := vertices[i + 1];
        g.ConnectVertices(v1,v2);
    }
    
    result := g.FindConnectedComponents();
    
    v1 := vertices[49];
    v2 := vertices[51];
    
    if disjoint.Find(result[v1.Identifier()]) == disjoint.Find(result[v2.Identifier()]) {
        t.Errorf("Two vertices were unexpectedly in the same connected component");
    }
    
}