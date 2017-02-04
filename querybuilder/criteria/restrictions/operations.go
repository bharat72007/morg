package restrictions

import c "github.com/morg/querybuilder/commons"

func (restriction *Restriction) NotEqual(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_NotEq)
}

func (restriction *Restriction) Equal(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_Eq)
}

func (restriction *Restriction) Gt(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_Gt)
}

func (restriction *Restriction) Gte(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_Gte)
}

func (restriction *Restriction) Lte(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_Lte)
}

func (restriction *Restriction) Lt(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_Lt)
}

func (restriction *Restriction) Like(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_Like)
}

func (restriction *Restriction) ILike(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, c.Type_Operation_ILike)
}
