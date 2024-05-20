package pagination

import (
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Pagination struct {
	TotalRecords int64       `json:"total_records"`
	TotalPages   int         `json:"total_pages"`
	Records      interface{} `json:"items"`
}

func parseParams(c echo.Context) (page, limit int) {
	page = 1
	limit = 20

	if pageParam := c.QueryParam("page"); pageParam != "" {
		p, err := strconv.Atoi(pageParam)
		if err == nil && p > 0 {
			page = p
		}
	}

	if limitParam := c.QueryParam("limit"); limitParam != "" {
		l, err := strconv.Atoi(limitParam)
		if err == nil && l > 0 {
			limit = l
		}
	}

	return page, limit
}

func Paginate(c echo.Context, db *gorm.DB, records interface{}) (*Pagination, error) {
	page, limit := parseParams(c)

	var totalRecords int64
	if err := db.Model(records).Count(&totalRecords).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * limit

	if err := db.Limit(limit).Offset(offset).Find(records).Error; err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(totalRecords) / float64(limit)))

	pagination := &Pagination{
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		Records:      records,
	}

	return pagination, nil
}
