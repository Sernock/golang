package main

import (
	"time"
)

func main() {
	for i := 0; i < len(dbData); i++ {
		go dbCall(i)
	}
	time.Sleep(3 * time.Second)
}
