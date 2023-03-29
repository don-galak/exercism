package grep

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

// var flagMap = map[string]bool{"-n": true, "-l": true, "-i":true, "-v":true}

func Search(pattern string, flags, files []string) []string {
	result := make([]string, 0)
	flagMap := make(map[string]bool, len(flags))
	for _, f := range flags {
		flagMap[f] = true
	}

	// fmt.Printf("pattern: %s\nflags %s\nfiles: %v\n", pattern, flags, files)
	fileContents := make(map[string]string, len(files))
	for _, file := range files {
		if fc, ok := os.ReadFile(file); ok == nil {
			fileContents[file] = string(fc)
		}
	}

	for file, fc := range fileContents {
		for i, line := range strings.Split(fc, "\n") {
			if flagMap["-i"] {
				r, _ := regexp.Compile("(?i)" + pattern)
				if r.MatchString(line) {
					result = append(result, line)
				}
			}

			if strings.Contains(line, pattern) {
				switch {
				case flagMap["-n"]:
					line = strconv.Itoa(i+1) + ":" + line
				case flagMap["-l"]:
					line = file
					// case flagMap["-x"]:

				}

				result = append(result, line)
			}
		}
	}

	// fmt.Printf("files contents: \n%s\n\n", fileContents)

	return result
}
