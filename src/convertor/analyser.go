package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

type Instructions struct {
	Name     string
	Function func(string, *int) string
}

func analyseFunctionContent(instructionsList []*Instructions, function *FunctionStruct) int {
	var new_line string
	added := false
	indentation := 0

	for _, line := range function.Content {
		added = false
		if line == "" {
			continue
		}
		for _, instruction := range instructionsList {
			if strings.Contains(line, instruction.Name) {
				new_line = instruction.Function(line, &indentation)
				if new_line != "" {
					return 1
				}
				function.ConvertedContent = append(function.ConvertedContent, new_line)
				added = true
				break
			}
		}
		if added == false && strings.Contains(line, "->") {
			if analyseVariable(line, &indentation) != "" {
				return 1
			}
			function.ConvertedContent = append(function.ConvertedContent, new_line)
			added = true
		}
		if added == false && strings.Contains(line, "(") {
			if analyseFunction(line, &indentation) != "" {
				return 1
			}
			function.ConvertedContent = append(function.ConvertedContent, new_line)
			added = true
		}
		if added == false {
			fmt.Fprintf(os.Stderr, "%sERROR: Can't convert the line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
			return 1
		}
	}
	return 0
}

func analyseCode(code *CodeStruct) int {
	var instructionsList []*Instructions = []*Instructions{
		&Instructions{
			Name:     "retourner:",
			Function: analyseReturn,
		},
		&Instructions{
			Name:     "tant que:",
			Function: analyseWhile,
		},
		&Instructions{
			Name:     "sinon si:",
			Function: analyseElseIf,
		},
		&Instructions{
			Name:     "si:",
			Function: analyseIf,
		},
		&Instructions{
			Name:     "else:",
			Function: analyseElse,
		},
		&Instructions{
			Name:     ":",
			Function: analyseInitVariable,
		},
	}

	for _, function := range code.FunctionsList {
		if analyseFunctionContent(instructionsList, function) != 0 {
			//return 1
		}
	}
	if analyseFunctionContent(instructionsList, code.MainFunction) != 0 {
		return 1
	}
	return 0
}
