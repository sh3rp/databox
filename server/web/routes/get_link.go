package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sh3rp/databox/common"
	"github.com/sh3rp/databox/msg"
	"github.com/sh3rp/databox/server/web/io"
)

func (r *RouterBase) GetLink(c *gin.Context) {
	id := c.Param("id")

	links, err := r.DB.GetLinksByBoxId(msg.Key{Id: id, Type: msg.Key_BOX})

	if err == nil {
		io.Respond(c, common.SUCCESS, links)
	} else {
		io.Respond(c, common.E_DB_LINK_NOT_FOUND, nil)
	}

}
