package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetLimit get limit from request
func GetLimit(context *gin.Context) int {
	limit, err := strconv.Atoi(context.Query("limit"))
	if err != nil || limit == 0 {
		limit = 10
	}

	return limit
}

//GetPage number from request
func GetPage(context *gin.Context) int {
	page, err := strconv.Atoi(context.Query("page"))
	if err != nil {
		page = 0
	}

	return page
}
