#!/bin/bash
go build -ldflags="-s -w "  #-gcflags '-m -m'   #-a

#go tool pprof -http=":8081" iploc mem.prof


#  with noref build
#2019/09/12 18:00:55 loaded! count:2753541
# with ref build
#2019/09/12 18:03:33 loaded! count:2753541
# with no slice
#2019/09/13 00:01:33 loaded! count:2753541
# with indirect address
#2019/09/18 15:30:29 loaded! count:2753539


# GODEBUG=gctrace=1  ${BIN}



# GODEBUG=allocfreetrace=1 ${BIN}

