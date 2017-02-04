package criteria

import (
	"fmt"
	c "github.com/morg/querybuilder/commons"
	pj "github.com/morg/querybuilder/criteria/projections"
	cr "github.com/morg/querybuilder/criteria/restrictions"
	"reflect"
	"strings"
)

type Query struct {
	Projection string
	Filter     *Criteria
	From       string
}

func (query *Query) AddCriteria(criteria *Criteria) *Query {
	query.Filter = criteria
	return query
}

func NewQuery() *Query {
	return &Query{}
}

func (query *Query) Project(fields ...string) *Query {
	query.Projection = strings.Join(fields, ",")
	return query
}

func (query *Query) Transform() {
	var tokens []string = make([]string, 0)
	var orderbyflag bool = false
	var groupbyflag bool = false
	var whereflag bool = false
	//var ent string
	if query.Projection != "" {
		tokens = append(tokens, c.Keyword_SELECT, query.Projection)
	} else {
		tokens = append(tokens, c.Keyword_SELECT)
	}
	if len(query.Filter.Projections) > 0 {
		for i := 0; i < len(query.Filter.Projections); i++ {
			switch query.Filter.Projections[i].(type) {
			case *pj.Projection:
				tokens = append(tokens, query.Filter.Projections[i].(*pj.Projection).Tostring())
			}
		}
	}
	if len(query.Filter.Projections) > 0 {
		for i := 0; i < len(query.Filter.Projections); i++ {
			switch query.Filter.Projections[i].(type) {
			case *pj.AliasProjection:
				tokens = append(tokens, query.Filter.Projections[i].(*pj.AliasProjection).Tostring())
			}
		}
	}
	if query.Filter.Entity != "" {
		tokens = append(tokens, c.Keyword_From, reflect.TypeOf(query.Filter.Entity).Elem().Name())
	}
	if len(query.Filter.Restrictions) > 0 {
		//TODO
		for i := 0; i < len(query.Filter.Restrictions); i++ {
			switch t := query.Filter.Restrictions[i].(type) {
			case *cr.Restriction:
				if !whereflag {
					tokens = append(tokens, c.Keyword_Where)
					whereflag = true
				}
				if query.Filter.Combiners[i] == c.Keyword_Not {
					tokens = append(tokens, "( "+c.Keyword_Not, query.Filter.Restrictions[i].(*cr.Restriction).Tostring(), ")")
				} else if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*cr.Restriction).Tostring())
				} else {
					tokens = append(tokens, query.Filter.Combiners[i], query.Filter.Restrictions[i].(*cr.Restriction).Tostring())
				}

			case *cr.BetweenRestriction:
				if !whereflag {
					tokens = append(tokens, c.Keyword_Where)
					whereflag = true
				}
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*cr.BetweenRestriction).Tostring())
				} else {
					tokens = append(tokens, query.Filter.Combiners[i], query.Filter.Restrictions[i].(*cr.BetweenRestriction).Tostring())
				}

			case *cr.InRestriction:
				if !whereflag {
					tokens = append(tokens, c.Keyword_Where)
					whereflag = true
				}
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*cr.InRestriction).Tostring())
				} else {
					tokens = append(tokens, query.Filter.Combiners[i], query.Filter.Restrictions[i].(*cr.InRestriction).Tostring())
				}

			case *cr.LimitRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*cr.LimitRestriction).Tostring())
			case *cr.OffsetRestriction:
				tokens = append(tokens, query.Filter.Restrictions[i].(*cr.OffsetRestriction).Tostring())
			case *cr.GroupByRestriction:
				if !groupbyflag {
					tokens = append(tokens, c.Keyword_GroupBy)
					groupbyflag = true
				}
				tokens = append(tokens, query.Filter.Restrictions[i].(*cr.GroupByRestriction).Tostring())
			case *cr.OrderByRestriction:
				if !orderbyflag {
					tokens = append(tokens, c.Keyword_OrderBy)
					orderbyflag = true
				}
				tokens = append(tokens, query.Filter.Restrictions[i].(*cr.OrderByRestriction).Tostring())

			default:
				fmt.Printf("%T", t)
			}
		}
	}
	fmt.Println(tokens)
}
