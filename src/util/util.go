package util

import (
	"fmt"
	"math/rand"
	"time"
)

// ProcessError ...
func ProcessError(err error) {
	if err != nil {
		panic(err)
		//log.Fatalln(err)
	}
}

// RecoverError ...
func RecoverError() {
	err := recover()
	fmt.Println(err)
}

//GetRandomName ...
func GetRandomName() string {
	alpha := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "z", "y", "z",
	}

	seedValue := time.Now().Unix()
	rand.Seed(seedValue)

	name := ""
	for i := 0; i < 10; i++ {
		name = name + alpha[rand.Intn(len(alpha))]
	}

	return name
}

// GetRandomUIntNumber ...
func GetRandomUIntNumber(maxNumber int32) int32 {
	seedValue := time.Now().Unix()
	rand.Seed(seedValue)
	return rand.Int31n(maxNumber + 1)
}
