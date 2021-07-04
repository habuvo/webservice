package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/habuvo/webservice/pkg/handlers"
	"github.com/habuvo/webservice/pkg/servers"
)

func TestRouter(t *testing.T) {
	r := servers.NewBaseRouter(handlers.NewBundle(nil))
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/create")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /create is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Post(ts.URL+"/create", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code for /create is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusMethodNotAllowed)
	}

	res, err = http.Get(ts.URL + "/not-exists")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code for /not-exist is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}
