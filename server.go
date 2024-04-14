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

}

func doRegisterPusher(username string)(typeDefs.AuthProcedure){
	user := typeDefs.Pusher{};
	return &typeDefs.AuthRegisterPusherReply{user};
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


func main(){
	log.Println("Key: %v", generateKey(10));
}
