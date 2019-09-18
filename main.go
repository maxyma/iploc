package main

import (
//    "fmt"
//    "runtime"
    "log"
    "unsafe"
)

func main(){

    //runtime.LockOSThread()

    log.Println("loading...")

    //runtime.MemProfileRate = 1

    //mk_cpu_prof()

    root := load_dict()

    //stop_cpu_pro()

    log.Printf("loaded! count:%d\n", root.Count())
    //spew.Dump(root.child)

    //log.Printf("deepRight: %v\n", root.deepRight())


	log.Printf("%+v\n", root.SearchIP(FromStringToIP("0.0.0.0")))
	//log.Printf("%+v\n", root.SearchIP(FromStringToIP("1.2.3.5")))
	//log.Printf("%+v\n", root.SearchIP(FromStringToIP("1.2.3.4")))
	//log.Printf("%+v\n", root.SearchIP(FromStringToIP("1.2.3.5")))
	log.Printf("%+v\n", root.SearchIP(FromStringToIP("1.2.3.8")))
	log.Printf("%+v\n", root.SearchIP(FromStringToIP("255.255.0.0")))
	log.Printf("%+v\n", root.SearchIP(FromStringToIP("8.8.8.8")))
	log.Printf("%d\n", unsafe.Sizeof(Node{}))
	//log.Printf("%+v\n", root.SearchIP(FromStringToIP("1.2.3.3")))


    //mk_mem_prof().Close()

    <-make(chan int)
}

