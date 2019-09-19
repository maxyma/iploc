package main

import (
//    "fmt"
//    "runtime"
    "log"
    "unsafe"
    "iploc/dict"
)

func main(){

    //runtime.LockOSThread()

    log.Println("loading...")

    //runtime.MemProfileRate = 1

    //mk_cpu_prof()

    root := dict.Load()

    //stop_cpu_pro()

    log.Printf("loaded! count:%d\n", root.Count())
    //spew.Dump(root.child)

	log.Printf("%+v\n", dict.NewBytesIP([8]byte{8,1,0,2,0,3,0,4}))

	log.Printf("find %+v\n", root.SearchIP(dict.NewStringIP("0.0.0.0")))
	log.Printf("find %+v\n", root.SearchIP(dict.NewStringIP("1.2.3.5")))
	log.Printf("find %+v\n", root.SearchIP(dict.NewStringIP("223.2.3.4")))
	log.Printf("find %+v\n", root.SearchIP(dict.NewStringIP("33.2.3.5")))
	log.Printf("find %+v\n", root.SearchIP(dict.NewStringIP("4.4.4.4")))
	log.Printf("find %+v\n", root.SearchIP(dict.NewStringIP("255.255.0.0")))
	log.Printf("find %+v\n", root.SearchIP(dict.NewStringIP("8.8.8.8")))

	log.Printf("sizeof(Node{}) %d\n", unsafe.Sizeof(dict.Node{}))

    //mk_mem_prof().Close()

    <-make(chan int)
}

