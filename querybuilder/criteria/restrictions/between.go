package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type BetweenRestriction struct {
	column string
	lbound interface{}
	ubound interface{}
	rtype  string
}

func NewBetweenRestriction(col string, lvalue, uvalue interface{}, rtype string) *BetweenRestriction {
	return &BetweenRestriction{column: col, lbound: lvalue, ubound: uvalue, rtype: rtype}
}

func (restriction *Restriction) Between(col string, lvalue, uvalue interface{}) *BetweenRestriction {
	return NewBetweenRestriction(col, lvalue, uvalue, c.Type_Between)
}

func (restriction *BetweenRestriction) Tostring() string {
	stmt := make([]string, 0)
	stmt = append(stmt, c.Keyword_Between)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, restriction.column)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, "("+restriction.lbound.(string))
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, c.Keyword_And)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, restriction.ubound.(string)+")")
	return strings.Join(stmt, "")
}
