package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/braidsAuthTypes"
	"log"
)

var masterQueue []braidsAuthTypes.Message;

func init(){
	log.Println("SERVER INIT")
	pusherPath := "/etc/braidsPushers"
	pullerPath := "/etc/braidsPullers"

	pusher := braidsAuthTypes.Pusher{};
	puller := braidsAuthTypes.Puller{};

	status := altEthos.DirectoryCreate(pusherPath, &pusher, "all");
	if status != syscall.StatusOk && status != syscall.StatusExists {
		log.Fatalf("Error could not create %v %v\n", pusherPath, status);
	}
	log.Println("CREATE DIR1")
	status = altEthos.DirectoryCreate(pullerPath, &puller, "all");
	if status != syscall.StatusOk && status != syscall.StatusExists {
		log.Fatalf("Error could not create %v %v\n", pullerPath, status);
	}

	log.Println("CREATE DIR2")
	// Setup RPC
	braidsAuthTypes.SetupAuthRegisterPusher(doRegisterPusher);
	braidsAuthTypes.SetupAuthRegisterPuller(doRegisterPuller);

	log.Println("SETUP")
}

func doRegisterPusher(username string)(braidsAuthTypes.AuthProcedure){
	user := braidsAuthTypes.Pusher{};
	user.Username = username;
	user.Key = "123" //generateKey(40);
	filePath := "/etc/braidsPushers/" + username;
	status := altEthos.Write(filePath, &user);
	if status != syscall.StatusOk {
		log.Fatalf("Error writing auth for %v | status: %v", username, status);
	}
	return &braidsAuthTypes.AuthRegisterPusherReply{user};
}

func doRegisterPuller(username string)(braidsAuthTypes.AuthProcedure){
	user := braidsAuthTypes.Puller{};
	user.Username = username;
	user.Key = "123" //generateKey(40);
	filePath := "/etc/braidsPullers/" + username;
	status := altEthos.Write(filePath, &user);
	if status != syscall.StatusOk {
		log.Fatalf("Error writing auth for %v | status: %v", username, status);
	}
	return &braidsAuthTypes.AuthRegisterPullerReply{user};
}
/*
func doPush(user braidsAuthTypes.Pusher, data string)  (braidsAuthTypes.BrokerProcedure) {
	if(verifyPusher(user)){
		msg := braidsAuthTypes.Message{Id: generateKey(5), Data: data, CreatedAt: int64(syscall.GetTime())};
		masterQueue = append(masterQueue, msg);
		return &braidsAuthTypes.BrokerPushReply{syscall.StatusOk};
	}

	return &braidsAuthTypes.BrokerPushReply{syscall.StatusInvalidAuthentication};
}

func doPull(user braidsAuthTypes.Puller) (braidsAuthTypes.BrokerProcedure) {
	if(verifyPuller(user)){
		if len(masterQueue) == 0 {
			return &braidsAuthTypes.BrokerPullReply{braidsAuthTypes.Message{}, syscall.StatusInvalidLength};
		}
		msg := masterQueue[0];
		masterQueue = masterQueue[1:];
		return &braidsAuthTypes.BrokerPullReply{msg, syscall.StatusOk};
	}

	return &braidsAuthTypes.BrokerPullReply{braidsAuthTypes.Message{}, syscall.StatusInvalidAuthentication};
}
*/
func generateKey(length int) string{
	log.Print("Generating Key.....")
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

func verifyPusher(user braidsAuthTypes.Pusher) bool{
	storedUser := braidsAuthTypes.Pusher{};
	filePath := "/etc/braidsPushers/"+user.Username;
	status := altEthos.Read(filePath, &storedUser);
	if status != syscall.StatusOk {
		return false;
	}
	if(storedUser.Key != user.Key){
		return false;
	}
	return true;
}

func verifyPuller(user braidsAuthTypes.Puller) bool{
	storedUser := braidsAuthTypes.Puller{};
	filePath := "/etc/braidsPullers/"+user.Username;
	status := altEthos.Read(filePath, &storedUser);
	if status != syscall.StatusOk {
		return false;
	}
	if(storedUser.Key != user.Key){
		return false;
	}
	return true;
}

func main(){
	altEthos.LogToDirectory("application/braidsAuth")
	log.Println("Before Adv");
	listeningFd, status := altEthos.Advertise("braidsAuthTypes")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}
	log.Println("Adv Success")
	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted")

		authHandler := braidsAuthTypes.Auth{}
		altEthos.Handle(fd, &authHandler)
	}
}
