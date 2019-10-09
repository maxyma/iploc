package dict

import (
    "encoding/csv"
    "io"
    "io/ioutil"
    "bytes"
    "runtime/debug"
)

func Load(file string) (*Tree) {
    var (
        s []string
        buf []byte
        err error
    )

    gcp := debug.SetGCPercent(300)

    if buf,err = ioutil.ReadFile(file); err!=nil {
        panic(err)
    }

    csvreader := csv.NewReader(bytes.NewReader(buf))
    csvreader.Comma = '\t'
    csvreader.ReuseRecord = true
    csvreader.TrimLeadingSpace = true

    root := NewTree()
    root.Extend(bytes.Count(buf,[]byte{'\n'})*3)

    for {
        if s,err = csvreader.Read(); err!=nil {
            break
        }
        root.AppendRecord(NewRecord(s[0],s[1],s[2],s[3]))
    }

    if err!=nil && err!=io.EOF {
        panic(err)
    }

    root.Shrink()

    debug.SetGCPercent(gcp)
    debug.FreeOSMemory()

    return root
}

