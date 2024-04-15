Auth interface{
	RegisterPusher(username string) (user Pusher)
	RegisterPuller(username string) (user Puller)
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
	CreatedAt int64
}
