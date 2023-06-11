package main

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/createUser", PostUser)
	server.Handle("POST", "/api", server.AddMiddleware(HandleApi, CheckAuth(), Logger()))
	server.Listen()
}
