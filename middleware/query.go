package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func QueryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}

		if limit > 100 {
			limit = 100
		}

		offset := (page - 1) * limit
		search := c.Query("search")

		c.Set("page", page)
		c.Set("limit", limit)
		c.Set("offset", offset)
		c.Set("search", search)

		startDateStr := c.Query("startDate")
		endDateStr := c.Query("endDate")

		if startDateStr != "" {
			startDate, err := time.Parse("2006-01-02", startDateStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Failed",
					"info": gin.H{
						"code":    400,
						"message": "Invalid startDate format. Use YYYY-MM-DD",
					},
				})
				c.Abort()
				return
			}

			c.Set("startDate", startDate)
		}

		if endDateStr != "" {
			endDate, err := time.Parse("2006-01-02", endDateStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "Failed",
					"info": gin.H{
						"code":    400,
						"message": "Invalid endDate format. Use YYYY-MM-DD",
					},
				})
				c.Abort()
				return
			}

			endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)
			c.Set("endDate", endDate)
		}

		c.Next()
	}
}
