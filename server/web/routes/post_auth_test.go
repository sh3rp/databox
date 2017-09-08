package routes

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/auth"
	"github.com/sh3rp/databox/db"
	"github.com/sh3rp/databox/search"
	"github.com/sh3rp/databox/server/web/io"
	"github.com/stretchr/testify/assert"
)

func TestPostAuth(t *testing.T) {
	base := newRouterBase()
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/auth", base.PostAuth)
	data := []byte{}
	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(data))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	response := &io.Response{}
	json.NewDecoder(res.Body).Decode(response)
	assert.Equal(t, 0, response.Code)
}

func newRouterBase() *RouterBase {
	return &RouterBase{
		Auth:       auth.NewInMemoryAuthenticator(),
		TokenStore: auth.NewInMemoryTokenStore(1),
		DB:         db.NewInMemoryDB(),
		Search:     search.NewInMemorySearchEngine(),
	}
}
