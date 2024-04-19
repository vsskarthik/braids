package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/typeDefs"
	"log"
)

var masterQueue []typeDefs.Message;

func init(){
	pusherPath := "/etc/braidsPushers"
	pullerPath := "/etc/braidsPullers"

	pusher := typeDefs.Pusher{};
	puller := typeDefs.Puller{};

	status := altEthos.DirectoryCreate(pusherPath, &pusher, "all");
	if status != syscall.StatusOk && status != syscall.StatusExists {
		log.Fatalf("Error could not create %v %v\n", pusherPath, status);
	}

	status = altEthos.DirectoryCreate(pullerPath, &puller, "all");
	if status != syscall.StatusOk && status != syscall.StatusExists {
		log.Fatalf("Error could not create %v %v\n", pullerPath, status);
	}

	// Setup RPC
	typeDefs.SetupAuthRegisterPusher(doRegisterPusher);
	typeDefs.SetupAuthRegisterPuller(doRegisterPuller);
	typeDefs.SetupBrokerPush(doPush);
	typeDefs.SetupBrokerPull(doPull);

}

func doRegisterPusher(username string)(typeDefs.AuthProcedure){
	user := typeDefs.Pusher{};
	user.Username = username;
	user.Key = "123" //generateKey(40);
	filePath := "/etc/braidsPushers/" + username;
	status := altEthos.Write(filePath, &user);
	if status != syscall.StatusOk {
		log.Fatalf("Error writing auth for %v | status: %v", username, status);
	}
	return &typeDefs.AuthRegisterPusherReply{user};
}

func doRegisterPuller(username string)(typeDefs.AuthProcedure){
	user := typeDefs.Puller{};
	user.Username = username;
	user.Key = "123" //generateKey(40);
	filePath := "/etc/braidsPullers/" + username;
	status := altEthos.Write(filePath, &user);
	if status != syscall.StatusOk {
		log.Fatalf("Error writing auth for %v | status: %v", username, status);
	}
	return &typeDefs.AuthRegisterPullerReply{user};
}

func doPush(user typeDefs.Pusher, data string)  (typeDefs.BrokerProcedure) {
	if(verifyPusher(user)){
		msg := typeDefs.Message{Id: generateKey(5), Data: data, CreatedAt: int64(syscall.GetTime())};
		masterQueue = append(masterQueue, msg);
		return &typeDefs.BrokerPushReply{syscall.StatusOk};
	}

	return &typeDefs.BrokerPushReply{syscall.StatusInvalidAuthentication};
}

func doPull(user typeDefs.Puller) (typeDefs.BrokerProcedure) {
	if(verifyPuller(user)){
		if len(masterQueue) == 0 {
			return &typeDefs.BrokerPullReply{typeDefs.Message{}, syscall.StatusInvalidLength};
		}
		msg := masterQueue[0];
		masterQueue = masterQueue[1:];
		return &typeDefs.BrokerPullReply{msg, syscall.StatusOk};
	}

	return &typeDefs.BrokerPullReply{typeDefs.Message{}, syscall.StatusInvalidAuthentication};
}

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

func verifyPusher(user typeDefs.Pusher) bool{
	storedUser := typeDefs.Pusher{};
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

func verifyPuller(user typeDefs.Puller) bool{
	storedUser := typeDefs.Puller{};
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
	altEthos.LogToDirectory("application/braids")

	listeningFd, status := altEthos.Advertise("typeDefs")
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

		log.Println("new connection accepted")

		authHandler := typeDefs.Auth{}
		altEthos.Handle(fd, &authHandler)
	}
}
