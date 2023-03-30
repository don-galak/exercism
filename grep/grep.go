package grep

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

// var flagMap = map[string]bool{"-n": true, "-l": true, "-i":true, "-v":true}

func Search(pattern string, flags, files []string) []string {
	resultMap := make(map[string]bool)
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
				pattern = "(?i)" + pattern
			}

			r, _ := regexp.Compile(pattern)

			if r.MatchString(line) {
				switch {
				case flagMap["-l"]:
					line = file
				case !flagMap["-l"] && flagMap["-n"]:
					line = strconv.Itoa(i+1) + ":" + line
				}
				if len(files) > 1 && !flagMap["-l"] {
					line = file + ":" + line
				}

				// ! this wont output the correct result at all times
				// ! later when we append the keys in a slice, the result might not sorted
				// ! and some tests will fail
				resultMap[line] = true
			}
		}
	}
	//
	// fmt.Printf("files contents: \n%s\n\n", fileContents)
	result := make([]string, 0)
	for key := range resultMap {
		result = append(result, key)
	}

	return result
}
