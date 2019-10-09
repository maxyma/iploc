package main

import (
    "iploc/dict"
    "net/http"
    "encoding/json"
    "flag"
)

var (
    root *dict.Tree
)

func main(){
    //解析命令行参数
    serv_host := flag.String("host", "", "Hostname or IP")
    serv_port := flag.String("port", "8811", "Port to Listen on.")
    flag.Parse()
    //load dict
    root = dict.Load()
    //web server
    http.HandleFunc("/", forbid)
    http.HandleFunc("/favicon.ico", forbid)
    http.HandleFunc("/iploc", iploc)
    rl := NewReloader(*serv_host+":"+*serv_port)
    if err:=rl.Bind(); err==nil {
        rl.HttpServe(&http.Server{})
    } else {
        panic(err)
    }
}

func iploc(w http.ResponseWriter, req *http.Request){
    loc := root.SearchIP(dict.NewStringIP(req.URL.Query().Get("ip")))
    if loc!=nil {
        if out,err:=json.Marshal(NewNode2(loc)); err==nil {
            w.WriteHeader(http.StatusOK)
            w.Write(out)
            return
        }
    }
    w.WriteHeader(http.StatusInternalServerError)
}

func forbid(rw http.ResponseWriter,req *http.Request){
    rw.WriteHeader(http.StatusForbidden)
}

type Node2 struct{
    Left    string      `json:"rfrom"`
    Right   string      `json:"rto"`
    Value   string      `json:"loc"`
    Isp     string      `json:"isp"`
}

func NewNode2(loc []string) Node2 {
    return Node2{
        Left : loc[0],
        Value : loc[1],
        Isp : loc[2],
    }
}

