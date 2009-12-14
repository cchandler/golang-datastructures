package disjoint

type Element struct {
    Parent *Element;
    Value interface{};
}

func Makeset() (*Element)
{
    e := new(Element);
    e.Parent = e;
    return e;
}

func Find(e *Element) (*Element)
{
    if e.Parent == e {
        return e;
    }
    e.Parent = Find(e.Parent);
    return e.Parent;
}

func Union(e1,e2 *Element) ()
{
    root1 := Find(e1);
    root2 := Find(e2);
    root1.Parent = root2;
}