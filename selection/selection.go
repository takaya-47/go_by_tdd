package selection

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	// 先にレスポンスが帰ってきたURLを返却すればいいのでselectを使う
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	// 監視対象のチャネルを返す前にゴルーチンを予約。
	go func() {
		http.Get(url)
		close(ch) // Racer側でどれか一つのチャネルが閉じられたらRacer自体の処理が終了してしまうので、ping側では常にチャネルを閉じる必要がある。チャネルに値を送信する処理としてしまうと、Racer側で処理が終了した後に受信する相手がいなくなってしまい、残っとゴルーチンがゾンビ化する。
	}()

	// 一旦チャネルを返し、selectで監視に入る。あとはゴルーチンの処理が完了してチャネルが閉じられるのを待つだけ。
	return ch
}
