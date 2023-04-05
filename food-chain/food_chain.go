package foodchain

import (
	"bytes"
	"fmt"
)

type thing struct {
	animal    string
	statement string
}

func (t thing) getPhrase() string {
	return fmt.Sprintf("I know an old lady who swallowed a %s.\n%s", t.animal, t.statement)
}

var verseSlice = []thing{
	{},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her."},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
	{"horse", "She's dead, of course!"},
}

func composeLine(animal1, animal2 string) string {
	line := fmt.Sprintf("\nShe swallowed the %s to catch the %s", animal1, animal2)
	if animal1 != "bird" {
		return line + "."
	}
	return line + " that wriggled and jiggled and tickled inside her."
}

func Verse(v int) string {
	phrase := verseSlice[v].getPhrase()
	if v == 1 || v == 8 {
		return phrase
	}

	var b bytes.Buffer

	b.WriteString(phrase)

	for i := v; i > 1; i-- {
		b.WriteString(composeLine(verseSlice[i].animal, verseSlice[i-1].animal))
	}
	b.WriteString("\n" + verseSlice[1].statement)

	return b.String()
}

func Verses(start, end int) string {
	var verses bytes.Buffer

	i := start
	for ; i < end; i++ {
		verses.WriteString(Verse(i))
		verses.WriteString("\n\n")
	}
	verses.WriteString(Verse(i))

	return verses.String()
}

func Song() string {
	return Verses(1, 8)
}
