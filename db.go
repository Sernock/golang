package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func dbCall(i int) {
	delay := rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from database is:", dbData[i])
}
