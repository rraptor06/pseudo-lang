package convertor

import (
	"fmt"
	"os"
	"pseudo-lang/constant"
	"strings"
)

type Instructions struct {
	Name     string
	Function func(*FunctionStruct, string, *[]string) string
}

func analyseFunctionContent(instructionsList []*Instructions, function *FunctionStruct) int {
	var newLine string
	var indentationList []string
	added := false

	for _, line := range function.Content {
		added = false
		newLine = ""
		if line == "" {
			continue
		}
		for _, instruction := range instructionsList {
			if strings.Contains(line, instruction.Name) {
				newLine = instruction.Function(function, line, &indentationList)
				if newLine == "" {
					//return 1
				}
				function.ConvertedContent = append(function.ConvertedContent, newLine)
				added = true
				break
			}
		}
		if added == false && strings.Contains(line, "<-") {
			newLine = analyseVariable(function, line, &indentationList)
			if newLine == "" {
				//return 1
			}
			function.ConvertedContent = append(function.ConvertedContent, newLine)
			added = true
		}
		if added == false && strings.Contains(line, "(") {
			newLine = analyseFunction(function, line, &indentationList)
			if newLine == "" {
				//return 1
			}
			function.ConvertedContent = append(function.ConvertedContent, newLine)
			added = true
		}
		if added == false {
			fmt.Fprintf(os.Stderr, "%sERROR: Can't convert the line \"%s\" !\n%s", constant.ErrorColor, line, constant.ResetColor)
			//return 1
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
		//return 1
	}
	return 0
}
