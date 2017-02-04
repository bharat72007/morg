package projections

import c "github.com/morg/querybuilder/commons"

func (projection *Projection) Sum(col string) *Projection {
	return NewProjection(col, c.Type_Aggregate_Sum)
}

func (projection *Projection) Avg(col string) *Projection {
	return NewProjection(col, c.Type_Aggregate_Avg)
}

func (projection *Projection) Count(col string) *Projection {
	return NewProjection(col, c.Type_Aggregate_Count)
}

func (projection *Projection) Max(col string) *Projection {
	return NewProjection(col, c.Type_Aggregate_Max)
}

func (projection *Projection) Min(col string) *Projection {
	return NewProjection(col, c.Type_Aggregate_Min)
}
