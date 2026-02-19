package selection

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	t.Run("returns faster url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		// 可読性向上のためサーバー生成処理の直後に書くのが良い
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)
		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	fn := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}
	handler := http.HandlerFunc(fn) // fnをhttp.HandlerFuncという関数型にキャスト
	return httptest.NewServer(handler)
}
