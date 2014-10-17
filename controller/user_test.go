package controller

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestUserList(t *testing.T) {
	initDB()
	listUser := http.HandlerFunc(ListUser)
	test := GenerateHandleTester(t, listUser)
	w := test("GET", url.Values{})
	checkOKStatus(t, w)
	if !strings.Contains(w.Body.String(), "test@test.com") {
		t.Errorf("User list didn't return test user")
	}
}
