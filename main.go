package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"strconv"
	"github.com/streadway/amqp"
	"github.com/nutrun/lentil"
	"log"
)

func main() {
	fmt.Printf("Hello, world.\n")

	session, _ := mgo.Dial("localhost")
	c := session.DB("test").C("grades")
	n, _ := c.Count()

	fmt.Println("Number of grades: " + strconv.Itoa(n))

	connection, _ := amqp.Dial("localhost")
	defer connection.Close()

	conn, e := lentil.Dial("0.0.0.0:11300")
	if e != nil {
		log.Fatal(e)
	}
	jobId, e := conn.Put(0, 0, 60, []byte("hello"))
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("JOB ID: %d\n", jobId)

}
