package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type OffsetRestriction struct {
	offset int
	rtype  string
}

func NewOffsetRestriction(start int, rtype string) *OffsetRestriction {
	return &OffsetRestriction{offset: start, rtype: rtype}
}

func (restriction *OffsetRestriction) tostring() string {
	stmt := make([]string, 0)
	stmt = append(stmt, c.Keyword_Offset)
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, strconv.Itoa(restriction.offset))
	return strings.Join(stmt, "")
}

func (restriction *Restriction) Offset(start int) *OffsetRestriction {
	return NewOffsetRestriction(start, c.Type_Offset)
}
