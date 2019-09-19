package dict

//import     "fmt"

type Tree struct {
    root Node
    mapping []Node
}

func NewTree() (t *Tree) {
    return &Tree{}
}

func (t *Tree) SearchIP(ip IP) (IP) {
    var (
        p uint32
        depth int8
        found bool
    )
    paths := ip.ToPath()
    track := [8]uint32{}
    trackip := [8]byte{}

    for p = t.root.getChild(); p!=0; {
        if t.GetNode(p).value == paths[depth] {
            trackip[depth] = t.GetNode(p).value
            if depth+1 < 8 {
                p = t.GetNode(p).getChild()
            } else {
                found = true
                break
            }
            depth++
        } else if t.GetNode(p).value > paths[depth] {
            break
        } else {
            track[depth] = p
            p = t.GetNode(p).getNext()
        }
    }
    if !found {
        for i:=depth; i>=0; i-- {
            if track[i] != 0 {
                t.GetNode(track[i]).deepRight(t, byte(depth), &trackip)
                break
            }
        }
    }
    return NewBytesIP(trackip)
}

func (t *Tree) AppendIP(ip IP){
    t.root.appendIP(t, 0, ip.ToPath())
}

func (t *Tree) Count() (int) {
    return t.root.count(t)
}

func (t *Tree) Extend(size int) {
    t.mapping = make([]Node, 1, size)
}

func (t *Tree) AppendNode(n Node) (p uint32){
    p = uint32(len(t.mapping))
    t.mapping = append(t.mapping, n)
    return p
}

func (t *Tree) GetNode(p uint32) (*Node) {
    return &t.mapping[p]
}


