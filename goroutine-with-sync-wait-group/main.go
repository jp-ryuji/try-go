package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	// 開始時刻を記録
	startTime := time.Now()

	// foo
	loopNum := 3
	fooSleepTime := 2

	var wg sync.WaitGroup // WaitGroup を作成
	errorsCh := make(chan error, loopNum)

	fmt.Printf("foo is looping %d times\n", loopNum)

	for i := 0; i < loopNum; i++ {
		wg.Add(1)                // goroutine 開始時に wg.Add(1) を呼び出す
		go func(sleepTime int) { // goroutine を起動
			defer wg.Done() // goroutine 終了時に wg.Done() を呼び出す

			err := foo(sleepTime)
			if err != nil {
				errorsCh <- err
			}
		}(fooSleepTime)
	}

	wg.Wait() // 全ての goroutine が終了するまで待つ
	close(errorsCh)

	var errors []string
	for err := range errorsCh {
		errors = append(errors, err.Error())
	}

	if (len(errors)) > 0 {
		fmt.Printf("Errors: %v\n", strings.Join(errors, ", "))
	}

	// 経過時間を出力
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time taken: %s\n", elapsedTime)
}

func foo(sleepTime int) error {
	err := sleep("foo", sleepTime)
	if err != nil {
		return err
	}
	return nil
}

func sleep(fn string, sleepTime int) error {
	fmt.Printf("%s is sleeping for %d seconds\n", fn, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	return nil
}
