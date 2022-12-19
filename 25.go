//Реализовать собственную функцию sleep.
package main

import (
	"log"
	"time"
)

func main() {
	log.Println("До Sleep")
	sleep(time.Second * 2)
	log.Println("После Sleep")
}

func sleep(t time.Duration) {
	<-time.After(t)
}
