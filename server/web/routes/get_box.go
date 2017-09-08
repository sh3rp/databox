package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) GetBox(c *gin.Context) {
	boxes, err := r.DB.GetBoxes()
	if err != nil {
		c.JSON(200, io.Error(common.E_DB_BOX_NOT_FOUND, err))
	} else {
		c.JSON(200, io.Success(boxes))
	}
}
