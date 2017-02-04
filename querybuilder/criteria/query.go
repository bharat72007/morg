package criteria

import (
	c "github.com/morg/querybuilder/commons"
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
				if query.Filter.Combiners[i] == "NOT" {
					tokens = append(tokens, "( NOT", query.Filter.Restrictions[i].(*Restriction).tostring(), ")")
				} else if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*Restriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.Combiners[i], query.Filter.Restrictions[i].(*Restriction).tostring())
				}

			case *BetweenRestriction:
				if !whereflag {
					tokens = append(tokens, "WHERE")
					whereflag = true
				}
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*BetweenRestriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.Combiners[i], query.Filter.Restrictions[i].(*BetweenRestriction).tostring())
				}

			case *InRestriction:
				if !whereflag {
					tokens = append(tokens, "WHERE")
					whereflag = true
				}
				if i == 0 || i == len(query.Filter.Restrictions)-1 {
					tokens = append(tokens, query.Filter.Restrictions[i].(*InRestriction).tostring())
				} else {
					tokens = append(tokens, query.Filter.Combiners[i], query.Filter.Restrictions[i].(*InRestriction).tostring())
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
