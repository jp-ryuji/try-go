package main

import (
	"fmt"
	"time"
)

type FnOneResult struct {
	Result string
}

type FnTwoResult struct {
	Result string
}

func main() {
	errChan := make(chan error, 2)

	var fnOneResult *FnOneResult
	var fnTwoResult *FnTwoResult
	var err error

	go func() {
		fnOneResult, err = fnOne(2)
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	go func() {
		fnTwoResult, err = fnTwo(4)
		if err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(fnOneResult.Result)
	fmt.Println(fnTwoResult.Result)
}

func fnOne(sleepTime int) (*FnOneResult, error) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return &FnOneResult{Result: "fnOne"}, nil
}

func fnTwo(sleepTime int) (*FnTwoResult, error) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return &FnTwoResult{Result: "fnTwo"}, nil
}
