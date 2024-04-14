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

}


func main(){}
