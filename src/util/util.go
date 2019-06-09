package util

import "fmt"

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
