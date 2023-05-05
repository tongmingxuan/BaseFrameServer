// Package Utils /**
package Utils

import (
	"math"
)

// ClosureFunc 闭包
type ClosureFunc func() (int, interface{})

type Paginate struct {
	Page        int
	Limit       int
	LimitBefore int
	Where       interface{}
	List        interface{}
	count       int
	ClosureFunc ClosureFunc
}

func (Paginate *Paginate) Init(page, limit int, where interface{}) {
	Paginate.Page = page

	Paginate.Limit = limit

	Paginate.Where = where

	Paginate.LimitBefore = (page - 1) * limit
}

func (Paginate *Paginate) SetFunc(closureFunc ClosureFunc) {
	count, list := closureFunc()

	Paginate.count = count

	Paginate.List = list
}

func (Paginate *Paginate) GetPageList() map[string]interface{} {
	totalPage := math.Ceil(float64(Paginate.count) / float64(Paginate.Limit))

	return map[string]interface{}{
		"total_count": Paginate.count,
		"page":        Paginate.Page,
		"total_page":  int(totalPage),
		"list":        Paginate.List,
	}
}
