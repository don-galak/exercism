package foodchain

import (
	"bytes"
	"fmt"
)

const oldLadySwallowed = "I know an old lady who swallowed a "

type thing struct {
	animal    string
	statement string
}

func (t thing) getPhrase() string {
	return fmt.Sprintf("%s%s.\n%s", oldLadySwallowed, t.animal, t.statement)
}

var verseSlice = []thing{
	{},
	{animal: "fly", statement: "I don't know why she swallowed the fly. Perhaps she'll die."},
	{animal: "spider", statement: "It wriggled and jiggled and tickled inside her."},
	{animal: "bird", statement: "How absurd to swallow a bird!"},
	{animal: "cat", statement: "Imagine that, to swallow a cat!"},
	{animal: "dog", statement: "What a hog, to swallow a dog!"},
	{animal: "goat", statement: "Just opened her throat and swallowed a goat!"},
	{animal: "cow", statement: "I don't know how she swallowed a cow!"},
	{animal: "horse", statement: "She's dead, of course!"},
}

func Verse(v int) string {
	var b bytes.Buffer

	b.WriteString(verseSlice[v].getPhrase())

	for i := v; v > 1; v-- {
		b.WriteString(fmt.Sprintf("\nShe swallowed the %s to catch the %s.", verseSlice[i].animal, verseSlice[i-1].animal))
	}

	return b.String()
}

func Verses(start, end int) string {
	panic("Please implement the Verses function")
}

func Song() string {
	panic("Please implement the Song function")
}
