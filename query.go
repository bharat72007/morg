package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Query struct {
	Projection string
	Filter     *Criteria
	/*From       string*/
}

type Criteria struct {
	//can have list of Restriction
	Restrictions []interface{}
	combiners    []string
	Entity       interface{}
}

type Restriction struct {
	column   string
	value    interface{}
	resttype string
}

type BetweenRestriction struct {
	column   string
	lbound   interface{}
	ubound   interface{}
	resttype string
}

type InRestriction struct {
	column   string
	list     []interface{}
	resttype string
}

func CreateCriteria(class interface{}) *Criteria {
	return &Criteria{Restrictions: make([]interface{}, 0), combiners: make([]string, 0), Entity: class}
}

func NewBetweenRestriction(col string, lval, rval interface{}, rtype string) *BetweenRestriction {
	return &BetweenRestriction{column: col, lbound: lval, ubound: rval, resttype: rtype}
}

func NewInRestriction(col string, vals []interface{}, rtype string) *InRestriction {
	return &InRestriction{column: col, list: vals, resttype: rtype}
}

func NewRestriction(col string, val interface{}, rtype string) *Restriction {
	return &Restriction{column: col, value: val, resttype: rtype}
}

func NewQuery() *Query {
	return &Query{}
}

func (criteria *Criteria) add(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	if len(criteria.Restrictions) > 1 {
		criteria.combiners = append(criteria.combiners, "AND")
	}
	return criteria
}

func (criteria *Criteria) exclude(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	criteria.combiners = append(criteria.combiners, "NOT")
	return criteria
}

func (criteria *Criteria) addOr(restriction1, restriction2 interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction1)
	criteria.Restrictions = append(criteria.Restrictions, restriction2)
	criteria.combiners = append(criteria.combiners, "OR")
	return criteria
}

func (restriction *Restriction) tostring() string {

	strs := restriction.column + " " + restriction.resttype + " " + restriction.value.(string)
	return strs
}

func (restriction *BetweenRestriction) tostring() string {
	strs := "BETWEEN" + " " + restriction.column + " " + restriction.lbound.(string) + " " + "AND" + " " + restriction.ubound.(string)
	return strs
}

func (restriction *InRestriction) tostring() string {
	newlist := make([]string, 0)
	for _, v := range restriction.list {
		newlist = append(newlist, v.(string))
	}
	strs := restriction.column + " " + "IN" + " " + " (" + strings.Join(newlist, ",") + ")"
	return strs
}

func (restriction *Restriction) Equal(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "=")
}

func (restriction *Restriction) NotEqual(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "<>")
}

func (restriction *Restriction) Gt(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, ">")
}

func (restriction *Restriction) Gte(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, ">=")
}

func (restriction *Restriction) Lte(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "<=")
}

func (restriction *Restriction) Lt(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "<")
}

func (restriction *Restriction) Like(col string, val interface{}) *Restriction {
	return NewRestriction(col, val, "LIKE")
}

func (restriction *Restriction) Between(col string, val1, val2 interface{}) *BetweenRestriction {
	return NewBetweenRestriction(col, val1, val2, "BETWEEN")
}

func (restriction *Restriction) In(col string, val []interface{}) *InRestriction {
	return NewInRestriction(col, val, "IN")
}

func (query *Query) Project(fields ...string) *Query {
	query.Projection = strings.Join(fields, ",")
	return query
}

/*func (query *Query) Tables(tables ...string) *Query {
	query.From = strings.Join(tables, ",")
	return query
}*/

func (query *Query) AddCriteria(criteria *Criteria) *Query {
	query.Filter = criteria
	return query
}

func (query *Query) transform() {
	var tokens []string = make([]string, 0)
	//var ent string
	if query.Projection != "" {
		tokens = append(tokens, "SELECT", query.Projection)
	}
	if query.Filter.Entity != "" {
		tokens = append(tokens, "FROM", reflect.TypeOf(query.Filter.Entity).Elem().Name())
	}
	if len(query.Filter.Restrictions) > 0 {
		//TODO

		tokens = append(tokens, "WHERE")
		for i := 0; i < len(query.Filter.Restrictions); i++ {
			switch t := query.Filter.Restrictions[i].(type) {
			case *Restriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*Restriction).tostring())
				tokens = append(tokens, query.Filter.combiners[i])
			case *BetweenRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*BetweenRestriction).tostring())
				tokens = append(tokens, query.Filter.combiners[i])
			case *InRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*InRestriction).tostring())
				/*tokens = append(tokens, query.Filter.combiners[i])*/
			default:
				fmt.Printf("%T", t)
			}
		}

	}
	fmt.Println(tokens)
}

func main() {
	var student Student
	criteria := CreateCriteria(&student)
	var restriction Restriction
	ag := make([]interface{}, 0)
	ag = append(ag, "10")
	ag = append(ag, "20")
	criteria.add(restriction.Equal("age", "22")).add(restriction.Gt("rollno", "900")).add(restriction.NotEqual("age", "80")).add(restriction.Between("name", "00", "9")).add(restriction.In("age", ag))
	query := NewQuery()
	query.Project("name", "age", "rollno", "Student.name") /*.Tables("User", "Student")*/
	query.AddCriteria(criteria)
	query.transform()
}

type Student struct {
	name   string
	age    int
	rollno int
}
