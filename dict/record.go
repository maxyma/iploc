package dict

type Record struct {
    from string
    to string
    loc string
    isp string
}

func NewRecord(from,to,loc,isp string) (Record) {
    return Record { from,to,loc,isp }
}

