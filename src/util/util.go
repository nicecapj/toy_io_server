package util

import (
	"fmt"
	"math/rand"
	"time"
)

func ProcessError(err error) {
	if err != nil {
		panic(err)
		//log.Fatalln(err)
	}
}

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
