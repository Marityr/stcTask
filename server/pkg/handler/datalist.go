package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Datalist
// @Summary  Datalist
// @Description  Datalist
// @Tags     datalist
// @Accept   json
// @Produce  json
// @Param    key  query  string  false  "Key"
// @Router   /api/v1/datalist/ [get]
func (h *Handler) DataList(c *gin.Context) {
	data := c.Param("key")

	lists, err := h.services.DataList.GetKey(data)
	if err != nil {
		// h.logging.Error(err)
		c.AbortWithError(400, err)
		return
	}

	c.JSON(http.StatusOK, lists)
}
