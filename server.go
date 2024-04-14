package main

import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/typeDefs"
	"log"
)

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
	user.Key = generateKey(40);
	filePath := "/etc/braidsPullers/" + username;
	status := altEthos.Write(filePath, &user);
	if status != syscall.StatusOk {
		log.Fatalf("Error writing auth for %v | status: %v", username, status);
	}
	return &typeDefs.AuthRegisterPullerReply{user};
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
	log.Printf("status %v storeduser %v", status, storedUser);
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
	log.Printf("status %v storeduser %v", status, storedUser);
	if status != syscall.StatusOk {
		return false;
	}
	if(storedUser.Key != user.Key){
		return false;
	}
	return true;
}

func main(){
	user := typeDefs.Puller{Username: "karthik", Key: "123"}
	log.Printf("pusher reg: ", doRegisterPuller("karthik"));
	log.Printf("pusher verify: ", verifyPuller(user));
	// doRegisterPuller("karthik");
}
