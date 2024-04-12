Auth interface{
	RegisterPusher(username string) (user Pusher)
	RegisterPuller(username string) (user Puller)
}

Broker interface{
	Push(username string, key string, data string) (status Status)
	Pull(username string, key string) (data string, status Status)
}

Pusher struct{
	username string
	key string
}

Puller struct{
	username string
	key string
}