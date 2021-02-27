package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/services"
)

func FormInit(c *gin.Context) {
	matchData, err := services.GetFormInit(c.Request.Context(), c.Param("site_id"), c.Param("item_id"), c.Query("version"))

	if matchData != nil && err == nil {
		c.Header("Cache-Control", "max-age=600")
		c.JSON(http.StatusOK, matchData)
	} else {
		c.JSON(400, err)
	}
	return
}
