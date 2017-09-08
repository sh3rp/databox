package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) PostLink(c *gin.Context) {
	var box msg.Box
	err := c.BindJSON(&box)
	if err == nil {
		if box.Id == nil {
			newBox, err := r.DB.NewBox(box.Name, box.Description, []byte("password"))
			if err != nil {
				c.JSON(200, io.Error(common.E_DB_CREATE_BOX, err))
				return
			} else {
				c.JSON(200, io.Success(newBox))
				return
			}
		}
		c.JSON(200, io.Success(box))
	} else {
		c.JSON(200, io.Error(common.E_IO_DECODE_BOX, err))
	}
}
