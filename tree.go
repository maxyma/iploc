package main

//import     "fmt"

type Tree struct {
    root Node
    mapping []Node
}

func (t *Tree) SearchIP(ip MyIP) (MyIP) {
    var (
        p uint32
        depth int8
        found bool
        node *Node
    )
    paths := ip.ToPath()
    track := [8]uint32{}

    for p = t.root.child; p!=0; {
        if t.Get(p).value == paths[depth] {
            if depth+1 < 8 {
                p = t.Get(p).child
            } else {
                found = true
                break
            }
            depth++
        } else if t.Get(p).value > paths[depth] {
            break
        } else {
            track[depth] = p
            p = t.Get(p).next
        }
    }

    if p != 0 {
        node = t.Get(p)
    }
    if !found {
        for i:=depth; i>=0; i-- {
            if track[i] != 0 {
                node = t.Get(track[i]).deepRight(t)
                break
            }
        }
    }
    return FromUint32ToIP(node.ipv)
}

func (t *Tree) AppendIP(ip MyIP){
    n := t.root.appendIP(t, 0, ip.ToPath())
    n.ipv = ip.ToUint32()
}

func (t *Tree) Count() (int) {
    return t.root.count(t)
}

func (t *Tree) Extend(size int) {
    t.mapping = make([]Node, 1, size)
}

func (t *Tree) NewNode(n Node) (p uint32){
    p = uint32(len(t.mapping))
    t.mapping = append(t.mapping, n)
    return p
}

func (t *Tree) Get(p uint32) (*Node) {
    return &t.mapping[p]
}


