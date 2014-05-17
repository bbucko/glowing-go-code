package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"strconv"
	"github.com/pebbe/zmq4"
)

func main() {
	fmt.Printf("Hello, world.\n")

	session, _ := mgo.Dial("localhost")
	c := session.DB("test").C("grades")
	n, _ := c.Count()

	fmt.Println("Number of grades: " + strconv.Itoa(n))

	server, _ := zmq4.NewSocket(zmq4.PUSH)
	server.ServerAuthPlain("global")
	server.Bind("tcp://*:9000")
}
