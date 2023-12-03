package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPageAndPageSize(c *gin.Context) (int, int) {
	var (
		page        = c.Query("page")
		pageSize    = c.Query("page_size")
		pageInt     = 0
		pageSizeInt = 0
		err         error
	)

	pageInt, err = strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	pageSizeInt, err = strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 0
	}

	return pageInt, pageSizeInt
}
