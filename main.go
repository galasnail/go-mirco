package gomicro

func main() {

	server = NewServer()
	server.RegisterName(serviceName, service)
	server.Start("tcp", "127.0.0.1:0")
	serverAddr = server.Address()


}

