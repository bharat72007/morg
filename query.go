package morg

type Query struct {
	Projection []interface{}
	Filter     *Condition
	From       interface{}
}

type Condition struct {
	//can have list of Restriction
	Restrictions []Restriction
}

type Restriction struct {
	column interface{}
	value  interface{}
	typ    string
}

func NewQuery(table interface{}) *Query {
	return &Query{Projection: make([]string, 0), Filter: nil, From: table}
}

func (query *Query) addProjection(fields ...interface{}) *Query {
	numFields := len(fields)
	for i := 0; i < numFields; i++ {
		query.Projection = append(query.Projection, fields[i])
	}
	return &query
}

func (query *Query) addFilter(condition Condition) *Query {
	query.Filter = condition
	return &query
}

func NewRestriction(k, v interface{}) *Restriction {
	return &Restriction{column: k, value: v, typ: nil}
}

func NewCondition() *Condition {
	return &Condition{Restrictions: make([]Restriction, 0)}
}

func (condition *Condition) Equal(k, v interface{}) *Condition {
	restriction := NewRestriction(k, v)
	restriction.typ = "="
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) Gt(k, v interface{}) *Condition {
	restriction := NewRestriction(k, v)
	restriction.typ = ">"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) Lt(k, v interface{}) *Condition {
	restriction := NewRestriction(k, v)
	restriction.typ = "<"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) NotEqual(k, v interface{}) *Condition {
	restriction := NewRestriction(k, v)
	restriction.typ = "!="
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) Like(k, v interface{}) *Condition {
	restriction := NewRestriction(k, v)
	restriction.typ = "like"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) ILike(k, v interface{}) *Condition {
	restriction := NewRestriction(k, v)
	restriction.typ = "ilike"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) IsEmpty(k interface{}) *Condition {
	restriction := NewRestriction(k, nil)
	restriction.typ = "IsEmpty"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) IsNotEmpty(k interface{}) *Condition {
	restriction := NewRestriction(k, nil)
	restriction.typ = "IsNotEmpty"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) IsNull(k interface{}) *Condition {
	restriction := NewRestriction(k, nil)
	restriction.typ = "IsNull"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

func (condition *Condition) IsNotNull(k interface{}) *Condition {
	restriction := NewRestriction(k, nil)
	restriction.typ = "IsNotNull"
	condition.Restrictions = append(condition.Restrictions, restriction)
	return &condition
}

//Example1 :
// SELECT f1,f2 FROM tablename
//Create Query structure
// query := NewQuery("USER")
//query.addProjection("f1","f2")

//Example2 :
//SELECT f1,f2 FROM tablename WHERE f3 > 9
//query := NewQuery("USER")
//query.addProjection("f1","f2")
//con := NewCondition()
//con.Gt("f3","9")
//query.addFilter(con)

func examples() {

	//Example #1
	query := NewQuery("USER")
	query.addProjection("name", "age")

	//Example #2
	query1 := NewQuery("USER")
	query1.addProjection("name")
	cond := NewCondition()
	cond.Gt("age", 9)
	query.addFilter(cond)

}
