package routes

import (
	"testing"

	"github.com/sh3rp/databox/server/web/io"
)

func TestPostLinkSuccess(t *testing.T) {
	base := newRouterBase()
	token := authenticate(t, "admin", "password")
	test := &WSTest{
		Method:   "POST",
		Endpoint: "/box/:id/link",
		Service:  base.PostLink,
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

func TestPostLinkFieldFailure(t *testing.T) {

}
