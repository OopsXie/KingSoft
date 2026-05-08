package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	ReqData interface{} `json:"reqData"`
	Data    interface{} `json:"data"`
}

func main() {
	r := gin.Default()

	r.GET("/api/sum", func(c *gin.Context) {
		x, _ := c.GetQuery("x")
		y, _ := c.GetQuery("y")
		var result int

		if x != "" && y != "" {
			intX, errX := strconv.Atoi(x)
			intY, errY := strconv.Atoi(y)
			if errX == nil && errY == nil {
				result = intX + intY
			} else {
				c.JSON(400, Response{
					Code:    1,
					Message: "invalid query parameters",
					ReqData: gin.H{"x": x, "y": y},
					Data:    nil,
				})
				return
			}
		} else {
			c.JSON(400, Response{
				Code:    1,
				Message: "missing query parameters",
				ReqData: gin.H{"x": x, "y": y},
				Data:    nil,
			})
			return
		}

		c.JSON(200, Response{
			Code:    0,
			Message: "success",
			ReqData: gin.H{"x": x, "y": y},
			Data:    result,
		})
	})

	r.Run(":8080")
}
