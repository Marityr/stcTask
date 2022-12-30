package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Datalist
// @Summary  Datalist
// @Description  Datalist
// @Tags     datalist
// @Accept   json
// @Produce  json
// @Param    num  query  int  false  "Num"
// @Router   /api/v1/datalist/ [get]
func (h *Handler) DataList(c *gin.Context) {
	data := c.Query("num")

	num, err := strconv.Atoi(data)
	if err != nil {
		log.Println(err)
		c.JSON(406, gin.H{
			"data":  "",
			"error": err,
		})
		return
	}

	// TIPS query limit sqlite3 <1000
	if num > 990 {
		log.Println(errors.New("fetch limit exceeded, no more than 990 records at a time"))
		c.JSON(406, gin.H{
			"data":  "",
			"error": "fetch limit exceeded, no more than 990 records at a time",
		})
		return
	}

	lists, err := h.services.DataList.GetKey(num)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"data":  "",
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  lists,
		"error": "",
	})
}
