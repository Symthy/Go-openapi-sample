package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mtslzr/pokeapi-go/structs"
	"github.com/stretchr/testify/assert"
)

func TestSuccessToGetPokemon(t *testing.T) {

	raw, err := ioutil.ReadFile("./_test/306.json")
	if err != nil {
		t.Error(err)
	}
	expected := &structs.Pokemon{}
	err = json.Unmarshal(raw, expected)
	if err != nil {
		t.Error(err)
	}

	actual, err := GetPokemon("306")
	assert.NoError(t, err)
	assert.EqualValues(t, expected, actual)
}

func TestFailureToGet(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		resp := make(map[string]string)
		resp["message"] = "Bad Request"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
	}))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, strings.NewReader(""))
	if err != nil {
		assert.FailNow(t, "failed to build request")
	}
	actual, err := get(req)
	assert.EqualError(t, err, "Error Response. Status:400 Bad Request, Body:{\"message\":\"Bad Request\"}")
	assert.Nil(t, actual)
}

func TestHttpMockSample(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	got, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()

	if res.StatusCode != 200 {
		t.Errorf("GET %s: expected status code = %d; got %d", ts.URL, 200, res.StatusCode)
	}
	if string(got) != "Hello, client\n" {
		t.Errorf("expected body %v; got %v", "Hello, client", string(got))
	}
}
