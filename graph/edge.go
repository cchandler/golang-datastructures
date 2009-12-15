package graph

import "crypto/md5"
import "encoding/binary"
import "big"

type Edge struct {
    endpoint1 *Vertex;
    endpoint2 *Vertex;
    Weight int;
    id []byte;
    idAsInt *big.Int;
}

func newEdge(v1,v2 *Vertex) (*Edge) {
    e := new(Edge);
    if v1.Identifier() > v2.Identifier() {
        e.endpoint1 = v1;
        e.endpoint2 = v2;
    }
    else
    {
        e.endpoint1 = v2;
        e.endpoint2 = v1;
    }
    e.Weight = 0;
    
    hash := md5.New();
    binary.Write(hash, binary.BigEndian , e.endpoint1.Identifier());
    binary.Write(hash, binary.BigEndian , e.endpoint2.Identifier());
    e.id = hash.Sum();
    
    return e;
}

func (e *Edge) removeSelf() () {
    e.endpoint1.removeEdge(e);
    e.endpoint2.removeEdge(e);
}

func (e *Edge) Identifier() (string){
    if e.idAsInt != nil {
        return e.idAsInt.String();
    }
    bignum := big.NewInt(0);
    bignum.SetBytes(e.id);
    e.idAsInt = bignum;
    return bignum.String();
}

func (e *Edge) IdentifierAsBytes() ([]byte) {
    return e.id;
}

func (e *Edge) Endpoints() (v1,v2 *Vertex) {
    return e.endpoint1, e.endpoint2;
}