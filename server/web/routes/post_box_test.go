package routes

import (
	"testing"

	"github.com/sh3rp/databox/server/web/io"
)

func TestPostBoxSuccess(t *testing.T) {
	base := newRouterBase()
	token := authenticate(t, "admin", "password")
	test := &WSTest{
		Method:   "POST",
		Endpoint: "/box",
		Service:  base.PostBox,
		In: &io.NewBoxRequest{
			Token:       token,
			Name:        "testbox",
			Description: "Test box description",
			Password:    "password",
		},
		Tests: []func(*testing.T, *io.Response){
			testSuccess,
		},
	}
	runTest(t, test)
}

func TestPostBoxFieldFailure(t *testing.T) {

}
