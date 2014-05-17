package main

import (
	"fmt"
	"github.com/nutrun/lentil"
	"github.com/streadway/amqp"
	"labix.org/v2/mgo"
	"log"
	"strconv"
)

func main() {
	session, e := mgo.Dial("localhost")
	if e != nil {
		log.Fatal(e)
	}
	c := session.DB("test").C("grades")
	n, e := c.Count()
	if e != nil {
		log.Fatal("mongoDB: ", e)
	}

	fmt.Println("Number of grades: " + strconv.Itoa(n))

	connection, e := amqp.Dial("amqp://localhost")
	defer connection.Close()
	if e != nil {
		log.Fatal("rabbitMQ: ", e)
	}

	conn, e := lentil.Dial("localhost:11300")
	defer conn.Quit()
	if e != nil {
		log.Fatal("beanstalkd: ", e)
	}
	jobId, e := conn.Put(0, 0, 60, []byte("hello"))
	if e != nil {
		log.Fatal(e)
	}
	log.Printf("JOB ID: %d\n", jobId)

}
