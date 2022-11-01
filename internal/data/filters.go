package data

import (
	"math"
	"strings"
)

type Filters struct {
	Page     int
	PageSize int
	Sort     string
}

type Metadata struct {
	CurrentPage  int `json:"current_page"`
	PageSize     int `json:"page_size"`
	FirstPage    int `json:"first_page"`
	LastPage     int `json:"last_page"`
	TotalRecords int `json:"total_records"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(pageSize))),
		TotalRecords: totalRecords,
	}
}

func (f Filters) sortColumnReportQuery() string {
	s := strings.TrimPrefix(f.Sort, "-")
	if s == "sum" {
		return "service_price"
	} else {
		return "report_time"
	}
}

func (f Filters) sortColumnTransactionQuery() string {
	s := strings.TrimPrefix(f.Sort, "-")
	if s == "sum" {
		return "transaction_price"
	} else {
		return "transaction_time"
	}
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}

	return "ASC"
}

func (f Filters) limit() int {
	return f.PageSize
}

func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}
