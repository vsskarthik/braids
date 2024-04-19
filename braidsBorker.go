package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/braidsAuthTypes"
	"ethos/braidsBorkerTypes"
	"log"
)

var masterQueue []braidsAuthTypes.Message;

func init(){
	braidsBrokerTypes.SetupBrokerPush(doPush);
	braidsBrokerTypes.SetupBrokerPull(doPull);
}

func doPush(user braidsBrokerTypes.Pusher, data string)  (braidsBrokerTypes.BrokerProcedure) {
	/*if(verifyPusher(user)){
		msg := braidsAuthTypes.Message{Id: generateKey(5), Data: data, CreatedAt: int64(syscall.GetTime())};
		masterQueue = append(masterQueue, msg);
		return &braidsAuthTypes.BrokerPushReply{syscall.StatusOk};
	}*/

	return &braidsBrokerTypes.BrokerPushReply{syscall.StatusInvalidAuthentication};
}

func doPull(user braidsBrokerTypes.Puller) (braidsBrokerTypes.BrokerProcedure) {
	/*if(verifyPuller(user)){
		if len(masterQueue) == 0 {
			return &braidsAuthTypes.BrokerPullReply{braidsAuthTypes.Message{}, syscall.StatusInvalidLength};
		}
		msg := masterQueue[0];
		masterQueue = masterQueue[1:];
		return &braidsAuthTypes.BrokerPullReply{msg, syscall.StatusOk};
	}*/

	return &braidsAuthTypes.BrokerPullReply{braidsAuthTypes.Message{}, syscall.StatusInvalidAuthentication};
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
	altEthos.LogToDirectory("application/braidsBroker")
	listeningFd, status := altEthos.Advertise("braidsBorkerType")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}
	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("[BROKER] new connection accepted")

		authHandler := braidsAuthTypes.Auth{}
		altEthos.Handle(fd, &authHandler)
	}
}
