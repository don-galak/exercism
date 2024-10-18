package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"sort"
	s "strings"
)

const lenOfPadded = 30

type Team struct {
	Name   string
	Played int
	Won    int
	Drawn  int
	Lost   int
	Points int
}

func (team *Team) padWithSpaces(name string) {
	if team.Name == "" {
		team.Name = name
		addedMustBe := lenOfPadded - len(name)
		for i := 0; i < addedMustBe; i++ {
			team.Name += " "
		}
	}
}

type Teams []Team

func (teams Teams) Len() int {
	return len(teams)
}
func (teams Teams) Swap(i, j int) {
	teams[i], teams[j] = teams[j], teams[i]
}

func (teams Teams) Less(i, j int) bool {
	if teams[i].Points == teams[j].Points {
		return string(teams[i].Name[0]) < string(teams[j].Name[0])
	}

	return teams[i].Points > teams[j].Points
}

func Tally(reader io.Reader, writer io.Writer) error {
	isValidRegex := regexp.MustCompile(`[^@_]^.*draw|win|loss.*$`)
	endsWithOutcomeReg := regexp.MustCompile(`.*draw|win|loss$`)

	input, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	if !isValidRegex.MatchString(string(input)) {
		return errors.New("invalid input")
	}

	extracted := s.Split(string(input), "\n")
	teamMap := make(map[string]*Team)

	for _, m := range extracted {
		if endsWithOutcomeReg.MatchString(m) {
			values := s.Split(m, ";")
			home := values[0]
			away := values[1]
			outcome := values[2]

			if _, exists := teamMap[home]; !exists {
				teamMap[home] = &Team{}
			}

			if _, exists := teamMap[away]; !exists {
				teamMap[away] = &Team{}
			}

			teamMap[home].padWithSpaces(home)
			teamMap[away].padWithSpaces(away)

			switch outcome {
			case "win":
				teamMap[home].Won += 1
				teamMap[home].Points += 3
				teamMap[away].Lost += 1
			case "loss":
				teamMap[home].Lost += 1
				teamMap[away].Won += 1
				teamMap[away].Points += 3
			default:
				teamMap[home].Drawn += 1
				teamMap[away].Drawn += 1
				teamMap[home].Points += 1
				teamMap[away].Points += 1
			}
			teamMap[home].Played += 1
			teamMap[away].Played += 1
		}
	}

	var teams Teams
	for _, t := range teamMap {
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
