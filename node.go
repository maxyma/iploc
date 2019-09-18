package main

type Node struct {
    value byte
    child uint32
    next uint32
}

func (n *Node) deepRight(t *Tree, depth byte, trackip *[8]byte) (*Node) {
    if n.child == 0 {
        return n
    }
    var pre,p uint32
    for p=n.child; p!=0; p=t.Get(p).next {
        pre = p
    }
    trackip[depth] = t.Get(pre).value
    return t.Get(pre).deepRight(t, depth+1, trackip)
}

func (n *Node) appendIP(t *Tree, depth byte, paths [8]byte) (*Node) {
    val := paths[depth]
    var pre, p uint32
    if n.child==0 {
        p = t.NewNode(Node{value:val})
        n.child = p
    } else {
        for p=n.child; p!=0; p=t.Get(p).next {
            if t.Get(p).value == val {
                break
            } else if t.Get(p).value > val {
                np := t.NewNode(Node{value:val, next:p})
                if pre!=0 {
                    t.Get(pre).next = np
                } else {
                    n.child = np
                }
                p = np
                break
            }
            pre = p
        }
        if p==0 {
            p = t.NewNode(Node{value:val})
            t.Get(pre).next = p
        }
    }

    if depth++; depth < 8 {
        return t.Get(p).appendIP(t, depth, paths)
    } else {
        return t.Get(p)
    }
}

func (n *Node) count(t *Tree) (c int){
    var p uint32
    for p=n.child; p!=0; p=t.Get(p).next {
        c += 1
        c += t.Get(p).count(t)
    }
    return
}


