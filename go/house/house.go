package house

var starters = []string{
	"",
	"This is the house that Jack built.",
	"This is the malt",
	"This is the rat",
	"This is the cat",
	"This is the dog",
	"This is the cow with the crumpled horn",
	"This is the maiden all forlorn",
	"This is the man all tattered and torn",
	"This is the priest all shaven and shorn",
	"This is the rooster that crowed in the morn",
	"This is the farmer sowing his corn",
	"This is the horse and the hound and the horn",
}

var lines = []string{
	"that belonged to the farmer sowing his corn",
	"that kept the rooster that crowed in the morn",
	"that woke the priest all shaven and shorn",
	"that married the man all tattered and torn",
	"that kissed the maiden all forlorn",
	"that milked the cow with the crumpled horn",
	"that tossed the dog",
	"that worried the cat",
	"that killed the rat",
	"that ate the malt",
	"that lay in the house that Jack built.",
}

func Verse(v int) string {
	verse := starters[v]

	for i := v; i > 1; i-- {
		verse += "\n" + lines[len(lines)-i+1]
	}

	return verse
}

func Song() (s string) {
	i := 1
	for ; i < len(starters)-1; i++ {
		s += Verse(i) + "\n\n"
	}
	s += Verse(i)
	return
}
