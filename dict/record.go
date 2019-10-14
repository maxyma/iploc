package dict

type RecordLite struct {
    loc string
}

type Record struct {
    RecordLite
    from string
}

func NewRecord(from,to,loc string) (Record) {
    return Record { RecordLite{loc},from }
}

func (r Record) Lite() (RecordLite) {
    return RecordLite{ r.loc }
}

