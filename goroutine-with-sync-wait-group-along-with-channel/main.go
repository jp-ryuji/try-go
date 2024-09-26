package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	errChan := make(chan error, 2) // エラーチャネルを用意

	var fnOneRes *fnOneResult
	var fnTwoRes *fnTwoResult
	var err error

	fnOneSleepTime := 2

	wg.Add(1)
	go func(sleepTime int) {
		defer wg.Done()
		fnOneRes, err = fnOne(sleepTime)
		if err != nil {
			errChan <- err
			return
		}
	}(fnOneSleepTime)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fnTwoRes, err = fnTwo(5)
		if err != nil {
			errChan <- err
			return
		}
	}()

	wg.Wait()      // すべてのgoroutineの完了を待機
	close(errChan) // エラーチャネルをクローズ

	// エラーチャネルからエラーをチェック
	for err := range errChan {
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println(fnOneRes.result)
	fmt.Println(fnTwoRes.result)
}

type fnOneResult struct {
	result string
}

func fnOne(sleepTime int) (*fnOneResult, error) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return &fnOneResult{result: "fnOne"}, nil
}

type fnTwoResult struct {
	result string
}

func fnTwo(sleepTime int) (*fnTwoResult, error) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	return &fnTwoResult{result: "fnTwo"}, nil
}
