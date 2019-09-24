package dict

type Node struct {
    child uint16
    next uint16
    c8b byte        // 2^(8+16=24) = 16777216
    n8b byte        // 2^(8+16=24) = 16777216
    pack uint16     // 低4位存节点的值  高12位存IPLOC的指针
}

func NewNode(val byte, next uint32) (Node) {
    n := Node{}
    n.SetValue(val)
    n.setNext(next)
    return n
}

func (n *Node) deepRight(t *Tree, depth byte, trackip *[8]byte) (*Node) {
    trackip[depth] = n.GetValue()
    if n.getChild() == 0 {
        return n
    }
    var pre,p uint32
    for p=n.getChild(); p!=0; p=t.GetNode(p).getNext() {
        pre = p
    }
    return t.GetNode(pre).deepRight(t, depth+1, trackip)
}

func (n *Node) appendIP(t *Tree, depth byte, paths [8]byte) (*Node) {
    val := paths[depth]
    var pre, p uint32
    if n.getChild()==0 {
        p = t.AppendNode(NewNode(val,0))
        n.setChild(p)
    } else {
        for p=n.getChild(); p!=0; p=t.GetNode(p).getNext() {
            if t.GetNode(p).GetValue() == val {
                break
            } else if t.GetNode(p).GetValue() > val {
                np := t.AppendNode(NewNode(val,p))
                if pre!=0 {
                    t.GetNode(pre).setNext(np)
                } else {
                    n.setChild(np)
                }
                p = np
                break
            }
            pre = p
        }
        if p==0 {
            p = t.AppendNode(NewNode(val,0))
            t.GetNode(pre).setNext(p)
        }
    }

    if depth++; depth < 8 {
        return t.GetNode(p).appendIP(t, depth, paths)
    } else {
        return t.GetNode(p)
    }
}

func (n *Node) count(t *Tree) (c int){
    var p uint32
    for p=n.getChild(); p!=0; p=t.GetNode(p).getNext() {
        c += 1
        c += t.GetNode(p).count(t)
    }
    return
}

func (n *Node) getChild() uint32 {
    return uint32(n.c8b) << 16 | uint32(n.child)
}

func (n *Node) setChild(p uint32) {
    n.c8b = byte(p >> 16)
    n.child = uint16(p)
}

func (n *Node) getNext() uint32 {
    return uint32(n.n8b) << 16 | uint32(n.next)
}

func (n *Node) setNext(p uint32) {
    n.n8b = byte(p >> 16)
    n.next = uint16(p)
}

func (n *Node) SetValue(v byte) {
    n.pack = (n.pack >> 4 << 4) | uint16(v)
}

func (n *Node) GetValue() byte {
    return byte(n.pack) << 4 >> 4
}

func (n *Node) SetLoc(v uint16) {
    n.pack = (n.pack << 12 >> 12) | (v << 4)
}

func (n *Node) GetLoc() uint16 {
    return n.pack >> 4
}

