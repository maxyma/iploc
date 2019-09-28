package main

import (
    "iploc/dict"
    "net/http"
    "encoding/json"
)

var (
    root *dict.Tree
)

func init(){
    root = dict.Load()
}

func main(){
    http.HandleFunc("/", forbid)
    http.HandleFunc("/favicon.ico", forbid)
    http.HandleFunc("/iploc", iploc)
    http.ListenAndServe(":8088", nil)
}

func iploc(w http.ResponseWriter, req *http.Request){
    loc := root.SearchIP(dict.NewStringIP(req.URL.Query().Get("ip")))
    if out,err:=json.Marshal(NewNode2(loc)); err==nil {
        w.WriteHeader(http.StatusOK)
        w.Write(out)
    } else {
        w.WriteHeader(http.StatusInternalServerError)
    }
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

