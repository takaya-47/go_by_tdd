package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond) // 20ミリ秒あえて待つことで遅いサーバーをシミュレート
		w.WriteHeader(http.StatusOK)
	}
	handler := http.HandlerFunc(fn) // fnをHandlerFuncという関数型にキャスト
    slowServer := httptest.NewServer(handler)

	fn = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
	fastServer := httptest.NewServer(http.HandlerFunc(fn))

	slowUrl := slowServer.URL
	fastUrl := fastServer.URL

	want := fastUrl
	got := Racer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}