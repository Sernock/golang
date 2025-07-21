package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

// TODO

func main() {
	f, err := randomFLoat(4, 6)
	if err != nil {
		log.Fatalf("Ошибка генерации случайного значения: %s", err.Error())
	}

	fmt.Println("Случайное число", f)
}

func randomFLoat(min, max float64) (float64, error) {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return 0, err
	}

	num := binary.BigEndian.Uint64(b)

	f := float64(num) / float64(math.MaxUint64)

	return min + f*(max-min), nil
}
