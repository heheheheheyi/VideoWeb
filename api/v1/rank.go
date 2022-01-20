package v1

import (
	"VideoWeb/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DailyRank(c *gin.Context) {
	var data service.DailyRank
	res := data.GetRank()
	c.JSON(http.StatusOK, res)
}

func MonthlyRank(c *gin.Context) {
	var data service.MonthlyRank
	res := data.GetRank()
	c.JSON(http.StatusOK, res)
}
func ClickRank(c *gin.Context) {
	var data service.ClickRank
	res := data.GetRank()
	c.JSON(http.StatusOK, res)
}
