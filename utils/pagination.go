package utils

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
}

func Paginate(c *fiber.Ctx, db *gorm.DB, model interface{}, searchableFields []string, filterableFields []string) (*PaginatedResponse, error) {
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)
	sort := c.Query("sort", "id")
	order := c.Query("order", "asc")
	search := c.Query("search", "")

	offset := (page - 1) * pageSize

	query := db.Model(model)

	// === SEARCH ===
	if search != "" && len(searchableFields) > 0 {
		query = query.Where(func(tx *gorm.DB) *gorm.DB {
			for i, field := range searchableFields {
				if i == 0 {
					tx = tx.Where(field+" ILIKE ?", "%"+search+"%")
				} else {
					tx = tx.Or(field+" ILIKE ?", "%"+search+"%")
				}
			}
			return tx
		})
	}

	// === FILTER ===
	for _, field := range filterableFields {
		if value := c.Query(field); value != "" {
			query = query.Where(field+" = ?", value)
		}
	}

	// === SORT ===
	if sort != "" {
		query = query.Order(strings.ToLower(sort) + " " + strings.ToLower(order))
	}

	// === TOTAL DATA ===
	var total int64
	err := query.Count(&total).Error
	if err != nil {
		return nil, err
	}

	// === GET DATA ===
	err = query.Offset(offset).Limit(pageSize).Find(model).Error
	if err != nil {
		return nil, err
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))

	return &PaginatedResponse{
		Data:       model,
		Page:       page,
		PageSize:   pageSize,
		TotalRows:  total,
		TotalPages: totalPages,
	}, nil
}
