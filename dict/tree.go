package dict

//import "fmt"

type Tree struct {
    root Node
    mapping []Node
    records []Record
    iploc map[string]uint16
}

func NewTree() (t *Tree) {
    return &Tree{
        mapping : make([]Node,1),
        records : make([]Record,1),
        iploc : make(map[string]uint16),
    }
}

func (t *Tree) SearchIP(ip IP) (*Record) {
    if ip.Len() == 0 {
        return nil
    }
    var (
        p uint32
        depth int8
        found bool
        node *Node
    )
    paths := ip.ToPath()
    track := [8]uint32{}
    trackip := [8]byte{}

    for p = t.root.getChild(); p!=0; {
        if t.GetNode(p).GetValue() == paths[depth] {
            trackip[depth] = t.GetNode(p).GetValue()
            if depth+1 < 8 {
                p = t.GetNode(p).getChild()
            } else {
                found = true
                break
            }
            depth++
        } else if t.GetNode(p).GetValue() > paths[depth] {
            break
        } else {
            track[depth] = p
            p = t.GetNode(p).getNext()
        }
    }
    if !found {
        for i:=depth; i>=0; i-- {
            if track[i] != 0 {
                node = t.GetNode(track[i]).deepRight(t, byte(i), &trackip)
                break
            }
        }
    } else {
        node = t.GetNode(p)
    }
    return &t.records[node.GetLoc()]
}

func (t *Tree) AppendRecord(r Record) {
    ip := NewStringIP(r.from)
    if ip.Len() > 0 {
        node := t.root.appendIP(t, 0, ip.ToPath())
        if v,ok:=t.iploc[r.loc]; ok==true {
            node.SetLoc(v)
        } else {
            idx := uint16(len(t.records))
            node.SetLoc(idx)
            t.iploc[r.loc] = idx
            t.records = append(t.records, r)
        }
    }
}

func (t *Tree) Count() (int) {
    return t.root.count(t)
}

func (t *Tree) Extend(size int) {
    t.mapping = make([]Node, 1, size)
}

func (t *Tree) Shrink() {
    a := make([]Node, len(t.mapping))
    copy(a, t.mapping)
    t.mapping = a

    b := make([]Record, len(t.records))
    copy(b, t.records)
    t.records = b

    t.iploc = nil
}

func (t *Tree) AppendNode(n Node) (p uint32){
    p = uint32(len(t.mapping))
    t.mapping = append(t.mapping, n)
    return p
}

func (t *Tree) GetNode(p uint32) (*Node) {
    return &t.mapping[p]
}


