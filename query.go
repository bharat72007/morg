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
	Projections  []interface{}
}

type Restriction struct {
	column string
	value  interface{}
	rtype  string
}

type Projection struct {
	column string
	ptype  string
}

type AliasProjection struct {
	column string
	ptype  string
	alias  string
}

type BetweenRestriction struct {
	column string
	lbound interface{}
	ubound interface{}
	rtype  string
}

type InRestriction struct {
	column string
	list   []interface{}
	rtype  string
}

type LimitRestriction struct {
	rowscount int
	rtype     string
}

type OffsetRestriction struct {
	offset int
	rtype  string
}

type OrderByRestriction struct {
	column string
	order  string
	rtype  string
}

type GroupByRestriction struct {
	column string
	rtype  string
}

func CreateCriteria(class interface{}) *Criteria {
	return &Criteria{Restrictions: make([]interface{}, 0), combiners: make([]string, 0), Entity: class}
}

func NewBetweenRestriction(col string, lval, rval interface{}, rtype string) *BetweenRestriction {
	return &BetweenRestriction{column: col, lbound: lval, ubound: rval, rtype: rtype}
}

func NewAliasProjection(col, aliasname string, rtype string) *AliasProjection {
	return &AliasProjection{column: col, alias: aliasname, ptype: rtype}
}

func NewOrderByRestriction(col, order, rtype string) *OrderByRestriction {
	return &OrderByRestriction{column: col, order: order, rtype: rtype}
}

func NewGroupByRestriction(col, rtype string) *GroupByRestriction {
	return &GroupByRestriction{column: col, rtype: rtype}
}

func NewLimitRestriction(rows int, rtype string) *LimitRestriction {
	return &LimitRestriction{rowscount: rows, rtype: rtype}
}

func NewOffsetRestriction(start int, rtype string) *OffsetRestriction {

	return &OffsetRestriction{offset: start, rtype: rtype}
}
func NewInRestriction(col string, vals []interface{}, rtype string) *InRestriction {
	return &InRestriction{column: col, list: vals, rtype: rtype}
}

func NewRestriction(col string, val interface{}, rtype string) *Restriction {
	return &Restriction{column: col, value: val, rtype: rtype}
}

func NewProjection(col, ptype string) *Projection {
	return &Projection{column: col, ptype: ptype}
}

func NewQuery() *Query {
	return &Query{}
}

func (criteria *Criteria) add(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	criteria.combiners = append(criteria.combiners, "AND")
	return criteria
}

func (criteria *Criteria) addP(projection interface{}) *Criteria {
	criteria.Projections = append(criteria.Projections, projection)
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
	strs := restriction.column + " " + restriction.rtype + " " + restriction.value.(string)
	return strs
}

func (projection *AliasProjection) tostring() string {
	strs := projection.column + "  AS " + projection.alias
	return strs
}

func (projection *Projection) tostring() string {
	if projection.ptype == "addcolumn" {
		return projection.column
	} else if projection.ptype == "distinct" {
		return projection.ptype + " " + projection.column
	} else {
		return projection.ptype + "( " + projection.column + " )"
	}
}

func (restriction *OrderByRestriction) tostring() string {
	strs := restriction.column + " " + restriction.order
	return strs
}

