package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type LimitRestriction struct {
	rowscount int
	rtype     string
}

func NewLimitRestriction(rows int, rtype string) *LimitRestriction {
	return &LimitRestriction{rowscount: rows, rtype: rtype}
}

func (restriction *LimitRestriction) tostring() string {
	stmt := make([]string, 0)
	stmt = append(stmt, c.Keyword_Limit)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, strconv.Itoa(restriction.rowscount))
	return strings.Join(stmt, "")
}

func (restriction *Restriction) Limit(rows int) *LimitRestriction {
	return NewLimitRestriction(rows, c.Type_Limit)
}
