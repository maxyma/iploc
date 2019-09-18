package main

import (
    "encoding/csv"
    "io"
    "io/ioutil"
    "bytes"
    "runtime/debug"
)

func load_dict() (*Tree) {
    var (
        s []string
        buf []byte
        err error
    )

    gcp := debug.SetGCPercent(300)

    if buf,err = ioutil.ReadFile("texts.txt"); err!=nil {
        panic(err)
    }

    csvreader := csv.NewReader(bytes.NewReader(buf))
    csvreader.Comma = '\t'
    csvreader.ReuseRecord = true
    csvreader.TrimLeadingSpace = true

    root := &Tree{}
    root.Extend(bytes.Count(buf,[]byte{'\n'})+1)

    for {
        if s,err = csvreader.Read(); err!=nil {
            break
        }
        root.AppendIP(FromStringToIP(s[0]))
    }

    if err!=nil && err!=io.EOF {
        panic(err)
    }

    debug.SetGCPercent(gcp)
    debug.FreeOSMemory()

    return root
}