func (restriction *GroupByRestriction) tostring() string {
	strs := restriction.column
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

func (projection *Projection) Distinct(col string) *Projection {
	return NewProjection(col, "distinct")
}

func (projection *AliasProjection) Alias(col, aliasname string) *AliasProjection {
	return NewAliasProjection(col, aliasname, "alias")
}

func (projection *Projection) Sum(col string) *Projection {
	return NewProjection(col, "sum")
}

func (projection *Projection) Avg(col string) *Projection {
	return NewProjection(col, "avg")
}

func (projection *Projection) Count(col string) *Projection {
	return NewProjection(col, "count")
}

func (projection *Projection) Max(col string) *Projection {
	return NewProjection(col, "max")
}

func (projection *Projection) Min(col string) *Projection {
	return NewProjection(col, "min")
}

func (projection *Projection) Column(col string) *Projection {
	return NewProjection(col, "addcolumn")
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

func (restriction *Restriction) OrderBy(col, order string) *OrderByRestriction {
	return NewOrderByRestriction(col, order, "ORDER")
}

func (restriction *Restriction) GroupBy(col string) *GroupByRestriction {
	return NewGroupByRestriction(col, "GROUPBY")
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

//TODO
/*func (criteria *Criteria) List() []interface{} {
	newquery := Query{Projection: "*", From: reflect.TypeOf(criteria.Entity).Elem().Name()}
	newquery.transform()
	return nil
}*/

func main() {
	//Example #1
	var student Student = Student{name: "abc", age: 9, rollno: 100}
	criteria := CreateCriteria(&student)
	var restriction Restriction
	ag := make([]interface{}, 0)
	ag = append(ag, "10")
	ag = append(ag, "20")
	var projection Projection
	var aliasprojection AliasProjection

	criteria.exclude(restriction.Equal("name", "mane")).Or(restriction.Equal("age", "22")).add(restriction.Equal("age", "22")).add(restriction.Gt("rollno", "900")).add(restriction.NotEqual("age", "80")).add(restriction.Between("name", "00", "9")).add(restriction.In("age", ag)).add(restriction.Limit(10)).add(restriction.Offset(2)).add(restriction.GroupBy("name")).add(restriction.GroupBy("age")).add(restriction.OrderBy("age", "ASC")).add(restriction.OrderBy("name", "DESC"))
	//(NOT name = "mane") AND age = 22 AND rollno > 900 AND age <> 80 AND (re OR rs)
	criteria.addP(projection.Sum("age")).addP(projection.Count("*")).addP(projection.Distinct("name")).addP(aliasprojection.Alias("name", "StudentName"))
	query := NewQuery()
	query.Project("name", "age", "rollno", "Student.name") /*.Tables("User", "Student")*/
	query.AddCriteria(criteria)
	query.transform()

	//Example2 SELECT COUNT(*) FROM STUDENT WHERE age > 10
	criteria2 := CreateCriteria(&student)
	var restriction2 Restriction
	var projection2 Projection
	criteria2.addP(projection2.Count("*")).add(restriction2.Gt("age", "10"))
	query2 := NewQuery()
	query2.AddCriteria(criteria2)
	query2.Project()
	query2.transform()

	//Example3 SELECT count(*) AS count FROM STUDENT where age >10 and name like '%abc%'
	criteria3 := CreateCriteria(&student)
	var restriction3 Restriction
	var projection3 Projection
	var aliasprojection2 AliasProjection
	criteria3.addP(projection3.Count("*")).addP(aliasprojection2.Alias("", "count"))
	criteria3.add(restriction3.Gt("age", "10")).add(restriction3.Like("name", "%ab%"))
	query3 := NewQuery()
	query3.AddCriteria(criteria3)
	query3.transform()

	//Example4 SELECT count(*), DISTINCT name, age FROM Student
	criteria4 := CreateCriteria(&student)
	var projection4 Projection
	criteria4.addP(projection4.Count("*")).addP(projection4.Distinct("name")).addP(projection.Column("age"))
	query4 := NewQuery()
	query4.AddCriteria(criteria4)
	query4.transform()

	//TODO
	//Example SELECT * FROM Student
	/*	criteria5 := CreateCriteria(&student)
		criteria5.List()
	*/

	//Example5 SELECT * FROM Student GroupBy age
	criteria5 := CreateCriteria(&student)
	var projection5 Projection
	var restriction5 Restriction
	criteria5.addP(projection5.Column("*"))
	criteria5.add(restriction5.GroupBy("age"))
	query5 := NewQuery()
	query5.AddCriteria(criteria5)
	query5.transform()

	//Example6 SELECT name, age FROM Student Groupby age Orderby name DESC
	criteria6 := CreateCriteria(&student)
	var projection6 Projection
	var restriction6 Restriction
	criteria6.addP(projection6.Column("name")).addP(projection6.Column("age"))
	criteria6.add(restriction6.GroupBy("age")).add(restriction6.OrderBy("name", "DESC"))
	query6 := NewQuery()
	query6.AddCriteria(criteria6)
	query6.transform()

}

func (query *Query) transform() {
	var tokens []string = make([]string, 0)
	var orderbyflag bool = false
	var groupbyflag bool = false
	var whereflag bool = false
	//var ent string
	if query.Projection != "" {
		tokens = append(tokens, "SELECT", query.Projection)
	} else {
		tokens = append(tokens, "SELECT")
	}
	if len(query.Filter.Projections) > 0 {
		for i := 0; i < len(query.Filter.Projections); i++ {
			switch query.Filter.Projections[i].(type) {
			case *Projection:

				tokens = append(tokens, query.Filter.Projections[i].(*Projection).tostring())
			}
		}
	}
	if len(query.Filter.Projections) > 0 {
		for i := 0; i < len(query.Filter.Projections); i++ {
			switch query.Filter.Projections[i].(type) {
			case *AliasProjection:

				tokens = append(tokens, query.Filter.Projections[i].(*AliasProjection).tostring())
			}
		}
	}

	if query.Filter.Entity != "" {
		tokens = append(tokens, "FROM", reflect.TypeOf(query.Filter.Entity).Elem().Name())
	}
	if len(query.Filter.Restrictions) > 0 {
		//TODO
		for i := 0; i < len(query.Filter.Restrictions); i++ {
			switch t := query.Filter.Restrictions[i].(type) {
			case *Restriction:
				if !whereflag {
					tokens = append(tokens, "WHERE")
					whereflag = true
				}
				if query.Filter.combiners[i] == "NOT" {
					tokens = append(tokens, "( NOT", query.Filter.Restrictions[i].(*Restriction).tostring(), ")")
				} else if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*Restriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.combiners[i], query.Filter.Restrictions[i].(*Restriction).tostring())
				}

			case *BetweenRestriction:
				if !whereflag {
					tokens = append(tokens, "WHERE")
					whereflag = true
				}
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*BetweenRestriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.combiners[i], query.Filter.Restrictions[i].(*BetweenRestriction).tostring())
				}

			case *InRestriction:
				if !whereflag {
					tokens = append(tokens, "WHERE")
					whereflag = true
				}
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*InRestriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.combiners[i], query.Filter.Restrictions[i].(*InRestriction).tostring())
				}

			case *LimitRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*LimitRestriction).tostring())
			case *OffsetRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*OffsetRestriction).tostring())
			case *GroupByRestriction:
				if !groupbyflag {
					tokens = append(tokens, "GROUP BY")
					groupbyflag = true
				}
				tokens = append(tokens, query.Filter.Restrictions[i].(*GroupByRestriction).tostring())
			case *OrderByRestriction:
				if !orderbyflag {
					tokens = append(tokens, "ORDER BY")
					orderbyflag = true
				}
				tokens = append(tokens, query.Filter.Restrictions[i].(*OrderByRestriction).tostring())

			default:
				fmt.Printf("%T", t)
			}
		}

	}

	fmt.Println(tokens)
}

type Student struct {
	name   string
	age    int
	rollno int
}
