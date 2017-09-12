package routes

import (
	"encoding/json"
	"testing"

	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/server/web/io"
	"github.com/stretchr/testify/assert"
)

func TestPostAuthSuccess(t *testing.T) {
	base := newRouterBase()
	test := &WSTest{
		Method:   "POST",
		Endpoint: "/auth",
		Service:  base.PostAuth,
		In: &msg.AuthRequest{
			Username: "admin",
			Password: "password",
		},
		Tests: []func(*testing.T, *io.Response){
			testSuccess,
			func(t *testing.T, res *io.Response) {
				token := &msg.Token{}
				data, err := json.Marshal(res.Data)
				assert.Nil(t, err)
				err = json.Unmarshal(data, token)
				assert.Nil(t, err)
				assert.NotNil(t, token)
			},
		},
	}
	runTest(t, test)
}

func TestPostAuthUsernameFailure(t *testing.T) {
	base := newRouterBase()
	test := &WSTest{
		Method:   "POST",
		Endpoint: "/auth",
		Service:  base.PostAuth,
		In: &msg.AuthRequest{
			Username: "asdf",
			Password: "password",
		},
		Tests: []func(*testing.T, *io.Response){
			func(t *testing.T, res *io.Response) {
				assert.Equal(t, common.E_IO_INVALID_LOGIN, res.Code)
				assert.Nil(t, res.Data)
			},
		},
	}
	runTest(t, test)
}

func TestPostAuthPasswordFailure(t *testing.T) {
	base := newRouterBase()
	test := &WSTest{
		Method:   "POST",
		Endpoint: "/auth",
		Service:  base.PostAuth,
		In: &msg.AuthRequest{
			Username: "admin",
			Password: "asdf",
		},
		Tests: []func(*testing.T, *io.Response){
			func(t *testing.T, res *io.Response) {
				assert.Equal(t, common.E_IO_INVALID_LOGIN, res.Code)
				assert.Nil(t, res.Data)
			},
		},
	}
	runTest(t, test)
}
