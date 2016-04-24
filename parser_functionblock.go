package govorto

import (
	"strings"
)

func ParseFunctionBlock(lines []string, m *VortoModel) {
	begin := lines[0]
	begin = strings.TrimSuffix(begin, "{")
	name := strings.TrimSpace(begin)[14:]
	fb := &FunctionBlockModel{
		Name: name,
	}

	blockBuf := []string{}
	parseBlock := false
	doParse := false
	for _, line := range lines[1 : len(lines)-1] {
		if parseBlock {
			blockBuf = append(blockBuf, line)
			if strings.TrimSpace(line) == "}" {
				parseBlock = false
				doParse = true
			}
		} else {
			if strings.HasPrefix(line, "status") || strings.HasPrefix(line, "operations") {
				blockBuf = []string{line}
				parseBlock = true
			}
		}

		if doParse {
			ParseFunctionBlockBuf(blockBuf, fb)
			doParse = false
			blockBuf = []string{}
		}

	}
	m.FunctionBlock = fb
}

func ParseFunctionBlockBuf(lines []string, m *FunctionBlockModel) {
	line := lines[0]

	switch {
	case strings.HasPrefix(line, "status"):
		ParseFunctionBlockStatus(lines, m)
		break

	case strings.HasPrefix(line, "operations"):
		ParseFunctionBlockOperations(lines, m)
		break
	}
}

func ParseFunctionBlockStatus(lines []string, m *FunctionBlockModel) {
	lines = lines[1:len(lines)-1]

	for _, line := range lines {
		s := &FunctionBlockStatus{}

		s.Description = line[strings.Index(line, "\"")+1: strings.LastIndex(line, "\"")]

		frag := strings.Split(strings.TrimSpace(line[:strings.Index(line, "\"")]), " ")
		if frag[0] == "mandatory" {
			s.Mandatory = true
			s.Name = frag[1]
			s.DataType = frag[3]
		} else {
			s.Mandatory = false
			s.Name = frag[0]
			s.DataType = frag[2]
		}

		m.Status = append(m.Status, s)
	}
}

func ParseFunctionBlockOperations(lines []string, m *FunctionBlockModel) {
	lines = lines[1:len(lines)-1]

	for _, line := range lines {
		s := &FunctionBlockOperation{}

		s.Description = line[strings.Index(line, "\"")+1: strings.LastIndex(line, "\"")]

		frag := strings.Split(strings.TrimSpace(line[:strings.Index(line, "\"")]), " ")
		s.Name = frag[0]
		s.DataType = frag[2]

		m.Operations = append(m.Operations, s)
	}
}
