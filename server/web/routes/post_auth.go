package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) PostAuth(c *gin.Context) {
	var auth *msg.AuthRequest
	err := c.BindJSON(&auth)
	if err != nil {
		io.Respond(c, common.E_IO_AUTH_PARSE, nil)
	} else {
		if r.Auth.Authenticate(auth.Username, []byte(auth.Password)) {
			token := r.TokenStore.GenerateToken(auth.Username)
			if token == nil {
				io.Respond(c, common.E_IO_TOKEN_CREATION, nil)
			} else {
				io.Respond(c, common.SUCCESS, token)
			}
		} else {
			io.Respond(c, common.E_IO_INVALID_LOGIN, nil)
		}
	}
}
