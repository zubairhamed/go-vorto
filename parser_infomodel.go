package govorto

import (
	"strings"
)

func ParseInfoModel(lines []string, m *VortoModel) {
	begin := lines[0]
	begin = strings.TrimSuffix(begin, "{")
	name := strings.TrimSpace(begin)[10:]
	im := &InformationModel{
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
			if strings.HasPrefix(line, "functionblocks") {
				blockBuf = []string{line}
				parseBlock = true
			}
		}

		if doParse {
			ParseInformationModelBlockBuf(blockBuf, im)
			doParse = false
			blockBuf = []string{}
		}

	}
	m.InfoModel = im
}

func ParseInformationModelBlockBuf(lines []string, m *InformationModel) {
	line := lines[0]

	switch {
	case strings.HasPrefix(line, "functionblocks"):
		ParseInformationModelFunctionBlocks(lines, m)
		break
	}
}

func ParseInformationModelFunctionBlocks(lines []string, m *InformationModel) {
	lines = lines[1:len(lines)-1]
	for _, line := range lines {
		s := &FunctionBlock{}
		s.Description = line[strings.Index(line, "\"")+1: strings.LastIndex(line, "\"")]
		frag := strings.Split(strings.TrimSpace(line[:strings.Index(line, "\"")]), " ")

		s.Name = frag[0]
		s.DataType = frag[2]

		m.FunctionBlocks = append(m.FunctionBlocks, s)
	}
}
