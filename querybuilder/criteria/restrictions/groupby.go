package restrictions

import (
	c "github.com/morg/querybuilder/commons"
)

type GroupByRestriction struct {
	column string
	rtype  string
}

func NewGroupByRestriction(col, rtype string) *GroupByRestriction {
	return &GroupByRestriction{column: col, rtype: rtype}
}

func (restriction *Restriction) GroupBy(col string) *GroupByRestriction {
	return NewGroupByRestriction(col, c.Type_Group)
}

func (restriction *GroupByRestriction) Tostring() string {
	return restriction.column
}
