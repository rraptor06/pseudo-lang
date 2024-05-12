package convertor

import "strings"

func GetVariablesType(line string, inFuction bool) string {
	supplement := " "

	for index := 0; index < len(line); index++ {
		char := line[index]
		if char == ' ' || char == '\t' {
			continue
		}
		if inFuction && strings.HasPrefix(line[index:], "(adresse)") {
			supplement += "&"
			index += 8
			continue
		}
		if (inFuction && strings.HasPrefix(line[index:], "(valeur)")) || (!inFuction && strings.HasPrefix(line[index:], "(adresse)")) {
			supplement += "*"
			index += 7
			if !inFuction {
				index++
			}
			continue
		}
		if strings.TrimRight(line[index:], " \t") == "entier" {
			return "int" + supplement
		}
		if strings.TrimRight(line[index:], " \t") == "decimal" {
			return "double " + supplement
		}
		if strings.TrimRight(line[index:], " \t") == "vide" {
			return "void " + supplement
		}
		break
	}
	return ""
}

func IsValidVariableName(name string) bool {
	for _, char := range name {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < '0' || char > '9') && char != '_' {
			return false
		}
	}
	return true
}
