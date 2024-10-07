package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type singleton struct{}

var singletonInstance *singleton

func getInstance() *singleton {
	if singletonInstance == nil {
		// double check locking
		lock.Lock()
		defer lock.Unlock()

		if singletonInstance == nil {
			fmt.Println("new instance")
			singletonInstance = &singleton{}
		}

		fmt.Println("old instance")
	}

	return singletonInstance
}

// prototype pattern is about creating one, and only one instance of a struct
// useful in creating database connections
