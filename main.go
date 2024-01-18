package main

import "github.com/DiegoSantos95/UsersExampleGin/server"

func main() {
	server := server.New()
	server.Start()
}
