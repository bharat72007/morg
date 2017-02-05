package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	t "github.com/morg/querybuilder/criteria/types"
	"strings"
)

type BetweenRestriction struct {
	column string
	lbound interface{}
	ubound interface{}
	rtype  string
}

func NewBetweenRestriction(col string, lvalue, uvalue interface{}, rtype string) *BetweenRestriction {
	t.CreateTypeMemInstance()
	t.TypeMemInstance.Fields = append(t.TypeMemInstance.Fields, t.NewTypeInfo(lvalue, c.Findtype(lvalue)))
	t.TypeMemInstance.Fields = append(t.TypeMemInstance.Fields, t.NewTypeInfo(uvalue, c.Findtype(uvalue)))
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
	/*stmt = append(stmt, "("+restriction.lbound.(string))*/
	stmt = append(stmt, "(?")
	stmt = append(stmt, c.Type_Space)
	stmt = append(stmt, c.Keyword_And)
	stmt = append(stmt, c.Type_Space)
	/*stmt = append(stmt, restriction.ubound.(string)+")")*/
	stmt = append(stmt, "?)")
	return strings.Join(stmt, "")
}
