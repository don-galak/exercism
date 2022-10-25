package proverb

import "fmt"

const template = "For want of a %s the %s was lost."
const last = "And all for the want of a %s."

func Proverb(rhyme []string) (out []string) {
	if len(rhyme) == 0 {
		return
	}
	for i := 1; i < len(rhyme); i++ {
		out = append(out, fmt.Sprintf(template, rhyme[i-1], rhyme[i]))
	}
	out = append(out, fmt.Sprintf(last, rhyme[0]))
	return
}
