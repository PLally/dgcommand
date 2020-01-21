package dgcommand

import "strings"

func parseCommand(cmd string) []string {
	i := 0
	var args []string
	start := 0
	end := 0
	for i < len(cmd) {
		char := cmd[i]
		if char == '"' {
			i++
			end = strings.Index( cmd[i:],`"`)
			if end == -1 {
				end = len(cmd) - 1
			} else {
				end = i+end
			}
			args = append(args, cmd[i:end])
			i = end+1
			start = i+1
		} else if char == ' ' {
			args = append(args, cmd[start:i])
			start = i+1
		}
		i++
	}
	if i <= len(cmd) {
		args = append(args, cmd[start:i])
	}

	return args
}

func firstWord(s string) string {
	i := strings.IndexByte(s, " "[0])
	if i == -1 {
		return s
	} else {
		return s[:i]
	}
}
