package govorto

import (
	"regexp"
	"strings"
)

func ParseDatatype(lines []string, m *VortoModel) {
	begin := lines[0]
	begin = strings.TrimSuffix(begin, "{")
	enumName := strings.TrimSpace(begin)[5:]

	enumVals := lines[1 : len(lines)-1]

	dt := &DataType{
		Fields: []*DataTypeField{},
	}
	expr, _ := regexp.Compile("(.*) \"(.*)\"")
	for _, v := range enumVals {
		vs := expr.FindStringSubmatch(v)

		tf := &DataTypeField{
			Name:        vs[1],
			Description: vs[2],
		}

		dt.Fields = append(dt.Fields, tf)
	}

	dt.Name = enumName

	m.DataType = dt
}
