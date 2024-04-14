Auth interface{
	RegisterPusher(username string) (user Pusher)
	RegisterPuller(username string) (user Puller)
}

Broker interface{
	Push(username string, key string, data string) (status Status)
	Pull(username string, key string) (data string, status Status)
}

Pusher struct{
	Username string
	Key string
}

Puller struct{
	Username string
	Key string
}

Message struct{
	Id string
	Data string
	CreatedAt string
}
