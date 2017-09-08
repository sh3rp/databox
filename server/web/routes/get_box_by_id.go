package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) GetBoxById(c *gin.Context) {
	id := c.Param("id")

	box, err := r.DB.GetBoxById(msg.Key{Id: id, Type: msg.Key_BOX})

	if err == nil {
		c.JSON(200, io.Success(box))
	} else {
		c.JSON(200, io.Error(common.E_DB_BOX_NOT_FOUND, err))
	}
}
