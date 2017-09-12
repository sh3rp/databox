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
		io.Respond(c, common.E_DB_BOX_NOT_FOUND, nil)
		return
	}

	linkId := c.Param("linkId")

	link, err := r.DB.GetLinkById(msg.Key{Id: linkId, BoxId: boxId, Type: msg.Key_LINK})

	if err != nil {
		io.Respond(c, common.E_DB_LINK_NOT_FOUND, nil)
		return
	} else {
		io.Respond(c, common.SUCCESS, link)
	}

}
