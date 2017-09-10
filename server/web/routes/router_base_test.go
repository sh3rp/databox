package routes

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/server/web/io"
	"github.com/stretchr/testify/assert"
)

type WSTest struct {
	Method   string
	Endpoint string
	Service  func(*gin.Context)
	In       interface{}
	Tests    []func(*testing.T, *io.Response)
}

func runTest(t *testing.T, test *WSTest) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	switch test.Method {
	case "GET":
		r.GET(test.Endpoint, test.Service)
	case "POST":
		r.POST(test.Endpoint, test.Service)
	}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(test.In)
	req := httptest.NewRequest(test.Method, test.Endpoint, bytes.NewReader(buf.Bytes()))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	response := &io.Response{}
	json.NewDecoder(res.Body).Decode(response)
	for _, test := range test.Tests {
		test(t, response)
	}
}

func newRouterBase() *RouterBase {
	base := &RouterBase{
		Auth:       auth.NewInMemoryAuthenticator(),
		TokenStore: auth.NewInMemoryTokenStore(1),
		DB:         db.NewInMemoryDB(),
		Search:     search.NewInMemorySearchEngine(),
	}

	base.Auth.AddUser("admin", []byte("password"))

	return base
}

func testSuccess(t *testing.T, res *io.Response) {
	assert.Equal(t, common.SUCCESS, res.Code)
	assert.Equal(t, common.SUCCESS_MSG, res.Message)
}
