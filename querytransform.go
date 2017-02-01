package morg

import (
	"strings"
)

type SQLQuery struct {
	query string
}

func NewSQLQuery() *SQLQuery {
	return &SQLQuery{query: nil}
}

func (query *Query) transform() *SQLQuery {
	var trquery string
	if query.Select {
		trquery = "SELECT"
	}
	trquery = trquery + strings.Join(query.Projection, ",")
	trquery = trquery + query.From
	if query.Filter != nil {
		trquery = trquery + "WHERE" + strings.Join(query.Filter.Restrictions, " ")
	}
	sqlquery := NewSQLQuery()
	sqlquery.query = trquery
	return &sqlquery
}
