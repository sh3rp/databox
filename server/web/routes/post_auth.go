package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) PostAuth(c *gin.Context) {
	var auth *msg.AuthRequest
	err := c.BindJSON(&auth)
	if err != nil {
		c.JSON(200, io.Error(common.E_IO_AUTH_PARSE, err))
	} else {
		if r.Auth.Authenticate(auth.Username, []byte(auth.Password)) {
			token := r.TokenStore.GenerateToken(auth.Username)
			if token == nil {
				c.JSON(200, io.Error(common.E_IO_TOKEN_CREATION, errors.New("Error generating token")))
			} else {
				c.JSON(200, io.Success(token))
			}
		} else {
			c.JSON(200, io.Error(common.E_IO_INVALID_LOGIN, errors.New("Invalid authentication")))
		}
	}
}
