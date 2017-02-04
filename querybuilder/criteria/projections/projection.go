package projections

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type Projection struct {
	column string
	ptype  string
}

func NewProjection(col, ptype string) *Projection {
	return &Projection{column: col, ptype: ptype}
}

func (projection *Projection) Tostring() string {
	if projection.ptype == c.Type_Project {
		return projection.column
	} else if projection.ptype == c.Type_Distinct {
		stmt := make([]string, 0)
		stmt = append(stmt, projection.ptype)
		stmt = append(stmt, c.Type_Space)
		stmt = append(stmt, projection.column)
		return strings.Join(stmt, "")
	} else {
		stmt := make([]string, 0)
		stmt = append(stmt, projection.ptype)
		stmt = append(stmt, "(")
		stmt = append(stmt, projection.column)
		stmt = append(stmt, ")")
		return strings.Join(stmt, "")
	}
}

func (projection *Projection) Distinct(col string) *Projection {
	return NewProjection(col, c.Type_Distinct)
}

func (projection *Projection) Column(col string) *Projection {
	return NewProjection(col, c.Type_Project)
}
