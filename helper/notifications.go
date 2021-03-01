package helper

import (
	"fmt"
)

func errFatalNotification(err error) {
	if err != nil {
		panic(err)
	}
}

func errNotification(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func stringNotification(s string) {
	fmt.Println(s)
}
