package http_utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// getUint64Param returns *uint param. On error returns error and nil if param not exists
func getUint64Param(param string, c *gin.Context) (*uint, error) {
	stringParam, ok := c.GetQuery(param)
	if ok {
		paramUint64, err := strconv.ParseUint(stringParam, 10, 32)
		if err != nil {
			return nil, err
		}
		result := uint(paramUint64)
		return &result, nil
	} else {
		return nil, nil
	}
}
