package main

import (
	"fmt"
	"strings"
)

type Query struct {
	Projection string
	Filter     *Criteria
	From       string
}

type Criteria struct {
	//can have list of Restriction
	Restrictions []Restriction
}

type Restriction struct {
	column   string
	value    interface{}
	resttype string
}

func CreateCriteria() *Criteria {
	return &Criteria{Restrictions: make([]Restriction, 0)}
}

func NewRestriction(col string, val interface{}, rtype string) *Restriction {
	return &Restriction{column: col, value: val, resttype: rtype}
}

func NewQuery() *Query {
	return &Query{}
}

func (criteria *Criteria) add(restriction *Restriction) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, *restriction)
	return criteria
}

/*func (criteria *Criteria) addOr(restriction1, restriction2 *Restriction) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, *restriction1)
	criteria.Restrictions = append(criteria.Restrictions, *restriction2)
	return criteria
}*/

func (restriction Restriction) tostring() string {

	strs := restriction.column + " " + restriction.resttype + " " + restriction.value.(string)
	return strs
}

func (restriction Restriction) Equal(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "=")
}

func (restriction Restriction) NotEqual(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "<>")
}

func (restriction Restriction) Gt(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, ">")
}

func (restriction Restriction) Gte(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, ">=")
}

func (restriction Restriction) Lte(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "<=")
}

func (restriction Restriction) Lt(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "<")
}

func (restriction Restriction) Like(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "LIKE")
}

/*func (restriction Restriction) Between(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "BETWEEN")
}*/

func (query *Query) Project(fields ...string) *Query {
	query.Projection = strings.Join(fields, ",")
	return query
}

func (query *Query) Tables(tables ...string) *Query {
	query.From = strings.Join(tables, ",")
	return query
}

func (query *Query) AddCriteria(criteria *Criteria) *Query {
	query.Filter = criteria
	return query
}

func (query *Query) transform() {
	var tokens []string = make([]string, 0)
	if query.Projection != "" {
		tokens = append(tokens, "SELECT", query.Projection)
	}
	if query.From != "" {
		tokens = append(tokens, "FROM", query.From)
	}
	if len(query.Filter.Restrictions) > 0 {
		//TODO

		for i := 0; i < len(query.Filter.Restrictions); i++ {
			if i == 0 {
				tokens = append(tokens, "WHERE", query.Filter.Restrictions[i].tostring())
			} else {
				tokens = append(tokens, "AND", query.Filter.Restrictions[i].tostring())
			}
		}

	}
	fmt.Println(tokens)
}

func main() {

	query := NewQuery()
	query.Project("name", "age", "rollno", "Student.name").Tables("User", "Student")
	//query.transform()
	criteria := CreateCriteria()
	res := Restriction{}
	/*ag := make([]int, 0)
	ag = append(ag, 10)
	ag = append(ag, 20)*/
	criteria.add(res.Equal("age", "22")).add(res.Gt("rollno", "900")).add(res.NotEqual("age", "80")) /*.add(res.Between("age", ag))*/
	query.AddCriteria(criteria)
	query.transform()
}
