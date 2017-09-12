package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) PostBox(c *gin.Context) {
	var req io.NewBoxRequest
	err := c.BindJSON(&req)
	if err == nil {
		if req.Name == "" {
			io.Respond(c, common.E_BOX_INVALID_NAME, nil)
		} else if req.Description == "" {
			io.Respond(c, common.E_BOX_INVALID_DESCRIPTION, nil)
		} else if req.Password == "" {
			io.Respond(c, common.E_BOX_INVALID_PASSWORD, nil)
		}
		newBox, err := r.DB.NewBox(req.Name, req.Description, []byte(req.Password))
		if err != nil {
			io.Respond(c, common.E_DB_CREATE_BOX, nil)
			return
		} else {
			io.Respond(c, common.SUCCESS, newBox)
			return
		}
	} else {
		io.Respond(c, common.E_IO_DECODE_BOX, nil)
	}
}
