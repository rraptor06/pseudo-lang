package convertor

// IsIndented Check if a line is indented
//
/* line: The line to check */
//
// Returns true if the line is indented, false otherwise
func IsIndented(line string) bool {
	space := 0

	for _, char := range line {
		if char != '\t' && char != ' ' {
			return false
		}
		if char == '\t' {
			space += 4
		} else {
			space++
		}
		if space >= 4 {
			break
		}
	}
	return true
}

func RemoveIndentation(line string) string {
	index := 0
	spaces := 0

	for _, char := range line {
		if char != '\t' && char != ' ' {
			break
		}
		if char == '\t' {
			spaces += 4
		} else {
			spaces++
		}
		index++
		if spaces >= 4 {
			break
		}
	}
	return line[index:]
}
