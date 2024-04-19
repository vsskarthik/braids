package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/braidsAuthTypes"
	"log"
)

var masterQueue []braidsAuthTypes.Message;

func init(){
	pusherPath := "/etc/braidsPushers"
	pullerPath := "/etc/braidsPullers"

	pusher := braidsAuthTypes.Pusher{};
	puller := braidsAuthTypes.Puller{};

	status := altEthos.DirectoryCreate(pusherPath, &pusher, "all");
	if status != syscall.StatusOk && status != syscall.StatusExists {
		log.Fatalf("Error could not create %v %v\n", pusherPath, status);
	}
	status = altEthos.DirectoryCreate(pullerPath, &puller, "all");
	if status != syscall.StatusOk && status != syscall.StatusExists {
		log.Fatalf("Error could not create %v %v\n", pullerPath, status);
	}

	// Setup RPC
	braidsAuthTypes.SetupAuthRegisterPusher(doRegisterPusher);
	braidsAuthTypes.SetupAuthRegisterPuller(doRegisterPuller);
	braidsAuthTypes.SetupAuthVerifyPusher(doVerifyPusher);
	braidsAuthTypes.SetupAuthVerifyPuller(doVerifyPuller);
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

func doVerifyPusher(user braidsAuthTypes.Pusher) (braidsAuthTypes.AuthProcedure){
	storedUser := braidsAuthTypes.Pusher{};
	filePath := "/etc/braidsPushers/"+user.Username;
	status := altEthos.Read(filePath, &storedUser);
	if status != syscall.StatusOk {
		return &braidsAuthTypes.AuthVerifyPusherReply{syscall.StatusInvalidAuthentication};
	}
	if(storedUser.Key != user.Key){
		return &braidsAuthTypes.AuthVerifyPusherReply{syscall.StatusInvalidAuthentication};
	}
	return &braidsAuthTypes.AuthVerifyPusherReply{syscall.StatusOk};
}

func doVerifyPuller(user braidsAuthTypes.Puller) (braidsAuthTypes.AuthProcedure){
	storedUser := braidsAuthTypes.Puller{};
	filePath := "/etc/braidsPullers/"+user.Username;
	status := altEthos.Read(filePath, &storedUser);
	if status != syscall.StatusOk {
		return &braidsAuthTypes.AuthVerifyPullerReply{syscall.StatusInvalidAuthentication};
	}
	if(storedUser.Key != user.Key){
		return &braidsAuthTypes.AuthVerifyPullerReply{syscall.StatusInvalidAuthentication};
	}
	return &braidsAuthTypes.AuthVerifyPullerReply{syscall.StatusOk};
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
