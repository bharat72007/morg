package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
)

/*type A struct {
	name string
	id   int `dep:"DEP"`
}

func main() {
	a := A{name: "abcd", id: 22}
	st := reflect.TypeOf(a)
	num := reflect.TypeOf(a).NumField()
	for i := 0; i < num; i++ {
		fmt.Println(st.Field(i).Tag.Get("dep"))
	}

	//	fmt.Println(reflect.TypeOf(a).Field(1).Tag)
}
*/

/*type User struct {
	id   int    `crud:"pk, colname:Id"`
	name string `crud:"colname:Name"`
}

func main() {
	user := User{id: 12, name: "abcde"}
	st := reflect.TypeOf(user)
	num := st.NumField()
	for i := 0; i < num; i++ {
		fmt.Println(st.Field(i).Tag.Get("crud"))
	}
}


*/

type ORM struct {
	TableName  string
	Dbname     string
	ColumnName map[string]string //If column name is not present, then fieldname is taken for column names.
	PK         string            //Assuming Only One Field as PK, i.e No composite Key
	NotNull    []string
	Unique     []string
}

var orm ORM

type User struct {
	id   int    `orm:"pk;tblname:Id;notnull;unique"`
	name string `orm:"tblname:Name"`
}

func main() {
	user := User{id: 11, name: "Cecil"}
	st := reflect.TypeOf(user)
	fieldnums := st.NumField()
	orm = ORM{TableName: "", Dbname: "", ColumnName: make(map[string]string), PK: "", NotNull: make([]string, fieldnums), Unique: make([]string, fieldnums)}

	for i := 0; i < fieldnums; i++ {
		fmt.Println(st.Field(i).Tag.Get("orm"))
		parseTag(st.Field(i).Tag.Get("orm"), st.Field(i).Name)
	}
	fmt.Printf("Primary Key is: %v \n", orm.PK)
	for k, v := range orm.ColumnName {
		fmt.Printf("Field: %v\n ", k)
		fmt.Printf("Column Name: %v \n ", v)
	}
	fmt.Println(orm.Unique)
	fmt.Println(orm.NotNull)
}

func parseTag(tag string, fieldname string) {
	tokens := strings.Split(tag, ";")
	for _, token := range tokens {
		if strings.Contains(token, ":") {
			if t := strings.Split(token, ":"); t[0] == "tblname" {
				orm.ColumnName[fieldname] = t[1]
			}

		} else {
			if token == "pk" {
				orm.PK = fieldname
			} else if token == "notnull" {
				orm.NotNull = append(orm.NotNull, fieldname)
			} else if token == "unique" {
				orm.Unique = append(orm.Unique, fieldname)
			}
		}
	}
}

func get(user User) {
	//"SELECT * FROM UserTable WHERE Id = 5 limit 1"
	tableName := orm.TableName
	pk := orm.ColumnName[orm.PK]
	query := fmt.Sprintf("SELECT * FROM %v WHERE %v = ?", tableName, pk)
	var args []interface{}
	args = append(args, user.id)
	rs, err := sql.DB.Prepare(query, args)
}

func delete(user User) {
	//"SELECT * FROM UserTable WHERE Id = 5 limit 1"
	tableName := orm.TableName
	pk := orm.ColumnName[orm.PK]
	query := fmt.Sprintf("DELETE FROM %v WHERE %v = ?", tableName, pk)
	var args []interface{}
	args = append(args, user.id)
	rs, err := sql.DB.Prepare(query, args)
}
