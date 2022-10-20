package tournament

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"sort"
	s "strings"
)

type Team struct {
	Name   string
	Played int
	Won    int
	Drawn  int
	Lost   int
	Points int
}

type Teams []Team

func (teams Teams) Len() int {
	return len(teams)
}
func (teams Teams) Swap(i, j int) {
	teams[i], teams[j] = teams[j], teams[i]
}

func (teams Teams) Less(i, j int) bool {
	if teams[i].Name == teams[j].Name {
		return string(teams[i].Name[0]) > string(teams[j].Name[0])
	}

	return teams[i].Points > teams[j].Points
}

func Tally(reader io.Reader, writer io.Writer) error {
	endsWith_OutcomeReg := regexp.MustCompile(`.*draw|win|loss$`)

	input, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	etled := s.Split(string(input), "\n")
	mapped := make(map[string]*Team)
	const lenOfPadded = 30

	for _, m := range etled {
		if endsWith_OutcomeReg.MatchString(m) {
			values := s.Split(m, ";")
			home := values[0]
			away := values[1]
			outcome := values[2]

			if _, exists := mapped[home]; !exists {
				mapped[home] = &Team{}
			}

			if _, exists := mapped[away]; !exists {
				mapped[away] = &Team{}
			}

			if mapped[home].Name == "" {
				mapped[home].Name = home

				addedMustBe := lenOfPadded - len(home)
				for i := 0; i < addedMustBe; i++ {
					mapped[home].Name += " "
				}
			}

			switch outcome {
			case "win":
				mapped[home].Won += 1
				mapped[home].Points += 3
				mapped[away].Lost += 1
			case "loss":
				mapped[home].Lost += 1
				mapped[away].Won += 1
				mapped[away].Points += 3
			default:
				mapped[home].Drawn += 1
				mapped[away].Drawn += 1
				mapped[home].Points += 1
				mapped[away].Points += 1
			}
			mapped[home].Played += 1
			mapped[away].Played += 1
		}
	}

	var teams Teams
	for _, t := range mapped {
		teams = append(teams, *t)
	}

	sort.Sort(Teams(teams))
	final := "Team                           | MP |  W |  D |  L |  P\n"

	for _, team := range teams {
		final += fmt.Sprintf("%s |  %d |  %d |  %d |  %d |  %d\n", team.Name, team.Played, team.Won, team.Drawn, team.Lost, team.Points)
	}
	writer.Write(bytes.NewBufferString(final).Bytes())

	return nil
}
