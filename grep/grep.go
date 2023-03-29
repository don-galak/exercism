package grep

import "fmt"

func Search(pattern string, flags, files []string) []string {
	fmt.Printf("pattern: %s\nflags %s\nfiles: %v\n\n", pattern, flags, files)

	return []string{}
}
