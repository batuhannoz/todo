package main

import "log"

func main() {
	err := StartServer(3000)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Provider Service Listening :3000")
}
