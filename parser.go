package govorto

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ParseNamespace(txt string, m *VortoModel) {
	m.Namespace = txt[10:]
}

func ParseVersion(txt string, m *VortoModel) {
	m.Version = txt[8:]
}

func ParseDisplayName(txt string, m *VortoModel) {
	m.Displayname = strings.Trim(txt[12:], "\"")
}

func ParseDescription(txt string, m *VortoModel) {
	m.Description = strings.Trim(txt[12:], "\"")
}
func ParseCategory(txt string, m *VortoModel) {
	m.Category = txt[9:]
}

func ParseUsing(txt string, m *VortoModel) {
	m.Using = append(m.Using, txt[6:])
}

func ParseBuf(bufArr []string, m *VortoModel) {
	buf := bufArr[0]

	switch {
	case strings.HasPrefix(buf, "namespace"):
		ParseNamespace(buf, m)
		break

	case strings.HasPrefix(buf, "version"):
		ParseVersion(buf, m)
		break

	case strings.HasPrefix(buf, "displayname"):
		ParseDisplayName(buf, m)
		break

	case strings.HasPrefix(buf, "description"):
		ParseDescription(buf, m)
		break

	case strings.HasPrefix(buf, "category"):
		ParseCategory(buf, m)
		break

	case strings.HasPrefix(buf, "using"):
		ParseUsing(buf, m)
		break

	case strings.HasPrefix(buf, "enum"):
		ParseDatatype(bufArr, m)
		break

	case strings.HasPrefix(buf, "functionblock"):
		ParseFunctionBlock(bufArr, m)
		break

	case strings.HasPrefix(buf, "infomodel"):
		ParseInfoModel(bufArr, m)
		break
	}
}

func ParseFile(filename string) *VortoModel {
	m := NewVortoModel()
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	for idx, line := range lines {
		txt := line

		if strings.HasPrefix(txt, "namespace") ||
		strings.HasPrefix(txt, "version") ||
		strings.HasPrefix(txt, "displayname") ||
		strings.HasPrefix(txt, "description") ||
		strings.HasPrefix(txt, "category") ||
		strings.HasPrefix(txt, "using") {
			ParseBuf([]string{txt}, m)
		} else {
			if strings.HasPrefix(txt, "infomodel") ||
			strings.HasPrefix(txt, "enum") ||
			strings.HasPrefix(txt, "functionblock") {
				ParseBuf(lines[idx:len(lines)], m)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m
}
