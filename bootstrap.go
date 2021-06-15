package main

import (
	"bytes"
	"custom-runtime/api"
	"os"
)

func main() {
	runtimeAddr := os.Getenv("AWS_LAMBDA_RUNTIME_API")
	api := api.NewAPI("http://" + runtimeAddr)

	for {
		err := api.Next()
		if err != nil {
			panic(err)
		}

		out := handler()
		r := bytes.NewReader([]byte(out))
		err = api.Response(r)
		if err != nil {
			panic(err)
		}
	}
}

func handler() string {
	return "it works"
}
