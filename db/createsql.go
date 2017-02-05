package db

import (
	"fmt"
	ty "github.com/morg/querybuilder/criteria/types"
	"strings"
)

func CreateSQL(tokens []string, arguments []ty.TypeInfo) {
	var query string = strings.Join(tokens, " ")
	fmt.Println(query)
	args := make([]interface{}, 0)
	if arguments == nil {
		return
	}
	for _, value := range arguments {
		args = append(args, value.Value)
	}
	fmt.Println(args)
}
