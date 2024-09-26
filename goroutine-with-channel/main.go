package main

import (
	"fmt"
	"time"
)

type Result struct {
	Result string
	Err    error
}

// foo と bar という別々の処理（中身は似ているが）が同時に実行される。
// それぞれ channel に結果を送信し、両方の受信を待って先に進む。
func main() {
	// 開始時刻を記録
	startTime := time.Now()

	// foo
	fooCh := make(chan Result)
	fooSleepTime := 2

	go func(sleepTime int) {
		result, err := foo(sleepTime)
		fooCh <- Result{result, err}
	}(fooSleepTime)

	// bar
	barCh := make(chan Result)
	barSleepTime := 5

	go func(sleepTime int) {
		result, err := bar(sleepTime)
		barCh <- Result{result, err}
	}(barSleepTime)

	// 結果を受け取る
	fooResult, barResult := <-fooCh, <-barCh
	if fooResult.Err != nil {
		fmt.Println(fooResult.Err)
	}
	if barResult.Err != nil {
		fmt.Println(barResult.Err)
	}

	// 経過時間を出力
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total time taken: %s\n", elapsedTime)
}

func foo(sleepTime int) (string, error) {
	awake, err := sleep("foo", sleepTime)
	if err != nil {
		return "", err
	}
	return awake, nil
}

func bar(sleepTime int) (string, error) {
	awake, err := sleep("bar", sleepTime)
	if err != nil {
		return "", err
	}
	return awake, nil
}

func sleep(fn string, sleepTime int) (string, error) {
	fmt.Printf("%s is sleeping for %d seconds\n", fn, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)

	return fmt.Sprintf("%s is awake", fn), nil
}
