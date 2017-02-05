package criteria

import (
	c "github.com/morg/querybuilder/commons"
)

type Criteria struct {
	//can have list of Restriction
	Restrictions []interface{}
	Combiners    []string
	Entity       interface{}
	Projections  []interface{}
}

func CreateCriteria(class interface{}) *Criteria {
	return &Criteria{Restrictions: make([]interface{}, 0), Combiners: make([]string, 0), Entity: class}
}

func (criteria *Criteria) Add(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	criteria.Combiners = append(criteria.Combiners, c.Keyword_And)

	return criteria
}

func (criteria *Criteria) AddP(projection interface{}) *Criteria {
	criteria.Projections = append(criteria.Projections, projection)
	return criteria
}

func (criteria *Criteria) Exclude(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	criteria.Combiners = append(criteria.Combiners, c.Keyword_Not)
	return criteria
}

func (criteria *Criteria) Or(restriction interface{}) *Criteria {
	criteria.Restrictions = append(criteria.Restrictions, restriction)
	criteria.Combiners = append(criteria.Combiners, c.Keyword_Or)
	return criteria
}
