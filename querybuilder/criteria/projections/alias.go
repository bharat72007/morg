package projections

import (
	c "github.com/morg/querybuilder/commons"
	"strings"
)

type AliasProjection struct {
	column string
	ptype  string
	alias  string
}

func NewAliasProjection(col, aliasname string, rtype string) *AliasProjection {
	return &AliasProjection{column: col, alias: aliasname, ptype: rtype}
}

func (projection *AliasProjection) Alias(col, aliasname string) *AliasProjection {
	return NewAliasProjection(col, aliasname, c.Type_Alias)
}

func (projection *AliasProjection) Tostring() string {
	rslt := make([]string, 0)
	rslt = append(rslt, projection.column)
	rslt = append(rslt, c.Type_Space)
	rslt = append(rslt, c.Keyword_Alias)
	rslt = append(rslt, c.Type_Space)
	rslt = append(rslt, projection.alias)
	return strings.Join(rslt, "")
}
