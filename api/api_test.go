package api_test

import (
	"custom-runtime/api"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewApi(t *testing.T) {
	t.Run("Fetches the invocation event", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/2018-06-01/runtime/invocation/next", func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Lambda-Runtime-Aws-Request-Id", "1")

			fmt.Fprintln(rw, "Hello")
		})

		srv := httptest.NewServer(mux)
		defer srv.Close()

		api := api.NewAPI(srv.URL)

		err := api.Next()
		if err != nil {
			t.Fatal("not expecting error, got", err)
		}
	})

	t.Run("Fails if the requestId header is not set", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/2018-06-01/runtime/invocation/next", func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Set("Foo", "1")

			fmt.Fprintln(rw, "Hello")
		})

		srv := httptest.NewServer(mux)
		defer srv.Close()

		api := api.NewAPI(srv.URL)

		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("Expected panic, did not")
			}
		}()

		api.Next()
	})
}
