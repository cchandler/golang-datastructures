package disjoint

import("testing")

func TestMakeSet(t *testing.T) {
    e1 := Makeset();
    if e1.Parent != e1 {
        t.Errorf("Incorrect parent in Maketset");
    }
}

func TestUnion(t *testing.T) {
    e1 := Makeset();
    e2 := Makeset();
    
    Union(e1,e2);
    
    if Find(e1) != e2 {
        t.Errorf("Incorrect parent after a union");
    }
}

func TestPathCompression(t *testing.T) {
    e1 := Makeset();
    e2 := Makeset();
    e3 := Makeset();
    
    Union(e2,e1);
    Union(e3,e2);
    
    if e3.Parent != e1 {
        t.Errorf("Path was incorrectly compressed after 2 unions");
    }
}