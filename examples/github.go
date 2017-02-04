package examples

import (
	cr "github.com/morg/querybuilder/criteria"
	pj "github.com/morg/querybuilder/criteria/projections"
	rs "github.com/morg/querybuilder/criteria/restrictions"
)

type Student struct {
	name   string
	age    int
	rollno int
}

func main() {
	//Example #1
	var student Student = Student{name: "abc", age: 9, rollno: 100}
	criteria := cr.CreateCriteria(&student)
	var restriction rs.Restriction
	ag := make([]interface{}, 0)
	ag = append(ag, "10")
	ag = append(ag, "20")
	var projection pj.Projection
	var aliasprojection pj.AliasProjection

	criteria.Exclude(restriction.Equal("name", "mane")).Or(restriction.Equal("age", "22")).Add(restriction.Equal("age", "22")).Add(restriction.Gt("rollno", "900")).Add(restriction.NotEqual("age", "80")).Add(restriction.Between("name", "00", "9")).Add(restriction.In("age", ag)).Add(restriction.Limit(10)).Add(restriction.Offset(2)).Add(restriction.GroupBy("name")).Add(restriction.GroupBy("age")).Add(restriction.OrderBy("age", "ASC")).Add(restriction.OrderBy("name", "DESC"))
	//(NOT name = "mane") AND age = 22 AND rollno > 900 AND age <> 80 AND (re OR rs)
	criteria.AddP(projection.Sum("age")).AddP(projection.Count("*")).AddP(projection.Distinct("name")).AddP(aliasprojection.Alias("name", "StudentName"))
	query := cr.NewQuery()
	query.Project("name", "age", "rollno", "Student.name") /*.Tables("User", "Student")*/
	query.AddCriteria(criteria)
	query.Transform()

	//Example2 SELECT COUNT(*) FROM STUDENT WHERE age > 10
	criteria2 := cr.CreateCriteria(&student)
	var restriction2 rs.Restriction
	var projection2 pj.Projection
	criteria2.AddP(projection2.Count("*")).Add(restriction2.Gt("age", "10"))
	query2 := cr.NewQuery()
	query2.AddCriteria(criteria2)
	query2.Project()
	query2.Transform()

	//Example3 SELECT count(*) AS count FROM STUDENT where age >10 and name like '%abc%'
	criteria3 := cr.CreateCriteria(&student)
	var restriction3 rs.Restriction
	var projection3 pj.Projection
	var aliasprojection2 pj.AliasProjection
	criteria3.AddP(projection3.Count("*")).AddP(aliasprojection2.Alias("", "count"))
	criteria3.Add(restriction3.Gt("age", "10")).Add(restriction3.Like("name", "%ab%"))
	query3 := cr.NewQuery()
	query3.AddCriteria(criteria3)
	query3.Transform()

	//Example4 SELECT count(*), DISTINCT name, age FROM Student
	criteria4 := cr.CreateCriteria(&student)
	var projection4 pj.Projection
	criteria4.AddP(projection4.Count("*")).AddP(projection4.Distinct("name")).AddP(projection.Column("age"))
	query4 := cr.NewQuery()
	query4.AddCriteria(criteria4)
	query4.Transform()

	//TODO
	//Example SELECT * FROM Student
	/*	criteria5 := CreateCriteria(&student)
		criteria5.List()
	*/

	//Example5 SELECT * FROM Student GroupBy age
	criteria5 := cr.CreateCriteria(&student)
	var projection5 pj.Projection
	var restriction5 rs.Restriction
	criteria5.AddP(projection5.Column("*"))
	criteria5.Add(restriction5.GroupBy("age"))
	query5 := cr.NewQuery()
	query5.AddCriteria(criteria5)
	query5.Transform()

	//Example6 SELECT name, age FROM Student Groupby age Orderby name DESC
	criteria6 := cr.CreateCriteria(&student)
	var projection6 pj.Projection
	var restriction6 rs.Restriction
	criteria6.AddP(projection6.Column("name")).AddP(projection6.Column("age"))
	criteria6.Add(restriction6.GroupBy("age")).Add(restriction6.OrderBy("name", "DESC"))
	query6 := cr.NewQuery()
	query6.AddCriteria(criteria6)
	query6.Transform()
}
