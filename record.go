package main

type Record struct {
    from uint32
    to uint32
    loc []byte
    //isp []byte
}

//  type Mapping struct {
//      mapping map[uint32]Record
//  }
//  
//  func NewMapping() (Mapping) {
//      return Mapping{make(map[uint32]Record)}
//  }
//  
//  func (t *Mapping) Append(from, to, loc, isp string) {
//      //ip := FromStringToIP(from)
//      //n := t.appendIP(0, ip.ToPath())
//      //n.ip = ip.ToUint32()
//  
//  //    if _,ok := mapping[ip.ToUint32()]; ok!=true {
//  //        to := FromStringToIP(to)
//  //        mapping[ip.ToUint32()] = Record{
//  //            ip.ToUint32(),
//  //            to.ToUint32(),
//  //            []byte(loc),
//  //            []byte(isp),
//  //        }
//  //    }
//  }

