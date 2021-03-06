package common

const (
	Operation_Gt    = ">"
	Operation_Lt    = "<"
	Operation_Lte   = "=<"
	Operation_Gte   = ">="
	Operation_Eq    = "="
	Operation_NotEq = "<>"
	Operation_Like  = "LIKE"
	Operation_ILike = "ILIKE"

	Keyword_GroupBy  = "GROUP BY"
	Keyword_OrderBy  = "ORDER BY"
	Keyword_SELECT   = "SELECT"
	Keyword_Alias    = "AS"
	Keyword_Limit    = "LIMIT"
	Keyword_Offset   = "OFFSET"
	Keyword_Distinct = "DISTINCT"
	Keyword_Between  = "BETWEEN"
	Keyword_And      = "AND"
	Keyword_Not      = "NOT"
	Keyword_Or       = "OR"
	Keyword_In       = "IN"
	Keyword_Where    = "WHERE"
	Keyword_From     = "FROM"

	Type_Alias           = "ALIAS"
	Type_Limit           = "Limit"
	Type_Offset          = "Offset"
	Type_Project         = "PROJECT"
	Type_Space           = " "
	Type_AllColumns      = "*"
	Type_Between         = "BETWEEN"
	Type_Group           = "GROUP"
	Type_Order           = "ORDER"
	Type_In              = "In"
	Type_Distinct        = "DISTINCT"
	Type_Aggregate_Min   = "Min"
	Type_Aggregate_Count = "Count"
	Type_Aggregate_Avg   = "Avg"
	Type_Aggregate_Sum   = "Sum"
	Type_Aggregate_Max   = "Max"
	Type_Operation_Gt    = "Gt"
	Type_Operation_Lt    = "Lt"
	Type_Operation_Gte   = "Gte"
	Type_Operation_Lte   = "Lte"
	Type_Operation_Eq    = "Equal"
	Type_Operation_NotEq = "NotEq"
	Type_Operation_Like  = "Like"
	Type_Operation_ILike = "ILike"

	Aggregate_Sum   = "SUM"
	Aggregate_Avg   = "AVG"
	Aggregate_Count = "COUNT"
	Aggregate_Min   = "MIN"
	Aggregate_Max   = "MAX"
)

func Findtype(value interface{}) string {

	switch value.(type) {
	case string:
		return "string"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case int:
		return "int"
	case int64:
		return "int64"
	default:
		return "interface{}"
	}
}
