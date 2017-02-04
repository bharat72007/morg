package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type Restriction struct {
	column string
	value  interface{}
	rtype  string
}

func NewRestriction(col string, val interface{}, rtype string) *Restriction {
	return &Restriction{column: col, value: val, rtype: rtype}
}

func (restriction *Restriction) Tostring() string {
	stmt := make([]string, 0)
	stmt = append(stmt, restriction.column)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, restriction.value.(string))
	return strings.Join(stmt, "")
}
