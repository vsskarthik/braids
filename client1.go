package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/braidsAuthTypes"
	"ethos/defined"
	"log"
)

func init() {
	braidsAuthTypes.SetupAuthRegisterPusherReply(doRegisterPusherReply);
}

func doRegisterPusherReply(user braidsAuthTypes.Pusher) (braidsAuthTypes.AuthProcedure){
	log.Println("Registered Pusher: ", user);
	return nil;
}

func callRpc(rpc defined.Rpc){
	fd, status := altEthos.IpcRepeat("braidsAuthTypes", "", nil)
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
	altEthos.LogToDirectory("test/braidsClient");
	log.Println("INIT CLIENT")

	log.Println("CALLING REG PUSHER")
	callRpc(&braidsAuthTypes.AuthRegisterPusher{"karthik"});




}
