package dict

type RecordLite struct {
    loc string
    isp string
}

type Record struct {
    RecordLite
    from string
    to string
}

func NewRecord(from,to,loc,isp string) (Record) {
    return Record { RecordLite{loc,isp},from,to }
}

func (r Record) Lite() (RecordLite) {
    return RecordLite{ r.loc, r.isp }
}

