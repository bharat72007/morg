package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type InRestriction struct {
	column string
	list   []interface{}
	rtype  string
}

func (restriction *InRestriction) Tostring() string {
	stmt := make([]string, 0)
	stmt = append(stmt, restriction.column)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, c.Keyword_In)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, "(")
	elemlist := make([]string, 0)
	for _, v := range restriction.list {
		elemlist = append(elemlist, v.(string))
	}
	stmt = append(stmt, strings.Join(elemlist, ","))
	stmt = append(stmt, ")")
	return strings.Join(stmt, "")
}

func NewInRestriction(col string, vals []interface{}, rtype string) *InRestriction {
	return &InRestriction{column: col, list: vals, rtype: rtype}
}

func (restriction *Restriction) In(col string, val []interface{}) *InRestriction {
	return NewInRestriction(col, val, c.Type_In)
}
