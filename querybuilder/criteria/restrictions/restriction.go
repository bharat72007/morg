package restrictions

import (
	c "github.com/morg/querybuilder/commons"
	t "github.com/morg/querybuilder/criteria/types"
	"strings"
)

type Restriction struct {
	column string
	value  interface{}
	rtype  string
}

func NewRestriction(col string, val interface{}, rtype string) *Restriction {
	t.CreateTypeMemInstance()
	t.TypeMemInstance.Fields = append(t.TypeMemInstance.Fields, t.NewTypeInfo(val, c.Findtype(val)))
	return &Restriction{column: col, value: val, rtype: rtype}
}

func (restriction *Restriction) Tostring() string {
	var rtypeval string
	stmt := make([]string, 0)
	stmt = append(stmt, restriction.column)
	stmt = append(stmt, c.Type_Space)
	if restriction.rtype == c.Type_Operation_Gt {
		rtypeval = c.Operation_Gt
	} else if restriction.rtype == c.Type_Operation_Eq {
		rtypeval = c.Operation_Eq
	} else if restriction.rtype == c.Type_Operation_NotEq {
		rtypeval = c.Operation_NotEq
	} else if restriction.rtype == c.Type_Operation_Gte {
		rtypeval = c.Operation_Gte
	} else if restriction.rtype == c.Type_Operation_Lt {
		rtypeval = c.Operation_Lt
	} else if restriction.rtype == c.Type_Operation_Lte {
		rtypeval = c.Operation_Lte
	} else if restriction.rtype == c.Type_Operation_Like {
		rtypeval = c.Operation_Like
	} else if restriction.rtype == c.Type_Operation_ILike {
		rtypeval = c.Operation_ILike
	}
	stmt = append(stmt, rtypeval)
	stmt = append(stmt, c.Type_Space)
	/*stmt = append(stmt, restriction.value.(string))*/
	stmt = append(stmt, "?")
	return strings.Join(stmt, "")
}
