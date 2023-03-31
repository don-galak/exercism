package grep

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func addToResult(results *[]string, line string) {
	for _, r := range *results {
		if r == line {
			return
		}
	}
	*results = append(*results, line)
}

func Search(pattern string, flags, files []string) []string {
	result := make([]string, 0)
	flagMap := make(map[string]bool, len(flags))
	for _, f := range flags {
		flagMap[f] = true
	}
	caseInsensitiveMatchWholeLine := flagMap["-i"] && flagMap["-x"]
	matchWholeLine := !flagMap["-i"] && flagMap["-x"]
	caseInsensitive := flagMap["-i"] && !flagMap["-x"]
	defaultMatch := !flagMap["-i"] && !flagMap["-x"]
	inverse := flagMap["-v"]

	fileContents := make(map[string]string, len(files))
	for _, file := range files {
		if fc, ok := os.ReadFile(file); ok == nil {
			fileContents[file] = string(fc)
		}
	}

	for file, fc := range fileContents {
		for i, line := range strings.Split(fc, "\n") {
			matches := false

			if (caseInsensitiveMatchWholeLine && strings.EqualFold(pattern, line)) || matchWholeLine && pattern == line || defaultMatch && strings.Contains(line, pattern) {
				matches = true
			} else if r, _ := regexp.Compile("(?i)" + pattern); r.MatchString(line) && caseInsensitive {
				matches = true
			}

			if matches && !inverse {
				switch {
				case flagMap["-l"]:
					line = file
				case !flagMap["-l"] && flagMap["-n"]:
					line = strconv.Itoa(i+1) + ":" + line
				}
				if len(files) > 1 && !flagMap["-l"] {
					line = file + ":" + line
				}

				addToResult(&result, line)
			}
		}
	}

	return result
}
