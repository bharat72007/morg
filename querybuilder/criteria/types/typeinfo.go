package types

var TypeMemInstance *TypeMem

type TypeMem struct {
	Fields []TypeInfo
}

type TypeInfo struct {
	Value interface{}
	Type  string
}

func NewTypeMem() *TypeMem {
	return &TypeMem{Fields: make([]TypeInfo, 0)}
}

func NewTypeInfo(value interface{}, typeinfo string) TypeInfo {
	return TypeInfo{Value: value, Type: typeinfo}
}

func init() {
	TypeMemInstance = NewTypeMem()
}
