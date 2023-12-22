package pagination

import (
	"errors"
	"math"
	"net/url"
	"strconv"
)

type Pagination struct {
	Page      int `form:"page" json:"page"`
	PerPage   int `form:"per_page" json:"per_page"`
	Limit     int `form:"-" json:"-"`
	Offset    int `form:"-" json:"-"`
	Total     int `form:"total" json:"total"`
	TotalPage int `form:"total_page" json:"total_page"`
}

func (p *Pagination) Transform(query url.Values) error {
	pageQuery := query.Get("page")
	if pageQuery == "" {
		pageQuery = "1"
	}
	perPageQuery := query.Get("per_page")
	if perPageQuery == "" {
		perPageQuery = "10"
	}
	page, err := strconv.ParseInt(pageQuery, 10, 32)
	if err != nil {
		return errors.New("page is not a number")
	}

	perPage, err := strconv.ParseInt(perPageQuery, 10, 32)
	if err != nil {
		return errors.New("per_page is not a number")
	}

	limit := perPage
	offset := (page - 1) * perPage

	p.Page = int(page)
	p.PerPage = int(perPage)
	p.Limit = int(limit)
	p.Offset = int(offset)
	return nil
}

func (p *Pagination) Finish(count int) {
	p.Total = count
	p.TotalPage = int(math.Ceil(float64(count) / float64(p.PerPage)))
}
