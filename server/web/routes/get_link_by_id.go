package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) GetLinkById(c *gin.Context) {
	boxId := c.Param("id")

	_, err := r.DB.GetBoxById(msg.Key{Id: boxId, Type: msg.Key_BOX})

	if err != nil {
		c.JSON(200, io.Error(common.E_DB_BOX_NOT_FOUND, err))
		return
	}

	linkId := c.Param("linkId")

	link, err := r.DB.GetLinkById(msg.Key{Id: linkId, BoxId: boxId, Type: msg.Key_LINK})

	if err != nil {
		c.JSON(200, io.Error(common.E_DB_LINK_NOT_FOUND, err))
		return
	} else {
		c.JSON(200, io.Success(link))
	}

}
