package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Query struct {
	Projection string
	Filter     *Criteria
	From       string
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

type LimitRestriction struct {
	rowscount int
	resttype  string
}

type OffsetRestriction struct {
	offset   int
	resttype string
}

func CreateCriteria(class interface{}) *Criteria {
	return &Criteria{Restrictions: make([]interface{}, 0), combiners: make([]string, 0), Entity: class}
}

func NewBetweenRestriction(col string, lval, rval interface{}, rtype string) *BetweenRestriction {
	return &BetweenRestriction{column: col, lbound: lval, ubound: rval, resttype: rtype}
}

func NewLimitRestriction(rows int, rtype string) *LimitRestriction {
	return &LimitRestriction{rowscount: rows, resttype: rtype}
}

func NewOffsetRestriction(start int, rtype string) *OffsetRestriction {

	return &OffsetRestriction{offset: start, resttype: rtype}
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
	criteria.combiners = append(criteria.combiners, "AND")
	return criteria
}

func (criteria *Criteria) exclude(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	criteria.combiners = append(criteria.combiners, "NOT")
	return criteria
}

func (criteria *Criteria) Or(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	criteria.combiners = append(criteria.combiners, "OR")
	return criteria
}

func (restriction *Restriction) tostring() string {

	strs := restriction.column + " " + restriction.resttype + " " + restriction.value.(string)
	return strs
}

func (restriction *LimitRestriction) tostring() string {
	strs := "LIMIT" + " " + strconv.Itoa(restriction.rowscount)
	return strs
}

func (restriction *OffsetRestriction) tostring() string {
	strs := "OFFSET" + " " + strconv.Itoa(restriction.offset)
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

func (restriction *Restriction) Limit(rows int) *LimitRestriction {
	return NewLimitRestriction(rows, "LIMIT")
}

func (restriction *Restriction) Offset(start int) *OffsetRestriction {
	return NewOffsetRestriction(start, "OFFSET")
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

func (criteria *Criteria) List() []interface{} {
	newquery := Query{Projection: "*", From: reflect.TypeOf(criteria.Entity).Elem().Name()}
	newquery.transform()
	return nil
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
				if query.Filter.combiners[i] == "NOT" {
					tokens = append(tokens, "( NOT", query.Filter.Restrictions[i].(*Restriction).tostring(), ")")
				} else if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*Restriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.combiners[i], query.Filter.Restrictions[i].(*Restriction).tostring())
				}

			case *BetweenRestriction:
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*BetweenRestriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.combiners[i], query.Filter.Restrictions[i].(*BetweenRestriction).tostring())
				}

			case *InRestriction:
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*InRestriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.combiners[i], query.Filter.Restrictions[i].(*InRestriction).tostring())
				}

			case *LimitRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*LimitRestriction).tostring())
			case *OffsetRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*OffsetRestriction).tostring())

			default:
				fmt.Printf("%T", t)
			}
		}

	}

	fmt.Println(tokens)
}

func main() {
	var student Student = Student{name: "abc", age: 9, rollno: 100}
	criteria := CreateCriteria(&student)
	var restriction Restriction
	ag := make([]interface{}, 0)
	ag = append(ag, "10")
	ag = append(ag, "20")
	criteria.exclude(restriction.Equal("name", "mane")).Or(restriction.Equal("age", "22")).add(restriction.Equal("age", "22")).add(restriction.Gt("rollno", "900")).add(restriction.NotEqual("age", "80")).add(restriction.Between("name", "00", "9")).add(restriction.In("age", ag)).add(restriction.Limit(10)).add(restriction.Offset(2))
	//(NOT name = "mane") AND age = 22 AND rollno > 900 AND age <> 80 AND (re OR rs)
	query := NewQuery()
	query.Project("name", "age", "rollno", "Student.name") /*.Tables("User", "Student")*/
	query.AddCriteria(criteria)
	query.transform()

	//Another Example
	//criteria2 := CreateCriteria(&student)
	//This will return the list of all records. criteria2.List() ==>SELECT * FROM Student

}

type Student struct {
	name   string
	age    int
	rollno int
}
