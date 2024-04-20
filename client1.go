package main

import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/braidsAuthTypes"
	"ethos/braidsBrokerTypes"
	"ethos/defined"
	"log"
)

func init() {
	braidsAuthTypes.SetupAuthRegisterPusherReply(doRegisterPusherReply);
	braidsAuthTypes.SetupAuthRegisterPullerReply(doRegisterPullerReply);
	braidsBrokerTypes.SetupBrokerPushReply(doPushReply);
	braidsBrokerTypes.SetupBrokerPullReply(doPullReply);
}

func doRegisterPusherReply(user braidsAuthTypes.Pusher) (braidsAuthTypes.AuthProcedure){
	log.Println("Registered Pusher: ", user);
	return nil;
}

func doRegisterPullerReply(user braidsAuthTypes.Puller) (braidsAuthTypes.AuthProcedure){
	log.Println("Registered Puller: ", user);
	return nil;
}

func doPushReply(status syscall.Status) (braidsBrokerTypes.BrokerProcedure){
	log.Println("Status After Calling Push: ", status)
	return nil;
}

func doPullReply(msg braidsBrokerTypes.Message, status syscall.Status) (braidsBrokerTypes.BrokerProcedure){
	log.Println("PULL CALL REPLY", msg, status);
	return nil;
}

func callRpc(rpcType string, rpc defined.Rpc){
	fd, status := altEthos.IpcRepeat(rpcType, "", nil)
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
	callRpc("braidsAuthTypes", &braidsAuthTypes.AuthRegisterPusher{"karthik"});

	log.Println("CALLING PUSH")
	user := braidsBrokerTypes.Pusher{"karthik", "123"};
	callRpc("braidsBrokerTypes", &braidsBrokerTypes.BrokerPush{user, "TST"});

	log.Println("CALLING REG PULLER")
	callRpc("braidsAuthTypes", &braidsAuthTypes.AuthRegisterPuller{"karthik"});

	log.Println("CALLING PULL")
	user1 := braidsBrokerTypes.Puller{"karthik", "123"};
	callRpc("braidsBrokerTypes", &braidsBrokerTypes.BrokerPull{user1});

}
