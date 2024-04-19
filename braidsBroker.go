package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/braidsAuthTypes"
	"ethos/braidsBrokerTypes"
	"ethos/defined"
	"log"
)

var masterQueue []braidsBrokerTypes.Message;
var authStatus syscall.Status;

func init(){
	braidsBrokerTypes.SetupBrokerPush(doPush);
	braidsBrokerTypes.SetupBrokerPull(doPull);
	braidsAuthTypes.SetupAuthVerifyPusherReply(doAfterVerifyPusher);
	braidsAuthTypes.SetupAuthVerifyPullerReply(doAfterVerifyPuller);
}

func doAfterVerifyPusher(status syscall.Status) (braidsAuthTypes.AuthProcedure) {
	authStatus = status;
	return nil;
}

func doAfterVerifyPuller(status syscall.Status) (braidsAuthTypes.AuthProcedure) {
	authStatus = status;
	return nil;
}

func doPush(user braidsBrokerTypes.Pusher, data string)  (braidsBrokerTypes.BrokerProcedure) {
	authUser := braidsAuthTypes.Pusher{ user.Username,  user.Key}
	callRpc(&braidsAuthTypes.AuthVerifyPusher{authUser});
	log.Println("PUSHER_AUTH_VERIFY_CALL ", authUser, " STATUS: ", authStatus);
	if(authStatus == syscall.StatusOk){
		msg := braidsBrokerTypes.Message{Id: generateKey(5), Data: data, CreatedAt: int64(syscall.GetTime())};
		masterQueue = append(masterQueue, msg);
		log.Println("PUSH_QUEUE_STATE ", masterQueue);
		return &braidsBrokerTypes.BrokerPushReply{syscall.StatusOk};
	}
	return &braidsBrokerTypes.BrokerPushReply{syscall.StatusInvalidAuthentication};
}

func doPull(user braidsBrokerTypes.Puller) (braidsBrokerTypes.BrokerProcedure) {
	authUser := braidsAuthTypes.Puller{ user.Username,  user.Key}
	callRpc(&braidsAuthTypes.AuthVerifyPuller{authUser});
	log.Println("PULLER_AUTH_VERIFY_CALL ", authUser, " STATUS: ", authStatus);
	if(authStatus == syscall.StatusOk){
		if len(masterQueue) == 0 {
			return &braidsBrokerTypes.BrokerPullReply{braidsBrokerTypes.Message{}, syscall.StatusInvalidLength};
		}
		msg := masterQueue[0];
		masterQueue = masterQueue[1:];
		log.Println("PULL_QUEUE_STATE ", masterQueue);
		return &braidsBrokerTypes.BrokerPullReply{msg, syscall.StatusOk};
	}

	return &braidsBrokerTypes.BrokerPullReply{braidsBrokerTypes.Message{}, syscall.StatusInvalidAuthentication};
}

func generateKey(length int) string{
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomString := make([]byte, length)
	seed := syscall.GetTime();
	for i:=0; i<length; i++{
		seed += (seed * 1103515245 + 12345) % (1 << 31)
		index := int(int64(seed) % int64(len(charset)));
		randomString[i] = charset[index]
	}
	log.Println("Done")
	return string(randomString)
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
	listeningFd, status := altEthos.Advertise("braidsBrokerTypes")
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


		brokerHandler := braidsBrokerTypes.Broker{}
		altEthos.Handle(fd, &brokerHandler)
	}
}
