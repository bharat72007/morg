package transform

import (
	"fmt"
	"strings"
)

func transform(query *Query) {
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
