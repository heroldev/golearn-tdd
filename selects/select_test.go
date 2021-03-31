package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(test *testing.T) {

	test.Run("test the racer", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q but wanted %q", got, want)
		}

	})

	test.Run("test post 10 second halt", func(t *testing.T) {

		server := makeDelayedServer(1 * time.Second)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 5*time.Millisecond)

		if err == nil {
			t.Error("wanted error but didn't get one")
		}

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
