package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/jijeshmohan/goseed/model"
)

var once sync.Once

type HandleTester func(
	method string,
	params url.Values,
) *httptest.ResponseRecorder

// Given the current test runner and an http.Handler, generate a
// HandleTester which will test its given input against the
// handler.

func GenerateHandleTester(
	t *testing.T,
	handleFunc http.Handler,
) HandleTester {

	// Given a method type ("GET", "POST", etc) and
	// parameters, serve the response against the handler and
	// return the ResponseRecorder.

	return func(
		method string,
		params url.Values,
	) *httptest.ResponseRecorder {

		req, err := http.NewRequest(
			method,
			"",
			strings.NewReader(params.Encode()),
		)
		if err != nil {
			t.Errorf("%v", err)
		}
		req.Header.Set(
			"Content-Type",
			"application/x-www-form-urlencoded; param=value",
		)
		w := httptest.NewRecorder()
		handleFunc.ServeHTTP(w, req)
		return w
	}
}

func initDB() {
	once.Do(func() {
		os.Remove("../test.db")
		model.InitDbWithName("../test.db")
		user := model.User{Name: "test", Email: "test@test.com"}
		model.CreateNewUser(&user)
	})
}

func checkOKStatus(t *testing.T, w *httptest.ResponseRecorder) {
	if w.Code != http.StatusOK {
		t.Errorf("Http status was %v", http.StatusText(w.Code))
	}
}
