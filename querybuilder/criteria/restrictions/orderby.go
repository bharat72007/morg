package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type OrderByRestriction struct {
	column string
	order  string
	rtype  string
}

func NewOrderByRestriction(col, order, rtype string) *OrderByRestriction {
	return &OrderByRestriction{column: col, order: order, rtype: rtype}
}

func (restriction *OrderByRestriction) tostring() string {
	stmt = make([]string, 0)
	stmt = append(stmt, restriction.column)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, restriction.order)
	return strings.Join(stmt, "")
}

func (restriction *Restriction) OrderBy(col, order string) *OrderByRestriction {
	return NewOrderByRestriction(col, order, c.Type_Order)
}
