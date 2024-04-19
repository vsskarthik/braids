Broker interface{
	Push(user Pusher, data string) (status Status)
	Pull(user Puller) (msg Message, status Status)
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
