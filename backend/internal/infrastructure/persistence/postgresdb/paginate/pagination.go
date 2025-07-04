package paginate

import (
	"math"

	"gorm.io/gorm"
)


type Pagination struct {
	Limit      int         `json:"limit,omitempty"`
	Page       int         `json:"page,omitempty"`
	Sort       string      `json:"sort,omitempty"`
	TotalRows  int64       `json:"totalRows"`
	TotalPages int         `json:"totalPages"`
	Rows       interface{} `json:"rows"`
	Filter     []string    `json:"filter"`
	FilterArgs [][]interface{}    `json:"filterArgs"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
    var totalRows int64

    countQuery := db.Model(value)
    for i, filter := range pagination.Filter {
        countQuery = countQuery.Where(filter, pagination.FilterArgs[i]...)
    }

    countQuery.Count(&totalRows)

    pagination.TotalRows = totalRows
    totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
    pagination.TotalPages = totalPages

    return func(db *gorm.DB) *gorm.DB {
        scopedDB := db

        for i, filter := range pagination.Filter {
            scopedDB = scopedDB.Where(filter, pagination.FilterArgs[i]...)
        }

        return scopedDB.
            Offset(pagination.GetOffset()).
            Limit(pagination.GetLimit()).
            Order(pagination.GetSort())
    }
}