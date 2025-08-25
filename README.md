# try-go

## 並行性と並列性

並行性と並列性は関連しているが異なる概念です。

### 並行性 (Concurrency)

- **並行性**は一度に複数のことを処理することです。
- 構造に関する概念で、複数のタスクを処理するためのプログラム設計方法です。
- Goでは、ゴルーチンとチャネルによって並行性を実現します。
- 並行性は必ずしも並列実行を意味するわけではありません。1つのCPUコアでタスク切り替えを行う場合もあります。

### 並列性 (Parallelism)

- **並列性**は複数のことを同時に実行することです。
- 実行に関する概念で、複数のタスクを同時に実行することです。
- 並列性には複数の処理ユニット（コア）が必要です。
- すべての並列プログラムは並行性を持ちますが、すべての並行プログラムが並列であるとは限りません。

### Goにおける並行性と並列性

Goではゴルーチンによって並行性を実現し、複数のCPUコアが利用可能であれば並列に実行されます。Goランタイムのスケジューラがゴルーチンを管理し、利用可能なプロセッサコアに配分します。

## Concurrency

[Share Memory By Communicating](https://go.dev/blog/codelab-share)

- Go's concurrency primitives - `goroutines` and `channels`
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
