package grep

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func hasValue(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func compileRegexp(pattern string, flags []string) (*regexp.Regexp, error) {
	matchWholeLine := hasValue(flags, "-x")
	if matchWholeLine {
		pattern = fmt.Sprintf("^%s$", pattern)
	}
	caseInsensitive := hasValue(flags, "-i")
	if caseInsensitive {
		pattern = "(?i)" + pattern
	}
	return regexp.Compile(pattern)
}

func Search(pattern string, flags, files []string) []string {
	result := make([]string, 0)
	withMultipleFiles := len(files) > 1
	withLineNumbers := hasValue(flags, "-n")
	withFileNames := hasValue(flags, "-l")
	invert := hasValue(flags, "-v")
	exp, e := compileRegexp(pattern, flags)
	if e != nil {
		log.Fatal(e)
		return nil
	}

	for _, file := range files {
		f, e := os.Open(file)
		if e != nil {
			log.Fatal(e)
			return nil
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		lineNum := 0
		for b, _, e := reader.ReadLine(); e == nil; b, _, e = reader.ReadLine() {
			line := string(b)
			lineNum++
			if exp.MatchString(line) != invert {
				if withLineNumbers {
					line = strconv.Itoa(lineNum) + ":" + line
				}
				if withMultipleFiles && !withFileNames {
					line = file + ":" + line
				}
				if withFileNames && !hasValue(result, file) {
					result = append(result, file)
				} else if !withFileNames {
					result = append(result, line)
				}
			}
		}
	}
	return result
}
