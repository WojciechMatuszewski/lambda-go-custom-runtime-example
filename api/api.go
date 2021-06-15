package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type API struct {
	endpoint  string
	requestId string
}

func NewAPI(endpoint string) API {
	return API{endpoint: endpoint, requestId: ""}
}

func (a *API) Next() error {
	endpoint := fmt.Sprintf("%v/2018-06-01/runtime/invocation/next", a.endpoint)
	res, err := http.Get(endpoint)
	if err != nil {
		return errors.Wrap(err, "failed to get the next event")
	}
	defer res.Body.Close()

	requestId := res.Header.Get("Lambda-Runtime-Aws-Request-Id")
	if requestId == "" {
		panic(errors.New("Missing requestId from the response"))
	}

	a.requestId = requestId
	return nil
}

func (a API) Response(r io.Reader) error {
	if a.requestId == "" {
		return errors.New("missing requestId")
	}

	endpoint := fmt.Sprintf("%v/2018-06-01/runtime/invocation/%v/response", a.endpoint, a.requestId)
	res, err := http.Post(endpoint, "text/plain", r)
	if err != nil {
		return errors.Wrap(err, "failed to send the response")
	}
	defer res.Body.Close()

	return nil
}

func (a API) Error() {}
