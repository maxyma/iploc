package dict

type Node struct {
    child uint16
    next uint16
    value byte
    c8b byte        // 2^(8+16=24) = 16777216
    n8b byte        // 2^(8+16=24) = 16777216
    _ byte          // blank
}

func NewNode(val byte, next uint32) (Node) {
    n := Node{value:val}
    n.setNext(next)
    return n
}

func (n *Node) deepRight(t *Tree, depth byte, trackip *[8]byte) (*Node) {
    if n.getChild() == 0 {
        return n
    }
    var pre,p uint32
    for p=n.getChild(); p!=0; p=t.GetNode(p).getNext() {
        pre = p
    }
    trackip[depth] = t.GetNode(pre).value
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
            if t.GetNode(p).value == val {
                break
            } else if t.GetNode(p).value > val {
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

func (t *Node) getChild() uint32 {
    return uint32(t.c8b) << 16 | uint32(t.child)
}

func (t *Node) setChild(p uint32) {
    t.c8b = byte(p >> 16)
    t.child = uint16(p)
}

func (t *Node) getNext() uint32 {
    return uint32(t.n8b) << 16 | uint32(t.next)
}

func (t *Node) setNext(p uint32) {
    t.n8b = byte(p >> 16)
    t.next = uint16(p)
}

