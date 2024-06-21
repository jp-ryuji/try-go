package main

import (
	"fmt"
	"time"
)

func main() {
	// 開始時刻を記録
	startTime := time.Now()

	// foo
	loopNum := 3
	fmt.Printf("foo is looping %d times\n", loopNum)
	for i := 0; i < loopNum; i++ {
		err := foo(2)
		if err != nil {
			fmt.Println(err)
			continue
		}
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
