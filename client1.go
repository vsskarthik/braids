package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/typeDefs"
	"ethos/defined"
	"log"
)

func init() {

}


func callRpc(rpc defined.Rpc){
	fd, status := altEthos.IpcRepeat("typeDefs", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	status = altEthos.ClientCall(fd, rpc)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}
}


func main(){
	altEthos.LogToDirectory




}
