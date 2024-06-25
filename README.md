# try-go

## Concurrency

[Share Memory By Communicating](https://go.dev/blog/codelab-share)

- Go’s concurrency primitives - `goroutines` and `channels`
  - Go encourages the use of channels to pass references to data between goroutines. This approach ensures that only one goroutine has access to the data at a given time.
- Do not communicate by sharing memory; instead, share memory by communicating.

### goroutines

go 文で関数を実行すると起動する。

### channels

チャネル( Channel )型は、チャネルオペレータの <- を用いて値の送受信ができる通り道です。\
ref: <https://go-tour-jp.appspot.com/concurrency/2>

```txt
ch <- v    // v をチャネル ch へ送信する
v := <-ch  // ch から受信した変数を v へ割り当てる

// 上記の例で、受信した変数を割り当てる必要が無ければ（同期を待つだけであれば）、以下でも良い。
<-ch
```

#### 戻り値として返す場合

戻り値として channel を返す場合は、`<-chan string` （読み取り専用の channel）という形式とする。

```go
func foo() <-chan string {
  ch := make(chan string)
  :
  return ch
}
```
